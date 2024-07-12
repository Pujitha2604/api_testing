package analyze

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestParseNewmanReport(t *testing.T) {
	// Setup temporary Newman report file
	reportData := `
	{
		"run": {
			"executions": [
				{
					"item": {
						"name": "Test Endpoint",
						"request": {
							"method": "GET",
							"url": {
								"path": ["employee", "123"]
							}
						}
					},
					"response": {
						"code": 200
					}
				},
				{
					"item": {
						"name": "Invalid Endpoint",
						"request": {
							"method": "POST",
							"url": {
								"path": ["invalid"]
							}
						}
					},
					"response": {
						"code": 404
					}
				}
			]
		}
	}
	`

	tempFile, err := ioutil.TempFile("", "newman_report*.json")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	_, err = tempFile.WriteString(reportData)
	if err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	tempFile.Close()

	// Run the function
	endpoints, err := parseNewmanReport(tempFile.Name())
	if err != nil {
		t.Fatalf("Error parsing Newman report: %v", err)
	}

	// Validate the results
	expectedEndpoints := map[string]int{
		"/employee/123": 200,
		"/invalid":      404,
	}

	if len(endpoints) != len(expectedEndpoints) {
		t.Fatalf("Expected %d endpoints, got %d", len(expectedEndpoints), len(endpoints))
	}

	for path, code := range expectedEndpoints {
		if endpoints[path] != code {
			t.Errorf("Expected status code %d for path %s, got %d", code, path, endpoints[path])
		}
	}

	// Test error case: File not found
	_, err = parseNewmanReport("nonexistent-file.json")
	if err == nil {
		t.Error("Expected error for nonexistent file, but got nil")
	}

	// Test error case: Invalid JSON format
	invalidData := "invalid JSON data"
	invalidFile, _ := ioutil.TempFile("", "invalid_newman_report*.json")
	invalidFile.Close()
	ioutil.WriteFile(invalidFile.Name(), []byte(invalidData), 0644)

	_, err = parseNewmanReport(invalidFile.Name())
	if err == nil {
		t.Error("Expected error for invalid JSON format, but got nil")
	}
}
