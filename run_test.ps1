# Define paths
$HOST_DIR = "C:\Users\Rekanto\Desktop\employee-service"
$API_TESTING_DIR = "C:\Users\Rekanto\Desktop\api_testing"
$CONTAINER_DIR = "/app/files"

# Print contents of the host directory
Write-Output "Contents of the host directory:"
Get-ChildItem -Path $HOST_DIR | Format-List

# Run Docker container to execute the run.ps1 script
docker run --rm -v "${HOST_DIR}:${CONTAINER_DIR}" -v /var/run/docker.sock:/var/run/docker.sock -w "${CONTAINER_DIR}" test-tools /bin/bash -c "chmod +x run.ps1 && echo 'Running run.ps1' && ./run.ps1"

# Run Docker container to build and execute the main.go file of the api_testing service
docker run --rm -v "${HOST_DIR}:${CONTAINER_DIR}/employee-service" -v "${API_TESTING_DIR}:${CONTAINER_DIR}/api_testing" -w "${CONTAINER_DIR}/api_testing" golang:1.22 /bin/bash -c "go run main.go"
