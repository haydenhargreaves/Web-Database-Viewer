# Requirements v002 — run 20261125-query-history-01

Source: ticket `../ticket.md`; clarification `../clarifications.md` (2026-11-25 verification-only HTMX authorization). S0/current HEAD 9ebc8c67e24c79392989eb533f98195370580feb. Revision: architect-002; previous architect-001 did not persist a requirements revision.

## Objective and non-goals
Review and reuse recent SQL per selected connection, locally in the browser. No saved/pinned queries, folders, history search, import/export, synchronization, editor replacement, server history endpoint or unrelated cleanup.

| ID | Source | Mandatory requirement and observable acceptance criterion | Verification |
|---|---|---|---|
| R-001 | ticket behavior, AC 1–2 | Each actually submitted manual or delayed Auto-Run POST with non-whitespace SQL adds one entry, even on query failure; empty/whitespace adds none. Keep full submitted text unchanged. | Real HTMX browser request-count and payload tests, P-003/P-005 |
| R-002 | ticket behavior, AC 3 | Each selected connection displays only its own history; switching away/back and reloading restores it. No valid connection means no entry or cross-connection display. | Browser selection, replacement and reload tests, P-002/P-005 |
| R-003 | ticket behavior | Entries display newest first with readable timestamps and SQL preview; full SQL remains accessible and loadable. | Browser long/multiline and timestamp checks, P-003/P-005 |
| R-004 | ticket behavior, AC 4 | Selecting loads full SQL but sends no query; history select/remove/clear must not trigger Auto-Run, including an already delayed input. | Actual HTMX delayed-input timing tests, P-003/P-005 |
| R-005 | ticket behavior, AC 5 | Remove one entry or clear only the current connection; controls visibly identify their purpose. | Browser delete/clear/isolation tests, P-003/P-005 |
| R-006 | ticket behavior, AC 7 | Retain at most 50 entries per connection, discarding oldest upon each new entry; repeated SQL submissions remain separate. | Browser 51 submissions and duplicates tests, P-003/P-005 |
| R-007 | ticket behavior, AC 6, security | Persist only SQL and timestamp per entry under a stable non-secret connection identifier in localStorage; never persist connection URL, credentials, result data or other connection details as history metadata. Display that SQL can be sensitive and history is stored in this browser, with an obvious clear action. SQL itself may contain secrets. | Inspect storage keys/values using sentinel connection data and render warning, P-002/P-005 |
| R-008 | ticket behavior, AC 8 | Invalid/malformed/unavailable storage, including denied access and quota failure, cannot prevent query execution or page usability; history may be unavailable. | Browser injection of throwing storage operations and malformed data, P-003/P-005 |
| R-009 | ticket behavior, AC 8 | Panel near editor/results is usable in light/dark themes and narrow widths, keyboard accessible with meaningful button labels and visible focus. | Browser keyboard, computed-style and viewport checks, P-004/P-005 |
| R-010 | ticket AC 9; existing invariant | Existing Go tests pass and query, selector, editor toggling and tree-driven input behavior remain compatible; no new server history persistence or endpoint. | Canonical suite and real HTMX regression checks, P-005 |

## Compatibility and protected behavior
Never change query execution semantics just to record history; history observes requests and fails open. Preserve existing connection submit payload and server session logic, HTMX selector/editor replacements, and tree helpers that intentionally dispatch input. Store no connection value: current `option.value` contains URL and driver. Execute no SQL during test. Unrelated developer and Council state is protected.

## Assumptions and questions
Connection display names, not option values, are the server's current session identity. Hash the exact displayed name for storage key; this is a stable pseudonymous identifier, not encryption, and cannot protect low-entropy names from guessing. Names may be recreated or reused: history follows name, since no durable independent ID exists; this matches existing session identity. An empty/no-connection placeholder is not an identity. A query containing credentials is still recorded as SQL per ticket; warn explicitly. Browser tabs can race on localStorage read-modify-write; no cross-tab transactional guarantee is requested. Real HTMX asset and browser verification prerequisites remain to be obtained, not waived; see plan. No blocking product question identified.
