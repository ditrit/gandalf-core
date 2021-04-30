package models

import "github.com/jinzhu/gorm"

type KeyValue struct {
	gorm.Model
	Value              string
	KeyID              uint
	Key                Key
	LogicalComponentID uint
	//LogicalConnector   LogicalConnector
}
