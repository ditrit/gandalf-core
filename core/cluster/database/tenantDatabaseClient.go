//Package database :
package database

import (
	"log"

	"github.com/ditrit/gandalf/core/models"

	"github.com/jinzhu/gorm"
)

// NewTenantDatabaseClient : Tenant database client constructor.
func NewTenantDatabaseClient(tenant, databasePath string) (tenantDatabaseClient *gorm.DB, err error) {

	databaseFullPath := databasePath + "/" + tenant + ".db"

	tenantDatabaseClient, err = gorm.Open("sqlite3", databaseFullPath)
	if err != nil {
		log.Println("failed to connect database")
	}

	return

	/*databaseFullPath := databasePath + "/" + tenant + ".db"

	 	if _, err := os.Stat(databaseFullPath); err == nil {
			tenantDatabaseClient, err = gorm.Open("sqlite3", databaseFullPath)

		} else if os.IsNotExist(err) {
			tenantDatabaseClient, err = gorm.Open("sqlite3", databaseFullPath)
			if err != nil {
				log.Println("failed to connect database")
			}

			InitTenantDatabase(tenantDatabaseClient)
		}

		//DemoPopulateTenantDatabase(tenantDatabaseClient) */

	return
}

// InitTenantDatabase : Tenant database init.
func InitTenantDatabase(tenantDatabaseClient *gorm.DB) (err error) {
	tenantDatabaseClient.AutoMigrate(&models.Aggregator{}, &models.Connector{}, &models.Application{}, &models.Event{}, &models.Command{}, &models.Config{}, &models.ConnectorConfig{}, &models.ConnectorType{}, &models.ConnectorCommand{}, &models.ConnectorEvent{}, &models.ConnectorProduct{}, &models.Action{}, &models.PermissionAction{}, &models.PermissionCommand{}, &models.PermissionEvent{}, &models.Resource{}, &models.Role{}, &models.User{})

	return
}

//DemoCreateAggregator
func DemoCreateAggregator(tenantDatabaseClient *gorm.DB) {
	tenantDatabaseClient.Create(&models.Aggregator{Name: "Aggregator1", Secret: "TATA"})
}

//DemoCreateConnector
func DemoCreateConnector(tenantDatabaseClient *gorm.DB) {
	tenantDatabaseClient.Create(&models.Connector{Name: "Connector1", Secret: "TOTO"})
}

//DemoCreateConnectorType
func DemoCreateConnectorType(tenantDatabaseClient *gorm.DB) {
	tenantDatabaseClient.Create(&models.ConnectorType{Name: "Utils"})
	tenantDatabaseClient.Create(&models.ConnectorType{Name: "Workflow"})
	tenantDatabaseClient.Create(&models.ConnectorType{Name: "Gitlab"})
	tenantDatabaseClient.Create(&models.ConnectorType{Name: "Azure"})
}

//DemoCreateRole
func DemoCreateRole(tenantDatabaseClient *gorm.DB) {
	tenantDatabaseClient.Create(&models.Role{Name: "Role1"})
	tenantDatabaseClient.Create(&models.Role{Name: "Role2"})
}

//DemoCreateUser1
func DemoCreateUser1(tenantDatabaseClient *gorm.DB) {

	var Role1 models.Role
	tenantDatabaseClient.Where("name = ?", "Role1").First(&Role1)

	user := models.NewUser("User1", "User1", "User1", Role1)
	tenantDatabaseClient.Create(&user)
}

//DemoCreateUser2
func DemoCreateUser2(tenantDatabaseClient *gorm.DB) {

	var Role2 models.Role
	tenantDatabaseClient.Where("name = ?", "Role2").First(&Role2)

	user := models.NewUser("User2", "User2", "User2", Role2)
	tenantDatabaseClient.Create(&user)
}

//DemoCreateProductUtils
func DemoCreateProductUtils(tenantDatabaseClient *gorm.DB) {

	var ConnectorTypeUtils models.ConnectorType
	tenantDatabaseClient.Where("name = ?", "Utils").First(&ConnectorTypeUtils)

	tenantDatabaseClient.Create(&models.ConnectorProduct{Name: "Custom", Version: "1", ConnectorType: ConnectorTypeUtils})
	tenantDatabaseClient.Create(&models.ConnectorProduct{Name: "Custom", Version: "2", ConnectorType: ConnectorTypeUtils})
}

