HOST_DIR="/home/ec2-user/web-service"
CONTAINER_DIR="/app/files"
 
echo "Contents of the host directory:"
ls -la "$HOST_DIR"
 
docker run --rm -v "$HOST_DIR:$CONTAINER_DIR" -v /var/run/docker.sock:/var/run/docker.sock -w "$CONTAINER_DIR" test-tools /bin/bash -c "echo 'Contents of the container directory:' && ls -la $CONTAINER_DIR"
 
docker run --rm -v "$HOST_DIR:$CONTAINER_DIR" -v /var/run/docker.sock:/var/run/docker.sock -w "$CONTAINER_DIR" test-tools /bin/bash -c "chmod +x apitests.sh && echo 'Running apitests.sh:' && ./apitests.sh"