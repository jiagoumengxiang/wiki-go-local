# Codebase Concerns

**Analysis Date:** 2026-04-05

## Tech Debt

**Debug Print Statements in Production Code:**
- Issue: Multiple `fmt.Printf` statements used for debugging that should be replaced with proper logging
- Files: `internal/handlers/files.go`, `internal/handlers/metadata_api.go`
- Impact: Debug output in production can leak sensitive information and clutter logs
- Fix approach: Replace all `fmt.Printf` calls with structured logging using `log` package or a proper logging framework

**Code Duplication in Session Management:**
- Issue: `CleanupExpiredSessions` in session_store.go duplicates file save logic (lines 106-130)
- Files: `internal/auth/session_store.go`
- Impact: Maintenance burden - changes to save logic need to be made in multiple places
- Fix approach: Extract save logic into a private helper function that can be called from both `SaveSessions` and `CleanupExpiredSessions`

**Large File Complexity:**
- Issue: files.go is 1395 lines with multiple responsibilities (upload, list, delete, rename, serve)
- Files: `internal/handlers/files.go`
- Impact: Difficult to understand, test, and maintain; violates single responsibility principle
- Fix approach: Split into smaller focused modules: `upload.go`, `listing.go`, `deletion.go`, `renaming.go`, `serving.go`

**Deprecated Package Usage:**
- Issue: Uses `io/ioutil` package which is deprecated since Go 1.16
- Files: `internal/auth/session_store.go` (lines 5, 73)
- Impact: Package will be removed in future Go versions, requiring migration
- Fix approach: Replace `ioutil.ReadFile` with `os.ReadFile` throughout codebase

**Migration Code in Production:**
- Issue: Migration functions run on every startup but handle one-time schema changes
- Files: `internal/migration/role_migration.go`, `internal/migration/fix_config.go`
- Impact: Unnecessary overhead; migration logic should be separate from core application
- Fix approach: Create standalone migration tool that users run explicitly when upgrading

## Known Bugs

**Ignored File Read Errors:**
- Symptoms: Silent failures when reading markdown files for comments
- Files: `internal/handlers/page.go` (lines 228, 233)
- Trigger: When filesystem I/O fails during comment loading
- Workaround: None - errors are silently ignored
- Impact: Comments may fail to load without user awareness

**Hardcoded HTTP Server Binding:**
- Symptoms: Server always binds to "0.0.0.0" regardless of config
- Files: `internal/config/config.go` (line 95)
- Trigger: Always - config value is overridden
- Workaround: Modify source code to change binding address
- Impact: Cannot bind to localhost only in production for security

## Security Considerations

**Unsafe CSP Policy:**
- Risk: Content-Security-Policy uses 'unsafe-inline' and 'unsafe-eval' which weaken security
- Files: `internal/routes/routes.go` (lines 55, 59)
- Current mitigation: Currently in Report-Only mode (line 80), not enforced
- Recommendations: Remove 'unsafe-inline' and 'unsafe-eval', enforce CSP, use nonces for inline scripts

**Panic on Resource Initialization:**
- Risk: Application crashes if embedded filesystem setup fails
- Files: `internal/resources/resources.go` (lines 26, 46, 55)
- Current mitigation: None - panics are not caught
- Recommendations: Return errors instead of panicking, handle gracefully in main()

**Backup File Cleanup:**
- Risk: No automatic cleanup of old backup files
- Files: `internal/handlers/backup.go`
- Current mitigation: Manual deletion via admin interface
- Recommendations: Add retention policy (e.g., keep last 10 backups, max 30 days), implement automatic cleanup

**Session File Permissions:**
- Risk: Session store files created with 0o600 permissions but directory with 0o700
- Files: `internal/auth/session_store.go` (line 22)
- Current mitigation: Restrictive permissions on files
- Recommendations: Document that temp/sessions.json directory should be secured, add validation on startup

## Performance Bottlenecks

**Unconstrained Backup Operations:**
- Problem: Backup operation walks entire filesystem without throttling
- Files: `internal/handlers/backup.go` (lines 93, 125)
- Cause: No rate limiting or resource limits during zip creation
- Impact: Can cause high CPU/memory usage on large wikis, slow server response
- Improvement path: Add progress tracking already present, implement chunking for large wikis, add resource limits

**Multiple File System Reads per Request:**
- Problem: Each page request may read multiple files (markdown, metadata, comments, versions)
- Files: `internal/handlers/page.go`
- Cause: No caching layer for frequently accessed documents
- Impact: Slow page load times, high I/O on file server
- Improvement path: Add in-memory cache for document content with TTL, implement cache invalidation on edits

**Inefficient Comment Loading:**
- Problem: Comments are read from filesystem on every page load
- Files: `internal/handlers/page.go` (line 233)
- Cause: No caching of comment data
- Impact: Unnecessary filesystem I/O for pages with comments
- Improvement path: Cache comment metadata, invalidate on comment create/delete

