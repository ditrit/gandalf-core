package dao

import (
	"errors"

	"github.com/ditrit/gandalf/core/cluster/api/utils"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListConfigurationAggregator(database *gorm.DB) (configurationAggregators []models.ConfigurationAggregator, err error) {
	err = database.Find(&configurationAggregators).Error

	return
}

func CreateConfigurationAggregator(database *gorm.DB, configurationAggregator models.ConfigurationAggregator) (err error) {
	admin, err := utils.GetStateGandalf(database)
	if err == nil {
		if admin {
			err = database.Create(&configurationAggregator).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}

func ReadConfigurationAggregator(database *gorm.DB, id int) (configurationAggregator models.ConfigurationAggregator, err error) {
	err = database.First(&configurationAggregator, id).Error

	return
}

func UpdateConfigurationAggregator(database *gorm.DB, configurationAggregator models.ConfigurationAggregator) (err error) {
	err = database.Save(&configurationAggregator).Error

	return
}

func DeleteConfigurationAggregator(database *gorm.DB, id int) (err error) {
	admin, err := utils.GetStateGandalf(database)
	if err == nil {
		if admin {
			var configurationAggregator models.ConfigurationAggregator
			err = database.Delete(&configurationAggregator, id).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}
