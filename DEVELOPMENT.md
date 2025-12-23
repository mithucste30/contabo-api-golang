# Development Guide

## Local Development (Before Publishing)

To use this SDK locally before publishing to GitHub:

### In your project's go.mod:

```go
module your-project

go 1.21

require github.com/mithucste30/contabo-api-golang v0.0.0

replace github.com/mithucste30/contabo-api-golang => /Users/kahf/works/mithucste30/contabo-sdk-golang
```

### Or use the command:

```bash
cd /path/to/your/project
go mod edit -replace github.com/mithucste30/contabo-api-golang=/Users/kahf/works/mithucste30/contabo-sdk-golang
go mod tidy
```

## Publishing to GitHub

Once you're ready to publish:

### 1. Initialize Git Repository

```bash
cd /Users/kahf/works/mithucste30/contabo-sdk-golang
git init
git add .
git commit -m "Initial commit: Contabo Go SDK"
```

### 2. Push to GitHub

Your repository is already set up at:
- https://github.com/mithucste30/contabo-api-golang

Just commit and push your changes:

```bash
git add .
git commit -m "Update module path and fix imports"
git push origin main
```

### 4. Create a Release (Optional but Recommended)

```bash
# Tag the first version
git tag v0.1.0
git push origin v0.1.0
```

### 5. Use in Other Projects

Once published, users can install with:

```bash
go get github.com/mithucste30/contabo-api-golang@latest
```

Or for a specific version:

```bash
go get github.com/mithucste30/contabo-api-golang@v0.1.0
```

## Testing

### Run Examples Locally

```bash
export CONTABO_CLIENT_ID="your-client-id"
export CONTABO_CLIENT_SECRET="your-client-secret"
export CONTABO_API_USER="your-email"
export CONTABO_API_PASSWORD="your-password"

cd examples
go run main.go
```

### Run Tests (when added)

```bash
go test ./...
```

## Versioning

This project follows [Semantic Versioning](https://semver.org/):
- v0.x.x - Initial development
- v1.x.x - First stable release
- MAJOR.MINOR.PATCH

### Version Bump Examples:

```bash
# Bug fixes
git tag v0.1.1
git push origin v0.1.1

# New features (backward compatible)
git tag v0.2.0
git push origin v0.2.0

# Breaking changes
git tag v1.0.0
git push origin v1.0.0
```
