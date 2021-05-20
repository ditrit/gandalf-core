package cli

import (
	"strconv"

	"github.com/ditrit/gandalf/core/models"
)

// ConnectorProductService :
type ConnectorProductService struct {
	client *Client
}

// List :
func (as *ConnectorProductService) List(token string) ([]models.ConnectorProduct, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/connectorproduct/", token, nil)
	if err != nil {
		return nil, err
	}
	var connectorProducts []models.ConnectorProduct
	err = as.client.do(req, &connectorProducts)
	return connectorProducts, err
}

// Create :
func (as *ConnectorProductService) Create(token string, resource models.ConnectorProduct) error {
	req, err := as.client.newRequest("POST", "/auth/gandalf/connectorproduct/", token, resource)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Read :
func (as *ConnectorProductService) Read(token string, id int) (*models.ConnectorProduct, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/connectorproduct/"+strconv.Itoa(id), token, nil)
	if err != nil {
		return nil, err
	}
	var resource models.ConnectorProduct
	err = as.client.do(req, &resource)
	return &resource, err
}

// Read :
func (as *ConnectorProductService) ReadByName(token string, name string) (*models.ConnectorProduct, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/connectorproduct/"+name, token, nil)
	if err != nil {
		return nil, err
	}
	var resource models.ConnectorProduct
	err = as.client.do(req, &resource)
	return &resource, err
}

// Update :
func (as *ConnectorProductService) Update(token string, id int, resource models.ConnectorProduct) error {
	req, err := as.client.newRequest("PUT", "/auth/gandalf/connectorproduct/"+strconv.Itoa(id), token, resource)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Delete :
func (as *ConnectorProductService) Delete(token string, id int) error {
	req, err := as.client.newRequest("DELETE", "/auth/gandalf/connectorproduct/"+strconv.Itoa(id), token, nil)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}
