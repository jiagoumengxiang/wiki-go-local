/**
 * Utility functions for the Wiki-Go application
 */

/**
 * Get the current document path from the URL
 * @returns {string} The standardized document path
 */
function getCurrentDocPath() {
    // Log the raw path for debugging
    console.log("Raw pathname for path processing:", window.location.pathname);

    const isHomepage = window.location.pathname === '/';
    if (isHomepage) {
        console.log("Using homepage path");
        return 'pages/home';
    }

    // For versions, we need to keep the full path structure
    let path = window.location.pathname;

    // Remove leading slash
    if (path.startsWith('/')) {
        path = path.substring(1);
    }

    // Remove trailing slash if it exists
    if (path.endsWith('/')) {
        path = path.substring(0, path.length - 1);
    }

    // If /document exists at the end, remove it for old directory structure
    // Old structure: /-opencode/ttttt/document -> -opencode/ttttt
    // New structure: /-opencode/ttttt.md -> -opencode/ttttt
    const lastSegment = path.split('/').pop();
    if (lastSegment === 'document') {
        path = path.substring(0, path.length - '/document'.length);
        console.log("Removed '/document' from path for old directory structure");
    }

    // Note: We now use .md file structure (e.g., .opencode/a.md), so .md files should keep their .md extension
    // This ensures attachments are stored in the same directory as the .md file
    // When accessing /.opencode/a.md, docPath should be ".opencode/a.md", not ".opencode/a"
    // Do NOT remove the .md extension - the backend needs it to identify .md file structure

    console.log("Processed document path:", path);
    return path || 'pages/home';
}

/**
 * Get the directory path from a document path
 * If docPath ends with .md, it strips the filename and returns the directory
 * @param {string} docPath - The document path
 * @returns {string} The directory path
 */
function getDocDir(docPath) {
    if (!docPath || docPath === '' || docPath === '/') {
        return '';
    }
    // If the path ends with .md, extract the directory
    if (docPath.endsWith('.md')) {
        const lastSlash = docPath.lastIndexOf('/');
        if (lastSlash !== -1) {
            return docPath.substring(0, lastSlash);
        }
        return '';
    }
    return docPath;
}

// Make function available globally
window.getCurrentDocPath = getCurrentDocPath;
window.getDocDir = getDocDir;