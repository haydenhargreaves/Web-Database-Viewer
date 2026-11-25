# Independent plan review

Run 20261125-query-history-01; invocation critic-002; revision review. Plan v005 SHA-256 fbf2fe149c0e32e3d3a939d841b0d5f5e029fc8a880e79a5d2e866b18853959e; requirements plan/requirements-v005.md SHA-256 c876635b1d791b3a80dea91ddb783f72c6f6bba427d147b5599c51759e04e12a. Supplied S0 and observed HEAD 9ebc8c67e24c79392989eb533f98195370580feb; no application changes observed, run metadata untracked. Hashes independently computed.

## Verdict: REJECTED
Planning readiness only: selected option text is not a faithful representation of all accepted connection names, so the proposed display key can differ from the response key. No implementation or tests are certified.

## Evidence inspected
Ticket, clarifications, baseline and prior finding report; `internal/database/query.go` resolves session map and executes SQL; `connect.go` and `create.go` set current name; `internal/templates/connections.go` interpolates unescaped names into option HTML; `internal/templates/query.go` and `web/templates/index.html` supply swapped and delayed query controls; `web/static/scripts/tree.js` dispatches input; `theme.js` and index access storage. `go.mod` and `tools/styles/compile.sh` ground tooling. Baseline reports 11 Go tests passing; not executed in this review. Plan and requirements hashes match assignment.

## Requirement coverage
SATISFIED denotes adequate implementation and verification planning only, not actual acceptance.

| ID | Planned behavior / steps | Verification | Outcome |
|---|---|---|---|
| R-001 | P-001/P-002 response-correlated SQL capture; omit unknowable response | P-004 manual, auto, failure, races | SATISFIED |
| R-002 | P-001/P-002 hash selected option text and actual request identity | P-004 switches, HTTP, shared-cookie race | UNSATISFIED: C002-1 |
| R-003 | P-003 timestamp, preview and full SQL | P-004 multiline and ordering | SATISFIED |
| R-004 | P-003 suppress pending delayed request without suppressing real input | P-004 real HTMX timing | SATISFIED |
| R-005 | P-003 isolated remove and clear | P-004 duplicates, isolation and keyboard | SATISFIED |
| R-006 | P-002 bounded sorted entries, duplicate preservation | P-004 51, races, cross-tab limitation | SATISFIED |
| R-007 | P-001–P-003 pseudonymous keys, SQL/timestamp only, warning | P-004 vectors, sentinels, XSS | SATISFIED |
| R-008 | P-003 fail-open storage and theme | P-004 exceptions, malformed data, query POST | SATISFIED |
| R-009 | P-003 accessible responsive themed panel | P-004 keyboard, focus, styles, width | SATISFIED |
| R-010 | P-001/P-004 preserve query behavior and selector; additive header | canonical Go suite and HTMX regression | SATISFIED |

## Blocking finding

### C002-1 — HIGH — selector-derived identity can mismatch the actual connection (R-002)
Plan Design/P-002 hashes the selected option's text while the server hashes the exact session name. `internal/templates/connections.go` inserts the name into HTML without escaping, and `internal/database/create.go` accepts arbitrary names. For example a name `A &amp; B` is parsed into option text `A & B`; the server hashes the original literal name, but the panel hashes the decoded text. A real query is stored under one key and its selected connection displays another, including after reload. Similar name/placeholder ambiguity is possible. This is a concrete existing input, not a hypothetical browser capability issue. Specify a reliable non-secret selected-connection identifier corresponding exactly to the session name even for HTML-sensitive and whitespace names, and distinguish an actual connection from the placeholder/missing current without changing selector semantics. The plan already allows scoped selector markup modification if needed; preserve compatibility and avoid storing or exposing URL/credentials as history metadata. Add real browser and server tests for literal entity-like names, whitespace/Unicode and placeholder-like names, demonstrating recording, display, switch and reload use the same identity. Verify unchanged selector behavior and secrecy.

## Prior blockers
| ID | Disposition | Evidence |
|---|---|---|
| C001-1 | Resolved at planning scope | v005 Design replaces secure-context-only Web Crypto with synchronous JS and Go digest; P-004 requires capability-disabled HTTP-equivalent test, isolation and secrecy. |
| C001-2 | Resolved at planning scope | v005 P-001/P-002 correlates each XHR with a server-resolved request-session header; clarifications explicitly permit omission of lost/unidentifiable responses without retry; P-004 covers switch and shared-cookie races. C002-1 is a separate display identity defect. |

## Scope, feasibility and verification
Bounded path scope, protected S0, response header and fail-open behavior are otherwise credible. Digest is explicitly not encryption; cross-tab lost updates and late response after clear are documented limitations. Real HTMX 2.0.1 provenance, browser and canonical Go suite remain downstream verification prerequisites; none was run here. No optional observations or external blockers.
