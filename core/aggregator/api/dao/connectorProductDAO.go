package dao

import (
	"errors"
	"fmt"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListConnectorProduct(database *gorm.DB) (connectorProducts []models.ConnectorProduct, err error) {
	err = database.Find(&connectorProducts).Error

	return
}

func CreateConnectorProduct(database *gorm.DB, connectorProduct models.ConnectorProduct) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			err = database.Create(&connectorProduct).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}

func ReadConnectorProduct(database *gorm.DB, id int) (connectorProduct models.ConnectorProduct, err error) {
	err = database.First(&connectorProduct, id).Error

	return
}

func ReadConnectorProductByName(database *gorm.DB, name string) (connectorProduct models.ConnectorProduct, err error) {
	fmt.Println("DAO")
	err = database.Where("name = ?", name).First(&connectorProduct).Error
	fmt.Println(err)
	fmt.Println(connectorProduct)
	return
}

func UpdateConnectorProduct(database *gorm.DB, connectorProduct models.ConnectorProduct) (err error) {
	err = database.Save(&connectorProduct).Error

	return
}

func DeleteConnectorProduct(database *gorm.DB, id int) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			var connectorProduct models.ConnectorProduct
			err = database.Delete(&connectorProduct, id).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}
