package dao

import (
	"fmt"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListPivot(database *gorm.DB) (pivots []models.Pivot, err error) {
	err = database.Find(&pivots).Error

	return
}

func ReadPivot(database *gorm.DB, id int) (pivot models.Pivot, err error) {
	err = database.First(&pivot, id).Error

	return
}

func ReadPivotByName(database *gorm.DB, name string) (pivot models.Pivot, err error) {
	fmt.Println("DAO")
	err = database.Where("name = ?", name).First(&pivot).Error
	fmt.Println(err)
	fmt.Println(pivot)
	return
}
