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

var Film = newFilmTable("sakila", "film", "")

type filmTable struct {
	mysql.Table

	// Columns
	FilmID             mysql.ColumnInteger
	Title              mysql.ColumnString
	Description        mysql.ColumnString
	ReleaseYear        mysql.ColumnInteger
	LanguageID         mysql.ColumnInteger
	OriginalLanguageID mysql.ColumnInteger
	RentalDuration     mysql.ColumnInteger
	RentalRate         mysql.ColumnFloat
	Length             mysql.ColumnInteger
	ReplacementCost    mysql.ColumnFloat
	Rating             mysql.ColumnString
	SpecialFeatures    mysql.ColumnString
	LastUpdate         mysql.ColumnTimestamp

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type FilmTable struct {
	filmTable

	NEW filmTable
}

// AS creates new FilmTable with assigned alias
func (a FilmTable) AS(alias string) *FilmTable {
	return newFilmTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new FilmTable with assigned schema name
func (a FilmTable) FromSchema(schemaName string) *FilmTable {
	return newFilmTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new FilmTable with assigned table prefix
func (a FilmTable) WithPrefix(prefix string) *FilmTable {
	return newFilmTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new FilmTable with assigned table suffix
func (a FilmTable) WithSuffix(suffix string) *FilmTable {
	return newFilmTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newFilmTable(schemaName, tableName, alias string) *FilmTable {
	return &FilmTable{
		filmTable: newFilmTableImpl(schemaName, tableName, alias),
		NEW:       newFilmTableImpl("", "new", ""),
	}
}

func newFilmTableImpl(schemaName, tableName, alias string) filmTable {
	var (
		FilmIDColumn             = mysql.IntegerColumn("film_id")
		TitleColumn              = mysql.StringColumn("title")
		DescriptionColumn        = mysql.StringColumn("description")
		ReleaseYearColumn        = mysql.IntegerColumn("release_year")
		LanguageIDColumn         = mysql.IntegerColumn("language_id")
		OriginalLanguageIDColumn = mysql.IntegerColumn("original_language_id")
		RentalDurationColumn     = mysql.IntegerColumn("rental_duration")
		RentalRateColumn         = mysql.FloatColumn("rental_rate")
		LengthColumn             = mysql.IntegerColumn("length")
		ReplacementCostColumn    = mysql.FloatColumn("replacement_cost")
		RatingColumn             = mysql.StringColumn("rating")
		SpecialFeaturesColumn    = mysql.StringColumn("special_features")
		LastUpdateColumn         = mysql.TimestampColumn("last_update")
		allColumns               = mysql.ColumnList{FilmIDColumn, TitleColumn, DescriptionColumn, ReleaseYearColumn, LanguageIDColumn, OriginalLanguageIDColumn, RentalDurationColumn, RentalRateColumn, LengthColumn, ReplacementCostColumn, RatingColumn, SpecialFeaturesColumn, LastUpdateColumn}
		mutableColumns           = mysql.ColumnList{TitleColumn, DescriptionColumn, ReleaseYearColumn, LanguageIDColumn, OriginalLanguageIDColumn, RentalDurationColumn, RentalRateColumn, LengthColumn, ReplacementCostColumn, RatingColumn, SpecialFeaturesColumn, LastUpdateColumn}
	)

	return filmTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		FilmID:             FilmIDColumn,
		Title:              TitleColumn,
		Description:        DescriptionColumn,
		ReleaseYear:        ReleaseYearColumn,
		LanguageID:         LanguageIDColumn,
		OriginalLanguageID: OriginalLanguageIDColumn,
		RentalDuration:     RentalDurationColumn,
		RentalRate:         RentalRateColumn,
		Length:             LengthColumn,
		ReplacementCost:    ReplacementCostColumn,
		Rating:             RatingColumn,
		SpecialFeatures:    SpecialFeaturesColumn,
		LastUpdate:         LastUpdateColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