## Fragile Areas

**Backup Job Management:**
- Files: `internal/handlers/backup.go` (lines 42-44)
- Why fragile: Backup jobs stored in in-memory map, lost on server restart
- Safe modification: Migrate to persistent job tracking (database or file), add job recovery on startup
- Test coverage: No tests for backup job state persistence across restarts

**Configuration File Parsing:**
- Files: `internal/config/config.go`
- Why fragile: Schema changes can break existing configs, relies on manual migration
- Safe modification: Implement versioning in config file, auto-migration between versions
- Test coverage: Limited tests for config parsing edge cases

**Emoji Data Loading:**
- Files: `internal/goldext/emoji.go` (line 38)
- Why fragile: Silent failure if emoji data is corrupted or missing
- Safe modification: Add error handling, fallback to emoji-less rendering, log warnings
- Test coverage: No tests for missing/corrupted emoji data

**Import Process:**
- Files: `internal/handlers/import.go` (line 139)
- Why fragile: Import runs in goroutine with no progress tracking or cancellation
- Safe modification: Add job tracking similar to backups, implement cancellation, improve error reporting
- Test coverage: No tests for import process reliability

## Scaling Limits

**In-Memory Session Storage:**
- Current capacity: All sessions stored in memory + persisted to single JSON file
- Limit: Memory grows linearly with active users, file locking on concurrent writes
- Scaling path: Migrate to Redis or dedicated session store for horizontal scaling

**File-Based Storage:**
- Current capacity: All content stored on local filesystem
- Limit: Single-server deployment, no built-in replication or backup redundancy
- Scaling path: Add S3/minio backend option, implement storage abstraction layer

**Single-Process Server:**
- Current capacity: Single Go process handles all requests
- Limit: CPU-bound on a single core, no automatic sharding of requests
- Scaling path: Add graceful shutdown support, allow multiple processes behind load balancer

## Dependencies at Risk

**gopkg.in/yaml.v3:**
- Risk: Using third-party YAML parser instead of standard library
- Impact: yaml package has security vulnerabilities, adds dependency surface
- Migration plan: Consider migrating to standard library encoding/yaml when available, keep version updated

**Embedded Vendor Directory:**
- Risk: Vendor directory contains 15000+ lines of third-party code
- Impact: Large binary size, outdated dependencies not updated
- Migration plan: Use Go modules properly, remove vendor directory from git, ensure go.sum is committed

## Missing Critical Features

**Rate Limiting:**
- Problem: No rate limiting on API endpoints or authentication attempts
- Blocks: Protection against brute force attacks, API abuse
- Impact: Vulnerable to denial-of-service attacks

**Audit Logging:**
- Problem: No audit trail for admin actions, configuration changes, or data modifications
- Blocks: Compliance requirements, security incident investigation
- Impact: Cannot track who made changes or when

**Health Check Endpoint:**
- Problem: No /health endpoint for load balancer or monitoring probes
- Blocks: Container orchestration integration, monitoring systems
- Impact: Difficult to detect service failures in production

**Graceful Shutdown:**
- Problem: Server exits immediately on SIGTERM, may drop in-flight requests
- Blocks: Zero-downtime deployments, proper resource cleanup
- Impact: May corrupt data during deployment

## Test Coverage Gaps

**Handler Functions:**
- What's not tested: All HTTP handlers in `internal/handlers/` (25 files)
- Files: `internal/handlers/*.go`
- Risk: Request handling bugs, missing authentication checks, incorrect error responses
- Priority: High - handlers are the core of the application

**Authentication Flow:**
- What's not tested: Session management, login/logout, password verification
- Files: `internal/auth/auth.go`, `internal/auth/session_store.go`
- Risk: Authentication bypass, session hijacking, credential leakage
- Priority: High - security critical

**File Operations:**
- What's not tested: File upload, download, deletion, content type detection
- Files: `internal/handlers/files.go` (1395 lines)
- Risk: Security vulnerabilities, data loss, incorrect content type handling
- Priority: High - security and data integrity critical

**Configuration Loading:**
- What's not tested: Config parsing, validation, default values
- Files: `internal/config/config.go`
- Risk: Application crashes on invalid config, silent misconfiguration
- Priority: Medium - affects startup reliability

**Migration Logic:**
- What's not tested: Role migration, config fix logic
- Files: `internal/migration/*.go`
- Risk: Data corruption during upgrade, migration failures
- Priority: Medium - affects upgrade path

**Template Rendering:**
- What's not tested: HTML template rendering, markdown conversion
- Files: `internal/handlers/*.go`
- Risk: Incorrect HTML output, XSS vulnerabilities, broken rendering
- Priority: Medium - affects user experience

**Integration Tests:**
- What's not tested: End-to-end workflows (create/edit/delete documents, comments, uploads)
- Risk: System-level bugs, broken workflows after refactoring
- Priority: Medium - ensures overall functionality

---

*Concerns audit: 2026-04-05*