//DemoCreateProductWorkflow
func DemoCreateProductWorkflow(tenantDatabaseClient *gorm.DB) {

	var ConnectorTypeWorkflow models.ConnectorType
	tenantDatabaseClient.Where("name = ?", "Workflow").First(&ConnectorTypeWorkflow)

	tenantDatabaseClient.Create(&models.ConnectorProduct{Name: "Custom", Version: "1", ConnectorType: ConnectorTypeWorkflow})
	tenantDatabaseClient.Create(&models.ConnectorProduct{Name: "Custom", Version: "2", ConnectorType: ConnectorTypeWorkflow})
}

//DemoCreateApplicationUtils
func DemoCreateApplicationUtils(tenantDatabaseClient *gorm.DB) {
	var AggregatorUtils models.Aggregator
	var ConnectorUtils models.Connector
	var ConnectorTypeUtils models.ConnectorType

	tenantDatabaseClient.Where("name = ?", "Aggregator1").First(&AggregatorUtils)
	tenantDatabaseClient.Where("name = ?", "Connector1").First(&ConnectorUtils)
	tenantDatabaseClient.Where("name = ?", "Utils").First(&ConnectorTypeUtils)

	tenantDatabaseClient.Create(&models.Application{Name: "Application1",
		Aggregator:    AggregatorUtils,
		Connector:     ConnectorUtils,
		ConnectorType: ConnectorTypeUtils})
}

//DemoCreateApplicationWorkflow
func DemoCreateApplicationWorkflow(tenantDatabaseClient *gorm.DB) {
	var AggregatorWorkflow models.Aggregator
	var ConnectorWorkflow models.Connector
	var ConnectorTypeWorkflow models.ConnectorType

	tenantDatabaseClient.Where("name = ?", "Aggregator2").First(&AggregatorWorkflow)
	tenantDatabaseClient.Where("name = ?", "Connector2").First(&ConnectorWorkflow)
	tenantDatabaseClient.Where("name = ?", "Workflow").First(&ConnectorTypeWorkflow)

	tenantDatabaseClient.Create(&models.Application{Name: "Application2",
		Aggregator:    AggregatorWorkflow,
		Connector:     ConnectorWorkflow,
		ConnectorType: ConnectorTypeWorkflow})

}

//DemoCreateApplicationAzure
func DemoCreateApplicationAzure(tenantDatabaseClient *gorm.DB) {
	var AggregatorAzure models.Aggregator
	var ConnectorAzure models.Connector
	var ConnectorTypeAzure models.ConnectorType

	tenantDatabaseClient.Where("name = ?", "Aggregator3").First(&AggregatorAzure)
	tenantDatabaseClient.Where("name = ?", "Connector3").First(&ConnectorAzure)
	tenantDatabaseClient.Where("name = ?", "Gitlab").First(&ConnectorTypeAzure)

	tenantDatabaseClient.Create(&models.Application{Name: "Application3",
		Aggregator:    AggregatorAzure,
		Connector:     ConnectorAzure,
		ConnectorType: ConnectorTypeAzure})
}

//DemoCreateApplicationGitlab
func DemoCreateApplicationGitlab(tenantDatabaseClient *gorm.DB) {
	var AggregatorGitlab models.Aggregator
	var ConnectorGitlab models.Connector
	var ConnectorTypeGitlab models.ConnectorType

	tenantDatabaseClient.Where("name = ?", "Aggregator4").First(&AggregatorGitlab)
	tenantDatabaseClient.Where("name = ?", "Connector4").First(&ConnectorGitlab)
	tenantDatabaseClient.Where("name = ?", "Azure").First(&ConnectorTypeGitlab)

	tenantDatabaseClient.Create(&models.Application{Name: "Application4",
		Aggregator:    AggregatorGitlab,
		Connector:     ConnectorGitlab,
		ConnectorType: ConnectorTypeGitlab})
}

