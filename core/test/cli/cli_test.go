package cli

import (
	"fmt"
	"testing"

	"github.com/ditrit/gandalf/core/cli"
	"github.com/ditrit/gandalf/core/models"
)

func TestCreateDomain(t *testing.T) {
	const (
		token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjY3MzQyMTg1MTUwODg2NzA3MywiTmFtZSI6IkFkbWluaXN0cmF0b3IxIiwiRW1haWwiOiJBZG1pbmlzdHJhdG9yMSIsIlRlbmFudCI6IiIsImV4cCI6MTYzMTU4MjUwMX0.dYSuQzl27yE4wJFrP3H-Ck6ZOKEGxQE66nGX4UVStLA"
	)
	cliClient := cli.NewClient("http://localhost:9203")

	t.Log("DOMAIN.TEST >> SUCCESS")
	name := "test"
	parentName := "root"

	// Good Scheme
	domain := models.Domain{Name: name}
	err := cliClient.DomainService.Create(token, domain, parentName)
	t.Log(err)

	// Same Name
	t.Log("DOMAIN.TEST >> FAIL - DOMAIN NAME ALREADY EXISTS")
	domain = models.Domain{Name: name}
	err = cliClient.DomainService.Create(token, domain, parentName)
	t.Log(err)

	t.Log("DOMAIN.TEST >> FAIL - WRONG PARENT NAME")
	// Wrong Parent
	parentName = ""

	domain = models.Domain{Name: name}
	err = cliClient.DomainService.Create(token, domain, parentName)
	t.Log(err)

	t.Log("DOMAIN.TEST >> FAIL - SAME PARENT NAME THAN DOMAIN NAME")
	parentName = "test"

	domain = models.Domain{Name: name}
	err = cliClient.DomainService.Create(token, domain, parentName)
	t.Log(err)
}
func TestUpdateDomain(t *testing.T) {
	const (
		token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjY3MzQyMTg1MTUwODg2NzA3MywiTmFtZSI6IkFkbWluaXN0cmF0b3IxIiwiRW1haWwiOiJBZG1pbmlzdHJhdG9yMSIsIlRlbmFudCI6IiIsImV4cCI6MTYzMTU4MjUwMX0.dYSuQzl27yE4wJFrP3H-Ck6ZOKEGxQE66nGX4UVStLA"
	)
	cliClient := cli.NewClient("http://localhost:9203")

	name := "test"
	newName := "other"

	t.Log("DOMAIN.TEST >> SUCCESS")

	oldDomain, err := cliClient.DomainService.ReadByName(token, name)
	if err == nil {
		domain := models.Domain{Name: newName}
		err = cliClient.DomainService.Update(token, int(oldDomain.ID), domain)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}

	t.Log("DOMAIN.TEST >> FAILED - Same Domain Name")
	newName = "test"
	if err == nil {
		domain := models.Domain{Name: newName}
		err = cliClient.DomainService.Update(token, int(oldDomain.ID), domain)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}

	t.Log("DOMAIN.TEST >> FAILED - Do not exists")
	name = "neo"
	if err == nil {
		domain := models.Domain{Name: newName}
		err = cliClient.DomainService.Update(token, int(oldDomain.ID), domain)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}

func TestDeleteDomain(t *testing.T) {
	const (
		token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjY3MzQyMTg1MTUwODg2NzA3MywiTmFtZSI6IkFkbWluaXN0cmF0b3IxIiwiRW1haWwiOiJBZG1pbmlzdHJhdG9yMSIsIlRlbmFudCI6IiIsImV4cCI6MTYzMTU4MjUwMX0.dYSuQzl27yE4wJFrP3H-Ck6ZOKEGxQE66nGX4UVStLA"
	)
	cliClient := cli.NewClient("http://localhost:9203")

	t.Log("DOMAIN.TEST >> SUCCESS")
	name := "test"
	oldDomain, err := cliClient.DomainService.ReadByName(token, name)
	if err == nil {
		err = cliClient.DomainService.Delete(token, int(oldDomain.ID))
		if err != nil {
			t.Log(err)
		}
	} else {
		t.Log(err)
	}

	/*
		name := args[0]
		fmt.Printf("gandalf cli delete domain called with domain=%s\n", name)
		configurationCli := cmodels.NewConfigurationCli()
		cliClient := cli.NewClient(configurationCli.GetEndpoint())

		oldDomain, err := cliClient.DomainService.ReadByName(configurationCli.GetToken(), name)
		if err == nil {
			err = cliClient.DomainService.Delete(configurationCli.GetToken(), int(oldDomain.ID))
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	*/

	t.Log("DOMAIN.TEST >> FAIL: REMOVE ALREADY REMOVED DOAMIN OR A DOMAIN THAT DO NOT EXISTS")
	oldDomain, err = cliClient.DomainService.ReadByName(token, name)
	if err == nil {
		err = cliClient.DomainService.Delete(token, int(oldDomain.ID))
		if err != nil {
			t.Log(err)
		}
	} else {
		t.Log(err)
	}
}

func TestCreateResourceType__Pivot(t *testing.T) {
	const (
		token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjY3MzQyMTg1MTUwODg2NzA3MywiTmFtZSI6IkFkbWluaXN0cmF0b3IxIiwiRW1haWwiOiJBZG1pbmlzdHJhdG9yMSIsIlRlbmFudCI6IiIsImV4cCI6MTYzMTU4MjUwMX0.dYSuQzl27yE4wJFrP3H-Ck6ZOKEGxQE66nGX4UVStLA"
	)
	cliClient := cli.NewClient("http://localhost:9203")

	name := "pivotTest"
	pivotProductConnectorName := "utils"
	typeName := "pivot"

	t.Log("TRY WITH PIVOT")
	// Pivot

	t.Log("PIVOT.TEST >> SUCCESS")
	if typeName == "pivot" {
		pivot, err := cliClient.PivotService.ReadByName(token, pivotProductConnectorName)

		if err == nil {
			resourceType := models.ResourceType{Name: name, Pivot: *pivot}
			err := cliClient.ResourceTypeService.Create(token, resourceType)
			if err != nil {
				t.Log(err)
			}
		} else {
			t.Log(err)
		}
	} else {
		t.Log("Error: must be connectorProduct or pivot.")
	}

	t.Log("PIVOT.TEST >> INCORRECT TYPE")
	// Pivot - TRY 2 INCORRECT TYPE
	typeName = "incorrect"
	if typeName == "pivot" {
		pivot, err := cliClient.PivotService.ReadByName(token, pivotProductConnectorName)

		if err == nil {
			resourceType := models.ResourceType{Name: name, Pivot: *pivot}
			err := cliClient.ResourceTypeService.Create(token, resourceType)
			if err != nil {
				t.Log(err)
			}
		} else {
			t.Log(err)
		}
	} else {
		t.Log("Error: must be connectorProduct or pivot.")
	}
	t.Log("PIVOT.TEST >> INCORRECT TYPE NAME (== PIVOT && PIVOT.NAME != utils)")
	typeName = "pivot"
	pivotProductConnectorName = "not_utils"
	if typeName == "pivot" {
		pivot, err := cliClient.PivotService.ReadByName(token, pivotProductConnectorName)

		if err == nil {
			resourceType := models.ResourceType{Name: name, Pivot: *pivot}
			err := cliClient.ResourceTypeService.Create(token, resourceType)
			if err != nil {
				t.Log(err)
			}
		} else {
			t.Log(err)
		}
	} else {
		t.Log("Error: must be connectorProduct or pivot.")
	}
}
func TestCreateResourceType__ProductConnector(t *testing.T) {
	const (
		token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjY3MzQyMTg1MTUwODg2NzA3MywiTmFtZSI6IkFkbWluaXN0cmF0b3IxIiwiRW1haWwiOiJBZG1pbmlzdHJhdG9yMSIsIlRlbmFudCI6IiIsImV4cCI6MTYzMTU4MjUwMX0.dYSuQzl27yE4wJFrP3H-Ck6ZOKEGxQE66nGX4UVStLA"
	)
	cliClient := cli.NewClient("http://localhost:9203")

	// ERROR - 1-1A
	// PRODUCT CONNECTOR PART //
	t.Log("TRY WITH PRODUCT_CONNECTOR")
	t.Log("PRODUCT_CONNECTOR.TEST >> SUCCESS")
	// Product Connector - Correct
	name := "productConnectorTest"
	pivotProductConnectorName := "UtilsCustom1.0"
	typeName := "productConnector"
	if typeName == "productConnector" {

		productConnector, err := cliClient.ProductConnectorService.ReadByName(token, pivotProductConnectorName)
		if err == nil {
			resourceType := models.ResourceType{Name: name, ProductConnector: *productConnector}
			err := cliClient.ResourceTypeService.Create(token, resourceType)
			if err != nil {
				t.Log(err)
			}
		} else {
			t.Log(err)
		}
	} else {
		t.Error("Error: must be connectorProduct or pivot.")
	}

	// Product Connector - TRY 2 INCORRECT TYPE
	t.Log("PRODUCT_CONNECTOR.TEST >> INCORRECT TYPE")
	pivotProductConnectorName = "incorrect"
	if typeName == "productConnector" {

		productConnector, err := cliClient.ProductConnectorService.ReadByName(token, pivotProductConnectorName)
		if err == nil {
			resourceType := models.ResourceType{Name: name, ProductConnector: *productConnector}
			err := cliClient.ResourceTypeService.Create(token, resourceType)
			if err != nil {
				t.Log(err)
			}
		} else {
			t.Log(err)
		}
	} else {
		t.Log("Error: must be connectorProduct or pivot.")
	}

	t.Log("PRODUCT_CONNECTOR.TEST >> INCORRECT TYPE NAME (== CONNECTOR_PRODUCT && CONNECTOR_PRODUCT.NAME != UtilsCustom1.0)")
	typeName = "productConnector"
	pivotProductConnectorName = "not_utils"
	if typeName == "productConnector" {

		productConnector, err := cliClient.ProductConnectorService.ReadByName(token, pivotProductConnectorName)
		if err == nil {
			resourceType := models.ResourceType{Name: name, ProductConnector: *productConnector}
			err := cliClient.ResourceTypeService.Create(token, resourceType)
			if err != nil {
				t.Log(err)
			}
		} else {
			t.Log(err)
		}
	} else {
		t.Log("Error: must be connectorProduct or pivot.")
	}
}

func TestUpdateResourceType__Pivot(t *testing.T) {

	const (
		token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjY3MzQyMTg1MTUwODg2NzA3MywiTmFtZSI6IkFkbWluaXN0cmF0b3IxIiwiRW1haWwiOiJBZG1pbmlzdHJhdG9yMSIsIlRlbmFudCI6IiIsImV4cCI6MTYzMTU4MjUwMX0.dYSuQzl27yE4wJFrP3H-Ck6ZOKEGxQE66nGX4UVStLA"
	)
	cliClient := cli.NewClient("http://localhost:9203")

	t.Log("PIVOT.TEST >> UPDATE - SUCCESS")
	name := "pivotTest"
	newName := "neoPivotTest"
	typeName := "pivot"
	pivotProductConnectorName := "utils"

	oldResourceType, err := cliClient.ResourceTypeService.ReadByName(token, name)
	if err == nil {

		if typeName == "pivot" {
			pivot, err := cliClient.PivotService.ReadByName(token, pivotProductConnectorName)
			if err == nil {
				resourceType := models.ResourceType{Name: newName, Pivot: *pivot}
				err = cliClient.ResourceTypeService.Update(token, int(oldResourceType.ID), resourceType)
				fmt.Println(err)
			} else {
				t.Log(err)
				fmt.Println("ERROR: CANNOT FIND SPECIFIED PIVOT")
			}

		} else {
			t.Log(err)
		}
	}

	pivotProductConnectorName = "not_utils"
	t.Log("PIVOT.TEST - INCORRECT PIVOT.NAME")
	if err == nil {

		if typeName == "pivot" {
			pivot, err := cliClient.PivotService.ReadByName(token, pivotProductConnectorName)
			if err == nil {
				resourceType := models.ResourceType{Name: newName, Pivot: *pivot}
				err = cliClient.ResourceTypeService.Update(token, int(oldResourceType.ID), resourceType)
				fmt.Println(err)
			} else {
				t.Log(err)
				fmt.Println("ERROR: CANNOT FIND SPECIFIED PIVOT")
			}

		} else {
			t.Log(err)
		}
	}

	t.Log("PIVOT.TEST >> UPDATE - BASE NAME DOES NOT EXISTS")
	name = "unknown"
	pivotProductConnectorName = "utils"

	if typeName == "productConnector" {
		productConnector, err := cliClient.ProductConnectorService.ReadByName(token, pivotProductConnectorName)
		if err == nil {
			resourceType := models.ResourceType{Name: newName, ProductConnector: *productConnector}
			err = cliClient.ResourceTypeService.Update(token, int(oldResourceType.ID), resourceType)
			fmt.Println(err)
		}
	} else {
		t.Log(err)
		fmt.Println("ERROR: CANNOT FIND SPECIFIED PRODUCTCONNECTOR")
	}
}
func TestUpdateResourceType__ProductConnector(t *testing.T) {

	const (
		token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjY3MzQyMTg1MTUwODg2NzA3MywiTmFtZSI6IkFkbWluaXN0cmF0b3IxIiwiRW1haWwiOiJBZG1pbmlzdHJhdG9yMSIsIlRlbmFudCI6IiIsImV4cCI6MTYzMTU4MjUwMX0.dYSuQzl27yE4wJFrP3H-Ck6ZOKEGxQE66nGX4UVStLA"
	)
	cliClient := cli.NewClient("http://localhost:9203")

	t.Log("PRODUCT_CONNECTOR.TEST >> UPDATE - SUCCESS")

	name := "productConnectorName"
	newName := "neoProductConnectorTest"
	typeName := "productConnector"
	pivotProductConnectorName := "UtilsCustom1.0"

	oldResourceType, err := cliClient.ResourceTypeService.ReadByName(token, name)
	if typeName == "productConnector" {
		productConnector, err := cliClient.ProductConnectorService.ReadByName(token, pivotProductConnectorName)
		if err == nil {
			resourceType := models.ResourceType{Name: newName, ProductConnector: *productConnector}
			err = cliClient.ResourceTypeService.Update(token, int(oldResourceType.ID), resourceType)
			t.Log(err)
		} else {
			t.Log(err)
		}
	} else {
		t.Log(err)
		fmt.Println("ERROR: CANNOT FIND SPECIFIED PRODUCTCONNECTOR")
	}

	t.Log("PRODUCT_CONNECTOR.TEST >> UPDATE - INCORRECT PRODUCT_CONNECTOR.NAME")
	pivotProductConnectorName = "not_correct"

	if typeName == "productConnector" {
		productConnector, err := cliClient.ProductConnectorService.ReadByName(token, pivotProductConnectorName)
		if err == nil {
			resourceType := models.ResourceType{Name: newName, ProductConnector: *productConnector}
			err = cliClient.ResourceTypeService.Update(token, int(oldResourceType.ID), resourceType)
			t.Log(err)
		} else {
			t.Log(err)
		}
	} else {
		t.Log(err)
		fmt.Println("ERROR: CANNOT FIND SPECIFIED PRODUCTCONNECTOR")
	}

	t.Log("PRODUCT_CONNECTOR.TEST >> UPDATE - BASE NAME DOES NOT EXISTS")
	name = "unknown"
	pivotProductConnectorName = "UtilsCustom1.0"

	if typeName == "productConnector" {
		productConnector, err := cliClient.ProductConnectorService.ReadByName(token, pivotProductConnectorName)
		if err == nil {
			resourceType := models.ResourceType{Name: newName, ProductConnector: *productConnector}
			err = cliClient.ResourceTypeService.Update(token, int(oldResourceType.ID), resourceType)
			fmt.Println(err)
		}
	} else {
		t.Log(err)
		fmt.Println("ERROR: CANNOT FIND SPECIFIED PRODUCTCONNECTOR")
	}
}

func TestDeleteResourceType__ProductConnector(t *testing.T) {

	const (
		token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjY3MzQyMTg1MTUwODg2NzA3MywiTmFtZSI6IkFkbWluaXN0cmF0b3IxIiwiRW1haWwiOiJBZG1pbmlzdHJhdG9yMSIsIlRlbmFudCI6IiIsImV4cCI6MTYzMTU4MjUwMX0.dYSuQzl27yE4wJFrP3H-Ck6ZOKEGxQE66nGX4UVStLA"
	)
	cliClient := cli.NewClient("http://localhost:9203")

	t.Log("PRODUCT_CONNECTOR.TEST >> DELETE - SUCCESS")

	name := "neoProductConnectorTest"

	ResourceType, err := cliClient.ResourceTypeService.ReadByName(token, name)
	if err == nil {
		err = cliClient.ResourceTypeService.Delete(token, int(ResourceType.ID))
		t.Log(err)
	} else {
		t.Log(err)
	}

	t.Log("PRODUCT_CONNECTOR.TEST >> DELETE - FAIL")

	name = "productConnectorNameThatDoNotExists"

	ResourceType, err = cliClient.ResourceTypeService.ReadByName(token, name)
	if err == nil {
		err = cliClient.ResourceTypeService.Delete(token, int(ResourceType.ID))
		t.Log(err)
	} else {
		t.Log(err)
	}
}

func TestDeleteResourceType__Pivot(t *testing.T) {

	const (
		token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjY3MzQyMTg1MTUwODg2NzA3MywiTmFtZSI6IkFkbWluaXN0cmF0b3IxIiwiRW1haWwiOiJBZG1pbmlzdHJhdG9yMSIsIlRlbmFudCI6IiIsImV4cCI6MTYzMTU4MjUwMX0.dYSuQzl27yE4wJFrP3H-Ck6ZOKEGxQE66nGX4UVStLA"
	)
	cliClient := cli.NewClient("http://localhost:9203")

	t.Log("PIVOT.TEST >> DELETE - SUCCESS")

	name := "neoPivotTest"

	ResourceType, err := cliClient.ResourceTypeService.ReadByName(token, name)
	if err == nil {
		err = cliClient.ResourceTypeService.Delete(token, int(ResourceType.ID))
		t.Log(err)
	} else {
		t.Log(err)
	}

	t.Log("PIVOT.TEST >> DELETE - FAIL")

	name = "pivotNameThatDoNotExists"

	ResourceType, err = cliClient.ResourceTypeService.ReadByName(token, name)
	if err == nil {
		err = cliClient.ResourceTypeService.Delete(token, int(ResourceType.ID))
		t.Log(err)
	} else {
		t.Log(err)
	}
}

func TestCreateEventType_Pivot(t *testing.T) {
	name := "created_eventType_pivot"
	schema := "test"
	pivotProductConnectorName := "utils"
	typeName := "pivot"

	const (
		token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjY3MzQyMTg1MTUwODg2NzA3MywiTmFtZSI6IkFkbWluaXN0cmF0b3IxIiwiRW1haWwiOiJBZG1pbmlzdHJhdG9yMSIsIlRlbmFudCI6IiIsImV4cCI6MTYzMTU4MjUwMX0.dYSuQzl27yE4wJFrP3H-Ck6ZOKEGxQE66nGX4UVStLA"
	)
	cliClient := cli.NewClient("http://localhost:9203")

	t.Log("PIVOT.TEST >> CREATE - SUCCESS")

	if typeName == "pivot" {
		pivot, err := cliClient.PivotService.ReadByName(token, pivotProductConnectorName)

		if err == nil {
			eventType := models.EventType{Name: name, Schema: schema, Pivot: *pivot}
			err := cliClient.EventTypeService.Create(token, eventType)
			if err != nil {
				t.Log(err)
			}
		} else {
			t.Log(err)
		}
	}
	t.Log("PIVOT.TEST >> CREATE - FAIL: LOGICAL COMPONENT NAME")

	pivotProductConnectorName = "no_utils"

	if typeName == "pivot" {
		pivot, err := cliClient.PivotService.ReadByName(token, pivotProductConnectorName)

		if err == nil {
			eventType := models.EventType{Name: name, Schema: schema, Pivot: *pivot}
			err := cliClient.EventTypeService.Create(token, eventType)
			if err != nil {
				t.Log(err)
			}
		} else {
			t.Log(err)
		}
	}
	typeName = "nothing"
	pivotProductConnectorName = "utils"

	t.Log("PIVOT.TEST >> CREATE - FAIL: INCORRECT TYPENAME")

	pivotProductConnectorName = "no_UtilsCustom1.0"

	if typeName == "pivot" {
		pivot, err := cliClient.PivotService.ReadByName(token, pivotProductConnectorName)

		if err == nil {
			eventType := models.EventType{Name: name, Schema: schema, Pivot: *pivot}
			err := cliClient.EventTypeService.Create(token, eventType)
			if err != nil {
				t.Log(err)
			}
		} else {
			t.Log(err)
		}
	}

}

func TestCreateEventType_ProductConnector(t *testing.T) {
	name := "created_eventType_productConnector"
	schema := "test"
	pivotProductConnectorName := "UtilsCustom1.0"
	typeName := "productConnector"

	const (
		token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjY3MzQyMTg1MTUwODg2NzA3MywiTmFtZSI6IkFkbWluaXN0cmF0b3IxIiwiRW1haWwiOiJBZG1pbmlzdHJhdG9yMSIsIlRlbmFudCI6IiIsImV4cCI6MTYzMTU4MjUwMX0.dYSuQzl27yE4wJFrP3H-Ck6ZOKEGxQE66nGX4UVStLA"
	)
	cliClient := cli.NewClient("http://localhost:9203")

	t.Log("PRODUCT_CONNECTOR.TEST >> CREATE - SUCCESS")

	if typeName == "productConnector" {

		productConnector, err := cliClient.ProductConnectorService.ReadByName(token, pivotProductConnectorName)
		if err == nil {
			eventType := models.EventType{Name: name, Schema: schema, ProductConnector: *productConnector}
			err := cliClient.EventTypeService.Create(token, eventType)
			if err != nil {
				t.Log(err)
			}
		} else {
			t.Log(err)
		}
	}

	t.Log("PRODUCT_CONNECTOR.TEST >> CREATE - FAIL: LOGICAL COMPONENT NAME")

	pivotProductConnectorName = "no_UtilsCustom1.0"

	if typeName == "productConnector" {
		pivot, err := cliClient.PivotService.ReadByName(token, pivotProductConnectorName)

		if err == nil {
			eventType := models.EventType{Name: name, Schema: schema, Pivot: *pivot}
			err := cliClient.EventTypeService.Create(token, eventType)
			if err != nil {
				t.Log(err)
			}
		} else {
			t.Log(err)
		}
	}
	typeName = "nothing"
	pivotProductConnectorName = "UtilsCustom1.0"

	t.Log("PRODUCT_CONNECTOR.TEST >> CREATE - FAIL: INCORRECT TYPENAME")

	pivotProductConnectorName = "no_UtilsCustom1.0"

	if typeName == "productConnector" {
		pivot, err := cliClient.PivotService.ReadByName(token, pivotProductConnectorName)

		if err == nil {
			eventType := models.EventType{Name: name, Schema: schema, Pivot: *pivot}
			err := cliClient.EventTypeService.Create(token, eventType)
			if err != nil {
				t.Log(err)
			}
		} else {
			t.Log(err)
		}
	}
}
func TestUpdateEventType_Pivot(t *testing.T) {
	name := "created_eventType_pivot"
	newName := "updated_eventType_pivot"
	schema := "testSchema"
	pivotProductConnectorName := "utils"
	typeName := "pivot"

	const (
		token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjY3MzQyMTg1MTUwODg2NzA3MywiTmFtZSI6IkFkbWluaXN0cmF0b3IxIiwiRW1haWwiOiJBZG1pbmlzdHJhdG9yMSIsIlRlbmFudCI6IiIsImV4cCI6MTYzMTU4MjUwMX0.dYSuQzl27yE4wJFrP3H-Ck6ZOKEGxQE66nGX4UVStLA"
	)
	cliClient := cli.NewClient("http://localhost:9203")

	t.Log("PIVOT.TEST >> UPDATE - SUCCESS")
	oldEventType, err := cliClient.EventTypeService.ReadByName(token, name)
	if err == nil {
		if typeName == "pivot" {
			pivot, err := cliClient.PivotService.ReadByName(token, pivotProductConnectorName)
			if err == nil {
				eventType := models.EventType{Name: newName, Schema: schema, Pivot: *pivot}
				err := cliClient.EventTypeService.Update(token, int(oldEventType.ID), eventType)
				if err != nil {
					t.Log(err)
				}
			} else {
				t.Log(err)
			}
		}
	} else {
		t.Log(err)
	}

	t.Log("EVENTTYPE >> PIVOT.UPDATE - INCORRECT PRODUCT_CONNECTOR.NAME")
	pivotProductConnectorName = "not_utils"
	if err == nil {
		if typeName == "pivot" {
			pivot, err := cliClient.PivotService.ReadByName(token, pivotProductConnectorName)
			if err == nil {
				eventType := models.EventType{Name: newName, Schema: schema, Pivot: *pivot}
				err := cliClient.EventTypeService.Update(token, int(oldEventType.ID), eventType)
				if err != nil {
					t.Log(err)
				}
			} else {
				t.Log(err)
			}
		}
	} else {
		t.Log(err)
	}
	t.Log("EVENTTYPE >> PIVOT.TEST >> UPDATE - BASE NAME DOES NOT EXISTS")
	name = "DoesNotExists"
	pivotProductConnectorName = "utils"
	if err == nil {
		if typeName == "pivot" {
			pivot, err := cliClient.PivotService.ReadByName(token, pivotProductConnectorName)
			if err == nil {
				eventType := models.EventType{Name: newName, Schema: schema, Pivot: *pivot}
				err := cliClient.EventTypeService.Update(token, int(oldEventType.ID), eventType)
				if err != nil {
					t.Log(err)
				}
			} else {
				t.Log(err)
			}
		}
	} else {
		t.Log(err)
	}

	t.Log("EVENTTYPE >> PIVOT.TEST >> UPDATE - TYPENAME INCORRECT")
	name = "created_eventType_pivot"
	pivotProductConnectorName = "utils"
	typeName = "undefined"

	if err == nil {
		if typeName == "pivot" {
			pivot, err := cliClient.PivotService.ReadByName(token, pivotProductConnectorName)
			if err == nil {
				eventType := models.EventType{Name: newName, Schema: schema, Pivot: *pivot}
				err := cliClient.EventTypeService.Update(token, int(oldEventType.ID), eventType)
				if err != nil {
					t.Log(err)
				}
			} else {
				t.Log(err)
			}
		} else {
			t.Log(err)
		}
	} else {
		t.Log(err)
	}
}

func TestUpdateEventType_ProductConnector(t *testing.T) {
	name := "created_eventType_productConnector"
	newName := "updated_eventType_productConnector"
	schema := "test"
	pivotProductConnectorName := "UtilsCustom1.0"
	typeName := "productConnector"

	const (
		token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjY3MzQyMTg1MTUwODg2NzA3MywiTmFtZSI6IkFkbWluaXN0cmF0b3IxIiwiRW1haWwiOiJBZG1pbmlzdHJhdG9yMSIsIlRlbmFudCI6IiIsImV4cCI6MTYzMTU4MjUwMX0.dYSuQzl27yE4wJFrP3H-Ck6ZOKEGxQE66nGX4UVStLA"
	)
	cliClient := cli.NewClient("http://localhost:9203")

	t.Log("EVENTTYPE >> PRODUCT_CONNECTOR.TEST >> UPDATE - SUCCESS")
	oldEventType, err := cliClient.EventTypeService.ReadByName(token, name)
	if typeName == "productConnector" {
		productConnector, err := cliClient.ProductConnectorService.ReadByName(token, pivotProductConnectorName)
		if err == nil {
			eventType := models.EventType{Name: newName, Schema: schema, ProductConnector: *productConnector}
			err := cliClient.EventTypeService.Update(token, int(oldEventType.ID), eventType)
			if err != nil {
				t.Log(err)
			}
		} else {
			t.Log(err)
		}
	} else {
		t.Log(err)
	}

	t.Log("EVENTTYPE >> PRODUCT_CONNECTOR.TEST >> UPDATE - INCORRECT PRODUCT_CONNECTOR.NAME")
	pivotProductConnectorName = "not_utils"

	if typeName == "productConnector" {
		productConnector, err := cliClient.ProductConnectorService.ReadByName(token, pivotProductConnectorName)
		if err == nil {
			eventType := models.EventType{Name: newName, Schema: schema, ProductConnector: *productConnector}
			err := cliClient.EventTypeService.Update(token, int(oldEventType.ID), eventType)
			if err != nil {
				t.Log(err)
			}
		} else {
			t.Log(err)
		}
	} else {
		t.Log(err)
	}

	t.Log("EVENTTYPE >> PRODUCT_CONNECTOR.TEST >> UPDATE - BASE NAME DOES NOT EXISTS")
	name = "DoesNotExists"
	pivotProductConnectorName = "UtilsCustom1.0"

	if typeName == "productConnector" {
		productConnector, err := cliClient.ProductConnectorService.ReadByName(token, pivotProductConnectorName)
		if err == nil {
			eventType := models.EventType{Name: newName, Schema: schema, ProductConnector: *productConnector}
			err := cliClient.EventTypeService.Update(token, int(oldEventType.ID), eventType)
			if err != nil {
				t.Log(err)
			}
		} else {
			t.Log(err)
		}
	} else {
		t.Log(err)
	}

	t.Log("EVENTTYPE >> PRODUCT_CONNECTOR.TEST >> UPDATE - TYPENAME INCORRECT")
	name = "created_eventType_productConnector"
	pivotProductConnectorName = "UtilsCustom1.0"
	typeName = "undefined"

	if typeName == "productConnector" {
		productConnector, err := cliClient.ProductConnectorService.ReadByName(token, pivotProductConnectorName)
		if err == nil {
			eventType := models.EventType{Name: newName, Schema: schema, ProductConnector: *productConnector}
			err := cliClient.EventTypeService.Update(token, int(oldEventType.ID), eventType)
			if err != nil {
				t.Log(err)
			}
		} else {
			t.Log(err)
		}
	} else {
		t.Log(err)
	}
}

func TestDeleteEventType__Pivot(t *testing.T) {
	const (
		token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjY3MzQyMTg1MTUwODg2NzA3MywiTmFtZSI6IkFkbWluaXN0cmF0b3IxIiwiRW1haWwiOiJBZG1pbmlzdHJhdG9yMSIsIlRlbmFudCI6IiIsImV4cCI6MTYzMTU4MjUwMX0.dYSuQzl27yE4wJFrP3H-Ck6ZOKEGxQE66nGX4UVStLA"
	)
	cliClient := cli.NewClient("http://localhost:9203")

	name := "update_eventType_pivot"

	t.Log("EVENTTYPE >> PIVOT.TEST >> DELETE - SUCCESS")

	oldEventType, err := cliClient.EventTypeService.ReadByName(token, name)
	if err == nil {
		err = cliClient.EventTypeService.Delete(token, int(oldEventType.ID))
		if err != nil {
			t.Log(err)
		}
	} else {
		t.Log(err)
	}

	t.Log("EVENTTYPE >> PIVOT.TEST >> DELETE - FAIL")

	name = "DoNotExists"
	if err == nil {
		err = cliClient.EventTypeService.Delete(token, int(oldEventType.ID))
		if err != nil {
			t.Log(err)
		}
	} else {
		t.Log(err)
	}
}

func TestDeleteEventType__ProductConnector(t *testing.T) {
	const (
		token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjY3MzQyMTg1MTUwODg2NzA3MywiTmFtZSI6IkFkbWluaXN0cmF0b3IxIiwiRW1haWwiOiJBZG1pbmlzdHJhdG9yMSIsIlRlbmFudCI6IiIsImV4cCI6MTYzMTU4MjUwMX0.dYSuQzl27yE4wJFrP3H-Ck6ZOKEGxQE66nGX4UVStLA"
	)
	cliClient := cli.NewClient("http://localhost:9203")

	name := "update_eventType_productConnector"

	t.Log("EVENTTYPE >> PRODUCT_CONNECTOR.TEST >> DELETE - SUCCESS")

	oldEventType, err := cliClient.EventTypeService.ReadByName(token, name)
	if err == nil {
		err = cliClient.EventTypeService.Delete(token, int(oldEventType.ID))
		if err != nil {
			t.Log(err)
		}
	} else {
		t.Log(err)
	}

	name = "DoNotExists"

	t.Log("EVENTTYPE >> PRODUCT_CONNECTOR.TEST >> DELETE - FAIL")

	if err == nil {
		err = cliClient.EventTypeService.Delete(token, int(oldEventType.ID))
		if err != nil {
			t.Log(err)
		}
	} else {
		t.Log(err)
	}
}

// WIP
func TestCreateResource__Pivot(t *testing.T) {
	const (
		token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjY3MzQyMTg1MTUwODg2NzA3MywiTmFtZSI6IkFkbWluaXN0cmF0b3IxIiwiRW1haWwiOiJBZG1pbmlzdHJhdG9yMSIsIlRlbmFudCI6IiIsImV4cCI6MTYzMTU4MjUwMX0.dYSuQzl27yE4wJFrP3H-Ck6ZOKEGxQE66nGX4UVStLA"
	)
	cliClient := cli.NewClient("http://localhost:9203")

	name := "testResource"
	logicalComponentName := "connector5"
	domainName := "testDomain2"
	resourceTypeName := "tutu"
	parentName := "root"

	// Create new domain
	domain := models.Domain{Name: domainName}
	err := cliClient.DomainService.Create(token, domain, parentName)
	t.Log(err)

	t.Log("RESOURCE >> PIVOT.TEST >> CREATE - SUCCESS")

	// Create new resource by pivot, using new domain
	logicalComponent, err := cliClient.LogicalComponentService.ReadByName(token, logicalComponentName)
	t.Log(err)
	if err == nil {
		domain, err := cliClient.DomainService.ReadByName(token, domainName)
		if err == nil {
			resourceType, err := cliClient.ResourceTypeService.ReadByName(token, resourceTypeName)
			if err == nil {
				resource := models.Resource{Name: name, LogicalComponent: *logicalComponent, Domain: *domain, ResourceType: *resourceType}
				err = cliClient.ResourceService.Create(token, resource)
				if err != nil {
					t.Log(err)
				}
			} else {
				t.Log(err)
			}
		} else {
			t.Log(err)
		}
	} else {
		t.Log(err)
	}

	logicalComponentName = "not_utils"

	t.Log("RESOURCE >> PIVOT.TEST >> CREATE - FAIL: WRONG logicalComponentName")
	if err == nil {
		domain, err := cliClient.DomainService.ReadByName(token, domainName)
		if err == nil {
			resourceType, err := cliClient.ResourceTypeService.ReadByName(token, resourceTypeName)
			if err == nil {
				resource := models.Resource{Name: name, LogicalComponent: *logicalComponent, Domain: *domain, ResourceType: *resourceType}
				err = cliClient.ResourceService.Create(token, resource)
				if err != nil {
					t.Log((err))
				}
			} else {
				t.Log((err))
			}
		} else {
			t.Log((err))
		}
	} else {
		t.Log((err))
	}
	logicalComponentName = "utils"
	resourceTypeName = "failed_pivot"

	t.Log("RESOURCE >> PIVOT.TEST >> CREATE - FAIL: WRONG resourceTypeName")
	if err == nil {
		domain, err := cliClient.DomainService.ReadByName(token, domainName)
		if err == nil {
			resourceType, err := cliClient.ResourceTypeService.ReadByName(token, resourceTypeName)
			if err == nil {
				resource := models.Resource{Name: name, LogicalComponent: *logicalComponent, Domain: *domain, ResourceType: *resourceType}
				err = cliClient.ResourceService.Create(token, resource)
				if err != nil {
					t.Log((err))
				}
			} else {
				t.Log((err))
			}
		} else {
			t.Log((err))
		}
	} else {
		t.Log((err))
	}
}

func TestCreateResource__ProductConnector(t *testing.T) {
	const (
		token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjY3MzQyMTg1MTUwODg2NzA3MywiTmFtZSI6IkFkbWluaXN0cmF0b3IxIiwiRW1haWwiOiJBZG1pbmlzdHJhdG9yMSIsIlRlbmFudCI6IiIsImV4cCI6MTYzMTU4MjUwMX0.dYSuQzl27yE4wJFrP3H-Ck6ZOKEGxQE66nGX4UVStLA"
	)
	cliClient := cli.NewClient("http://localhost:9203")

	name := "testResource"
	logicalComponentName := "UtilsCustom1.0"
	domainName := "test"
	resourceTypeName := "productConnector"

	t.Log("RESOURCE >> PRODUCT_CONNECTOR.TEST >> CREATE - SUCCESS")

	logicalComponent, err := cliClient.LogicalComponentService.ReadByName(token, logicalComponentName)
	fmt.Println("err")
	fmt.Println(err)
	if err == nil {
		domain, err := cliClient.DomainService.ReadByName(token, domainName)
		if err == nil {
			resourceType, err := cliClient.ResourceTypeService.ReadByName(token, resourceTypeName)
			if err == nil {
				resource := models.Resource{Name: name, LogicalComponent: *logicalComponent, Domain: *domain, ResourceType: *resourceType}
				err = cliClient.ResourceService.Create(token, resource)
				if err != nil {
					t.Log((err))
				}
			} else {
				t.Log((err))
			}
		} else {
			t.Log((err))
		}
	} else {
		t.Log((err))
	}

	logicalComponentName = "not_UtilsCustom1.0"

	t.Log("RESOURCE >> PRODUCT_CONNECTOR.TEST >> CREATE - FAIL: WRONG logicalComponentName")
	if err == nil {
		domain, err := cliClient.DomainService.ReadByName(token, domainName)
		if err == nil {
			resourceType, err := cliClient.ResourceTypeService.ReadByName(token, resourceTypeName)
			if err == nil {
				resource := models.Resource{Name: name, LogicalComponent: *logicalComponent, Domain: *domain, ResourceType: *resourceType}
				err = cliClient.ResourceService.Create(token, resource)
				if err != nil {
					t.Log((err))
				}
			} else {
				t.Log((err))
			}
		} else {
			t.Log((err))
		}
	} else {
		t.Log((err))
	}
	logicalComponentName = "UtilsCustom1.0"
	resourceTypeName = "failed_pivot"

	t.Log("RESOURCE >> PRODUCT_CONNECTOR.TEST >> CREATE - FAIL: WRONG resourceTypeName")
	if err == nil {
		domain, err := cliClient.DomainService.ReadByName(token, domainName)
		if err == nil {
			resourceType, err := cliClient.ResourceTypeService.ReadByName(token, resourceTypeName)
			if err == nil {
				resource := models.Resource{Name: name, LogicalComponent: *logicalComponent, Domain: *domain, ResourceType: *resourceType}
				err = cliClient.ResourceService.Create(token, resource)
				if err != nil {
					t.Log((err))
				}
			} else {
				t.Log((err))
			}
		} else {
			t.Log((err))
		}
	} else {
		t.Log((err))
	}
}
func TestCreateRole(t *testing.T) {
	const (
		token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjY3MzQyMTg1MTUwODg2NzA3MywiTmFtZSI6IkFkbWluaXN0cmF0b3IxIiwiRW1haWwiOiJBZG1pbmlzdHJhdG9yMSIsIlRlbmFudCI6IiIsImV4cCI6MTYzMTU4MjUwMX0.dYSuQzl27yE4wJFrP3H-Ck6ZOKEGxQE66nGX4UVStLA"
	)
	cliClient := cli.NewClient("http://localhost:9203")
	name := "testRole"

	t.Log("ROLE >> TEST 1/2 >> CREATE - SUCCESS")
	role := models.Role{Name: name}
	err := cliClient.RoleService.Create(token, role)
	if err != nil {
		t.Log(err)
	}

	t.Log("ROLE >> TEST 2/2 >> CREATE - FAIL: DOUBLE SAME NAME")
	role = models.Role{Name: name}
	err = cliClient.RoleService.Create(token, role)
	if err != nil {
		t.Log(err)
	}
}

func TestUpdateRole(t *testing.T) {
	const (
		token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjY3MzQyMTg1MTUwODg2NzA3MywiTmFtZSI6IkFkbWluaXN0cmF0b3IxIiwiRW1haWwiOiJBZG1pbmlzdHJhdG9yMSIsIlRlbmFudCI6IiIsImV4cCI6MTYzMTU4MjUwMX0.dYSuQzl27yE4wJFrP3H-Ck6ZOKEGxQE66nGX4UVStLA"
	)
	cliClient := cli.NewClient("http://localhost:9203")
	name := "testRole"
	newName := "newTestRole"

	t.Log("ROLE >> TEST >> UPDATE - SUCCESS")
	oldRole, err := cliClient.RoleService.ReadByName(token, name)
	if err == nil {
		role := models.Role{Name: newName}
		err = cliClient.RoleService.Update(token, int(oldRole.ID), role)
		if err != nil {
			t.Log(err)
		}
	} else {
		t.Log(err)
	}

	t.Log("ROLE >> TEST >> UPDATE - FAIL: NAME DO NOT EXISTS")
	name = "DONOTEXIST"

	if err == nil {
		role := models.Role{Name: newName}
		err = cliClient.RoleService.Update(token, int(oldRole.ID), role)
		if err != nil {
			t.Log(err)
		}
	} else {
		t.Log(err)
	}

	t.Log("ROLE >> TEST >> UPDATE - FAIL: SAME NAME THAN BASE NAME")
	name = "newTestRole"

	if err == nil {
		role := models.Role{Name: newName}
		err = cliClient.RoleService.Update(token, int(oldRole.ID), role)
		if err != nil {
			t.Log(err)
		}
	} else {
		t.Log(err)
	}
}
func TestDeleteRole(t *testing.T) {
	const (
		token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjY3MzQyMTg1MTUwODg2NzA3MywiTmFtZSI6IkFkbWluaXN0cmF0b3IxIiwiRW1haWwiOiJBZG1pbmlzdHJhdG9yMSIsIlRlbmFudCI6IiIsImV4cCI6MTYzMTU4MjUwMX0.dYSuQzl27yE4wJFrP3H-Ck6ZOKEGxQE66nGX4UVStLA"
	)
	cliClient := cli.NewClient("http://localhost:9203")

	t.Log("ROLE >> TEST >> UPDATE - SUCCESS")
	name := "testRole"
	oldRole, err := cliClient.RoleService.ReadByName(token, name)
	if err == nil {
		err = cliClient.RoleService.Delete(token, int(oldRole.ID))
		if err != nil {
			t.Log(err)
		}
	} else {
		t.Log(err)
	}

	t.Log("ROLE >> TEST >> UPDATE - FAIL: NAME DOES NOT EXISTS")
	name = "DONOTEXIST"
	if err == nil {
		err = cliClient.RoleService.Delete(token, int(oldRole.ID))
		if err != nil {
			t.Log(err)
		}
	} else {
		t.Log(err)
	}
}
func TestCreateUser(t *testing.T) {
	const (
		token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjY3MzQyMTg1MTUwODg2NzA3MywiTmFtZSI6IkFkbWluaXN0cmF0b3IxIiwiRW1haWwiOiJBZG1pbmlzdHJhdG9yMSIsIlRlbmFudCI6IiIsImV4cCI6MTYzMTU4MjUwMX0.dYSuQzl27yE4wJFrP3H-Ck6ZOKEGxQE66nGX4UVStLA"
	)
	cliClient := cli.NewClient("http://localhost:9203")

	name := "neo"
	email := "testEmail4"
	password := "password"

	t.Log("USER >> TEST >> CREATE - SUCCESS")

	user := models.NewUser(name, email, password)
	err := cliClient.UserService.Create(token, user)
	if err != nil {
		fmt.Println(err)
	}

	t.Log("USER >> TEST >> CREATE - FAIL: SAME NAME ALREADY EXISTS")
	user = models.NewUser(name, email, password)
	err = cliClient.UserService.Create(token, user)
	if err != nil {
		fmt.Println(err)
	}
}

func TestUpdateUser(t *testing.T) {
	const (
		token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjY3MzQyMTg1MTUwODg2NzA3MywiTmFtZSI6IkFkbWluaXN0cmF0b3IxIiwiRW1haWwiOiJBZG1pbmlzdHJhdG9yMSIsIlRlbmFudCI6IiIsImV4cCI6MTYzMTU4MjUwMX0.dYSuQzl27yE4wJFrP3H-Ck6ZOKEGxQE66nGX4UVStLA"
	)
	cliClient := cli.NewClient("http://localhost:9203")

	name := "neo"
	newName := "neo2"
	email := "testEmail4"
	password := "password"

	t.Log("USER >> TEST >> UPDATE - SUCCESS")
	oldUser, err := cliClient.UserService.ReadByName(token, name)
	if err == nil {
		user := models.NewUser(newName, email, password)
		err = cliClient.UserService.Update(token, int(oldUser.ID), user)
		if err != nil {
			t.Log(err)
		}
	} else {
		t.Log(err)
	}

	name = "notExists"

	t.Log("USER >> TEST >> UPDATE - FAIL: WRONG NAME")
	oldUser, err = cliClient.UserService.ReadByName(token, name)
	if err == nil {
		user := models.NewUser(newName, email, password)
		err = cliClient.UserService.Update(token, int(oldUser.ID), user)
		if err != nil {
			t.Log(err)
		}
	} else {
		t.Log(err)
	}

	name = "user"
	password = "notExists"

	t.Log("USER >> TEST >> UPDATE - FAIL: WRONG PASSWORD")
	oldUser, err = cliClient.UserService.ReadByName(token, name)
	if err == nil {
		user := models.NewUser(newName, email, password)
		err = cliClient.UserService.Update(token, int(oldUser.ID), user)
		if err != nil {
			t.Log(err)
		}
	} else {
		t.Log(err)
	}
}

func TestDeleteUser(t *testing.T) {
	const (
		token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjY3MzQyMTg1MTUwODg2NzA3MywiTmFtZSI6IkFkbWluaXN0cmF0b3IxIiwiRW1haWwiOiJBZG1pbmlzdHJhdG9yMSIsIlRlbmFudCI6IiIsImV4cCI6MTYzMTU4MjUwMX0.dYSuQzl27yE4wJFrP3H-Ck6ZOKEGxQE66nGX4UVStLA"
	)
	cliClient := cli.NewClient("http://localhost:9203")

	t.Log("USER >> TEST >> DELETE - FAIL: DO NOT EXISTS")
	name := "UserDoNotExists"
	oldUser, err := cliClient.UserService.ReadByName(token, name)
	if err == nil {
		err = cliClient.UserService.Delete(token, int(oldUser.ID))
		if err != nil {
			t.Log((err))
		}
	} else {
		t.Log((err))
	}
	name = "neo2"
	t.Log("USER >> TEST >> DELETE - SUCCESS")

	oldUser, err = cliClient.UserService.ReadByName(token, name)
	if err == nil {
		err = cliClient.UserService.Delete(token, int(oldUser.ID))
		if err != nil {
			t.Log((err))
		}
	} else {
		t.Log((err))
	}

	t.Log("USER >> TEST >> DELETE - FAIL: ALREADY DELETED")

	oldUser, err = cliClient.UserService.ReadByName(token, name)
	if err == nil {
		err = cliClient.UserService.Delete(token, int(oldUser.ID))
		if err != nil {
			t.Log((err))
		}
	} else {
		t.Log((err))
	}
}
func TestCreateEventTypeToPoll(t *testing.T) {
	const (
		token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjY3MzQyMTg1MTUwODg2NzA3MywiTmFtZSI6IkFkbWluaXN0cmF0b3IxIiwiRW1haWwiOiJBZG1pbmlzdHJhdG9yMSIsIlRlbmFudCI6IiIsImV4cCI6MTYzMTU4MjUwMX0.dYSuQzl27yE4wJFrP3H-Ck6ZOKEGxQE66nGX4UVStLA"
	)
	cliClient := cli.NewClient("http://localhost:9203")

	domainName := "testDomain"
	parentName := "root"

	//Create Domain
	domain := models.Domain{Name: domainName}
	err := cliClient.DomainService.Create(token, domain, parentName)
	if err != nil {
		t.Log((err))
	}

	// Create ResourceType
	typeName := "pivot"
	pivotProductConnectorName := "utils"
	resourceTypeName := "testResourceType"

	if typeName == "pivot" {
		pivot, err := cliClient.PivotService.ReadByName(token, pivotProductConnectorName)

		if err == nil {
			resourceType := models.ResourceType{Name: resourceTypeName, Pivot: *pivot}
			err := cliClient.ResourceTypeService.Create(token, resourceType)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	}

	// Create Resource
	resourceName := "testResource"
	logicalComponentName := "connector5"

	logicalComponent, err := cliClient.LogicalComponentService.ReadByName(token, logicalComponentName)
	if err == nil {
		domain, err := cliClient.DomainService.ReadByName(token, domainName)
		if err == nil {
			resourceType, err := cliClient.ResourceTypeService.ReadByName(token, resourceTypeName)
			if err == nil {
				resource := models.Resource{Name: resourceName, LogicalComponent: *logicalComponent, Domain: *domain, ResourceType: *resourceType}
				err = cliClient.ResourceService.Create(token, resource)
				if err != nil {
					t.Log((err))
				} else {
					t.Log((err))
				}
			} else {
				t.Log((err))
			}
		} else {
			t.Log((err))
		}
	} else {
		t.Log((err))
	}

	// Create EventType
	eventTypeName := "testEventType"
	schema := "testSchema"

	if typeName == "pivot" {
		pivot, err := cliClient.PivotService.ReadByName(token, pivotProductConnectorName)

		if err == nil {
			eventType := models.EventType{Name: eventTypeName, Schema: schema, Pivot: *pivot}
			err := cliClient.EventTypeService.Create(token, eventType)
			if err != nil {
				t.Log(err)
			}
		} else {
			t.Log(err)
		}
	} else {
		t.Log(err)
	}

	// Create normal EventTypeToPoll
	t.Log("EVENTTYPETOPOLL >> TEST >> CREATE - SUCCESS")
	resource, err := cliClient.ResourceService.ReadByName(token, resourceName)
	if err == nil {
		eventType, err := cliClient.EventTypeService.ReadByName(token, eventTypeName)
		if err == nil {
			eventTypeToPoll := models.EventTypeToPoll{Resource: *resource, EventType: *eventType}
			err := cliClient.EventTypeToPollService.Create(token, eventTypeToPoll)
			if err != nil {
				t.Log(err)
			}
		} else {
			t.Log(err)
		}
	} else {
		t.Log(err)
	}

	// Recreate the normal EventTypeToPoll
	t.Log("EVENTTYPETOPOLL >> TEST >> CREATE - FAIL: ALREADY EXISTS (SAME NAME)")
	resource, err = cliClient.ResourceService.ReadByName(token, resourceName)
	if err == nil {
		eventType, err := cliClient.EventTypeService.ReadByName(token, eventTypeName)
		if err == nil {
			eventTypeToPoll := models.EventTypeToPoll{Resource: *resource, EventType: *eventType}
			err := cliClient.EventTypeToPollService.Create(token, eventTypeToPoll)
			if err != nil {
				t.Log(err)
			}
		} else {
			t.Log(err)
		}
	} else {
		t.Log(err)
	}
}

func TestUpdateEventTypeToPoll(t *testing.T) {
	const (
		token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjY3MzQyMTg1MTUwODg2NzA3MywiTmFtZSI6IkFkbWluaXN0cmF0b3IxIiwiRW1haWwiOiJBZG1pbmlzdHJhdG9yMSIsIlRlbmFudCI6IiIsImV4cCI6MTYzMTU4MjUwMX0.dYSuQzl27yE4wJFrP3H-Ck6ZOKEGxQE66nGX4UVStLA"
	)
	cliClient := cli.NewClient("http://localhost:9203")

	t.Log("EVENTTYPETOPOLL - UPDATE: SUCCESS")
	resourceName := "testResource"
	eventTypeName := "testEventType"

	oldEventTypeToPoll, err := cliClient.EventTypeToPollService.ReadByName(token, eventTypeName)
	if err == nil {
		resource, err := cliClient.ResourceService.ReadByName(token, resourceName)
		if err == nil {
			eventType, err := cliClient.EventTypeService.ReadByName(token, eventTypeName)
			if err == nil {
				eventTypeToPoll := models.EventTypeToPoll{Resource: *resource, EventType: *eventType}
				err = cliClient.EventTypeToPollService.Update(token, int(oldEventTypeToPoll.ID), eventTypeToPoll)
				if err != nil {
					t.Log(err)
				}
			} else {
				t.Log(err)
			}
		} else {
			t.Log(err)
		}
	} else {
		t.Log(err)
	}

}
func TestDeleteEventTypeToPoll(t *testing.T) {
	const (
		token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjY3MzQyMTg1MTUwODg2NzA3MywiTmFtZSI6IkFkbWluaXN0cmF0b3IxIiwiRW1haWwiOiJBZG1pbmlzdHJhdG9yMSIsIlRlbmFudCI6IiIsImV4cCI6MTYzMTU4MjUwMX0.dYSuQzl27yE4wJFrP3H-Ck6ZOKEGxQE66nGX4UVStLA"
	)
	cliClient := cli.NewClient("http://localhost:9203")

	resourceName := "testResource"

	oldEventTypeToPoll, err := cliClient.EventTypeToPollService.ReadByName(token, resourceName)
	if err == nil {
		err = cliClient.EventTypeToPollService.Delete(token, int(oldEventTypeToPoll.ID))
		if err != nil {
			t.Log(err)
		}
	} else {
		t.Log(err)
	}
}

func TestCreateTenant(t *testing.T) {
	const (
		token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjY3MzQyMTg1MTUwODg2NzA3MywiTmFtZSI6IkFkbWluaXN0cmF0b3IxIiwiRW1haWwiOiJBZG1pbmlzdHJhdG9yMSIsIlRlbmFudCI6IiIsImV4cCI6MTYzMTU4MjUwMX0.dYSuQzl27yE4wJFrP3H-Ck6ZOKEGxQE66nGX4UVStLA"
	)
	cliClient := cli.NewClient("http://localhost:9203")

	name := "testTenant"

	tenant := models.Tenant{Name: name}
	login, password, err := cliClient.TenantService.Create(token, tenant)
	if err == nil {
		t.Log("login : " + login)
		t.Log("password : " + password)
	} else {
		t.Log(err)
	}
}