//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type FilmCategory struct {
	FilmID     uint16 `sql:"primary_key"`
	CategoryID uint8  `sql:"primary_key"`
	LastUpdate time.Time
}
