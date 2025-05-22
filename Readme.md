# Rage CLI

A powerful Go project scaffolder for modular clean architecture. Rage helps you bootstrap Go services and generate new modules following clean architecture principles.

## Features

- 🏗️ Project scaffolding with clean architecture structure
- 📦 Module generation with all clean architecture layers
- 🛠️ Handler, middleware, and router generation
- 🐳 Docker configuration generation
- 🔄 CI/CD pipeline generation (GitHub Actions & GitLab CI)
- 📚 API documentation generation

## Installation

### From GitHub Releases

1. Download the latest release for your platform from [GitHub Releases](https://github.com/klystro/rage/releases)
2. Extract the archive:
```bash
# For macOS/Linux
tar xzf rage_Darwin_arm64.tar.gz  # or rage_Darwin_x86_64.tar.gz for Intel Macs
xattr -d com.apple.quarantine rage # Required for macOS security

# For Windows
unzip rage_Windows_x86_64.zip
```

3. Move the binary to your PATH:
```bash
# For macOS/Linux
sudo mv rage /usr/local/bin/

# For Windows
# Move rage.exe to a directory in your PATH
```

4. Verify installation:
```bash
rage --help
```

## Usage

### Create a New Project
```bash
rage new myproject
```

### Generate a Module
```bash
rage generate module user --with-tests
```
This creates:
- Handler layer
- Repository layer
- Service layer
- UseCase layer
- Model definitions

### Generate Components
```bash
# Generate a handler
rage generate handler payment

# Generate middleware
rage generate middleware auth

# Generate router
rage generate router api
```

### Generate Infrastructure
```bash
# Generate Docker configuration
rage generate docker

# Generate CI/CD configuration
rage generate cicd --provider github

# Generate API documentation
rage generate docs
```

## Project Structure

```
.
├── cmd/                    # Command line applications
├── config/                 # Configuration files
├── internal/               # Private application code
│   ├── server/             # Server implementation
│   └── middleware/         # Middleware components
├── pkg/                    # Public library code
└── modules/                # Business modules
    └── user/               # Example module
        ├── handler/        # HTTP handlers
        ├── model/          # Domain models
        ├── repository/     # Data access layer
        ├── service/        # Business logic
        └── usecase/        # Use cases
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'feat: add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

For support, please open an issue in the GitHub repository.

## Acknowledgments

- Clean Architecture principles by Robert C. Martin
- Go project layout standards
- Cobra CLI framework

        