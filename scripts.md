`goreleaser release --snapshot --clean`

# Show help and available commands
rage --help

# Generate a new module
rage generate module <module-name> [--with-tests]

# Generate a handler
rage generate handler <handler-name>

# Generate middleware
rage generate middleware <middleware-name>

# Generate router
rage generate router <router-name>

# Generate API documentation
rage generate docs

# Generate Docker configuration
rage generate docker

# Generate CI/CD configuration
rage generate cicd [--provider github|gitlab]

# For macOS/Linux
tar xzf rage_Darwin_arm64.tar.gz  # Replace with your downloaded file

# For Windows
unzip rage_Windows_x86_64.zip     # Replace with your downloaded file

# For macOS/Linux
sudo mv rage /usr/local/bin/

# For Windows
# Move rage.exe to a directory in your PATH

rage --help