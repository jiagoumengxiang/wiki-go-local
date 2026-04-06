---
phase: quick
plan: 260406-fah-remove-default-password-prompt
type: execute
wave: 1
depends_on: []
files_modified:
  - internal/resources/templates/base.html
  - internal/resources/static/js/auth.js
  - internal/resources/static/css/theme.css
  - internal/resources/static/css/navigation.css
  - internal/resources/static/css/print.css
autonomous: true
requirements: []
must_haves:
  truths:
    - "No default password warning banner appears on any page"
    - "No JavaScript code checks for default password on page load"
    - "No CSS styles reference password warning banner"
  artifacts:
    - path: "internal/resources/templates/base.html"
      provides: "HTML templates without password warning banner"
      contains: "!password-warning-banner"
    - path: "internal/resources/static/js/auth.js"
      provides: "Auth JavaScript without default password check"
      contains: "!checkDefaultPassword()"
    - path: "internal/resources/static/css/theme.css"
      provides: "Theme CSS without password warning styles"
      contains: "!.password-warning-banner"
  key_links: []
---

<objective>
Remove the default password warning banner and all related functionality from the frontend.

Purpose: The default password reminder is no longer needed in this environment.
Output: Clean frontend without password warning prompts.
</objective>

<execution_context>
@$HOME/.config/opencode/get-shit-done/workflows/execute-plan.md
@$HOME/.config/opencode/get-shit-done/templates/summary.md
</execution_context>

<context>
@.planning/STATE.md
@internal/resources/templates/base.html
@internal/resources/static/js/auth.js
@internal/resources/static/css/theme.css
@internal/resources/static/css/navigation.css
@internal/resources/static/css/print.css
</context>

<tasks>

<task type="auto">
  <name>Task 1: Remove default password warning banner and related code</name>
  <files>
    internal/resources/templates/base.html
    internal/resources/static/js/auth.js
    internal/resources/static/css/theme.css
    internal/resources/static/css/navigation.css
    internal/resources/static/css/print.css
  </files>
  <action>
Remove all default password warning functionality from the frontend:

1. **HTML Template (internal/resources/templates/base.html)**:
   - Remove lines 93-96 containing the password warning banner div:
     ```html
     <!-- Password warning banner (will be shown if default password is in use) -->
     <div id="password-warning-banner" class="password-warning-banner" style="display: none;">
         <i class="fa fa-lg fa-exclamation-triangle" aria-hidden="true"></i> Change the default admin password.
     </div>
     ```

2. **JavaScript (internal/resources/static/js/auth.js)**:
   - Remove line 71: `checkDefaultPassword();`
   - NOTE: Do not remove the `checkDefaultPassword` function definition or export to avoid breaking any external dependencies, but the call that triggers it on page load must be removed.

3. **CSS - Theme (internal/resources/static/css/theme.css)**:
   - Remove lines 145-185 containing:
     - `.password-warning-banner` selector and styles
     - `body.has-password-warning` selector and padding styles
     - `body.has-password-warning .sidebar` and `.hamburger` override styles

4. **CSS - Navigation (internal/resources/static/css/navigation.css)**:
   - Remove lines 243-246: `body.has-password-warning .breadcrumbs` styles
   - Remove lines 313-318: `body.has-password-warning .hamburger` mobile override
   - Remove lines 413-416: `body.has-password-warning .breadcrumbs` desktop override

5. **CSS - Print (internal/resources/static/css/print.css)**:
   - Remove `.password-warning-banner,` from line 31 in the display:none selector list

All changes are removal-only (no new code added).
</action>
  <verify>
    <automated>grep -r "password-warning" internal/resources/templates/ internal/resources/static/ 2>/dev/null | grep -v "node_modules" || echo "No password-warning references found"</automated>
  </verify>
  <done>No references to password-warning banner, checkDefaultPassword() call, or related CSS styles exist in frontend code</done>
</task>

</tasks>

<verification>
Verify that:
1. No HTML elements with id="password-warning-banner" exist
2. No JavaScript calls to checkDefaultPassword() on page load
3. No CSS selectors for .password-warning-banner or body.has-password-warning
4. Application loads without password warning banner
</verification>

<success_criteria>
- All password warning banner code removed from frontend
- No JavaScript errors related to removed code
- Application functions normally without banner
</success_criteria>

<output>
After completion, create `.planning/quick/260406-fah-remove-default-password-prompt/260406-fah-SUMMARY.md`
</output>
