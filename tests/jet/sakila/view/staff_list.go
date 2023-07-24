//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package view

import (
	"github.com/go-jet/jet/v2/mysql"
)

var StaffList = newStaffListTable("sakila", "staff_list", "")

type staffListTable struct {
	mysql.Table

	// Columns
	ID      mysql.ColumnInteger
	Name    mysql.ColumnString
	Address mysql.ColumnString
	ZipCode mysql.ColumnString
	Phone   mysql.ColumnString
	City    mysql.ColumnString
	Country mysql.ColumnString
	Sid     mysql.ColumnInteger

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type StaffListTable struct {
	staffListTable

	NEW staffListTable
}

// AS creates new StaffListTable with assigned alias
func (a StaffListTable) AS(alias string) *StaffListTable {
	return newStaffListTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new StaffListTable with assigned schema name
func (a StaffListTable) FromSchema(schemaName string) *StaffListTable {
	return newStaffListTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new StaffListTable with assigned table prefix
func (a StaffListTable) WithPrefix(prefix string) *StaffListTable {
	return newStaffListTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new StaffListTable with assigned table suffix
func (a StaffListTable) WithSuffix(suffix string) *StaffListTable {
	return newStaffListTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newStaffListTable(schemaName, tableName, alias string) *StaffListTable {
	return &StaffListTable{
		staffListTable: newStaffListTableImpl(schemaName, tableName, alias),
		NEW:            newStaffListTableImpl("", "new", ""),
	}
}

func newStaffListTableImpl(schemaName, tableName, alias string) staffListTable {
	var (
		IDColumn       = mysql.IntegerColumn("ID")
		NameColumn     = mysql.StringColumn("name")
		AddressColumn  = mysql.StringColumn("address")
		ZipCodeColumn  = mysql.StringColumn("zip code")
		PhoneColumn    = mysql.StringColumn("phone")
		CityColumn     = mysql.StringColumn("city")
		CountryColumn  = mysql.StringColumn("country")
		SidColumn      = mysql.IntegerColumn("SID")
		allColumns     = mysql.ColumnList{IDColumn, NameColumn, AddressColumn, ZipCodeColumn, PhoneColumn, CityColumn, CountryColumn, SidColumn}
		mutableColumns = mysql.ColumnList{IDColumn, NameColumn, AddressColumn, ZipCodeColumn, PhoneColumn, CityColumn, CountryColumn, SidColumn}
	)

	return staffListTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:      IDColumn,
		Name:    NameColumn,
		Address: AddressColumn,
		ZipCode: ZipCodeColumn,
		Phone:   PhoneColumn,
		City:    CityColumn,
		Country: CountryColumn,
		Sid:     SidColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
