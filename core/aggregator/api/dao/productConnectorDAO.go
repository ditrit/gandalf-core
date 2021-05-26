package dao

import (
	"errors"
	"fmt"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListProductConnector(database *gorm.DB) (productConnectors []models.ProductConnector, err error) {
	err = database.Find(&productConnectors).Error

	return
}

func CreateProductConnector(database *gorm.DB, productConnector models.ProductConnector) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			err = database.Create(&productConnector).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}

func ReadProductConnector(database *gorm.DB, id int) (productConnector models.ProductConnector, err error) {
	err = database.First(&productConnector, id).Error

	return
}

func ReadProductConnectorByName(database *gorm.DB, name string) (productConnector models.ProductConnector, err error) {
	fmt.Println("DAO")
	err = database.Where("name = ?", name).First(&productConnector).Error
	fmt.Println(err)
	fmt.Println(productConnector)
	return
}

func UpdateProductConnector(database *gorm.DB, productConnector models.ProductConnector) (err error) {
	err = database.Save(&productConnector).Error

	return
}

func DeleteProductConnector(database *gorm.DB, id int) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			var productConnector models.ProductConnector
			err = database.Delete(&productConnector, id).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}
