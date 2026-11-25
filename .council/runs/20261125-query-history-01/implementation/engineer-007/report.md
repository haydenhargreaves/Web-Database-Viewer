# Implementation report — engineer-007

Run 20261125-query-history-01, revision invocation engineer-007. Status PARTIAL: final canonical and existing fixture rerun, and source inspection advanced, but several mandatory actual-page browser checks remain unexecuted. Plan v007 SHA-256 27db66662bacaadb9bd9cd34e3e836e5cc19a42c010de4837cb7c06628d55022; matching requirements plan/requirements-v007.md SHA-256 7d96c2bc60e939721674e5ed50a6b2470c0447a91fec261cfe14d180db7d1c8f; approved in reviews/plan/critic-004/report.md. Delegated and observed final HEAD 9ebc8c67e24c79392989eb533f98195370580feb, branch master; HEAD is not full working-tree identity. Ten application files matched assignment.json SHA-256 at invocation start and were not edited in this invocation. No original developer application changes are identified in S0 or prior reports; no staged application changes. S0 was clean. The S0 scope manifest captures originals and absence for new files. No rollback, scope expansion, dependency change, or intended historical artifact modification.

## Cumulative feature relative to S0

Council feature changes, not this invocation: modified internal/database/query.go (single session resolution and response identity), internal/templates/connections.go (exact-name digest and escaped option text), web/templates/index.html (history panel, startup storage guard), web/static/scripts/theme.js (storage guards); added internal/database/query_history_test.go, internal/templates/connections_history_test.go, web/static/scripts/query-history.js, web/tests/query-history-server.js, web/tests/query-history.html, web/tests/query-history.js. All are mode 0644, unstaged or untracked. CSS and router unchanged. Start and final hashes match the ten values in invocations/engineer-007/assignment.json. Existing synthetic fixture already supports explicitly assigned profile directory, so no fixture path edit was needed. Plan permits retaining existing implementation and fixing only demonstrated defects; none established in this invocation. No application/test changes made. Index unchanged by this invocation. Council metadata is untracked, including new invocation artifacts; prior metadata retained.

## Requirement mapping and acceptance

All R-001–R-010 remain UNVERIFIED overall: mandatory browser acceptance incomplete; statuses below refer only to evidence, not independent acceptance.

- R-001: query-history.js beforeRequest/afterRequest correlates XHR; fixture A1/A2 passes manual/auto, empty, SQL error, invalid/missing header, pending input and counts; I1 source inspection below.
- R-002: QueryCurrent/queryConnectionIdentity and templates.QueryHistoryID/ConnectionsList plus Go tests and fixture race. A3 actual selector swap/switch/back/reload/current-only clear and A4 nontrivial-name browser parsing remain unexecuted.
- R-003: render uses textContent, time display and full SQL; A5 fixture passes multiline, timestamps, ordering and load. A7 actual styled page visibility remains unexecuted.
- R-004: generation guard and fixture A2 pass pending select/remove/clear and later genuine input. A6 actual manual/auto swaps, tree input and manual run remain unexecuted.
- R-005: render and delegated clear handler; fixture tests remove/clear, but A3 isolation and A7 keyboard interaction remain unexecuted.
- R-006: read/write bounded to 50, sorted by submission time; A5 fixture passes 51 entries, repeated SQL, one removal and same-ID reversed completion.
- R-007: digest attribute/header, SQL/timestamp storage and warning; fixture checks serialized fields and literal XSS text; A7 warning visibility, A3 selector and I5 synthetic sentinel inspection remain incomplete.
- R-008: guarded read/write/remove and startup/toggle theme calls; fixture tests malformed JSON and storage getter/quota failure without blocking POST; A8 actual page/theme initialization and toggling under failure remain unexecuted.
- R-009: panel and focus classes exist; A7 real styles, themes, 320px width, keyboard/focus and activation remain unexecuted.
- R-010: additive identifier, no tracked router changes, canonical suite and real HTMX fixture pass; A3 and A6 remain unexecuted.

## Inspections

