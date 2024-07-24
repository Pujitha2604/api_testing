# Define paths
$HOST_DIR = "/home/ec2-user/web-service"
$CONTAINER_DIR = "/app/files"
 
# Print contents of the host directory
Write-Output "Contents of the host directory:"
Get-ChildItem -Path $HOST_DIR | Format-List
 
# Run Docker container to list contents of the container directory
docker run --rm -v "${HOST_DIR}:${CONTAINER_DIR}" -v /var/run/docker.sock:/var/run/docker.sock -w "${CONTAINER_DIR}" test-tools /bin/bash -c "echo 'Contents of the container directory:' && ls -la ${CONTAINER_DIR}"
 
# Run Docker container to execute the apitests.sh script
docker run --rm -v "${HOST_DIR}:${CONTAINER_DIR}" -v /var/run/docker.sock:/var/run/docker.sock -w "${CONTAINER_DIR}" test-tools /bin/bash -c "chmod +x run.ps1 && echo 'Running run.ps1' && ./run.ps1"