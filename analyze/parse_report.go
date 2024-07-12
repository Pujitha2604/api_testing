package analyze

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
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

func parseNewmanReport(reportPath string) (map[string]int, error) {
	data, err := os.ReadFile(reportPath)
	if err != nil {
		return nil, fmt.Errorf("error reading Newman report file: %v", err)
	}

	var report NewmanReport
	if err := json.Unmarshal(data, &report); err != nil {
		return nil, fmt.Errorf("error parsing Newman report JSON: %v", err)
	}

	endpoints := make(map[string]int)
	for _, exec := range report.Run.Executions {
		path := exec.Item.Request.URL.Path
		if len(path) > 0 && path[0] == "employee" && len(path) > 1 && path[1] != "" {
			path = path[:1]
		}
		pathStr := "/" + strings.Join(path, "/")
		endpoints[pathStr] = exec.Response.Code
	}

	return endpoints, nil
}
