# Coding Conventions

**Analysis Date:** 2026-04-05

## Naming Patterns

**Files:**
- Lowercase with underscores: `home.go`, `auth.go`, `password.go`, `frontmatter.go`
- Test files: `{source}_test.go` (e.g., `superscript_test.go`)
- Package directories: lowercase, single word (e.g., `handlers`, `auth`, `config`)

**Functions:**
- Exported functions: PascalCase (e.g., `LoadConfig`, `HashPassword`, `SanitizePath`, `GetSession`)
- Private functions: camelCase (e.g., `hashToken`, `getTemplate`, `clientIP`, `requireRole`)
- Methods: Follow same rules (exported = PascalCase, private = camelCase)

**Variables:**
- Local variables: camelCase (e.g., `username`, `password`, `filePath`, `configData`)
- Package-level exported: PascalCase (e.g., `ConfigFilePath`, `RoleAdmin`)
- Package-level private: lowercase (e.g., `sessions`, `cfg`, `loginBan`)

**Types/Structs:**
- Exported types: PascalCase (e.g., `User`, `Config`, `Session`, `Comment`, `Metadata`, `NavItem`, `PageData`)
- Fields: PascalCase (e.g., `Username`, `Password`, `Role`, `CreatedAt`, `LastModified`)

**Constants:**
- Exported constants: PascalCase (e.g., `RoleAdmin`, `RoleEditor`, `RoleViewer`, `defaultHomepageContent`)

## Code Style

**Formatting:**
- Standard `go fmt` formatting (no explicit config file)
- No external linting configuration (no `.golangci.yml`)
- Uses standard Go conventions

**Linting:**
- No explicit linting tool configured
- Relies on `go vet` and compiler checks

**Build:**
- Makefile-based build system (`Makefile`)
- Cross-platform builds via Makefile targets
- LDFLAGS for version injection and binary optimization

## Import Organization

**Order:**
1. Standard library imports (e.g., `fmt`, `log`, `net/http`, `os`, `strings`)
2. Third-party imports (e.g., `gopkg.in/yaml.v3`, `github.com/yuin/goldmark`)
3. Internal package imports (prefixed with module name `wiki-go/internal/...`)

**Blank lines:** Blank lines between each import group

**Path aliases:** No path aliases used - full package paths always used

## Error Handling

**Patterns:**
- Error wrapping with `%w`: `fmt.Errorf("failed to create directory %s: %w", dir, err)`
- Error returns: Multiple return values with error as last parameter
- Early returns on error: Check error immediately and return

**Logging errors:**
- `log.Printf("Error building navigation: %v", err)` - for non-fatal errors
- `log.Fatal("Error loading config:", err)` - for fatal errors
- Warning logs: `log.Printf("Warning: Failed to initialize session store: %v", err)`

**HTTP error responses:**
- `http.Error(w, "Error message", http.StatusInternalServerError)`
- Set status code explicitly: `w.WriteHeader(http.StatusNotFound)`
- JSON error responses with map: `json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "message": "..."})`

## Logging

**Framework:** Standard `log` package

**Patterns:**
- Use `log.Printf` for formatted logging
- Prefix error context in message: "Error building navigation: %v"
- Use "Warning:" prefix for non-critical issues
- Use `log.Fatal` for startup errors

**No structured logging:** Plain text logging only

## Comments

**When to Comment:**
- Above exported functions and types (doc comments)
- For complex logic sections
- To explain security considerations (e.g., in config comments)

**Function comments:** Present but not consistently comprehensive
- Single-line comments for simple functions
- Multi-line comments for complex functions
- Package-level comments: Not consistently present

**JSDoc/TSDoc:** N/A (Go uses different conventions)

**Example:**
```go
// HashPassword creates a bcrypt hash of password
func HashPassword(password string, passwordstrength int) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), passwordstrength)
    if err != nil {
        return "", err
    }
    return string(bytes), nil
}
```

## Function Design

**Size:** No explicit size guidelines observed
- Functions range from small (~20 lines) to large (~1400 lines)
- Handler functions tend to be larger (complex request/response logic)

**Parameters:**
- Named parameters for clarity
- Context not consistently passed (some handlers accept `http.Request`, others don't)
- Dependencies passed as parameters (e.g., `cfg *config.Config`, `w http.ResponseWriter`, `r *http.Request`)

**Return Values:**
- Multiple return values common: `(value, error)`
- Error as last return value when present
- Single return value for simple getters: `string`, `bool`

**Example pattern:**
```go
func LoadConfig(path string) (*Config, error) {
    // Set default values
    config := &Config{}

    // Read and parse
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }

    return config, nil
}
```

## Module Design

**Exports:** Exported symbols use PascalCase, private symbols use lowercase

**Barrel Files:** Not used (each package exports its own symbols directly)

**Package structure:**
- `internal/` directory for all application packages
- Each package in its own subdirectory
- `vendor/` for vendored dependencies
- Clear separation of concerns (auth, handlers, config, utils, etc.)

**Common patterns:**
- Package-level configuration variables (e.g., `var cfg *config.Config`)
- Initialization functions: `InitHandlers()`, `InitSessionStore()`
- Singleton-like patterns for shared state

## Struct Tags

**YAML:** `yaml:"field_name"` or `yaml:"field_name,omitempty"`
**JSON:** `json:"field_name"` or `json:"field_name,omitempty"`

**Example:**
```go
type User struct {
    Username string   `yaml:"username" json:"username"`
    Password string   `yaml:"password" json:"password,omitempty"`
    Role     string   `yaml:"role" json:"role"`
    Groups   []string `yaml:"groups,omitempty" json:"groups,omitempty"`
}
```

## Concurrency

**Patterns observed:**
- Mutex-based locking: `var mu sync.RWMutex`
- `mu.Lock()` and `mu.Unlock()` for protecting shared state
- `mu.RLock()` and `mu.RUnlock()` for read-only access
- Goroutines for background tasks (e.g., session cleanup)

**Example:**
```go
var (
    sessions = make(map[string]Session)
    mu       sync.RWMutex
)

func GetSession(r *http.Request) *Session {
    mu.Lock()
    defer mu.Unlock()
    // Access sessions safely
}
```

## HTTP Handler Patterns

**Signature:** Standard http.HandlerFunc pattern with config parameter:
```go
func PageHandler(w http.ResponseWriter, r *http.Request, cfg *config.Config) {
    // Handler logic
}
```

**Common operations:**
- Set headers: `w.Header().Set("Content-Type", "application/json")`
- Status codes: `w.WriteHeader(http.StatusOK)`
- JSON responses: `json.NewEncoder(w).Encode(data)`
- Redirects: `http.Redirect(w, r, "/login", http.StatusFound)`
- Error responses: `http.Error(w, message, statusCode)`

---

*Convention analysis: 2026-04-05*
