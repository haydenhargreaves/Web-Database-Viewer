# Requirements v005 — run 20261125-query-history-01

Input: `../ticket.md`, `../clarifications.md` and additive clarification supplied to architect-005; S0/current HEAD `9ebc8c67e24c79392989eb533f98195370580feb`. Revision replaces v002. Objective: review and reuse recent SQL for each connection in this browser. Non-goals: saved queries/folders, search, import/export, server history or synchronization, editor replacement, unrelated fixes.

All requirements below are mandatory; implementation acceptance remains unverified.

| ID | Source | Requirement and observable acceptance | Verification |
|---|---|---|---|
| R-001 | Ticket; clarification on lost responses | Each actual manual or delayed Auto-Run submission with non-whitespace SQL and an available valid effective-connection response identifier creates one entry preserving exact SQL, including SQL errors; whitespace creates none. A lost response or response without an established identity may omit it; never retry SQL or guess identity. | P-002/P-004 browser and Go failure/race cases |
| R-002 | Ticket; effective-ID clarification | Display separate histories for selected connection, restored on switch/reload. Attribute each recorded request to the server-resolved connection, including concurrent switches and shared-cookie tabs; never use stale selection as proof. No valid identity means no entry or foreign list. Works on supported HTTP and HTTPS. | P-001/P-002/P-004 identity, HTTP and concurrent-tab cases |
| R-003 | Ticket | Display newest first with readable timestamp and preview; complete multiline SQL remains accessible and loadable. | P-003/P-004 |
| R-004 | Ticket | Selecting loads SQL without executing; selection, removal and clear do not trigger even an already delayed Auto-Run; later real input works normally. | P-003/P-004 actual HTMX timing |
| R-005 | Ticket | Remove one chosen entry or clear current connection only, with clear accessible controls. | P-003/P-004 |
| R-006 | Ticket | At most 50 entries per connection, newest retained; duplicate submissions remain distinct. | P-002/P-004 51 and duplicate tests |
| R-007 | Ticket, security | History localStorage keys use stable non-secret identifier; history data contains SQL and timestamp only, no connection URL, password, result or other connection details as metadata. Warn SQL may be sensitive and stored in this browser, with obvious clear. SQL itself can contain secrets. | P-001/P-003/P-004 sentinel, XSS and UI tests |
| R-008 | Ticket | Malformed or unavailable storage must not prevent querying or page usability; history may be unavailable. | P-003/P-004 throwing and malformed storage |
| R-009 | Ticket | Panel near editor/results supports light/dark, narrow widths, keyboard use, meaningful labels and visible focus. | P-003/P-004 |
| R-010 | Ticket; existing invariant; additive clarification | Existing Go tests and query behavior, session and selector semantics, tree input and toggles remain compatible; only additive non-secret identifier on existing query responses, no history endpoint or server history persistence. | P-001/P-004 canonical Go and HTMX regression |

## Compatibility and assumptions
Request-time SQL is kept only in memory until its response establishes identity; no query is cancelled or retried for history. Identity is a SHA-256 digest of the exact existing session connection name, not a secret or encryption: low-entropy names may be guessed; names can be reused. SQL may itself include credentials, so the prohibition applies to metadata, not altered SQL. Selector option values contain URLs and must never be history identifiers. Storage is origin-scoped. Cross-tab localStorage read-modify-write is not transactional; concurrent writes can overwrite each other, and an outstanding query can finish after clear; see plan for deterministic best effort and limitations. No blocking product question remains; no promise of lossless simultaneous cross-tab updates is inferred. Trusted real HTMX asset and browser remain mandatory verification prerequisites, not planning blockers.
