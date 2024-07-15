package analyze

import (
	"encoding/json"
	"errors"
	"os"
	"path"
)

type NewmanReport struct {
	Run Run `json:"run"`
}

type Run struct {
	Executions []Execution `json:"executions"`
}

type Execution struct {
	Item     Item     `json:"item"`
	Response Response `json:"response"`
}

type Item struct {
	Name    string  `json:"name"`
	Request Request `json:"request"`
}

type Request struct {
	Method string `json:"method"`
	URL    URL    `json:"url"`
}

type URL struct {
	Path []string `json:"path"`
}

var osReadFile = os.ReadFile

func parseNewmanReport(filename string) (map[string]int, error) {
	data, err := osReadFile(filename)
	if err != nil {
		return nil, err
	}

	var report NewmanReport
	err = json.Unmarshal(data, &report)
	if err != nil {
		return nil, errors.New("error parsing Newman report JSON: " + err.Error())
	}

	endpoints := make(map[string]int)
	for _, execution := range report.Run.Executions {
		path := "/" + path.Join(execution.Item.Request.URL.Path...)
		endpoints[path] = execution.Response.Code
	}

	return endpoints, nil
}
