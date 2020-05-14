//Package database :
package database

import (
	"core/models"
	"log"

	"github.com/jinzhu/gorm"
)

var gandalfDatabaseClient *gorm.DB = nil

// NewGandalfDatabaseClient : Database client constructor.
func NewGandalfDatabaseClient(tenant, databasePath string) *gorm.DB {

	if gandalfDatabaseClient == nil {
		gandalfDatabaseClient, err := gorm.Open("sqlite3", databasePath+"/gandalf.db")

		if err != nil {
			log.Println("failed to connect database")
		}

		InitGandalfDatabase(gandalfDatabaseClient)
		DemoPopulateGandalfDatabase(gandalfDatabaseClient)
	}
	return gandalfDatabaseClient
}

// InitGandalfDatabase : Gandalf database init.
func InitGandalfDatabase(databaseClient *gorm.DB) (err error) {
	databaseClient.AutoMigrate(&models.Aggregator{}, &models.Application{},
		&models.ConnectorType{}, &models.Connector{}, &models.Event{}, &models.Command{})

	return
}

// DemoPopulateGandalfDatabase : Populate database demo.
func DemoPopulateGandalfDatabase(databaseClient *gorm.DB) {
	databaseClient.Create(&models.Aggregator{Name: "Aggregator1"})
	databaseClient.Create(&models.Aggregator{Name: "Aggregator2"})
	databaseClient.Create(&models.Aggregator{Name: "Aggregator3"})
	databaseClient.Create(&models.Aggregator{Name: "Aggregator4"})

	databaseClient.Create(&models.Connector{Name: "Connector1"})
	databaseClient.Create(&models.Connector{Name: "Connector2"})
	databaseClient.Create(&models.Connector{Name: "Connector3"})
	databaseClient.Create(&models.Connector{Name: "Connector4"})

	databaseClient.Create(&models.ConnectorType{Name: "Utils"})
	databaseClient.Create(&models.ConnectorType{Name: "Workflow"})
	databaseClient.Create(&models.ConnectorType{Name: "Gitlab"})
	databaseClient.Create(&models.ConnectorType{Name: "Azure"})

	var Aggregator models.Aggregator

	var Connector models.Connector

	var ConnectorType models.ConnectorType

	databaseClient.Where("name = ?", "Aggregator1").First(&Aggregator)
	databaseClient.Where("name = ?", "Connector1").First(&Connector)
	databaseClient.Where("name = ?", "Utils").First(&ConnectorType)

	databaseClient.Create(&models.Application{Name: "Application1",
		Aggregator:    "Aggregator1",
		Connector:     "Connector1",
		ConnectorType: "Utils"})

	databaseClient.Where("name = ?", "Aggregator2").First(&Aggregator)
	databaseClient.Where("name = ?", "Connector2").First(&Connector)
	databaseClient.Where("name = ?", "Workflow").First(&ConnectorType)

	databaseClient.Create(&models.Application{Name: "Application2",
		Aggregator:    "Aggregator2",
		Connector:     "Connector2",
		ConnectorType: "Workflow"})

	databaseClient.Where("name = ?", "Aggregator3").First(&Aggregator)
	databaseClient.Where("name = ?", "Connector3").First(&Connector)
	databaseClient.Where("name = ?", "Gitlab").First(&ConnectorType)

	databaseClient.Create(&models.Application{Name: "Application3",
		Aggregator:    "Aggregator3",
		Connector:     "Connector3",
		ConnectorType: "Gitlab"})

	databaseClient.Where("name = ?", "Aggregator4").First(&Aggregator)
	databaseClient.Where("name = ?", "Connector4").First(&Connector)
	databaseClient.Where("name = ?", "Azure").First(&ConnectorType)

	databaseClient.Create(&models.Application{Name: "Application4",
		Aggregator:    "Aggregator4",
		Connector:     "Connector4",
		ConnectorType: "Azure"})
}