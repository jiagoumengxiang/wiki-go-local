# Testing Patterns

**Analysis Date:** 2026-04-05

## Test Framework

**Runner:**
- Standard Go testing package (`testing`)
- No test framework config files

**Assertion Library:**
- Go's built-in testing assertions (`t.Errorf`, `t.Fatal`)

**Run Commands:**
```bash
go test              # Run all tests in current package
go test ./...        # Run all tests in all packages
go test -v           # Verbose output
go test -cover       # Run tests with coverage
```

**Note:** No test targets found in Makefile - tests must be run manually

## Test File Organization

**Location:**
- Co-located with source files (same package directory)
- Test file named `{source}_test.go` alongside `{source}.go`
- Only one test file found in codebase: `internal/goldext/superscript_test.go`

**Naming:**
- Test functions: `Test{FunctionName}` (e.g., `TestSuperscriptPreprocessor`)
- Subtests: Descriptive names in table-driven tests

**Structure:**
```
internal/
├── goldext/
│   ├── superscript.go
│   └── superscript_test.go
└── [other packages with no tests]
```

**Test coverage:** Extremely limited - only one test file found

## Test Structure

**Suite Organization:**
- Table-driven tests are the pattern
- Test cases defined as structs in a slice

**Pattern observed:**
```go
func TestSuperscriptPreprocessor(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
    }{
        {
            name:     "Basic superscript",
            input:    "This is a ^test^ of superscript.",
            expected: "This is a <sup>test</sup> of superscript.",
        },
        {
            name:     "Multiple superscripts",
            input:    "H^2^O and E=mc^2^ are formulas.",
            expected: "H<sup>2</sup>O and E=mc<sup>2</sup> are formulas.",
        },
        // ... more test cases
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := SuperscriptPreprocessor(tt.input, "")
            if result != tt.expected {
                t.Errorf("Expected: %q, got: %q", tt.expected, result)
            }
        })
    }
}
```

**Patterns:**
- **Setup:** Minimal - test cases defined inline
- **Teardown:** None observed
- **Assertion:** Direct comparison with `t.Errorf`

## Mocking

**Framework:** No mocking framework detected

**Patterns:**
- Tests are pure unit tests without external dependencies
- No interfaces used for mocking
- No mocking of HTTP handlers, databases, or external services
- Tests focus on pure functions with deterministic output

**What to Mock:**
- No guidelines found (mocking not used in codebase)

**What NOT to Mock:**
- Not applicable - mocking not used

## Fixtures and Factories

**Test Data:**
- Hardcoded in test cases (table-driven approach)
- No separate test fixtures directory

**Example from codebase:**
```go
tests := []struct {
    name     string
    input    string
    expected string
}{
    // Test cases defined inline
}
```

**Location:**
- Test data is inline in test files
- No external fixture files

## Coverage

**Requirements:** None enforced

**View Coverage:**
```bash
go test -cover           # Show coverage for package
go test -coverprofile=coverage.out    # Generate coverage profile
go tool cover -html=coverage.out     # View HTML coverage report
```

**Current coverage status:** Extremely low - only one test file covering one function

## Test Types

**Unit Tests:**
- Primary testing approach observed
- Tests individual functions in isolation
- No integration tests found
- No end-to-end tests found

**Integration Tests:**
- Not detected

**E2E Tests:**
- Not detected

## Common Patterns

**Table-driven tests:**
- Standard pattern used in the only test file
- Good for testing multiple scenarios
- Easy to add new test cases

**Subtests with t.Run():**
- Used for organizing test cases
- Provides clear test names in output

**String comparison:**
- Primary assertion method: `if result != tt.expected`
- Uses `%q` format for clear output on failure

## Missing Testing Infrastructure

**No test targets in Makefile:** Consider adding:
```makefile
test:
    go test ./...

test-cover:
    go test -cover ./...
```

**No CI test execution:** GitHub Actions workflows exist but focus on releases, not testing

**No test utilities:** No helper functions for test setup/teardown

## Recommendations

**Add tests for critical paths:**
- HTTP handlers (`internal/handlers/`)
- Authentication logic (`internal/auth/`)
- Configuration parsing (`internal/config/`)
- File operations and sanitization (`internal/utils/`)

**Consider testing framework:**
- No external framework needed - Go's testing is sufficient
- Consider `testify` for assertions if more complex tests are added
- Consider `httptest` for HTTP handler testing

**Integration tests:**
- Add tests for end-to-end workflows
- Test database operations (file-based in this case)
- Test session management

---

*Testing analysis: 2026-04-05*
