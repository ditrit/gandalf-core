package dao

import (
	"errors"
	"fmt"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListTenant(database *gorm.DB) (tenants []models.Tenant, err error) {
	err = database.Find(&tenants).Error

	return
}

func CreateTenant(database *gorm.DB, tenant models.Tenant) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			err = database.Create(&tenant).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}

func ReadTenant(database *gorm.DB, id int) (tenant models.Tenant, err error) {
	err = database.First(&tenant, id).Error

	return
}

func ReadTenantByName(database *gorm.DB, name string) (tenant models.tenant, err error) {
	fmt.Println("DAO")
	err = database.Where("name = ?", name).First(&tenant).Error
	fmt.Println(err)
	fmt.Println(tenant)
	return
}

func UpdateTenant(database *gorm.DB, tenant models.Tenant) (err error) {
	err = database.Save(&tenant).Error

	return
}

func DeleteTenant(database *gorm.DB, id int) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			var tenant models.Tenant
			err = database.Delete(&tenant, id).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}
