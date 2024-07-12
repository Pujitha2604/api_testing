package main

import (
	"api_testing/analyze"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Root directory not provided. Usage: go run main.go <rootDir>")
	}

	rootDir := os.Args[1]
	collectionFile := "C:/Users/Relanto/Downloads/api_testing/api_testing/collection.json"
	newmanReportPath := "C:/Users/Relanto/Downloads/api_testing/api_testing/newman-report.json"

	if err := analyze.RunNewman(collectionFile, newmanReportPath); err != nil {
		log.Fatalf("Error running Newman: %v", err)
	}

	analyze.Analysis(rootDir, newmanReportPath)
}