//DemoCreateConfigurationUtils
func DemoCreateConfigurationUtils(tenantDatabaseClient *gorm.DB) {

	var ConnectorTypeUtils models.ConnectorType
	var ConnectorUtilsCommands []models.ConnectorCommand
	var ConnectorUtilsEvents []models.ConnectorEvent

	tenantDatabaseClient.Create(&models.ConnectorType{Name: "Utils"})
	tenantDatabaseClient.Where("name = ?", "Utils").First(&ConnectorTypeUtils)

	tenantDatabaseClient.Create(&models.ConnectorCommand{Name: "SEND_AUTH_MAIL", Schema: `{"$schema":"http://json-schema.org/draft-04/schema#","$ref":"#/definitions/MailPayload","definitions":{"MailPayload":{"required":["Sender","Body","Receivers","Username","Password","Host"],"properties":{"Sender":{"type":"string"},"Body":{"type":"string"},"Receivers":{"items":{"type":"string"},"type":"array"},"Username":{"type":"string"},"Password":{"type":"string"},"Host":{"type":"string"}},"additionalProperties":false,"type":"object"}}}`})
	tenantDatabaseClient.Create(&models.ConnectorCommand{Name: "CREATE_FORM", Schema: `{"$schema":"http://json-schema.org/draft-04/schema#","$ref":"#/definitions/FormPayload","definitions":{"Field":{"required":["Name","HtmlType","Value"],"properties":{"Name":{"type":"string"},"HtmlType":{"type":"string"},"Value":{"additionalProperties":true}},"additionalProperties":false,"type":"object"},"FormPayload":{"required":["Fields"],"properties":{"Fields":{"items":{"$schema":"http://json-schema.org/draft-04/schema#","$ref":"#/definitions/Field"},"type":"array"}},"additionalProperties":false,"type":"object"}}}`})

	tenantDatabaseClient.Where("name IN (?)", []string{"SEND_AUTH_MAIL", "CREATE_FORM"}).Find(&ConnectorUtilsCommands)

	tenantDatabaseClient.Create(&models.ConnectorEvent{Name: "NEW_APP", Schema: `{"type":"string"}`})

	tenantDatabaseClient.Where("name IN (?)", []string{"NEW_APP"}).Find(&ConnectorUtilsEvents)

	tenantDatabaseClient.Create(&models.ConnectorConfig{Name: "ConnectorConfig1",
		ConnectorType:     ConnectorTypeUtils,
		Version:           1,
		ConnectorCommands: ConnectorUtilsCommands,
		ConnectorEvents:   ConnectorUtilsEvents})

	tenantDatabaseClient.Create(&models.ConnectorConfig{Name: "ConnectorConfig2",
		ConnectorType:     ConnectorTypeUtils,
		Version:           2,
		ConnectorCommands: ConnectorUtilsCommands,
		ConnectorEvents:   ConnectorUtilsEvents})

}

//DemoCreateConfigurationWorkflow
func DemoCreateConfigurationWorkflow(tenantDatabaseClient *gorm.DB) {
	var ConnectorWorkflowCommands []models.ConnectorCommand
	var ConnectorTypeWorkflow models.ConnectorType

	tenantDatabaseClient.Create(&models.ConnectorType{Name: "Workflow"})
	tenantDatabaseClient.Where("name = ?", "Workflow").First(&ConnectorTypeWorkflow)

	tenantDatabaseClient.Where("name IN (?)", []string{}).Find(&ConnectorWorkflowCommands)

	tenantDatabaseClient.Create(&models.ConnectorConfig{Name: "ConnectorConfig3",
		ConnectorType:     ConnectorTypeWorkflow,
		Version:           1,
		ConnectorCommands: ConnectorWorkflowCommands,
		ConnectorEvents:   []models.ConnectorEvent{}})
}

