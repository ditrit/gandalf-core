package cli

import (
	"github.com/ditrit/gandalf/core/models"
)

// LogicalComponentService :
type LogicalComponentService struct {
	client *Client
}

// List :
func (ls *LogicalComponentService) List(token string) ([]models.LogicalComponent, error) {
	req, err := ls.client.newRequest("GET", "/auth/gandalf/logicalcomponents/", token, nil)
	if err != nil {
		return nil, err
	}
	var logicalComponents []models.LogicalComponent
	err = ls.client.do(req, &logicalComponents)
	return logicalComponents, err
}

// ReadByName :
func (ls *LogicalComponentService) ReadByName(token string, name string) (*models.LogicalComponent, error) {
	req, err := ls.client.newRequest("GET", "/auth/gandalf/logicalcomponents/"+name, token, nil)
	if err != nil {
		return nil, err
	}
	var logicalComponent models.LogicalComponent
	err = ls.client.do(req, &logicalComponent)
	return &logicalComponent, err
}
