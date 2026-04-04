# Architecture

**Analysis Date:** 2026-04-05

## Pattern Overview

**Overall:** Layered HTTP Server with Flat-File Storage

**Key Characteristics:**
- Standard library `net/http` with `http.ServeMux` routing
- Handler-based request processing with middleware support
- No external database - all data stored as flat files
- Embedded static resources with runtime override capability
- Role-based authentication with session management
- Markdown preprocessing pipeline with custom extensions

## Layers

**HTTP Routing Layer:**
- Purpose: Route HTTP requests to handlers and apply cross-cutting middleware
- Location: `internal/routes/routes.go`
- Contains: Route definitions, middleware (CSP, cache control, security headers), static file serving
- Depends on: `internal/handlers`, `internal/auth`, `internal/config`, `internal/resources`
- Used by: `main.go` (via `routes.SetupRoutes()`)

**Request Handler Layer:**
- Purpose: Process HTTP requests and generate responses
- Location: `internal/handlers/*.go`
- Contains: Page handlers, API endpoints, authentication handlers, file management, search, comments
- Depends on: `internal/config`, `internal/auth`, `internal/utils`, `internal/types`, `internal/comments`, `internal/frontmatter`, `internal/i18n`
- Used by: `internal/routes`

**Authentication & Authorization Layer:**
- Purpose: Manage user sessions, authenticate requests, enforce access control
- Location: `internal/auth/*.go`, `internal/ban/*.go`, `internal/roles/*.go`
- Contains: Session management, password hashing, login attempt banning, role checking, access rules
- Depends on: `internal/config`, `internal/crypto`
- Used by: `internal/routes`, `internal/handlers`

**Configuration Layer:**
- Purpose: Load and manage application configuration
- Location: `internal/config/*.go`
- Contains: Config struct definition, YAML parsing, default values, file type validation
- Depends on: `gopkg.in/yaml.v3`, `internal/crypto`, `internal/roles`
- Used by: `main.go`, all internal packages

**Markdown Processing Layer:**
- Purpose: Preprocess and render markdown content
- Location: `internal/goldext/*.go`, `internal/frontmatter/*.go`
- Contains: Preprocessors for mermaid diagrams, emojis, videos, links, highlighting, frontmatter parsing
- Depends on: `github.com/yuin/goldmark`, `gopkg.in/yaml.v3`
- Used by: `internal/utils` (via `RenderMarkdown`)

**Data Storage Layer:**
- Purpose: Persist documents, attachments, comments, sessions, and config
- Location: File system (flat files in `data/` directory)
- Contains: Markdown documents in `data/documents/`, comments in `data/comments/`, sessions in `data/temp/`, config in `data/config.yaml`
- Depends on: Standard library `os`, `io`, `path/filepath`
- Used by: `internal/handlers`, `internal/comments`, `internal/auth`, `internal/config`

**Template Rendering Layer:**
- Purpose: Generate HTML responses from templates and data
- Location: `internal/handlers/template.go`, `internal/resources/templates/`
- Contains: Template functions, template caching, HTML templates
- Depends on: `internal/resources`, `internal/types`, `internal/utils`, `internal/i18n`, `internal/version`
- Used by: `internal/handlers`

**Internationalization Layer:**
- Purpose: Provide multi-language support for UI strings
- Location: `internal/i18n/*.go`, `internal/resources/langs/`
- Contains: Translation manager, placeholder processing, language file loading
- Depends on: `internal/config`, `internal/resources`
- Used by: `internal/handlers` (via template `t()` function)

**Static Asset Layer:**
- Purpose: Serve CSS, JS, fonts, images, and other static resources
- Location: `internal/static/*.go`, `internal/resources/static/`
- Contains: Static file serving, embedded resources, custom asset support
- Depends on: `internal/resources`
- Used by: `internal/routes`, `internal/handlers`

**Migration Layer:**
- Purpose: Handle configuration schema upgrades
- Location: `internal/migration/*.go`
- Contains: Role migration from boolean to role-based system
- Depends on: `gopkg.in/yaml.v3`, `internal/roles`
- Used by: `main.go` (during startup)

## Data Flow

**Page Request Flow:**

1. HTTP request received by `http.ServeMux` in `routes.SetupRoutes()`
2. CSP middleware adds security headers
3. Route matches request pattern (e.g., `/`, `/api/*`, `/static/*`)
4. Handler function invoked with request and response writer
5. Handler checks authentication via `auth.GetSession()` and `auth.CanAccessDocument()`
6. Handler builds navigation tree via `utils.BuildNavigation()` and `utils.FilterNavigation()`
7. Handler reads document content from filesystem (`data/documents/<path>/document.md`)
8. Markdown content processed through goldmark with custom preprocessors (`utils.RenderMarkdown()`)
9. Handler loads template with data (`types.PageData`)
10. Template rendered to HTML with i18n translations
11. Response sent with appropriate headers

