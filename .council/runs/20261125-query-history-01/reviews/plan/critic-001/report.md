# Independent plan review

Run 20261125-query-history-01; invocation critic-001; normal review; plan v002 SHA-256 ec02bd41d2f51f83db9de85fb58e0d24252a9670203f7ae77704e91bc8e371f6; requirements plan/requirements-v002.md SHA-256 5ba578e4cd41e1704674cc92a76590e016ea281fa209f303950d7f6e0f752112. Supplied S0 and observed HEAD 9ebc8c67e24c79392989eb533f98195370580feb; no application modifications observed, Council metadata untracked. Ticket, clarification and baseline.json reviewed.

## Verdict: REJECTED
Planning readiness only. Two material design gaps prevent approval; no implementation or test execution is certified.

## Evidence
`web/templates/index.html` supplies the selector change form and manual POST, and initial theme storage access. `internal/templates/query.go` defines swapped manual and delayed automatic query controls. `internal/templates/connections.go` renders names and URL-bearing option values; `internal/database/connect.go` switches session current name on a separate request. `internal/database/create.go` sets current on creation. `internal/http/router.go` supplies connection and query endpoints. `web/static/scripts/tree.js` dispatches input; `theme.js` accesses storage directly. `README.md:35` documents HTTP or HTTPS hostname access and README:95-109 documents local HTTP. `tools/styles/compile.sh` compiles checked-in CSS. Baseline reports 11 Go tests passing; this review did not run them.

## Requirement coverage
SATISFIED means implementation and verification coverage at planning scope only.

| ID | Planned behavior / steps | Proposed verification | Outcome |
|---|---|---|---|
| R-001 | request capture P-003; suspension during switch | manual, auto, errors, switch race P-005 | UNSATISFIED: C001-2 |
| R-002 | name digest, swap and async isolation P-002 | switch/reload/race P-005 | UNSATISFIED: C001-1, C001-2 |
| R-003 | text preview and full text P-003 | multiline/timestamp P-005 | SATISFIED |
| R-004 | pending delayed input suppression P-003 | each action, subsequent input and actual HTMX P-005 | SATISFIED |
| R-005 | isolated delete/clear P-003 | duplicates, isolation and keyboard P-005 | SATISFIED |
| R-006 | 50 newest, serial queue P-003 | 51, duplicates, rapid submissions P-005 | SATISFIED |
| R-007 | digest and SQL/timestamp only, warning P-002/P-003 | storage sentinel and XSS P-005 | SATISFIED |
| R-008 | fail-open storage/theme P-003/P-004 | thrown operations and malformed storage P-005 | SATISFIED |
| R-009 | panel, focus, theme, responsive P-004 | keyboard, styles, narrow screen P-005 | SATISFIED |
| R-010 | preserve query and selector, Go suite P-001/P-005 | actual HTMX regression and canonical command P-005 | SATISFIED at planning scope |

## Blocking findings

### C001-1 — HIGH: history unavailable on ordinary remote HTTP
Plan design requires Web Crypto SHA-256 in secure contexts and explicitly disables history when unavailable. README:35 supports `http(s)://hostname:port/v1/...`; non-localhost HTTP is not a secure context and lacks `crypto.subtle` in browsers. Thus an otherwise functional supported deployment loses the core feature permanently, not merely when storage fails. Resolve by specifying a stable non-secret identifier strategy working on supported HTTP as well as HTTPS without persisting connection details, and verification over a non-localhost insecure origin (or an equivalently faithful capability-disabled browser case). Confirm history isolation, persistence, and secret constraints there. Applies R-002 and indirectly all history behavior.

### C001-2 — HIGH: submitted queries dropped during connection switch
Plan design explicitly suspends history association during connection POST until response replacement. The form in index.html posts independently of manual/Auto-Run requests; neither disables the editor, so a user can submit an actual query in that interval. The plan deliberately drops it, contradicting R-001's each submitted query, and cannot establish the selected server session identity during overlapping requests. Resolve with an executable strategy for submissions during switching that preserves query behavior and correctly associates each recorded request, or obtain an authorized product resolution if correctness cannot be assured without changing contracts. Specify and verify delayed switch response, query before/during switch, failure, overlapping switches and the effective session identity without a live database. Applies R-001/R-002.

## Prior finding closure
ARCH-001: resolved at planning scope. Clarification explicitly authorizes exact HTMX 2.0.1 verification-only acquisition; P-001/P-005 require provenance, trusted identity comparison and real-HTMX tests for all three pending actions. Actual acquisition and tests remain downstream mandatory prerequisites, not accomplished by this review.

## Scope and verification conclusion
Bounded application paths, no new server persistence, protected S0, canonical Go suite, and real HTMX fixture are otherwise credible. Asset acquisition and trusted hash verification remain prerequisites for execution; no asset or test was run here. The loopback-only fixture must additionally address insecure-origin capability per C001-1 and effective-session races per C001-2. No nonblocking observations. No external blockers to revising the plan identified.