//DemoCreateConfigurationAzure
func DemoCreateConfigurationAzure(tenantDatabaseClient *gorm.DB) {
	var ConnectorAzureCommands []models.ConnectorCommand
	var ConnectorTypeAzure models.ConnectorType
	var ConnectorProductAzure models.ConnectorProduct

	tenantDatabaseClient.Create(&models.ConnectorType{Name: "Azure"})
	tenantDatabaseClient.Where("name = ?", "Azure").First(&ConnectorTypeAzure)

	tenantDatabaseClient.Create(&models.ConnectorProduct{Name: "Azure", Version: "0"})
	tenantDatabaseClient.Where("name = ?", "Azure").First(&ConnectorProductAzure)

	tenantDatabaseClient.Create(&models.ConnectorCommand{Name: "CREATE_VM_BY_JSON", Schema: `{"$schema":"http://json-schema.org/draft-04/schema#","$ref":"#/definitions/ComputeByJSONPayload","definitions":{"ComputeByJSONPayload":{"required":["ResourceGroupName","ResourceGroupLocation","DeploymentName","TemplateFile","ParametersFile"],"properties":{"ResourceGroupName":{"type":"string"},"ResourceGroupLocation":{"type":"string"},"DeploymentName":{"type":"string"},"TemplateFile":{"type":"string"},"ParametersFile":{"type":"string"}},"additionalProperties":false,"type":"object"}}}`})

	tenantDatabaseClient.Where("name IN (?)", []string{"CREATE_VM_BY_JSON"}).Find(&ConnectorAzureCommands)

	tenantDatabaseClient.Create(&models.ConnectorConfig{Name: "ConnectorConfig4",
		ConnectorType:     ConnectorTypeAzure,
		Version:           1,
		ConnectorProduct:  ConnectorProductAzure,
		ConnectorCommands: ConnectorAzureCommands,
		ConnectorEvents:   []models.ConnectorEvent{}})
}

//DemoCreateConfigurationGitlab
func DemoCreateConfigurationGitlab(tenantDatabaseClient *gorm.DB) {
	var ConnectorGitlabCommands []models.ConnectorCommand
	var ConnectorTypeGitlab models.ConnectorType
	var ConnectorProductGitlab models.ConnectorProduct

	tenantDatabaseClient.Create(&models.ConnectorType{Name: "Gitlab"})
	tenantDatabaseClient.Where("name = ?", "Gitlab").First(&ConnectorTypeGitlab)

	tenantDatabaseClient.Create(&models.ConnectorProduct{Name: "Gitlab", Version: "0"})
	tenantDatabaseClient.Where("name = ?", "Azure").First(&ConnectorProductGitlab)

	tenantDatabaseClient.Create(&models.ConnectorCommand{Name: "Gitlab1", Schema: ""})
	tenantDatabaseClient.Create(&models.ConnectorCommand{Name: "Gitlab2", Schema: ""})
	tenantDatabaseClient.Create(&models.ConnectorCommand{Name: "Gitlab3", Schema: ""})

	tenantDatabaseClient.Where("name IN (?)", []string{"Gitlab1", "Gitlab2", "Gitlab3"}).Find(&ConnectorGitlabCommands)

	tenantDatabaseClient.Create(&models.ConnectorConfig{Name: "ConnectorConfig5",
		ConnectorType:     ConnectorTypeGitlab,
		Version:           1,
		ConnectorProduct:  ConnectorProductGitlab,
		ConnectorCommands: ConnectorGitlabCommands,
		ConnectorEvents:   []models.ConnectorEvent{}})

}

// DemoPopulateTenantDatabase : Populate database demo.
func DemoPopulateTenantDatabase(tenantDatabaseClient *gorm.DB) {

	//CORE
	DemoCreateAggregator(tenantDatabaseClient)
	DemoCreateConnector(tenantDatabaseClient)
	DemoCreateConnectorType(tenantDatabaseClient)
	DemoCreateRole(tenantDatabaseClient)

	//PRODUCT
	DemoCreateProductUtils(tenantDatabaseClient)
	DemoCreateProductWorkflow(tenantDatabaseClient)

	//USER
	DemoCreateUser1(tenantDatabaseClient)
	DemoCreateUser2(tenantDatabaseClient)

	//APPLICATION
	DemoCreateApplicationUtils(tenantDatabaseClient)
	DemoCreateApplicationWorkflow(tenantDatabaseClient)
	DemoCreateApplicationAzure(tenantDatabaseClient)
	DemoCreateApplicationGitlab(tenantDatabaseClient)

	//CONFIGURATION
	//DemoCreateConfigurationUtils(tenantDatabaseClient)
	DemoCreateConfigurationWorkflow(tenantDatabaseClient)
	DemoCreateConfigurationAzure(tenantDatabaseClient)
	DemoCreateConfigurationGitlab(tenantDatabaseClient)

}