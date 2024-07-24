# Use a base image with Docker installed
FROM docker:20.10.24-dind

# Install necessary packages including bash, PowerShell, and icu-libs
RUN apk add --no-cache \
    py3-pip \
    docker-cli \
    bash \
    ca-certificates \
    less \
    ncurses-terminfo-base \
    tzdata \
    krb5-libs \
    libgcc \
    libintl \
    libssl1.1 \
    libstdc++ \
    lttng-ust \
    userspace-rcu \
    zlib \
    curl \
    icu-libs && \
    mkdir -p /opt/microsoft/powershell/7 && \
    curl -sSL https://github.com/PowerShell/PowerShell/releases/download/v7.3.6/powershell-7.3.6-linux-alpine-x64.tar.gz | tar zxf - -C /opt/microsoft/powershell/7 && \
    ln -s /opt/microsoft/powershell/7/pwsh /usr/bin/pwsh

# Set working directory
WORKDIR /app/files

# Copy scripts and other files
COPY . .
# Make sure your scripts are executable
RUN  chmod +x run_test.ps1

# Set entrypoint to use PowerShell
ENTRYPOINT ["pwsh", "-File", "/app/files/run.ps1"]
