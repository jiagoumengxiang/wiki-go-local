# Codebase Structure

**Analysis Date:** 2026-04-05

## Directory Layout

```
wiki-go-local/
├── main.go                     # Application entry point
├── go.mod                      # Go module definition
├── go.sum                      # Go dependency checksums
├── Makefile                    # Build commands
├── Dockerfile                  # Container build config
├── docker-compose-*.yml        # Docker Compose configurations
├── internal/                   # Application code
│   ├── auth/                   # Authentication and session management
│   ├── ban/                    # Login attempt banning
│   ├── comments/               # Comments system
│   ├── config/                 # Configuration management
│   ├── crypto/                 # Password hashing
│   ├── frontmatter/            # YAML frontmatter parsing
│   ├── goldext/                # Markdown preprocessing extensions
│   ├── handlers/               # HTTP request handlers
│   ├── i18n/                   # Internationalization
│   ├── migration/              # Configuration migration
│   ├── resources/              # Embedded static resources
│   ├── roles/                  # Role constants
│   ├── routes/                 # HTTP routing
│   ├── static/                 # Static asset management
│   ├── types/                  # Shared data structures
│   ├── utils/                  # Utility functions
│   └── version/                # Version tracking
├── .planning/                 # Planning documents
├── build/                     # Build artifacts
├── data/                      # Runtime data (created at runtime)
│   ├── config.yaml             # Configuration file
│   ├── documents/              # Wiki documents
│   ├── comments/               # Document comments
│   ├── temp/                  # Session store and bans
│   └── static/                # Custom static assets
└── vendor/                    # Vendored dependencies
```

## Directory Purposes

**main.go:**
- Purpose: Application entry point and initialization
- Contains: Server startup, config loading, migration, handler initialization
- Key files: `main.go`

**internal/auth:**
- Purpose: User authentication, session management, access control
- Contains: Session storage, cookie management, credential validation, access rule checking
- Key files: `auth.go`, `access.go`, `session_store.go`

**internal/ban:**
- Purpose: Prevent brute-force login attacks
- Contains: IP-based failure tracking with exponential backoff
- Key files: `ban.go`

**internal/comments:**
- Purpose: Document commenting system
- Contains: Comment creation, retrieval, deletion, markdown rendering
- Key files: `comments.go`

**internal/config:**
- Purpose: Configuration loading and management
- Contains: Config struct, YAML parsing, defaults, file type validation
- Key files: `config.go`, `filetypes.go`

**internal/crypto:**
- Purpose: Cryptographic operations
- Contains: Password hashing with bcrypt
- Key files: `password.go`

**internal/frontmatter:**
- Purpose: Parse and manipulate YAML frontmatter in markdown files
- Contains: Frontmatter extraction, parsing, adding
- Key files: `frontmatter.go`, `kanban.go`, `links.go`

**internal/goldext:**
- Purpose: Extend markdown rendering with custom preprocessors
- Contains: Preprocessors for mermaid, emojis, videos, links, highlighting, etc.
- Key files: `load.go`, `link.go`, `mermaid.go`, `emoji.go`, `direction.go`, `details.go`, `infoboxes.go`, `tasklist.go`, `toc.go`, `headinganchor.go`, `highlight.go`, `typography.go`, `superscript.go`, `subscript.go`, `scriptsanitize.go`, `shortcodes.go`

**internal/handlers:**
- Purpose: HTTP request handlers for pages and API endpoints
- Contains: Page rendering, authentication, file management, search, comments, backup, import
- Key files: `handlers.go`, `page.go`, `home.go`, `auth.go`, `files.go`, `editor.go`, `search.go`, `comments.go`, `backup.go`, `import.go`, `settings.go`, `users.go`, `access_rules.go`, `template.go`

**internal/i18n:**
- Purpose: Multi-language support for UI strings
- Contains: Translation manager, language file loading, placeholder processing
- Key files: `manager.go`, `embed.go`

**internal/migration:**
- Purpose: Handle configuration schema upgrades
- Contains: User role migration from boolean to role-based system
- Key files: `role_migration.go`

**internal/resources:**
- Purpose: Embedded static files (templates, CSS, JS, libraries, language files)
- Contains: `embed` directives, filesystem accessors
- Key files: `resources.go`
- Subdirectories:
  - `templates/`: HTML templates
  - `static/`: CSS, JS, libraries
  - `langs/`: Language JSON files

**internal/roles:**
- Purpose: Role constant definitions
- Contains: Role string constants (admin, editor, viewer)
- Key files: `roles.go`

**internal/routes:**
- Purpose: HTTP routing and middleware
- Contains: Route definitions, CSP middleware, cache control, static file serving
- Key files: `routes.go`

**internal/static:**
- Purpose: Static asset management
- Contains: Static file serving, embedded resources, custom asset support
- Key files: `static.go`

**internal/types:**
- Purpose: Shared data structures
- Contains: PageData, NavItem, BreadcrumbItem
- Key files: `types.go`

**internal/utils:**
- Purpose: Utility functions
- Contains: Path sanitization, navigation building, markdown rendering, formatting
- Key files: `utils.go`, `nav.go` (implied)

**internal/version:**
- Purpose: Version tracking
- Contains: Version string set at build time
- Key files: `version.go`

**data:**
- Purpose: Runtime data directory (created at runtime)
- Contains: Configuration, documents, comments, sessions, custom static assets
- Subdirectories:
  - `config.yaml`: Configuration file
  - `documents/`: Wiki documents (one directory per page with `document.md`)
  - `comments/`: Document comments (mirrors document structure)
  - `temp/`: Session store (`sessions.json`) and login bans (`login-bans.json`)
  - `static/`: Custom static assets (overrides embedded resources)

