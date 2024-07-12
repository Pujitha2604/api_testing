package analyze

import (
	"log"
)


func Analysis(rootDir string, newmanReportPath string) {
	allEndpoints := make(map[string]Endpoint)
	handlerFiles, err := getAllGoFiles(rootDir)
	if err != nil {
		log.Fatalf("Error retrieving handler files: %v", err)
	}

	for _, file := range handlerFiles {
		endpoints := analyzeFileForAPIEndpoints(file)
		for endpoint, endpointDetails := range endpoints {
			allEndpoints[endpoint] = Endpoint{
				Method: endpointDetails.Method,
				Path:   endpoint,
				Result: "Not Compared",
			}
		}
	}

	newmanEndpoints, err := parseNewmanReport(newmanReportPath)
	if err != nil {
		log.Fatalf("Error parsing Newman report: %v", err)
	}

	for endpoint, details := range allEndpoints {
		for newmanEndpoint, newmanStatus := range newmanEndpoints {
			if matchEndpoint(endpoint, newmanEndpoint) {
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
				break
			} else {
				allEndpoints[endpoint] = Endpoint{
					Method: details.Method,
					Path:   details.Path,
					Result: "Not Covered",
				}
			}
		}
	}

	printEndpointsTable(allEndpoints, newmanReportPath)
}