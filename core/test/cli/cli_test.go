package cli

import (
	"testing"

	"github.com/ditrit/gandalf/core/cli"
	"github.com/ditrit/gandalf/core/models"
)

func TestDomain(t *testing.T) {
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

func TestResourceTypePivot(t *testing.T) {
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
func TestResourceTypeProductConnector(t *testing.T) {
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
	pivotProductConnectorName := "utils"
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

	t.Log("PRODUCT_CONNECTOR.TEST >> INCORRECT TYPE NAME (== CONNECTOR_PRODUCT && CONNECTOR_PRODUCT.NAME != utils)")
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
