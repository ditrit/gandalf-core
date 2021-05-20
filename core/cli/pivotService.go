package cli

import (
	"strconv"

	"github.com/ditrit/gandalf/core/models"
)

// PivotService :
type PivotService struct {
	client *Client
}

// List :
func (ps *PivotService) List(token string) ([]models.Pivot, error) {
	req, err := ps.client.newRequest("GET", "/auth/gandalf/pivots/", token, nil)
	if err != nil {
		return nil, err
	}
	var pivots []models.Pivot
	err = ps.client.do(req, &pivots)
	return pivots, err
}

// Read :
func (ps *PivotService) Read(token string, id int) (*models.Pivot, error) {
	req, err := ps.client.newRequest("GET", "/auth/gandalf/pivots/"+strconv.Itoa(id), token, nil)
	if err != nil {
		return nil, err
	}
	var pivot models.Pivot
	err = ps.client.do(req, &pivot)
	return &pivot, err
}

// Read :
func (ps *PivotService) ReadByName(token string, name string) (*models.Pivot, error) {
	req, err := ps.client.newRequest("GET", "/auth/gandalf/pivots/"+name, token, nil)
	if err != nil {
		return nil, err
	}
	var pivot models.Pivot
	err = ps.client.do(req, &pivot)
	return &pivot, err
}
