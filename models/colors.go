package models

import "github.com/jinzhu/gorm"

//Color represents an object on db using gorm.
type Color struct {
	gorm.Model
	Hex string `json:"hex"`
	R   uint8  `json:"r"`
	G   uint8  `json:"g"`
	B   uint8  `json:"b"`
}
