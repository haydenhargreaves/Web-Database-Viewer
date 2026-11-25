# Implementation report — engineer-006

Run 20261125-query-history-01, revision invocation engineer-006; status PARTIAL. Approved plan v007 SHA-256 27db66662bacaadb9bd9cd34e3e836e5cc19a42c010de4837cb7c06628d55022, associated requirements plan/requirements-v007.md SHA-256 7d96c2bc60e939721674e5ed50a6b2470c0447a91fec261cfe14d180db7d1c8f; approval reviews/plan/critic-004/report.md. Delegated HEAD and observed final HEAD 9ebc8c67e24c79392989eb533f98195370580feb, branch master. HEAD alone is not the working-tree identity.

## Scope and provenance

The ten application and test file hashes matched invocations/engineer-006/assignment.json before edits. S0 was clean according to baseline and previous reports; engineer-001 through engineer-004 document earlier Council work and no overlapping original developer changes. No staged application changes were observed. The four modified tracked files relative to S0 are internal/database/query.go, internal/templates/connections.go, web/static/scripts/theme.js, web/templates/index.html. Six Council-created files are internal/database/query_history_test.go, internal/templates/connections_history_test.go, web/static/scripts/query-history.js, web/tests/query-history-server.js, web/tests/query-history.html and web/tests/query-history.js. This invocation changed only web/tests/query-history.js, final SHA-256 38bfbf7180a17c50f363a7a561fa5a9704b95651b6b2eaa40c4705e3d09fb926. All existing fixture assertions were retained. No product changes, rollback, index changes, or intentional historical artifact writes. Council metadata remains untracked. The assigned invocation directory was absent, checked for symlink ancestors, created and writable before edits.

New assertions verify exact multiline SQL load without POST and pending timer suppression plus later genuine input after both select and remove. A storage event refreshes the fixture after its prior synthetic getter failure so these checks observe the rendered list. No product defect was established. The fixture's existing real HTMX asset was read-only and its SHA-256 matched the plan; browser profile was created and removed within the assigned evidence directory. The synthetic tests do not replace actual page interaction.

## Acceptance and inspection matrix

| ID | Evidence and remaining work |
|---|---|
| A1 | PASS real HTMX 2.0.1 fixture tests manual/auto exact SQL, request counts, whitespace without entry, error, missing/invalid ID and no retry. |
| A2 | PASS fixture select/remove/clear cancel pending auto POST; genuine input after each sends exactly one; select loads without execution. |
| A3 | NOT VERIFIED actual selector outerHTML swap, switch/back/reload and current-only clear with two rendered connections. Fixture tests different-ID response race and isolated storage, but is not the required selector-path exercise. |
| A4 | PARTIAL Go identity/handler table tests pass names and membership; browser special-name parsing and switching not exercised. |
| A5 | PASS fixture 51/50 retention, distinct duplicates and single removal, same-ID reversed responses, literal multiline Unicode, readable timestamp, exact multiline load without POST. |
| A6 | NOT VERIFIED actual manual/auto editor swaps and tree-triggered input; A2 uses a synthetic combined control fixture rather than real swap. |
| A7 | NOT VERIFIED actual styled page keyboard/focus, themes and 320px viewport. Warning text exists in index.html:210; not a substitute for UI exercise. |
| A8 | PARTIAL fixture verifies malformed JSON, getter and quota failures allow query POST. Actual page/theme initialization and toggle under blocked storage remain untested. |
| A9 | PASS canonical Go suite final contents, 13 top-level tests and eight subtests, no skips, four no-test packages. S0 baseline reported 11 top-level tests; no exception asserted. |
| I1 | PARTIAL: QueryCurrent obtains execution parameters and name in one queryConnectionIdentity call; XHR WeakMap and validated response header/status attribute requests with no selector fallback/retry. Fixture covers error and invalid/missing identity; abort/timeout/lost lifecycle not fully established. |
| I2 | PARTIAL shared exact-name SHA-256 helper and escaped rendering inspected; Go tests passed; sentinel and browser selector check remain. |
| I3 | PARTIAL read normalizes SQL/timestamp, renderer uses textContent, fixture tests literal SQL and storage fields; browser warning visibility remains A7. |
| I4 | PARTIAL generation invalidation and pending tests inspected; actual editor swap/tree path remains A6. |
| I5 | PARTIAL no internal/http tracked diff; code writes SQL/timestamp and identifier is digest; comprehensive sentinel/route inspection remains. |
| I6 | PARTIAL read/write/remove and both theme initialization/toggle storage calls are guarded; runtime page/theme test remains A8. |

## Requirements mapping

All requirements remain UNVERIFIED for complete acceptance. R-001: XHR correlation in web/static/scripts/query-history.js and A1/A2, remaining I1. R-002: QueryCurrent/queryConnectionIdentity, templates.QueryHistoryID, ConnectionsList and race fixture, remaining A3/A4/I1/I2. R-003: render/textContent and A5, remaining A7/I3. R-004: input generation suppression and A2, remaining A6/I4. R-005: history remove/clear handlers and fixture, remaining A3/A6/A7. R-006: bounded sorted storage and A5, complete associated fixture evidence but overall requirements remain unverified pending holistic verification. R-007: digest, SQL/timestamp-only serialization, warning markup; remaining A7/I2/I3/I5. R-008: guarded history/theme storage and synthetic tests; remaining A8/I6. R-009: panel markup; remaining A7. R-010: Go suite and request fixture; remaining A3/A6/I1/I4/I5. These are implementation evidence, not independent acceptance.

## Verification

All commands run at project root, after final edit unless noted. Evidence in evidence/ directory.

- `GOTOOLCHAIN=local GOPROXY=off GOSUMDB=off go test -mod=vendor -count=1 -v ./...`: exit 0, 13 top-level tests and eight subtests, four packages no tests, no skips; go-test-final.log.
- `QUERY_HISTORY_EVIDENCE_DIR="$PWD/.council/runs/20261125-query-history-01/implementation/engineer-006/evidence" node web/tests/query-history-server.js`: exit 0, PASS 31 assertions with real Chromium, synthetic loopback and HTMX 2.0.1; browser-final.log.
- An intermediate browser run failed at 27 assertions because the test list had been cleared by the deliberately throwing storage getter. Added an explicit storage-event refresh after restoring the getter; final run passed. Intermediate output was superseded in browser-final.log; failure is recorded here. An earlier initial fixture run passed 26 assertions, evidence/browser.log.
- `git diff --check`: exit 0; only tracked diff is checked. No application changes occurred after final checks.
- HTMX asset hash verified as 6d4aaa4b0d3e8b4c91f8d97b92a361a19b1bd4544dea3f668fdc3e62a63995df. Browser asserts version 2.0.1. No dependency installed or external service used.

## Findings and handoff

ARCH-001, C001-1, C001-2 and C002-1 were resolved at planning scope only by critic-004; this report does not claim independent closure. Mandatory A3, A4 browser portion, A6, A7, A8 page/theme portion and I1–I6 remain incomplete, so this is PARTIAL, not acceptance or a waiver. Continue with approved in-scope fixture or documented browser inspection and rerun required final checks. Non-atomic cross-tab updates, pending response after clear and duplicate-URL ambiguity remain documented design limitations. No material design deviation was introduced.
