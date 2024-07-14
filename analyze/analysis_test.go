package analyze

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// // Mock for Analyzer interface
// type MockedAnalyzer struct {
// 	mock.Mock
// }

// func (m *MockedAnalyzer) getAllGoFiles(rootDir string) ([]string, error) {
// 	args := m.Called(rootDir)
// 	return args.Get(0).([]string), args.Error(1)
// }

// func (m *MockedAnalyzer) analyzeFileForAPIEndpoints(file string) map[string]EndpointDetails {
// 	args := m.Called(file)
// 	return args.Get(0).(map[string]EndpointDetails)
// }

// func (m *MockedAnalyzer) parseNewmanReport(newmanReportPath string) (map[string]int, error) {
// 	args := m.Called(newmanReportPath)
// 	return args.Get(0).(map[string]int), args.Error(1)
// }

// func TestAnalysis(t *testing.T) {
// 	// Initialize mock
// 	mockAnalyzer := new(MockedAnalyzer)

// 	// Setup expectations for getAllGoFiles mock
// 	mockAnalyzer.On("getAllGoFiles", "/path/to/handlers").Return([]string{"handler1.go", "handler2.go"}, nil)

// 	// Setup expectations for analyzeFileForAPIEndpoints mock
// 	mockAnalyzer.On("analyzeFileForAPIEndpoints", "handler1.go").Return(map[string]EndpointDetails{
// 		"/api/endpoint1": {Method: "GET"},
// 	})
// 	mockAnalyzer.On("analyzeFileForAPIEndpoints", "handler2.go").Return(map[string]EndpointDetails{
// 		"/api/endpoint2": {Method: "POST"},
// 	})

// 	// Setup expectations for parseNewmanReport mock
// 	mockAnalyzer.On("parseNewmanReport", "/path/to/newman/report.json").Return(map[string]int{
// 		"/api/endpoint1": 200,
// 	}, nil)

// 	// Call the function under test
// 	allEndpoints := Analysis("/path/to/handlers", "/path/to/newman/report.json", mockAnalyzer)

// 	// Assert the results
// 	expectedEndpoints := map[string]Endpoint{
// 		"/api/endpoint1": {Method: "GET", Path: "/api/endpoint1", Result: "Success"},
// 		"/api/endpoint2": {Method: "POST", Path: "/api/endpoint2", Result: "Not Covered"},
// 	}

// 	for _, endpoint := range expectedEndpoints {
// 		assert.Contains(t, allEndpoints, endpoint.Path)
// 		assert.Equal(t, expectedEndpoints[endpoint.Path].Method, allEndpoints[endpoint.Path].Method)
// 		assert.Equal(t, expectedEndpoints[endpoint.Path].Path, allEndpoints[endpoint.Path].Path)
// 		assert.Equal(t, expectedEndpoints[endpoint.Path].Result, allEndpoints[endpoint.Path].Result)
// 	}

// 	// Assert that all expectations were met
// 	mockAnalyzer.AssertExpectations(t)
// }
