package analyze
 
import (
    "io/ioutil"
    "os"
    "testing"
)
 
func TestAnalyzeFileForAPIEndpoints(t *testing.T) {
    testSource := `
package main
 
import (
    "net/http"
)
 
func main() {
    http.HandleFunc("/test", testHandler).Methods("GET")
}
 
func testHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
}
`
    // Create a temporary file with the test source code
    tempFile, err := ioutil.TempFile("", "*.go")
    if err != nil {
        t.Fatalf("Failed to create temp file: %v", err)
    }
    defer os.Remove(tempFile.Name())
 
    _, err = tempFile.WriteString(testSource)
    if err != nil {
        t.Fatalf("Failed to write to temp file: %v", err)
    }
    tempFile.Close()
 
    endpoints := analyzeFileForAPIEndpoints(tempFile.Name())
    if len(endpoints) != 1 {
        t.Fatalf("Expected 1 endpoint, got %d", len(endpoints))
    }
 
    expectedPath := "/test"
    endpoint, exists := endpoints[expectedPath]
    if !exists {
        t.Fatalf("Expected endpoint %s to exist", expectedPath)
    }
 
    if endpoint.Method != "GET" {
        t.Errorf("Expected method GET, got %s", endpoint.Method)
    }
 
    if endpoint.Path != expectedPath {
        t.Errorf("Expected path %s, got %s", expectedPath, endpoint.Path)
    }
 
    if endpoint.Result != "Not Compared" {
        t.Errorf("Expected result 'Not Compared', got %s", endpoint.Result)
    }
}
 
func TestExtractBasePath(t *testing.T) {
	tests := []struct {
		endpoint string
		expected string
	}{
		{"/test/path", "/test"},
		{"/test", "/test"},
		{"/test/path/extra", "/test"},
		{"/", "/"},
		{"test/path", "/test"},
	}

	for _, test := range tests {
		result := extractBasePath(test.endpoint)
		if result != test.expected {
			t.Errorf("extractBasePath(%q) = %v; want %v", test.endpoint, result, test.expected)
		}
	}
}

func TestMatchEndpoint(t *testing.T) {
	tests := []struct {
		handlerEndpoint string
		newmanEndpoint  string
		expected        bool
	}{
		{"/test", "/test", true},
		{"/test/path", "/test/path", true},
		{"/test/path", "/different/path", false},
		{"/test/path", "/test/path/extra", true},
		{"/", "/", true},
		{"/test/path", "/test/another", true},
	}

	for _, test := range tests {
		result := matchEndpoint(test.handlerEndpoint, test.newmanEndpoint)
		if result != test.expected {
			t.Errorf("matchEndpoint(%q, %q) = %v; want %v", test.handlerEndpoint, test.newmanEndpoint, result, test.expected)
		}
	}
}