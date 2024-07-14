package analyze

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockFileSystem struct {
	mock.Mock
}

func (m *MockFileSystem) ReadFile(filename string) ([]byte, error) {
	args := m.Called(filename)
	data, _ := args.Get(0).([]byte) // Ensure we always return a valid byte slice or nil
	return data, args.Error(1)
}

func TestParseNewmanReport(t *testing.T) {
	mockFS := new(MockFileSystem)

	// Mock data for a successful read and parse
	reportData := NewmanReport{
		Run: Run{
			Executions: []Execution{
				{
					Item: Item{
						Name: "Test 1",
						Request: Request{
							Method: "GET",
							URL: URL{
								Path: []string{"employee", "1"},
							},
						},
					},
					Response: Response{
						Code: 200,
					},
				},
				{
					Item: Item{
						Name: "Test 2",
						Request: Request{
							Method: "POST",
							URL: URL{
								Path: []string{"employee"},
							},
						},
					},
					Response: Response{
						Code: 201,
					},
				},
			},
		},
	}

	reportJSON, _ := json.Marshal(reportData)
	mockFS.On("ReadFile", "/newman-report.json").Return(reportJSON, nil)

	// Replace the real osReadFile with our mock
	oldReadFile := osReadFile
	osReadFile = mockFS.ReadFile
	defer func() { osReadFile = oldReadFile }()

	// Call the function
	endpoints, err := parseNewmanReport("/newman-report.json")

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, 2, len(endpoints))
	assert.Equal(t, 200, endpoints["/employee/1"])
	assert.Equal(t, 201, endpoints["/employee"])

	// Test case for JSON unmarshal error
	mockFS.On("ReadFile", "/invalid/json").Return([]byte("invalid json"), nil)
	_, err = parseNewmanReport("/invalid/json")
	assert.Error(t, err)
	assert.Equal(t, "error parsing Newman report JSON: invalid character 'i' looking for beginning of value", err.Error())

	mockFS.AssertExpectations(t)
}
