package dao

import (
	"fmt"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListProductConnector(database *gorm.DB) (productConnectors []models.ProductConnector, err error) {
	err = database.Find(&productConnectors).Error

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