## Key File Locations

**Entry Points:**
- `main.go`: Application entry point - initialization, config loading, server startup

**Configuration:**
- `internal/config/config.go`: Config struct and YAML parsing
- `internal/config/filetypes.go`: Allowed file extensions and MIME types
- `data/config.yaml`: Runtime configuration file

**Core Logic:**
- `internal/handlers/page.go`: Page rendering logic
- `internal/handlers/home.go`: Homepage handling
- `internal/handlers/editor.go`: Document editing
- `internal/handlers/auth.go`: Authentication endpoints
- `internal/handlers/files.go`: File upload and management
- `internal/handlers/search.go`: Search functionality

**Routing:**
- `internal/routes/routes.go`: Route definitions and middleware

**Authentication:**
- `internal/auth/auth.go`: Session management and access control
- `internal/ban/ban.go`: Login attempt banning
- `internal/roles/roles.go`: Role constants

**Markdown Processing:**
- `internal/goldext/load.go`: Preprocessor registration and ordering
- `internal/goldext/*.go`: Individual preprocessors
- `internal/frontmatter/frontmatter.go`: Frontmatter parsing

**Data Storage:**
- `data/documents/<path>/document.md`: Wiki content files
- `data/comments/<docpath>/<timestamp>_<username>.md`: Comment files
- `data/temp/sessions.json`: Session persistence
- `data/temp/login-bans.json`: Login ban tracking

**Testing:**
- `internal/goldext/superscript_test.go`: Superscript preprocessor tests

## Naming Conventions

**Files:**
- Lowercase with underscores: `auth.go`, `session_store.go`, `role_migration.go`
- Test files: `<name>_test.go` (e.g., `superscript_test.go`)

**Directories:**
- Lowercase: `auth`, `handlers`, `resources`, `goldext`

**Go Packages:**
- Match directory name: `package auth`, `package handlers`, `package config`

**Functions:**
- CamelCase: `GetSession`, `CreateDocumentHandler`, `LoadConfig`
- Private functions start with lowercase: `isValidCommentID`, `sanitizeUsername`

**Types/Structs:**
- PascalCase: `Session`, `PageData`, `NavItem`, `Config`

**Constants:**
- PascalCase or uppercase with underscores: `RoleAdmin`, `defaultHomepageContent`

**Interfaces:**
- PascalCase with suffix (if any): Not heavily used in this codebase

**Error Variables:**
- PascalCase: `ErrNotFound`, `ErrInvalidInput` (not heavily used)

**HTTP Handlers:**
- Suffix with "Handler": `LoginHandler`, `PageHandler`, `HomeHandler`

**Template Files:**
- Lowercase with hyphens: `base.html`, `login.html`, `sidebar.html`

**Markdown Documents:**
- Lowercase with hyphens/spaces in directories: `my-page/document.md`
- Filename always `document.md`

**Comments:**
- Format: `YYYYMMDDhhmmss_username.md` (e.g., `20260105123456_admin.md`)

## Where to Add New Code

**New Feature:**
- Primary code: `internal/handlers/<feature>.go` (new handler file or add to existing)
- Tests: `<feature>_test.go` in same directory as implementation
- Routes: Add route definition in `internal/routes/routes.go`
- Types: Add to `internal/types/types.go` if shared across handlers

**New Component/Module:**
- Implementation: `internal/<module>/` (create new directory)
- Package: `package <module>`

**New Markdown Extension:**
- Implementation: `internal/goldext/<extension>.go` (new preprocessor file)
- Registration: Add to `internal/goldext/load.go` init function

**New API Endpoint:**
- Handler: Add function to `internal/handlers/` (appropriate existing file or new)
- Routes: Add `mux.HandleFunc()` call in `internal/routes/routes.go`
- Middleware: Add middleware wrapper if auth required

**New Template:**
- HTML: Add file to `internal/resources/templates/`
- Usage: Reference in handler code via template rendering

**New Static Asset:**
- Embedded: Add to `internal/resources/static/` with `//go:embed` directive
- Custom: Document that users can add to `data/static/`

**New Utility Function:**
- Implementation: Add to `internal/utils/utils.go` or new file in `internal/utils/`
- Tests: Add `<name>_test.go` in `internal/utils/`

**New Configuration Field:**
- Struct: Add field to `Config` struct in `internal/config/config.go`
- Template: Update `GetConfigTemplate()` in `internal/config/config.go`
- Migration: Consider adding migration in `internal/migration/` if breaking change

**New Language Translation:**
- File: Add JSON file to `internal/resources/langs/<lang>.json`
- Copying: Handled automatically by `i18n.CopyLangsToStaticDir()`

## Special Directories

**internal/resources:**
- Purpose: Contains embedded static files using Go's `embed` directive
- Generated: No (source files)
- Committed: Yes
- Structure:
  - `templates/`: HTML templates (base.html, login.html, sidebar.html, etc.)
  - `static/`: CSS, JS, libraries (CodeMirror, FontAwesome, MathJax, Mermaid, Prism)
  - `langs/`: Translation JSON files

**data:**
- Purpose: Runtime data directory created at first run
- Generated: Yes (created on first startup)
- Committed: No (in .gitignore)
- Contains user content and configuration

**vendor:**
- Purpose: Vendored Go dependencies
- Generated: Yes (via `go mod vendor`)
- Committed: Yes (optional)

**build:**
- Purpose: Build artifacts (compiled binaries)
- Generated: Yes
- Committed: No

**.planning:**
- Purpose: Planning documents (this file and related)
- Generated: Yes
- Committed: Yes (for development workflow)

---

*Structure analysis: 2026-04-05*
