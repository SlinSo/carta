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

var FilmList = newFilmListTable("sakila", "film_list", "")

type filmListTable struct {
	mysql.Table

	// Columns
	Fid         mysql.ColumnInteger
	Title       mysql.ColumnString
	Description mysql.ColumnString
	Category    mysql.ColumnString
	Price       mysql.ColumnFloat
	Length      mysql.ColumnInteger
	Rating      mysql.ColumnString
	Actors      mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type FilmListTable struct {
	filmListTable

	NEW filmListTable
}

// AS creates new FilmListTable with assigned alias
func (a FilmListTable) AS(alias string) *FilmListTable {
	return newFilmListTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new FilmListTable with assigned schema name
func (a FilmListTable) FromSchema(schemaName string) *FilmListTable {
	return newFilmListTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new FilmListTable with assigned table prefix
func (a FilmListTable) WithPrefix(prefix string) *FilmListTable {
	return newFilmListTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new FilmListTable with assigned table suffix
func (a FilmListTable) WithSuffix(suffix string) *FilmListTable {
	return newFilmListTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newFilmListTable(schemaName, tableName, alias string) *FilmListTable {
	return &FilmListTable{
		filmListTable: newFilmListTableImpl(schemaName, tableName, alias),
		NEW:           newFilmListTableImpl("", "new", ""),
	}
}

func newFilmListTableImpl(schemaName, tableName, alias string) filmListTable {
	var (
		FidColumn         = mysql.IntegerColumn("FID")
		TitleColumn       = mysql.StringColumn("title")
		DescriptionColumn = mysql.StringColumn("description")
		CategoryColumn    = mysql.StringColumn("category")
		PriceColumn       = mysql.FloatColumn("price")
		LengthColumn      = mysql.IntegerColumn("length")
		RatingColumn      = mysql.StringColumn("rating")
		ActorsColumn      = mysql.StringColumn("actors")
		allColumns        = mysql.ColumnList{FidColumn, TitleColumn, DescriptionColumn, CategoryColumn, PriceColumn, LengthColumn, RatingColumn, ActorsColumn}
		mutableColumns    = mysql.ColumnList{FidColumn, TitleColumn, DescriptionColumn, CategoryColumn, PriceColumn, LengthColumn, RatingColumn, ActorsColumn}
	)

	return filmListTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		Fid:         FidColumn,
		Title:       TitleColumn,
		Description: DescriptionColumn,
		Category:    CategoryColumn,
		Price:       PriceColumn,
		Length:      LengthColumn,
		Rating:      RatingColumn,
		Actors:      ActorsColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
