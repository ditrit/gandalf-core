package cli

import (
	"strconv"

	"github.com/ditrit/gandalf/core/models"
)

// ProductConnectorService :
type ProductConnectorService struct {
	client *Client
}

// List :
func (as *ProductConnectorService) List(token string) ([]models.ProductConnector, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/productconnectors/", token, nil)
	if err != nil {
		return nil, err
	}
	var productConnectors []models.ProductConnector
	err = as.client.do(req, &productConnectors)
	return productConnectors, err
}

// Read :
func (as *ProductConnectorService) Read(token string, id int) (*models.ProductConnector, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/productconnectors/"+strconv.Itoa(id), token, nil)
	if err != nil {
		return nil, err
	}
	var resource models.ProductConnector
	err = as.client.do(req, &resource)
	return &resource, err
}

// Read :
func (as *ProductConnectorService) ReadByName(token string, name string) (*models.ProductConnector, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/productconnectors/"+name, token, nil)
	if err != nil {
		return nil, err
	}
	var resource models.ProductConnector
	err = as.client.do(req, &resource)
	return &resource, err
}
