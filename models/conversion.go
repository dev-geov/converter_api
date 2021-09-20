package models

import (
	"gorm.io/gorm"
)

type Conversion struct {
	gorm.Model
	Amount float64 `json:"amount"`
	From   string  `json:"from"`
	To     string  `json:"to"`
	Rate   float64 `json:"rate"`
}
