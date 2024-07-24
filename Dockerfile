# Use a base image with necessary tools
FROM docker:20.10.24-dind

# Install Docker Compose and Docker client
RUN apk add --no-cache py3-pip \
    && pip3 install docker-compose \
    && apk add --no-cache docker-cli

# Set working directory
WORKDIR /app/files

# Copy scripts and other files
COPY . .

# Make sure your scripts are executable
RUN chmod +x api_tests.sh
