# Council run 20261125-query-history-01

Status: COMPLETED. Feature: per-connection query history.

## Requirements
- R-001: SATISFIED — XHR correlation inspection; real-HTMX payload, timing, error and identity assertions.
- R-002: SATISFIED — Shared digest and session resolver, Go tests, actual-page switch and delayed-response checks.
- R-003: SATISFIED — Literal multiline SQL, full load, timestamps and ordering checked.
- R-004: SATISFIED — Pending action and subsequent input fixture checks; actual-page editor swap and tree input.
- R-005: SATISFIED — Duplicate removal, current-only clear, keyboard actions and connection isolation checked.
- R-006: SATISFIED — 51-entry cap, duplicates and reversed completion ordering checked.
- R-007: SATISFIED — SQL/timestamp serialization, actual Go identity and browser synthetic sentinel checks; warning present.
- R-008: SATISFIED — Storage guards inspected; malformed, getter and quota failure tests preserve querying.
- R-009: SATISFIED — Actual page checked at 320px with themes, keyboard actions, labels and focus ring.
- R-010: SATISFIED — Canonical Go suite, preserved selector and query paths, actual-page interaction checks and unchanged router.

Evidence: .council/runs/20261125-query-history-01/reviews/final/judge-001/report.md; .council/runs/20261125-query-history-01/reviews/implementation/inspector-005/report.md. Approved plan v007 by critic-004; implementation approved by inspector-005 and final state approved by judge-001. Original clarifications include omission when a query response is lost and its effective connection cannot be identified.

## Changed files
```
 M internal/database/query.go
 M internal/templates/connections.go
 M web/static/scripts/theme.js
 M web/templates/index.html
?? internal/database/query_history_test.go
?? internal/templates/connections_history_test.go
?? web/static/scripts/query-history.js
?? web/tests/query-history-server.js
?? web/tests/query-history.html
?? web/tests/query-history.js
```
No cleanup changes: simplifier-002 found no safe changes; S1 checkpoint .council/runs/20261125-query-history-01/checkpoints/s1-001/manifest.json; S2 identical. No rollback.

## Verification
Judge independently ran canonical `GOTOOLCHAIN=local GOPROXY=off GOSUMDB=off go test -mod=vendor -count=1 -v ./...`: PASS, 13 top-level tests and 8 subtests. Real HTMX browser fixture: PASS, 31 assertions. Actual page and CSS browser harness: PASS, synthetic endpoints; actual Go selector/query handler identity separately checked with synthetic SQLite: PASS. Baseline Go suite: PASS, 11 tests. Exact commands and evidence: .council/runs/20261125-query-history-01/reviews/final/judge-001/report.md. No claim of enforced browser network isolation; no production used.

Findings ARCH-001, C001-1, C001-2, C002-1 and P004-incomplete resolved by review; no open blockers. Cross-tab simultaneous localStorage writes may lose updates; pending responses may add after clearing, as documented in approved plan.

Counters: 6 plan revisions; 6 implementation revisions; 0 simplification retries; 0 final remediation cycles; 29 member invocations. S0 clean at 9ebc8c67e24c79392989eb533f98195370580feb; S0 snapshots and S1 checked; no staged changes or recorded developer work altered. No commits, pushes or deployments. Council artifacts are untracked; consider excluding .council locally from Git tracking. Developer retains final review and integration responsibility.
