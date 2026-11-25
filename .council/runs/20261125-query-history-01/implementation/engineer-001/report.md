# Implementation report — engineer-001

- Run: 20261125-query-history-01; invocation: engineer-001; stage IMPLEMENTING, normal mode.
- Status: PARTIAL. Implementation and initial tests exist, but the approved real-browser acceptance matrix is substantially incomplete. Do not treat this as implementation acceptance.
- Approved plan: `plan/v006.md`, SHA-256 `b545fb967d71800cb9da12a3b694be70061d22fc461cd4b5bbaf9354c7f3a486`; matching `plan/requirements-v006.md`, SHA-256 `5d0f39e23c2efa076c8e762fb99498c5a33fd87210e462cbdfb4e5cb338dbbf1`; approval `reviews/plan/critic-003/report.md`. Paths in this paragraph are relative to `.council/runs/20261125-query-history-01/`.
- Delegated input and final HEAD: `9ebc8c67e24c79392989eb533f98195370580feb`, branch `master`. Final feature identity is the unstaged working tree, not a commit. At start there were no staged or unstaged application changes and only untracked Council metadata. All scope files and modes matched the S0 manifest, including absence of new files. Existing work and index preserved; no rollback performed. Initial Python check could not run because python3 was unavailable; repeated successfully using Node. Scope files have recoverable S0 snapshots in the Council manifest.

## Changes relative to S0

All feature changes are from this invocation; no original developer application changes were observed. Existing files modified: `internal/database/query.go` (request-session identity and response header), `internal/templates/connections.go` (shared SHA-256 identity, escaped name and option identity), `web/templates/index.html` (panel, script, guarded early theme storage), `web/static/scripts/theme.js` (guarded theme storage). New files: `internal/database/query_history_test.go`, `internal/templates/connections_history_test.go`, `web/static/scripts/query-history.js`, `web/tests/query-history.html`, `web/tests/query-history.js`, `web/tests/query-history-server.js`. CSS was not changed. No history endpoint, server history persistence, dependency or schema added. The final index is unchanged; these files remain unstaged or untracked. The downloaded test asset and Go log are separate invocation-specific verification artifacts, not application dependencies. No historical artifacts intentionally overwritten.

Implementation choices: effective identity is SHA-256 of exact current map key, including empty actual members; synthetic missing-current and placeholder options have no ID. SQL is captured per HTMX XHR and stored only after a valid response header; display uses the selected option. Local storage contains SQL and timestamp only. Timestamps are made monotonically increasing within the tab (up to one millisecond beyond the previous timestamp per rapid submission) so same-millisecond request order survives response reordering without persisting sequence metadata; this minor choice should be reviewed against the plan's request-time/sequence wording. Rendering displays complete SQL instead of shortening it. Clear and removal reread storage or remove the current key. Delayed-input invalidation is applied on history actions; genuine subsequent input increments its generation. Theme access and storage failures are guarded. As planned, cross-tab storage writes are not atomic, and an in-flight response can add after clear.

## Requirement mapping

All requirement statuses remain UNVERIFIED because the mandatory P-004 acceptance matrix has not been completed, notwithstanding the partial evidence below.

| ID | Code and test mapping / remaining verification |
|---|---|
| R-001 | XHR capture and completion in `query-history.js`; basic real-HTMX manual/auto check in `web/tests/query-history.js`. SQL-error, whitespace, lost/invalid response, abort, timeout and reordered-response assertions missing. |
| R-002 | `QueryHistoryID`, `ConnectionsList`, `queryConnectionIdentity` and `QueryCurrent`; Go tests cover exact identity for representative names, empty member and missing current. Browser switch/reload, creation, duplicate URL, stale selector, shared-cookie races, and crypto-disabled cases missing. |
| R-003 | `render` emits full literal SQL and timestamp; browser check records text, but multiline, accessibility and reversed-response ordering tests missing. |
| R-004 | input generation and `beforeRequest` suppression in `query-history.js`; fixture checks select after pending input and subsequent input. Remove, clear, editor swap, actual request counts, manual/tree regression checks remain. |
| R-005 | per-entry remove and clear handlers; duplicate/isolation/keyboard/pending-response tests missing. |
| R-006 | storage writes capped at 50 and sorted by submission timestamp; 51 entries, duplicates, concurrent and reversed-completion checks missing. |
| R-007 | digest-only keys, SQL/timestamp serialization, warning and textContent rendering; sentinel, stored-data inspection, XSS, cross-browser identity tests missing. |
| R-008 | guarded localStorage operations and theme accesses; exception injection, malformed data, quota and query usability tests missing. |
| R-009 | panel uses existing light/dark utility classes and focus ring; browser computed styles, 320px layout, keyboard and focus tests missing. |
| R-010 | additive response/option identifiers; Go suite passed and fixture makes real HTMX requests. Full server contract, selector value/URL corner cases, toggles, tree and exact POST count tests missing. |

## Verification

- Canonical root `GOTOOLCHAIN=local GOPROXY=off GOSUMDB=off go test -mod=vendor -count=1 -v ./...`: PASS exit 0, 13 top-level tests, zero reported skips, four packages without tests; Go 1.26.8. Detailed output: `.council/runs/20261125-query-history-01/verification/engineer-001/go-test.log`. S0 reported 11 top-level tests and same Go version; no isolated baseline exceptions.
- `node --check` on application history script and the two fixture JS files: PASS exit 0. `git diff --check`: PASS exit 0.
- `node web/tests/query-history-server.js`: PASS exit 0, four assertions using Chromium 155 headless, ephemeral loopback and HTMX 2.0.1; tests version, manual storage, selection preventing a pending auto submission as observed in storage, and subsequent input recording. This is a small subset, not the approved full browser suite. An earlier execution failed after one assertion because test code clicked before HTMX initialized; waiting for page load fixed the fixture and the final run passed. Virtual-time fixture has not been independently validated against real-time CDP timing and does not assert network counts.
- HTMX asset fetched from `https://unpkg.com/htmx.org@2.0.1/dist/htmx.min.js` on 2026-11-25 into `.council/runs/20261125-query-history-01/verification/engineer-001/htmx.min.js`: SHA-256 `6d4aaa4b0d3e8b4c91f8d97b92a361a19b1bd4544dea3f668fdc3e62a63995df`; SHA-384 base64 `QWGpdj554B4ETpJJC9z+ZHJcA/i59TyjxEPXiiUgN2WmTyV5OEZWCD6gQhgkdpB/`. Independently fetched official tagged source `https://raw.githubusercontent.com/bigskysoftware/htmx/v2.0.1/dist/htmx.min.js` and jsDelivr package URL and compared SHA-256; both matched. Version asserted by fixture. A guessed v2 documentation hostname failed DNS; openssl was unavailable, so Node crypto computed SHA-384. No production resources accessed.
- Full P-004 browser matrix: NOT RUN; fixtures and assertions not yet implemented, not an environment/tooling claim. Essential remaining work includes real network request counts/timer validation and all cases identified above. No tests were weakened or skipped.

## Findings, blockers, recovery

Prior plan-review findings ARCH-001, C001-1, C001-2 and C002-1 were resolved only at planning scope by Critic, not independently closed here. Their implementation-level browser verification remains incomplete. No new Council finding IDs assigned. No material scope/design expansion was intentionally made. No safe-recovery issue or rollback; partial changes intentionally retained for continuation. A fresh Engineer continuation must complete the matrix, remediate any defects it exposes within approved scope, and rerun canonical and browser checks on final files before handing to Inspector. This report is not approval.
