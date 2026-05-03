# zero_control

`zero_control` is a small, deliberate experiment at the intersection of software, hardware, and collaboration. The project aims to build a Telegram bot in Go that is lightweight enough to run on a Raspberry Pi Zero, while treating the process itself as part of the work: a practical way to sharpen Go skills, understand the limits and quirks of the Pi Zero, and test how useful Codex can be as a real development partner.

At its core, this project is about constraint as a design tool. A tiny machine, a focused runtime, and a simple interface create the right conditions for learning by building something concrete. `zero_control` is not just about shipping a bot; it is about using a modest project to explore how clear code, limited hardware, and AI-assisted development behave when they are forced to work together.

## Project structure

```text
zero_control/
|-- cmd/
|   `-- zero_control/
|       `-- main.go
|-- doc/
|   |-- notes/
|   |   `-- go_telegram_bot_libraries.md
|   |-- decisions.md
|   |-- requirements.md
|   `-- user_stories.md
|-- internal/
|   |-- app/
|   |-- bot/
|   |-- config/
|   |-- device/
|   |-- logging/
|   `-- service/
|-- tests/
|   `-- integration/
|-- .env.sample
|-- .gitignore
|-- AGENTS.md
|-- go.mod
|-- LICENSE
|-- README.md
`-- TODO.md
```

- `cmd/zero_control/main.go` is the application entry point and keeps startup wiring out of the business logic.
- `internal/config` loads and validates runtime configuration such as the bot token.
- `internal/bot` holds Telegram-facing handlers, routing, and middleware.
- `internal/service` contains application use cases and orchestration logic.
- `internal/device` is reserved for Raspberry Pi and OS-facing operations.
- `internal/app` composes the application and coordinates startup.
- `internal/logging` centralizes logger construction.
- `tests/integration` is reserved for higher-level end-to-end checks once the bot starts doing real work.
- `doc/requirements.md` captures the functional and technical requirements as they become clearer.
- `doc/user_stories.md` records intended user-facing behavior and usage scenarios.
- `doc/decisions.md` documents implementation and architecture decisions, along with the reasoning behind them.
- `doc/notes/` stores research notes that inform technical choices.
- `.env.sample` shows the expected local configuration shape without exposing secrets.
- `README.md` contains the project overview and acts as the main entry point.
- `TODO.md` tracks short-term tasks and open work.
- `AGENTS.md` defines working rules and engineering expectations for AI-assisted collaboration in this repository.
