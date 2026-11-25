# Requirements v007 — run 20261125-query-history-01

Sources: original `../ticket.md`; user decisions and proportional reassessment in `../clarifications.md`. S0 HEAD and supplied code identity `9ebc8c67e24c79392989eb533f98195370580feb`; working-tree identity in `../invocations/architect-007/assignment.json`. Complete revision of v006: stable IDs retained, no requirement retired. Objective: review and reuse recent SQL per connection in this browser. Non-goals: named/pinned queries, folders, search, import/export, server history or synchronization, editor replacement and unrelated fixes. All requirements mandatory; acceptance remains UNVERIFIED.

| ID | Source | Requirement and observable acceptance | Verification |
|---|---|---|---|
| R-001 | Ticket; lost-response clarification | Each non-whitespace manual or Auto-Run submission with an established effective connection is recorded once with exact SQL, including SQL errors; whitespace creates no entry. A lost or unidentifiable response may be omitted; never guess or retry SQL. | P-004 A1, A2, I1 |
| R-002 | Ticket; effective-ID clarification | Switching connections displays separate history and switching back restores it; submissions are attributed to the effective request connection, including concurrent switches. Identifier is stable and non-secret; real accepted names including empty are distinct from placeholder/missing member. | P-004 A3, A4, I1, I2 |
| R-003 | Ticket | Newest first, readable timestamps and query text, full multiline SQL available and loadable. | P-004 A5, I3 |
| R-004 | Ticket | Selecting loads editor without executing; select, remove and clear must not trigger pending Auto-Run; subsequent genuine input works. | P-004 A2, A6, I4 |
| R-005 | Ticket | Remove individual entries and clear only current connection using accessible controls. | P-004 A3, A6, A7 |
| R-006 | Ticket | Per-connection history retains at most newest 50 entries; repeated submissions remain distinct. | P-004 A5 |
| R-007 | Ticket; security consideration | Browser localStorage history contains SQL and timestamp only, using a stable non-secret connection identifier; no URL, credential, result or other connection metadata. Clearly warn SQL may be sensitive and stored in browser, with obvious clear action. | P-004 A7, I2, I3, I5 |
| R-008 | Ticket | Unavailable or malformed storage cannot block query execution or page usability; history may be unavailable. | P-004 A8, I6 |
| R-009 | Ticket | Panel near editor/results works in light and dark and narrow viewport; keyboard accessible with meaningful labels and visible focus. | P-004 A7 |
| R-010 | Ticket and necessary existing invariants; effective-ID clarification | Preserve query execution, session, selector, toggles and tree behavior; only backward-compatible non-secret identity in existing responses/selector; no history endpoint or server persistence; existing Go tests pass. | P-004 A1, A3, A9, I1, I4, I5 |

Compatibility: SQL is stored verbatim even if it itself contains secrets; restriction concerns additional metadata. Per-response identity is authoritative, not an assumed selector value. Same exact connection name shares an origin-scoped history; digest is guessable, not encryption. Simultaneous cross-tab localStorage mutations may lose updates; pending response after clear can add later. Do not alter selector values or URL-based switching to resolve pre-existing duplicate-URL ambiguity. Protect existing tests and Council/developer changes. No destructive operations, new dependency or service.

Ordinary assumptions: real browser, real HTMX 2.0.1 and local synthetic data suffice; verification prerequisites and safe acquisition are described in plan. No unresolved product question. Mandatory checks remain required even if tooling unavailable. No implementation acceptance is claimed.
