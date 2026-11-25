# Implementation report — engineer-003

Run `20261125-query-history-01`, invocation `engineer-003`, revision mode. **PARTIAL**: the mandatory P-004 matrix remains incomplete; no independent acceptance is claimed. Delegated input identity and observed final HEAD: `9ebc8c67e24c79392989eb533f98195370580feb`, branch `master`; HEAD alone is not the feature working-tree identity. Exact plan `plan/v006.md` SHA-256 `b545fb967d71800cb9da12a3b694be70061d22fc461cd4b5bbaf9354c7f3a486`; matching `plan/requirements-v006.md` SHA-256 `5d0f39e23c2efa076c8e762fb99498c5a33fd87210e462cbdfb4e5cb338dbbf1`; approval `reviews/plan/critic-003/report.md`. These paths are run-relative.

## Scope, ownership and changes

S0 manifest and snapshots were inspected and snapshot hashes verified. At invocation start all ten feature file hashes matched `invocations/engineer-003/assignment.json`, all ten files mode 0644; index empty. S0 and previous reports identify no original developer application changes. Four prior tracked modifications: `internal/database/query.go`, `internal/templates/connections.go`, `web/templates/index.html`, `web/static/scripts/theme.js`; six prior additions: `internal/database/query_history_test.go`, `internal/templates/connections_history_test.go`, `web/static/scripts/query-history.js`, `web/tests/query-history.html`, `web/tests/query-history.js`, `web/tests/query-history-server.js`. CSS unchanged. Prior changes and design are detailed in `implementation/engineer-001/report.md` and `implementation/engineer-002/report.md`.

This invocation modified only the two previously added test scripts. `query-history-server.js` uses this invocation's verification directory for the ephemeral browser profile rather than writing to the prior invocation's historical directory; it records synthetic payloads, serves an invalid identifier and provides delayed, differently identified synthetic responses. `query-history.js` adds real HTMX assertions for exact payloads, invalid identity, 51 submissions, duplicate preservation and one-at-a-time removal, SQL/timestamp-only stored shape, literal multiline Unicode/XSS rendering and timestamp, malformed storage recovery, reverse response attribution and a pending response after clear. No application code changed. Final SHA-256: test JS `0e9dbde54bdef8b86a059096b45be96dede4a3b406ae945bcb20ebacd8ac0276`, server `137ce71ad614b6654619ee90e2fdbd77e9e3dee0e78390da78b948b4c0188534`.

No rollback, no staged files and no intentional historical artifact changes. The ten feature files remain unstaged or untracked, together with Council metadata. S0 snapshots and invocation assignment provide the recoverable representations. The browser removes only its own temporary profile. No new dependency, endpoint or architectural change. Prior engineer's monotonic timestamps may be up to one millisecond later than actual request time per closely spaced submission; noted previously for independent review. Simultaneous cross-tab localStorage mutations can lose updates, as approved.

## Requirement mapping and remaining evidence

Each remains **UNVERIFIED** against the complete mandatory P-004 acceptance matrix, despite the partial evidence. Details of prior Go and browser tests are in the preceding implementation reports.

| ID | Implementation and evidence | Remaining material cases |
|---|---|---|
| R-001 | HTMX XHR listeners in `query-history.js`; actual HTMX manual/auto, SQL error, whitespace, missing/invalid identity, exact payload, reverse-response attribution and request counts. | Abort, timeout and lost response; complete manual/auto cases. |
| R-002 | Go `QueryHistoryID`, `ConnectionsList`, `queryConnectionIdentity`, `QueryCurrent`; prior Go tests and synthetic response race. | Browser server-rendered special names, switch/back/reload for each, creation/outerHTML, duplicate URLs, delayed/failed/overlapping switches, shared-cookie stale selector, crypto-disabled HTTP isolation and persistence. |
| R-003 | `render` uses literal text and time; browser multiline Unicode/XSS and readable time assertions. | Long preview/full accessible text and reversed responses ordered within the same identifier. |
| R-004 | generation invalidation; real-timer load/remove/clear suppress pending POST and later input executes. | Repeat after editor swap, tree input and full manual regression. |
| R-005 | Per-entry remove and clear; duplicate removal and pending response after clear, other synthetic connection retained. | Full keyboard and separately switched connection isolation. |
| R-006 | Capped/sorted storage, actual 51 submissions and duplicate test. | Reversed completion ordering same identifier, sequential and overlapping two-tab operations, foreign keys. |
| R-007 | Server digest, limited serialized fields, literal SQL rendering and prior Go identity tests. | Full synthetic URL/password/driver/result sentinels in headers, attributes, keys and values, warning and browser/server digest agreement. |
| R-008 | Guarded history/theme storage; malformed JSON followed by successful request. | Throwing getter/getItem/setItem/removeItem, quota, non-array/invalid data and timestamps, theme/page usability. |
| R-009 | Existing themed panel and accessible control implementation. | Real keyboard Tab/Enter/Space, focus, both computed themes and 320px layout. |
| R-010 | Additive identifier, previous Go handler/template tests, canonical suite and actual HTMX payload/count tests. | Full selector/form/duplicate URL contracts, real toggle/tree inputs, other P-004 compatibility cases. |

## Verification on final application/test contents

- `GOTOOLCHAIN=local GOPROXY=off GOSUMDB=off go test -mod=vendor -count=1 -v ./...` at project root: PASS exit 0; 13 top-level passing tests, zero skips, four packages without tests. Sanitized full log `verification/engineer-003/go-test.log`. S0: 11 top-level passing tests; no approved baseline exception.
- `node --check web/tests/query-history-server.js && node --check web/tests/query-history.js && node --check web/static/scripts/query-history.js && git diff --check && node web/tests/query-history-server.js` at root: PASS exit 0, 23 assertions. Browser uses isolated local loopback, wall-clock Chromium and actual HTMX 2.0.1; fixture checks version. The fixture is still a small subset of the required matrix, not a substitute for it. An earlier intermediate run passed 21 assertions; final adds two race assertions.
- Asset read only: `verification/engineer-001/htmx.min.js` SHA-256 `6d4aaa4b0d3e8b4c91f8d97b92a361a19b1bd4544dea3f668fdc3e62a63995df`; official-source comparisons and acquisition provenance recorded in engineer-001 report. No network acquisition performed here.
- Full P-004: NOT RUN in full; assertions and fixtures for the cases listed above are not implemented. No tests were skipped or weakened to obtain the partial passing results. No independently assigned implementation findings; prior planning findings ARCH-001, C001-1, C001-2, C002-1 remain resolved at planning scope only. No new feature defect conclusively established in this invocation.

## Continuation

Complete each remaining browser/Go matrix case within approved scope, remediate any demonstrated defects, then rerun canonical and full browser checks on final files. The partial suite passing does not establish all mandatory acceptance. No material plan deviation is proposed; no external blocker identified for bounded continuation.
