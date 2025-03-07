package carta

import (
	"database/sql"
	"sort"
	"strings"

	"github.com/seambiz/strcase"
)

// column represents the ith struct field of this mapper where the column is to be mapped
type column struct {
	typ         *sql.ColumnType
	name        string
	columnIndex int
	i           fieldIndex
}

func allocateColumns(m *Mapper, columns map[string]column) error {
	presentColumns := map[string]column{}
	for cName, c := range columns {
		if m.IsBasic {
			candidate := getSingleColumnNameCandidate("", m.AncestorName)
			if cName == candidate {
				presentColumns[cName] = column{
					typ:         c.typ,
					name:        cName,
					columnIndex: c.columnIndex,
				}
				delete(columns, cName) // dealocate claimed column
			}
		} else {
			for i, field := range m.Fields {
				// can only allocate columns to basic fields
				if !field.IsMapped && isBasicType(field.Typ) {
					field := m.Fields[i]
					candidate := getSingleColumnNameCandidate(field.Name, m.AncestorName)
					if cName == candidate {
						presentColumns[cName] = column{
							typ:         c.typ,
							name:        cName,
							columnIndex: c.columnIndex,
							i:           i,
						}
						field.IsMapped = true
						delete(columns, cName) // dealocate claimed column
						break
					}
				}
			}
		}
	}
	m.PresentColumns = presentColumns

	columnIds := []int{}
	for _, column := range m.PresentColumns {
		if _, ok := m.SubMaps[column.i]; ok {
			continue
		}
		columnIds = append(columnIds, column.columnIndex)
	}
	sort.Ints(columnIds)
	m.SortedColumnIndexes = columnIds

	for i, subMap := range m.SubMaps {
		subMap.AncestorName = m.Fields[i].Name
		if err := allocateColumns(subMap, columns); err != nil {
			return err
		}
	}

	return nil
}

func getSingleColumnNameCandidate(fieldName string, ancestorName string) string {
	if fieldName == "" {
		return ancestorName
	} else if ancestorName == "" {
		return strcase.ToSnake(fieldName)
	}

	return strings.ToLower(ancestorName) + "." + strcase.ToSnake(fieldName)
}
