//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/mysql"
)

var Store = newStoreTable("sakila", "store", "")

type storeTable struct {
	mysql.Table

	// Columns
	StoreID        mysql.ColumnInteger
	ManagerStaffID mysql.ColumnInteger
	AddressID      mysql.ColumnInteger
	LastUpdate     mysql.ColumnTimestamp

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type StoreTable struct {
	storeTable

	NEW storeTable
}

// AS creates new StoreTable with assigned alias
func (a StoreTable) AS(alias string) *StoreTable {
	return newStoreTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new StoreTable with assigned schema name
func (a StoreTable) FromSchema(schemaName string) *StoreTable {
	return newStoreTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new StoreTable with assigned table prefix
func (a StoreTable) WithPrefix(prefix string) *StoreTable {
	return newStoreTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new StoreTable with assigned table suffix
func (a StoreTable) WithSuffix(suffix string) *StoreTable {
	return newStoreTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newStoreTable(schemaName, tableName, alias string) *StoreTable {
	return &StoreTable{
		storeTable: newStoreTableImpl(schemaName, tableName, alias),
		NEW:        newStoreTableImpl("", "new", ""),
	}
}

func newStoreTableImpl(schemaName, tableName, alias string) storeTable {
	var (
		StoreIDColumn        = mysql.IntegerColumn("store_id")
		ManagerStaffIDColumn = mysql.IntegerColumn("manager_staff_id")
		AddressIDColumn      = mysql.IntegerColumn("address_id")
		LastUpdateColumn     = mysql.TimestampColumn("last_update")
		allColumns           = mysql.ColumnList{StoreIDColumn, ManagerStaffIDColumn, AddressIDColumn, LastUpdateColumn}
		mutableColumns       = mysql.ColumnList{ManagerStaffIDColumn, AddressIDColumn, LastUpdateColumn}
	)

	return storeTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		StoreID:        StoreIDColumn,
		ManagerStaffID: ManagerStaffIDColumn,
		AddressID:      AddressIDColumn,
		LastUpdate:     LastUpdateColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
