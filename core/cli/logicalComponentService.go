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

//Upload :
/*
func (ls *LogicalComponentService) Upload(uri string, token string, name string, params map[string]string, paramName, path string) (error*http.Request, error*models.LogicalComponent, error) {
	req, err := ls.client.newRequest("POST", "/auth/gandalf/logicalcomponents/"+name, token, nil)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}
	file.Close()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, fi.Name())
	if err != nil {
		return nil, err
	}
	part.Write(fileContents)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	return http.NewRequest("POST", uri, body)
} */