I1: query.go:21–25,169–185 obtains conn/driver/name/member from one session snapshot and sets header for actual members; SQL processing continues with those parameters. query-history.js:104–132 captures SQL in a WeakMap keyed by actual XHR, deletes on afterRequest and only records on truthy status and valid 64-hex response header; there is no selector fallback or retry in that code. Network loss/status 0 therefore cannot record when afterRequest runs; abort/timeout handling inspected but not independently simulated. Failed selector requests reset switching if HTMX reports failed with the target ID; actual failed switch was not exercised.

I2: templates/connections.go:9–12,21–48 hashes exact name bytes with SHA-256, emits lowercase hex only for members and escapes displayed text; QueryCurrent invokes the same helper. Go tests compare representative names, real empty member and placeholder/missing membership; browser nontrivial-name path remains outstanding. No Web Crypto call in history script. Fixture uses loopback HTTP.

I3: query-history.js:17–25,46–62 normalizes read values to sql/timestamp and renders both as text; write stores capped entries. Fixture asserts SQL/timestamp-only fields, literal multiline Unicode and readable timestamp. index.html:208–213 contains sensitivity warning and clear. Actual visual verification remains A7.

I4: query-history.js:33,63–66,73–80,88–109 invalidates current generation on history actions and suppresses pending sql-element request while later input increments generation; tree.js:34–49,69–71 dispatches bubbling input; template query.go:7–40 defines manual and auto swaps. Existing fixture supports timer assertions, but real swaps and tree-triggered interaction still require A6.

I5: identifier key is fixed prefix plus digest; response header/attribute digest originates from name only. query-history.js serializes sql/timestamp, not URL, driver or result; SQL itself remains verbatim and may contain sensitive text. git diff -- internal/http empty; no new history endpoint observed. Explicit synthetic URL/password/driver/result sentinel exercise not completed, so I5 remains partial.

I6: query-history.js:17–30,95–100 guards getter, JSON parse, write and remove; non-array, non-string SQL/timestamp and invalid date read as empty. index.html:8–15 and theme.js:16–29 guard initialization, getItem and setItem; toggle applies theme after guarded write. Existing fixture tests throwing getter, setItem and malformed JSON with real query; page/theme runtime remains A8. These inspections do not replace browser acceptance.

## Verification

All commands at project root on final contents; no application edits occurred. Logs under implementation/engineer-007/evidence/.

- `GOTOOLCHAIN=local GOPROXY=off GOSUMDB=off go test -mod=vendor -count=1 -v ./...`: PASS exit 0, 13 top-level tests, eight subtests, no skips, four no-test packages; evidence/go-test.log. S0 baseline PASS 11 top-level tests; baseline reports Go 1.26.8 versus go.mod 1.25.7. No baseline exception asserted.
- `QUERY_HISTORY_EVIDENCE_DIR="$PWD/.council/runs/20261125-query-history-01/implementation/engineer-007/evidence" node web/tests/query-history-server.js`: PASS exit 0, 31 real Chromium/HTMX fixture assertions, real wall-clock, synthetic loopback; evidence/browser.log. This fixture does not establish actual selector, page/theme or editor swaps. Profile was created and removed by the existing fixture under the assigned evidence directory.
- HTMX asset SHA-256 6d4aaa4b0d3e8b4c91f8d97b92a361a19b1bd4544dea3f668fdc3e62a63995df, matches v007; fixture checks version 2.0.1. Provenance documented in engineer-001 report.
- `git diff --check`: PASS exit 0, tracked changes only.
- A3, A4 browser spot-check, A6, A7 and A8 page/theme runtime: NOT RUN; additional representative page harness or documented real-browser interaction still needed. No environmental blocker claimed. I5 explicit sentinel exercise remains incomplete. These checks were not waived or replaced by fixture output.

## Findings and handoff

ARCH-001, C001-1, C001-2, C002-1 were resolved at planning scope by critic-004 only; no independent implementation closure claimed. No new findings or product changes. Duplicate-URL switching ambiguity, non-atomic cross-tab writes and late response after clear are known approved limitations. Remaining work is the missing actual-browser checks and sentinel inspection, fix only demonstrated in-scope defects, then rerun canonical and fixture on final files. No plan deviation, rollback or destructive operation. Report/logs are invocation-specific, existing historical artifacts and index were not intentionally changed.
