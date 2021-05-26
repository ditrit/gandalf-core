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
func (as *ConnectorProductService) List(token string) ([]models.ProductConnector, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/productconnectors/", token, nil)
	if err != nil {
		return nil, err
	}
	var productConnectors []models.ProductConnector
	err = as.client.do(req, &productConnectors)
	return productConnectors, err
}

// Create :
func (as *ConnectorProductService) Create(token string, resource models.ProductConnector) error {
	req, err := as.client.newRequest("POST", "/auth/gandalf/productconnectors/", token, resource)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Read :
func (as *ConnectorProductService) Read(token string, id int) (*models.ProductConnector, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/productconnectors/"+strconv.Itoa(id), token, nil)
	if err != nil {
		return nil, err
	}
	var resource models.ProductConnector
	err = as.client.do(req, &resource)
	return &resource, err
}

// Read :
func (as *ConnectorProductService) ReadByName(token string, name string) (*models.ProductConnector, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/productconnectors/"+name, token, nil)
	if err != nil {
		return nil, err
	}
	var resource models.ProductConnector
	err = as.client.do(req, &resource)
	return &resource, err
}

// Update :
func (as *ConnectorProductService) Update(token string, id int, resource models.ProductConnector) error {
	req, err := as.client.newRequest("PUT", "/auth/gandalf/productconnectors/"+strconv.Itoa(id), token, resource)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Delete :
func (as *ConnectorProductService) Delete(token string, id int) error {
	req, err := as.client.newRequest("DELETE", "/auth/gandalf/productconnectors/"+strconv.Itoa(id), token, nil)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}
