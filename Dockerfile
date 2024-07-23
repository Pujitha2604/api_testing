FROM golang:1.22
 
RUN apt-get update && \
    apt-get install -y \
    docker.io \
    docker-compose \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*
 
RUN apt-get install -y \
    bash
 
WORKDIR /app
 
COPY . .
 
RUN go build -o myapp
 
ENTRYPOINT ["/app/myapp"]