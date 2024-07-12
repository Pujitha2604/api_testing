package analyze

// import (
// 	"bytes"
// 	"testing"

// 	"github.com/jedib0t/go-pretty/v6/table"
	
// )

// // MockTableWriter implements the table.Writer interface for capturing output
// type MockTableWriter struct {
// 	buf bytes.Buffer
// }

// func (m *MockTableWriter) SetOutputMirror(outputMirror table.WriterOptions) {}

// func (m *MockTableWriter) AppendHeader(row table.Row) {}

// func (m *MockTableWriter) AppendRow(row table.Row) {
// 	m.buf.WriteString("| ")
// 	for _, cell := range row {
// 		m.buf.WriteString(cell.(string))
// 		m.buf.WriteString(" | ")
// 	}
// 	m.buf.WriteString("\n")
// }

// func (m *MockTableWriter) Render() {
// 	// Do nothing in the mock
// }

// func TestPrintEndpointsTable(t *testing.T) {
// 	// Mock data
// 	allEndpoints := map[string]Endpoint{
// 		"/endpoint1": {Method: "GET", Path: "/endpoint1", Result: "Covered"},
// 		"/endpoint2": {Method: "POST", Path: "/endpoint2", Result: "Not Covered"},
// 	}

// 	// Capture output using MockTableWriter
// 	mockWriter := &MockTableWriter{}

// 	// Run the function
// 	newmanReportPath := "/path/to/newman/report.json"
// 	printEndpointsTable(allEndpoints, newmanReportPath)

// 	// Validate the output
// 	expectedOutput := `| # | METHOD | PATH      | RESULT      | SOURCE                   | 
// | 1 | GET    | /endpoint1| Covered     | /path/to/newman/report.json| 
// | 2 | POST   | /endpoint2| Not Covered |                           | 
// `

// 	actualOutput := mockWriter.buf.String()

// 	if actualOutput != expectedOutput {
// 		t.Errorf("Expected output:\n%s\n\nActual output:\n%s", expectedOutput, actualOutput)
// 	}
// }
