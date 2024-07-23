package main

import (
	"api_testing/analyze"
)

func main() {

	rootDir := "/home/ec2-user/web-service"
	
	newmanReportPath := "/home/ec2-user/web-service/collection/newman-report.json"

	analyze.Analysis(rootDir, newmanReportPath)
}
