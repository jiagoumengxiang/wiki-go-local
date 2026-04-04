# Technology Stack

**Analysis Date:** 2026-04-05

## Languages

**Primary:**
- Go 1.26.0 - Core application logic, HTTP server, handlers, authentication, markdown processing, and all business logic

**Secondary:**
- YAML - Configuration files (config.yaml)
- JSON - Session storage, API responses, manifest files

## Runtime

**Environment:**
- Go 1.26.0

**Package Manager:**
- Go Modules (go.mod, go.sum)
- Vendor directory present for vendored dependencies

## Frameworks

**Core:**
- net/http (standard library) - HTTP server and routing
- github.com/yuin/goldmark v1.7.16 - Markdown parsing and rendering
- gopkg.in/yaml.v3 v3.0.1 - YAML configuration parsing

**Testing:**
- None detected (standard library testing available but no test framework configured)

**Build/Dev:**
- Make - Multi-platform binary builds (linux_amd64, linux_arm64, windows_amd64, macos_arm64, etc.)
- Docker - Containerization with multi-stage builds
- GitHub Actions - CI/CD for automated builds

## Key Dependencies

**Critical:**
- github.com/yuin/goldmark v1.7.16 - Markdown to HTML conversion with custom extensions for wiki features
- gopkg.in/yaml.v3 v3.0.1 - Configuration file parsing (users, access rules, wiki settings)

**Infrastructure:**
- golang.org/x/crypto v0.48.0 - Cryptographic functions (SHA256 hashing for passwords/session tokens)
- github.com/gosimple/slug v1.15.0 - URL slug generation for documents
- golang.org/x/text v0.34.0 - Unicode text handling (indirect dependency)

**Utilities:**
- github.com/gosimple/unidecode v1.0.1 - Unicode to ASCII conversion (indirect dependency)

## Configuration

**Environment:**
- Configuration via YAML file (`data/config.yaml`)
- Environment variables: `CONFIGFILE` (optional override for config file path)
- No `.env` files used - all config in YAML

**Build:**
- Makefile with targets for multiple platforms
- `GOFLAGS=-mod=vendor` for vendored builds
- Build flags: `-ldflags="-X 'wiki-go/internal/version.Version=${VERSION}' -s -w -extldflags=-static -linkmode 'external'"`
- Build tags: `-tags netgo,usergo -trimpath -gcflags=all="-l -B -C"`

## Platform Requirements

**Development:**
- Go 1.26.0+
- Make (for multi-platform builds)
- Git (for versioning and build metadata)

**Production:**
- Linux, Windows, or macOS (multi-platform binaries available)
- Docker (optional, but supported)
- File system permissions for data directory
- SSL/TLS certificates (optional, for HTTPS)

---

*Stack analysis: 2026-04-05*
