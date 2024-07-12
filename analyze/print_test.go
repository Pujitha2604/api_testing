package analyze

import (
	"bytes"
	"os"
	"testing"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/stretchr/testify/mock"
)

// MockWriter is a mock implementation of Writer for capturing table output.
type MockWriter struct {
	mock.Mock
}

func (m *MockWriter) AppendHeader(row table.Row) {
	_ = m.Called(row)
}

func (m *MockWriter) AppendRow(row table.Row) {
	_ = m.Called(row)
}

func (m *MockWriter) Render() {
	_ = m.Called()
}

func TestPrintEndpointsTable(t *testing.T) {
	mockWriter := new(MockWriter)

	// Set expectations for methods called on MockWriter
	mockWriter.On("AppendHeader", []interface{}{table.Row{"#", "METHOD", "PATH", "RESULT", "SOURCE"}}).Once()
	mockWriter.On("AppendRow", []interface{}{table.Row{1, "GET", "/api/endpoint1", "Covered", "path/to/newman/report.json"}}).Once()
	mockWriter.On("AppendRow", []interface{}{table.Row{2, "POST", "/api/endpoint2", "Not Covered", ""}}).Once()
	mockWriter.On("Render").Once()

	// Replace os.Stdout with a buffer to capture output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Restore os.Stdout after the test
	defer func() {
		w.Close()
		os.Stdout = old
	}()

	// Define test data
	allEndpoints := map[string]Endpoint{
		"1": {Method: "GET", Path: "/api/endpoint1", Result: "Covered"},
		"2": {Method: "POST", Path: "/api/endpoint2", Result: "Not Covered"},
	}
	newmanReportPath := "path/to/newman/report.json"

	// Call the function with mocked dependencies
	go func() {
		printEndpointsTable(allEndpoints, newmanReportPath)
	}()

	// Read output from the buffer
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)

	// Assert expectations
	mockWriter.AssertExpectations(t)

	// Perform assertions on the captured output
	expectedOutput := `+-----+--------+------------------+--------+--------------------------+
|  #  | METHOD |       PATH       | RESULT |          SOURCE          |
+-----+--------+------------------+--------+--------------------------+
|  1  | GET    | /api/endpoint1   | Covered| path/to/newman/report.json|
|  2  | POST   | /api/endpoint2   |Not Covered|                          |
+-----+--------+------------------+--------+--------------------------+
`
	if buf.String() != expectedOutput {
		t.Errorf("Expected output:\n%s\nBut got:\n%s", expectedOutput, buf.String())
	}
}
