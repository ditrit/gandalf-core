package dao

import (
	"errors"
	"fmt"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListEventTypeToPoll(database *gorm.DB) (eventTypeToPolls []models.EventTypeToPoll, err error) {
	err = database.Find(&eventTypeToPolls).Error

	return
}

func CreateEventTypeToPoll(database *gorm.DB, eventTypeToPoll models.EventTypeToPoll) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			err = database.Create(&eventTypeToPoll).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}

func ReadEventTypeToPoll(database *gorm.DB, id int) (eventTypeToPoll models.EventTypeToPoll, err error) {
	err = database.First(&eventTypeToPoll, id).Error

	return
}

func ReadEventTypeToPollByName(database *gorm.DB, name string) (eventTypeToPoll models.EventTypeToPoll, err error) {
	fmt.Println("DAO")
	err = database.Where("name = ?", name).First(&eventTypeToPoll).Error
	fmt.Println(err)
	fmt.Println(eventTypeToPoll)
	return
}

func UpdateEventTypeToPoll(database *gorm.DB, eventTypeToPoll models.EventTypeToPoll) (err error) {
	err = database.Save(&eventTypeToPoll).Error

	return
}

func DeleteEventTypeToPoll(database *gorm.DB, id int) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			var eventTypeToPoll models.EventTypeToPoll
			err = database.Delete(&eventTypeToPoll, id).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}