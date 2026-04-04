# External Integrations

**Analysis Date:** 2026-04-05

## APIs & External Services

**Media Embedding:**
- YouTube - Video embed support (client-side, no API calls required)
  - Purpose: Allow embedding YouTube videos in wiki pages
  - Implementation: Custom markdown preprocessor (`internal/goldext/youtube.go`)
  - CSP policy: `frame-src 'self' https://*.youtube.com`
- Vimeo - Video embed support (client-side, no API calls required)
  - Purpose: Allow embedding Vimeo videos in wiki pages
  - Implementation: Custom markdown preprocessor (`internal/goldext/vimeo.go`)
  - CSP policy: `frame-src 'self' https://*.vimeo.com`
- External link metadata fetching
  - Purpose: Fetch page titles/descriptions when users add links to link collection
  - Implementation: HTTP client requests in `internal/handlers/links_api.go`

## Data Storage

**Databases:**
- None - Databaseless architecture

**File Storage:**
- Local filesystem only
  - Documents: `$WIKI_ROOT_DIR/documents/`
  - Attachments: `$WIKI_ROOT_DIR/files/`
  - Static assets: `$WIKI_ROOT_DIR/static/`
  - Sessions: `$WIKI_ROOT_DIR/temp/sessions.json`
  - Backups: `$WIKI_ROOT_DIR/temp/backups/`

**Caching:**
- None - No caching layer (file-based storage only)

## Authentication & Identity

**Auth Provider:**
- Custom implementation
  - Implementation: Password-based authentication with SHA256 hashing
  - Session management: JSON file-based session storage
  - Password hashing: `internal/crypto/password.go` using SHA256 with salt
  - Session tokens: Random 32-byte tokens hashed with SHA256
  - Cookie-based authentication with `SameSite=Lax` mode
  - Role-based access control: admin, editor, viewer
  - Group-based access control for path-level permissions
  - Config file: `data/config.yaml` stores users, roles, groups

## Monitoring & Observability

**Error Tracking:**
- None - No external error tracking service

**Logs:**
- Standard library logging (log package)
- Console output only
- No structured logging service

## CI/CD & Deployment

**Hosting:**
- Self-hosted (any server running the binary or Docker container)
- Deployment options:
  - Prebuilt binaries for Linux, Windows, macOS
  - Docker images with multi-stage builds
  - Docker Compose configurations for HTTP and SSL setups

**CI Pipeline:**
- GitHub Actions
  - Workflows:
    - `.github/workflows/release-docker.yml` - Docker image builds
    - `.github/workflows/release-docker-dev.yml` - Dev Docker builds
    - `.github/workflows/release-binaries.yml` - Binary builds for multiple platforms
    - `.github/workflows/release-binaries-dev.yml` - Dev binary builds
  - Automated releases on version tags

## Environment Configuration

**Required env vars:**
- None required (all configuration via YAML file)

**Optional env vars:**
- `CONFIGFILE` - Override default config file path (default: `data/config.yaml`)

**Secrets location:**
- `data/config.yaml` - Contains user passwords (SHA256 hashed), SSL cert paths
- No external secrets manager

## Webhooks & Callbacks

**Incoming:**
- None

**Outgoing:**
- None (link metadata fetching uses HTTP client for one-time requests on user action)

## Third-Party Content

**External Resources:**
- None loaded by default
- Optional user-configured media embeds:
  - YouTube videos (iframe embeds)
  - Vimeo videos (iframe embeds)
- CSP allows loading from:
  - `https://*.ytimg.com` (YouTube thumbnails)
  - `https://*.vimeocdn.com` (Vimeo thumbnails)
  - `data:` URLs (embedded images)

---

*Integration audit: 2026-04-05*
