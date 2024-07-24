package main

import (
	"api_testing/analyze"
)

func main() {

	rootDir := "/app/files/employee-service"
	
	newmanReportPath := "/app/files/employee-service/collection/newman-report.json"

	analyze.Analysis(rootDir, newmanReportPath)
}
