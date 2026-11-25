Please implement this ticket using the resources of the council. Do not leave any steps out.
Ticket: Add per-connection query history
Type: Feature  
Priority: Medium  
Summary: Let users review and rerun recent SQL queries for each database connection.
User story
As a user, I want to see queries I’ve run on the current connection, so I can quickly reuse or inspect past work without keeping a separate notes file.
Scope and behavior
- Add a Query History panel near the SQL editor or results area, consistent with the existing interface and light/dark themes.
- Record each submitted, non-empty SQL query when the user runs it manually or through Auto-Run.
- Associate entries with the selected connection so switching connections shows that connection’s history.
- Show newest entries first, with query text and a readable timestamp. Keep the full query available even when the preview is shortened.
- Selecting an entry loads its SQL into the editor but does not execute it. The user runs it explicitly.
- Provide actions to clear the current connection’s history and remove individual entries.
- Keep a bounded history: retain at most 50 entries per connection, dropping the oldest when adding another.
- Persist history in browser localStorage, not in server sessions or database storage. Do not include connection URLs, passwords, or result data in history.
- If local storage is unavailable or malformed, keep the page usable; history may be unavailable for that browser session, but query execution must still work.
Implementation notes
- Query submission currently goes through POST /v1/api/query from the controls in web/templates/index.html and internal/templates/query.go. The Auto-Run textarea uses HTMX’s delayed input trigger.
- The selected connection is represented in the page by #connected-database; the server also stores the current connection name in its session. Use a stable, non-secret connection identifier for client-side history keys. Avoid using connection URLs as identifiers or persisting them in history.
- Existing client-side behavior lives in web/static/scripts/; the page loads scripts from web/templates/index.html.
- Keep this feature client-side unless implementation reveals a compelling reason otherwise. Do not send query history to a new server endpoint.
- Prevent history selection and clear/remove actions from accidentally triggering Auto-Run. Loading a saved query should only populate the editor.
- Make the history panel keyboard accessible, with visible focus states and meaningful button labels.
Acceptance criteria
1. Running a non-empty query adds it to the current connection’s history, whether submitted manually or via Auto-Run.
2. Empty or whitespace-only submissions do not create history entries.
3. Switching connections shows that connection’s separate history; switching back restores the prior list.
4. Selecting an entry fills the SQL editor without executing the query.
5. Individual entries and the current connection’s entire history can be removed.
6. No connection credentials, URLs, query results, or other connection details are stored as history data.
7. History is capped at 50 entries per connection, with the newest retained.
8. History works in both themes and at narrow viewport widths, and does not block querying if browser storage fails.
9. Add or update tests where practical, and verify existing Go tests still pass.
Out of scope
- Named/pinned saved queries or folders
- Search across history
- Import/export or server-side synchronization
- SQL syntax highlighting or editor replacement
Security consideration: Query text itself can contain sensitive information. The UI should make clear that history is stored in this browser and provide an obvious way to clear it.
