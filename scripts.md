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

<br>




          
To test your Rage CLI locally without releasing, you can:

1. Build and run locally:
```bash
# From the project root
go build -o rage
```

2. Make it executable and available in your PATH temporarily:
```bash
# Make executable
chmod +x rage

# Add current directory to PATH temporarily
export PATH=$PATH:$(pwd)
```

3. Test different commands:
```bash
# Test help
rage --help

# Test module generation
rage generate module user --with-tests

# Test handler generation
rage generate handler payment

# Test infrastructure generation
rage generate docker
rage generate cicd --provider github
```

4. Check generated files:
```bash
# After generating a module
ls -R modules/user/

# After generating Docker files
ls Dockerfile docker-compose.yml .dockerignore

# After generating CI/CD
ls .github/workflows/ci.yml
```

5. Verify template rendering:
- Check if all variables are properly replaced
- Ensure file paths are correct
- Verify imports are working
- Check if the generated code compiles

6. Test the generated project:
```bash
# If you generated a new project
go mod tidy
go build ./...
go test ./...
```

If you encounter any issues:
- Check template paths in generator files
- Verify template syntax
- Check file permissions
- Look for any error messages in the output

Remember to clean up test files after verification:
```bash
# Remove test files
rm -rf modules/
rm -f Dockerfile docker-compose.yml .dockerignore
rm -rf .github/
```

This way, you can verify all functionality works correctly before creating a release.

<br>





          
I'll help you test your new changes locally with a different binary name. Here's how:

1. Build with a different name:
```bash
# From the project root
go build -o rage-s
```

2. Make it executable:
```bash
chmod +x rage-s
```

3. Move to a directory in your PATH (keeping your existing rage installation untouched):
```bash
sudo mv rage-s /usr/local/bin/
```

4. Verify the new command is available:
```bash
rage-s --help
```

Now you can:
- Use `rage` for your stable version from releases
- Use `rage-s` for testing your new changes

To test the new features:
```bash
# Test module generation
rage-s generate module user --with-tests

# Test infrastructure generation
rage-s generate docker
rage-s generate cicd --provider github
```

If you want to remove the test version later:
```bash
sudo rm /usr/local/bin/rage-s
```

This approach allows you to:
- Keep your stable version intact
- Test new changes safely
- Run both versions side by side
- Easily clean up the test version


<br>


# Build the binary
go build -o rage-s

# Create templates directory in the installation location
sudo mkdir -p /usr/local/bin/templates

# Copy binary and templates
sudo cp rage-s /usr/local/bin/
sudo cp -r templates/* /usr/local/bin/templates/

# Make executable
sudo chmod +x /usr/local/bin/rage-s