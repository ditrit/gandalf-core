//Package database :
package database

import (
	"gandalf-core/models"
	"log"

	"github.com/jinzhu/gorm"
)

// NewGandalfDatabaseClient : Database client constructor.
func NewGandalfDatabaseClient(databasePath string) *gorm.DB {

	gandalfDatabaseClient, err := gorm.Open("sqlite3", databasePath+"/gandalf.db")

	if err != nil {
		log.Println("failed to connect database")
	}

	InitGandalfDatabase(gandalfDatabaseClient)

	DemoPopulateGandalfDatabase(gandalfDatabaseClient)

	return gandalfDatabaseClient
}

// InitGandalfDatabase : Gandalf database init.
func InitGandalfDatabase(databaseClient *gorm.DB) (err error) {
	databaseClient.AutoMigrate(&models.ConnectorConfig{}, &models.ConnectorType{}, &models.ConnectorTypeCommand{})

	return
}

// DemoPopulateGandalfDatabase : Populate database demo.
func DemoPopulateGandalfDatabase(databaseClient *gorm.DB) {

	var ConnectorType models.ConnectorType
	var ConnectorTypeCommands []models.ConnectorTypeCommand

	databaseClient.Create(&models.ConnectorType{Name: "Utils"})
	databaseClient.Where("name = ?", "Utils").First(&ConnectorType)

	databaseClient.Create(&models.ConnectorTypeCommand{Name: "SEND_AUTH_MAIL", Schema: `{"$schema":"http://json-schema.org/draft-04/schema#","$ref":"#/definitions/MailPayload","definitions":{"MailPayload":{"required":["Sender","Body","Receivers","Identity","Username","Password","Host"],"properties":{"Sender":{"type":"string"},"Body":{"type":"string"},"Receivers":{"items":{"type":"string"},"type":"array"},"Identity":{"type":"string"},"Username":{"type":"string"},"Password":{"type":"string"},"Host":{"type":"string"}},"additionalProperties":false,"type":"object"}}}`})
	databaseClient.Create(&models.ConnectorTypeCommand{Name: "CREATE_FORM", Schema: `{"$schema":"http://json-schema.org/draft-04/schema#","$ref":"#/definitions/FormPayload","definitions":{"Field":{"required":["Name","HtmlType","Value"],"properties":{"Name":{"type":"string"},"HtmlType":{"type":"string"},"Value":{"additionalProperties":true}},"additionalProperties":false,"type":"object"},"FormPayload":{"required":["Fields"],"properties":{"Fields":{"items":{"$schema":"http://json-schema.org/draft-04/schema#","$ref":"#/definitions/Field"},"type":"array"}},"additionalProperties":false,"type":"object"}}}`})

	databaseClient.Where("name IN (?)", []string{"SEND_AUTH_MAIL", "CREATE_FORM"}).Find(&ConnectorTypeCommands)

	databaseClient.Create(&models.ConnectorConfig{Name: "ConnectorConfig1",
		ConnectorTypeID:       ConnectorType.ID,
		ConnectorTypeCommands: ConnectorTypeCommands})

	databaseClient.Create(&models.ConnectorType{Name: "Workflow"})
	databaseClient.Where("name = ?", "Workflow").First(&ConnectorType)

	databaseClient.Where("name IN (?)", []string{}).Find(&ConnectorTypeCommands)

	databaseClient.Create(&models.ConnectorConfig{Name: "ConnectorConfig2",
		ConnectorTypeID:       ConnectorType.ID,
		ConnectorTypeCommands: []models.ConnectorTypeCommand{}})

	databaseClient.Create(&models.ConnectorType{Name: "Gitlab"})
	databaseClient.Where("name = ?", "Gitlab").First(&ConnectorType)

	databaseClient.Create(&models.ConnectorTypeCommand{Name: "Gitlab1", Schema: ""})
	databaseClient.Create(&models.ConnectorTypeCommand{Name: "Gitlab2", Schema: ""})
	databaseClient.Create(&models.ConnectorTypeCommand{Name: "Gitlab3", Schema: ""})

	databaseClient.Where("name IN (?)", []string{"Gitlab1", "Gitlab2", "Gitlab3"}).Find(&ConnectorTypeCommands)

	databaseClient.Create(&models.ConnectorConfig{Name: "ConnectorConfig3",
		ConnectorTypeID:       ConnectorType.ID,
		ConnectorTypeCommands: ConnectorTypeCommands})

	databaseClient.Create(&models.ConnectorType{Name: "Azure"})
	databaseClient.Where("name = ?", "Azure").First(&ConnectorType)

	databaseClient.Create(&models.ConnectorTypeCommand{Name: "CREATE_VM_BY_JSON", Schema: ""})

	databaseClient.Where("name IN (?)", []string{"CREATE_VM_BY_JSON"}).Find(&ConnectorTypeCommands)

	databaseClient.Create(&models.ConnectorConfig{Name: "ConnectorConfig4",
		ConnectorTypeID:       ConnectorType.ID,
		ConnectorTypeCommands: ConnectorTypeCommands})
	databaseClient.Create(&models.ConnectorType{Name: "Workflow"})

}
