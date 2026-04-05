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

    // If .md exists in the path, remove it (some implementations add .md to URLs)
    if (path.endsWith('.md')) {
        path = path.substring(0, path.length - 3);
    }

    // If /document exists at the end, remove it for old directory structure
    // Old structure: /-opencode/ttttt/document -> -opencode/ttttt
    // New structure: /-opencode/ttttt.md -> -opencode/ttttt
    const lastSegment = path.split('/').pop();
    if (lastSegment === 'document') {
        path = path.substring(0, path.length - '/document'.length);
        console.log("Removed '/document' from path for old directory structure");
    }

    console.log("Processed document path:", path);
    return path || 'pages/home';
}

// Make function available globally
window.getCurrentDocPath = getCurrentDocPath;