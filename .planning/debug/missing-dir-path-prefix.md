---
status: fixing
trigger: "Issue: File preview shows 'file not found' when clicking on uploaded file. Log shows URL missing MDwiki/documents/ prefix"
created: 2026-04-06T00:00:00Z
updated: 2026-04-06T00:00:00Z
---

## Current Focus
hypothesis: "The generated file URLs are missing the full document path (MDwiki/documents/ prefix). This could be caused by getCurrentDocPath() returning wrong value, incorrect API call from frontend, or path reconstruction bug in ServeFileHandler."
test: "Added extensive logging throughout the file upload/listing/serving flow to trace where the path loses its prefix"
expecting: "Logs will show exactly where and how the path gets modified"
next_action: "Rebuild completed - ready for testing with actual application"

## Resolution
root_cause: "PENDING - Need to run application and observe logs to determine exact cause"
fix: "PENDING - Fix depends on root cause diagnosis"
verification: "PENDING"
files_changed: []

## Resolution
root_cause: 
fix:
verification:
files_changed: []

## Symptoms
expected: File preview should work when clicking on uploaded file
actual: File preview shows "file not found"
errors: URL missing MDwiki/documents/ prefix in file path
reproduction: Click on uploaded file in UI, observe 404 error
started: Unknown (user discovered it)

## Eliminated

## Evidence
- timestamp: 2026-04-06T00:00:00Z
  checked: internal/handlers/files.go - ListFilesHandler and ServeFileHandler
  found: ListFilesHandler creates URL: /api/files/{path}/{filename} where path should include MDwiki/documents/
  implication: URL should contain full path including MDwiki/documents/
  
  found: config.yaml shows root_dir="data", documents_dir="documents"
  implication: Files are stored at data/documents/{path}
  
  checked: internal/resources/static/js/utilities.js - getCurrentDocPath() function
  found: Function extracts path from window.location.pathname
  found: Removes leading slash, trailing slash, .md extension, and /document suffix
  implication: Should return full path like "MDwiki/documents/test"
  
  checked: internal/resources/static/js/file-utilities.js - file upload and listing
  found: Line 97 and 105 use getCurrentDocPath() to construct file URLs
  implication: Frontend constructs URLs as /api/files/${getCurrentDocPath()}/${filename}
  
  added: Enhanced logging to getCurrentDocPath() to trace path processing
  added: Enhanced logging to ListFilesHandler to trace URL generation
  added: Enhanced logging to UploadFileHandler to trace docPath from form
  
  action: Created test directory structure at data/documents/MDwiki/documents/test/
  action: Built application successfully (build/wiki-go-linux-amd64)
  next_action: "Test the application by accessing a page at /MDwiki/documents/test and observing logs"

## Resolution
root_cause:
fix:
verification:
files_changed: []
