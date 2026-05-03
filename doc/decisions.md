# Decisions

## Why record decisions

Write down key development decisions while the context is fresh. A short note today can save hours later by explaining what was chosen, what was rejected, and why the trade-off made sense at the time.

## Guidance

Use a lightweight Architecture Decision Record (ADR) style:

* Record decisions that affect architecture, data flow, public APIs, dependencies, deployment, security, or long-term maintenance.
* Write the decision when it is made, not after the context has faded.
* Prefer short entries that explain the context, decision, alternatives, and consequences.
* Include enough reasoning for a future maintainer to understand the trade-off.
* Do not document every small implementation detail; focus on choices that would be costly or confusing to rediscover.
* Update or supersede earlier decisions instead of silently rewriting history.

## Entry template

```markdown
### YYYY-MM-DD: Decision title

**Status:** Proposed | Accepted | Superseded

**Context:** What problem, constraint, or trade-off led to this decision?

**Decision:** What was chosen?

**Alternatives considered:** What other options were rejected, and why?

**Consequences:** What becomes easier, harder, riskier, or more constrained?
```

## Actual decisions

### 2026-05-03: Use `go-telegram/bot` as the initial Telegram library

**Status:** Accepted

**Context:** The project needs a Go library for building a Telegram bot that can run on a Raspberry Pi Zero. At this stage, the main goal is to start with a small, understandable implementation that supports learning, rapid iteration, and low operational complexity. The library comparison in `doc/notes/go_telegram_bot_libraries.md` showed several viable options, but they differ mainly in abstraction level, completeness, and ergonomics rather than core capability.

**Decision:** Start implementation with `go-telegram/bot` as the primary Telegram Bot API library.

**Alternatives considered:** `Telego` offers broader API coverage and more control, but it is a less focused starting point for an early prototype. `go-telegram-bot-api` is popular and simple, but the notes raised a maintenance concern that makes it a weaker default choice. Using raw `net/http` would maximize transparency and minimize dependencies, but it would also add avoidable boilerplate before the project has validated its basic behavior.

**Consequences:** This makes it easier to build an initial working bot quickly with a clean, modern handler-based API and minimal setup. It also keeps the early codebase small, which fits the Raspberry Pi Zero constraint and the learning goal. The trade-off is that a future migration may be needed if the project outgrows the library's abstraction or requires lower-level control or broader API coverage.
