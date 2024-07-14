package analyze

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintEndpointsTable(t *testing.T) {
	// Mock data
	allEndpoints := map[string]Endpoint{
		"1": {Method: "GET", Path: "/endpoint1", Result: "Covered"},
		"2": {Method: "POST", Path: "/endpoint2", Result: "Not Covered"},
	}
	newmanReportPath := "/path/to/newman-report.json"

	// Create a pipe to capture stdout
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to create pipe: %v", err)
	}

	// Save the original stdout
	oldStdout := os.Stdout

	// Set the stdout to our pipe writer
	os.Stdout = writer

	// Ensure we restore stdout after the test
	defer func() {
		os.Stdout = oldStdout
		writer.Close()
		reader.Close()
	}()

	// Run the function in a separate goroutine
	done := make(chan struct{})
	go func() {
		printEndpointsTable(allEndpoints, newmanReportPath)
		writer.Close()
		close(done)
	}()

	// Read the captured output
	var mockBuffer bytes.Buffer
	if _, err := io.Copy(&mockBuffer, reader); err != nil {
		t.Fatalf("Failed to read captured output: %v", err)
	}

	// Wait for the goroutine to finish
	<-done

	// Get the output from the buffer
	output := mockBuffer.String()

	// Assertions
	assert.Contains(t, output, "#")
	assert.Contains(t, output, "METHOD")
	assert.Contains(t, output, "PATH")
	assert.Contains(t, output, "RESULT")
	assert.Contains(t, output, "SOURCE")
	assert.Contains(t, output, "GET")
	assert.Contains(t, output, "/endpoint1")
	assert.Contains(t, output, "Covered")
	assert.Contains(t, output, "POST")
	assert.Contains(t, output, "/endpoint2")
	assert.Contains(t, output, "Not Covered")
}
