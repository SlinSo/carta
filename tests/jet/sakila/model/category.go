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

type Category struct {
	CategoryID uint8 `sql:"primary_key"`
	Name       string
	LastUpdate time.Time
}
