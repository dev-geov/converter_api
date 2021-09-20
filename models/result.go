package models

import (
	"gorm.io/gorm"
)

type Result struct {
	gorm.Model
	Value        float64     `json:"value"`
	CoinSymbol   string      `json:"coin_symbol"`
	ConversionID uint        `json:"conversion_id"`
	Conversion   *Conversion `json:"conversion" gorm:"ForeignKey:ConversionID;AssociationForeignKey:ID"`
}
