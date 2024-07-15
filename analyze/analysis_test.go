package analyze

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockFileRetriever struct {
	mock.Mock
}

func (m *MockFileRetriever) GetAllGoFiles(rootDir string) ([]string, error) {
	args := m.Called(rootDir)
	return args.Get(0).([]string), args.Error(1)
}

type MockFileAnalyzer struct {
	mock.Mock
}

func (m *MockFileAnalyzer) AnalyzeFileForAPIEndpoints(file string) map[string]Endpoint {
	args := m.Called(file)
	return args.Get(0).(map[string]Endpoint)
}

type MockNewmanParser struct {
	mock.Mock
}

func (m *MockNewmanParser) ParseNewmanReport(newmanReportPath string) (map[string]int, error) {
	args := m.Called(newmanReportPath)
	return args.Get(0).(map[string]int), args.Error(1)
}

type MockPrinter struct {
	mock.Mock
}

func (m *MockPrinter) PrintEndpointsTable(allEndpoints map[string]Endpoint, newmanReportPath string) {
	m.Called(allEndpoints, newmanReportPath)
}

func setUpMocks() (*MockFileRetriever, *MockFileAnalyzer, *MockNewmanParser, *MockPrinter) {
	fileRetriever := new(MockFileRetriever)
	fileAnalyzer := new(MockFileAnalyzer)
	newmanParser := new(MockNewmanParser)
	printer := new(MockPrinter)

	getAllGoFilesFunc = fileRetriever.GetAllGoFiles
	analyzeFileForAPIEndpointsFunc = fileAnalyzer.AnalyzeFileForAPIEndpoints
	parseNewmanReportFunc = newmanParser.ParseNewmanReport
	printEndpointsTableFunc = printer.PrintEndpointsTable

	return fileRetriever, fileAnalyzer, newmanParser, printer
}

func TestAnalysis(t *testing.T) {
	fileRetriever, fileAnalyzer, newmanParser, printer := setUpMocks()

	rootDir := "testDir"
	newmanReportPath := "testNewmanReport.json"

	fileRetriever.On("GetAllGoFiles", rootDir).Return([]string{"file1.go", "file2.go"}, nil)
	fileAnalyzer.On("AnalyzeFileForAPIEndpoints", "file1.go").Return(map[string]Endpoint{
		"/endpoint1": {Method: "GET", Path: "/endpoint1"},
	})
	fileAnalyzer.On("AnalyzeFileForAPIEndpoints", "file2.go").Return(map[string]Endpoint{
		"/endpoint2": {Method: "POST", Path: "/endpoint2"},
	})
	newmanParser.On("ParseNewmanReport", newmanReportPath).Return(map[string]int{
		"/endpoint1": 200,
		"/endpoint2": 500,
	}, nil)
	printer.On("PrintEndpointsTable", mock.Anything, newmanReportPath).Return()

	Analysis(rootDir, newmanReportPath)
	
	fileRetriever.AssertExpectations(t)
	fileAnalyzer.AssertExpectations(t)
	newmanParser.AssertExpectations(t)
	printer.AssertExpectations(t)
}

