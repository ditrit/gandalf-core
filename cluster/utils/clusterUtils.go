//Package utils :
package utils

import (
	"gandalf-core/database"
	"gandalf-core/models"
	"log"
	"shoset/msg"

	"github.com/jinzhu/gorm"
)

var gandalfDatabaseClient *gorm.DB = nil

// GetDatabaseClientByTenant : Cluster database client getter by tenant.
func GetDatabaseClientByTenant(tenant, databasePath string, mapDatabaseClient map[string]*gorm.DB) *gorm.DB {
	if _, ok := mapDatabaseClient[tenant]; !ok {
		mapDatabaseClient[tenant] = database.NewTenantDatabaseClient(tenant, databasePath)
	}

	return mapDatabaseClient[tenant]
}

// GetGandalfDatabaseClient : Database client constructor.
func GetGandalfDatabaseClient(databasePath string) *gorm.DB {

	if gandalfDatabaseClient == nil {
		gandalfDatabaseClient = database.NewGandalfDatabaseClient(databasePath)
	}
	return gandalfDatabaseClient
}

// GetApplicationContext : Cluster application context getter.
func GetApplicationContext(cmd msg.Command, client *gorm.DB) (applicationContext models.Application) {
	client.Where("connector_type = ?", cmd.GetContext()["ConnectorType"].(string)).First(&applicationContext)

	return
}

// GetConnectorConfiguration : Cluster application context getter.
func GetConnectorConfiguration(cmd msg.Command, client *gorm.DB) (applicationContext models.Application) {
	//client.Where("connector_type = ?", cmd.GetContext()["ConnectorType"].(string)).First(&applicationContext)

	return
}

// CaptureMessage : Cluster capture message function.
func CaptureMessage(message msg.Message, msgType string, client *gorm.DB) bool {
	ok := true

	switch msgType {
	case "cmd":
		currentMsg := models.FromShosetCommand(message.(msg.Command))
		client.Create(&currentMsg)
	case "evt":
		currentMsg := models.FromShosetEvent(message.(msg.Event))
		client.Create(&currentMsg)
	default:
		ok = false

		log.Println("Can't capture this message")
	}

	return ok
}
