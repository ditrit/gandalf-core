package cli

import (
	"fmt"
	"testing"

	"github.com/ditrit/gandalf/core/cli"
	"github.com/ditrit/gandalf/core/models"
)

func TestCreateDomain(t *testing.T) {
	const (
		token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjY2NDA1MDg1MTM4OTQ0MDAwMSwiTmFtZSI6IkFkbWluaXN0cmF0b3IyIiwiRW1haWwiOiJBZG1pbmlzdHJhdG9yMiIsIlRlbmFudCI6IiIsImV4cCI6MTYyODcyMjk3Mn0.6KTRZr9xl6rUqToWv_SUZypOVmwdRM4_sJhjRiEDpMU"
	)
	cliClient := cli.NewClient("http://localhost:9203")

	name := "test"
	parentName := ""

	// Good Scheme
	domain := models.Domain{Name: name}
	err := cliClient.DomainService.Create(token, domain, parentName)
	t.Log(err)

	parentName = "unknown"

	// Wrong Parent
	domain = models.Domain{Name: name}
	err = cliClient.DomainService.Create(token, domain, parentName)
	t.Log(err)
}

func TestCreateResourceType__Pivot(t *testing.T) {
	const (
		token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjY2NDA1MDg1MTM4OTQ0MDAwMSwiTmFtZSI6IkFkbWluaXN0cmF0b3IyIiwiRW1haWwiOiJBZG1pbmlzdHJhdG9yMiIsIlRlbmFudCI6IiIsImV4cCI6MTYyODcyMjk3Mn0.6KTRZr9xl6rUqToWv_SUZypOVmwdRM4_sJhjRiEDpMU"
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
		token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjY2NDA1MDg1MTM4OTQ0MDAwMSwiTmFtZSI6IkFkbWluaXN0cmF0b3IyIiwiRW1haWwiOiJBZG1pbmlzdHJhdG9yMiIsIlRlbmFudCI6IiIsImV4cCI6MTYyODcyMjk3Mn0.6KTRZr9xl6rUqToWv_SUZypOVmwdRM4_sJhjRiEDpMU"
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
		token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjY2NDA1MDg1MTM4OTQ0MDAwMSwiTmFtZSI6IkFkbWluaXN0cmF0b3IyIiwiRW1haWwiOiJBZG1pbmlzdHJhdG9yMiIsIlRlbmFudCI6IiIsImV4cCI6MTYyODcyMjk3Mn0.6KTRZr9xl6rUqToWv_SUZypOVmwdRM4_sJhjRiEDpMU"
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
		token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjY2NDA1MDg1MTM4OTQ0MDAwMSwiTmFtZSI6IkFkbWluaXN0cmF0b3IyIiwiRW1haWwiOiJBZG1pbmlzdHJhdG9yMiIsIlRlbmFudCI6IiIsImV4cCI6MTYyODcyMjk3Mn0.6KTRZr9xl6rUqToWv_SUZypOVmwdRM4_sJhjRiEDpMU"
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
			fmt.Println(err)
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
			fmt.Println(err)
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
		token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjY2NDA1MDg1MTM4OTQ0MDAwMSwiTmFtZSI6IkFkbWluaXN0cmF0b3IyIiwiRW1haWwiOiJBZG1pbmlzdHJhdG9yMiIsIlRlbmFudCI6IiIsImV4cCI6MTYyODcyMjk3Mn0.6KTRZr9xl6rUqToWv_SUZypOVmwdRM4_sJhjRiEDpMU"
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
		token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjY2NDA1MDg1MTM4OTQ0MDAwMSwiTmFtZSI6IkFkbWluaXN0cmF0b3IyIiwiRW1haWwiOiJBZG1pbmlzdHJhdG9yMiIsIlRlbmFudCI6IiIsImV4cCI6MTYyODcyMjk3Mn0.6KTRZr9xl6rUqToWv_SUZypOVmwdRM4_sJhjRiEDpMU"
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
		token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjY2NDA1MDg1MTM4OTQ0MDAwMSwiTmFtZSI6IkFkbWluaXN0cmF0b3IyIiwiRW1haWwiOiJBZG1pbmlzdHJhdG9yMiIsIlRlbmFudCI6IiIsImV4cCI6MTYyODcyMjk3Mn0.6KTRZr9xl6rUqToWv_SUZypOVmwdRM4_sJhjRiEDpMU"
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
		token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjY2NDA1MDg1MTM4OTQ0MDAwMSwiTmFtZSI6IkFkbWluaXN0cmF0b3IyIiwiRW1haWwiOiJBZG1pbmlzdHJhdG9yMiIsIlRlbmFudCI6IiIsImV4cCI6MTYyODcyMjk3Mn0.6KTRZr9xl6rUqToWv_SUZypOVmwdRM4_sJhjRiEDpMU"
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
