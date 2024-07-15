package analyze

import (
	"log"
)

var (
	getAllGoFilesFunc        = getAllGoFiles
	analyzeFileForAPIEndpointsFunc = analyzeFileForAPIEndpoints
	parseNewmanReportFunc    = parseNewmanReport
	printEndpointsTableFunc  = printEndpointsTable
)

func Analysis(rootDir string, newmanReportPath string) {
	allEndpoints := make(map[string]Endpoint)

	// Retrieve handler files
	handlerFiles, err := getAllGoFilesFunc(rootDir)
	if err != nil {
		log.Fatalf("Error retrieving handler files: %v", err)
	}

	// Analyze endpoints from handler files
	for _, file := range handlerFiles {
		endpoints := analyzeFileForAPIEndpointsFunc(file)
		for endpoint, endpointDetails := range endpoints {
			allEndpoints[endpoint] = Endpoint{
				Method: endpointDetails.Method,
				Path:   endpoint,
				Result: "Not Compared",
			}
		}
	}

	// Parse Newman report for existing endpoints and statuses
	newmanEndpoints, err := parseNewmanReportFunc(newmanReportPath)
	if err != nil {
		log.Fatalf("Error parsing Newman report: %v", err)
	}

	// Compare analyzed endpoints with Newman endpoints and set results
	for endpoint, details := range allEndpoints {
		newmanStatus, exists := newmanEndpoints[endpoint]
		if exists && matchEndpoint(endpoint, endpoint) {
			if newmanStatus == 200 {
				allEndpoints[endpoint] = Endpoint{
					Method: details.Method,
					Path:   details.Path,
					Result: "Success",
				}
			} else {
				allEndpoints[endpoint] = Endpoint{
					Method: details.Method,
					Path:   details.Path,
					Result: "Failure",
				}
			}
		} else {
			allEndpoints[endpoint] = Endpoint{
				Method: details.Method,
				Path:   details.Path,
				Result: "Not Covered",
			}
		}
	}

	// Print or process the final table of endpoints
	printEndpointsTableFunc(allEndpoints, newmanReportPath)
}