**Login Request Flow:**

1. POST to `/api/login` with username/password
2. `handlers.LoginHandler` checks login ban status via `ban.BanList.IsBanned()`
3. If not banned, `auth.ValidateCredentials()` checks password hash
4. If valid, `auth.CreateSession()` creates session and sets cookies
5. `ban.BanList.Clear()` clears failure tracking
6. JSON response with success status
7. If invalid, `ban.BanList.RegisterFailure()` tracks failure
8. After threshold, IP is banned with exponential backoff

**Document Save Flow:**

1. POST to `/api/save/<path>` with markdown content
2. Handler checks authentication and role via `auth.RequireRole()`
3. Handler reads existing document for version history
4. Old version copied to `versions/<timestamp>.md` if under max_versions limit
5. New content written to `data/documents/<path>/document.md`
6. JSON response with success status

**Comment Addition Flow:**

1. POST to `/api/comments/add/<docpath>` with markdown content
2. Handler checks authentication via `auth.GetSession()`
3. `comments.AddComment()` creates timestamped filename: `YYYYMMDDhhmmss_username.md`
4. Comment file written to `data/comments/<docpath>/<timestamp>_<username>.md`
5. JSON response with success status

**State Management:**
- Sessions: In-memory map persisted to JSON file (`data/temp/sessions.json`)
- Login bans: In-memory map persisted to JSON file (`data/temp/login-bans.json`)
- Configuration: Single YAML file (`data/config.yaml`)
- No database connection pooling or stateful servers

## Key Abstractions

**NavItem:**
- Purpose: Represents a node in the navigation tree (directory or document)
- Examples: `internal/types/types.go`, `internal/utils/nav.go` (implied)
- Pattern: Recursive tree structure with `Children []*NavItem`, used for building sidebar navigation

**Session:**
- Purpose: Represents an authenticated user session
- Examples: `internal/auth/auth.go`
- Pattern: Token-based with cookie storage, role and group metadata, expiration tracking

**PageData:**
- Purpose: Container for all data needed to render a page template
- Examples: `internal/types/types.go`
- Pattern: Single struct passed to template, includes navigation, content, breadcrumbs, config, user info

**Config:**
- Purpose: Centralized configuration for all application settings
- Examples: `internal/config/config.go`
- Pattern: Nested structs with YAML tags, defaults in code, persisted to file

**Preprocessor:**
- Purpose: Transform markdown content before rendering
- Examples: `internal/goldext/*.go` (mermaid, emoji, links, etc.)
- Pattern: Function registered in init(), processes AST or text in specific order

## Entry Points

**main.go:**
- Location: `main.go`
- Triggers: Application startup
- Responsibilities:
  - Parse command-line flags (`-configfile`)
  - Fix broken config files if needed
  - Migrate user roles from old to new schema
  - Load configuration
  - Initialize session store
  - Ensure homepage exists
  - Ensure static assets exist
  - Initialize handlers
  - Setup routes
  - Start HTTP server (HTTP or HTTPS based on config)

**goldext Package Initialization:**
- Location: `internal/goldext/load.go`
- Triggers: Import statement in `main.go` (blank import for side effects)
- Responsibilities: Register markdown preprocessors in correct order on package import

## Error Handling

**Strategy:** Explicit error checking with early returns

**Patterns:**
- Handler functions return HTTP errors via `http.Error()` with appropriate status codes
- File operations checked with `if err != nil` pattern
- Configuration errors logged with `log.Fatal()` for startup issues
- User-facing errors returned as JSON responses from API handlers
- Template errors return HTTP 500 status
- Authentication failures return 401 or 403 status codes

**Cross-Cutting Concerns**

**Logging:** Standard library `log` package, no structured logging

**Validation:** Path sanitization in `utils.SanitizePath()`, filename validation in `utils.IsValidFilename()`, comment ID validation

**Authentication:** Session-based with HTTP-only cookies, SHA256 token hashing, bcrypt password hashing

**Caching:**
- Template: In-memory singleton cache with `sync.Once`
- Static files: HTTP cache headers based on file type and versioning
- No response caching for dynamic content (no-cache headers set)

---

*Architecture analysis: 2026-04-05*
