package main

import (
	"api_testing/analyze"
)

func main() {
	// if len(os.Args) < 2 {
	// 	log.Fatal("Root directory not provided. Usage: go run main.go <rootDir>")
	// }

	rootDir := "/home/ec2-user/web-service"
	//collectionFile := "C:/Users/Rekanto/Desktop/api_testing/collection.json"
	newmanReportPath := "/home/ec2-user/web-service/collection/newman-report.json"

	// if err := analyze.RunNewman(collectionFile, newmanReportPath); err != nil {
	// 	log.Fatalf("Error running Newman: %v", err)
	// }

	analyze.Analysis(rootDir, newmanReportPath)
}
