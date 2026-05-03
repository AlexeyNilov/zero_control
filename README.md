# zero_control

`zero_control` is a small, deliberate experiment at the intersection of software, hardware, and collaboration. The project aims to build a Telegram bot in Go that is lightweight enough to run on a Raspberry Pi Zero, while treating the process itself as part of the work: a practical way to sharpen Go skills, understand the limits and quirks of the Pi Zero, and test how useful Codex can be as a real development partner.

At its core, this project is about constraint as a design tool. A tiny machine, a focused runtime, and a simple interface create the right conditions for learning by building something concrete. `zero_control` is not just about shipping a bot; it is about using a modest project to explore how clear code, limited hardware, and AI-assisted development behave when they are forced to work together.

## Project structure

```text
zero_control/
├── doc/
│   ├── decisions.md
│   ├── requirements.md
│   └── user_stories.md
├── .gitignore
├── AGENTS.md
├── LICENSE
├── README.md
└── TODO.md
```

- `README.md` contains the project overview and acts as the main entry point.
- `TODO.md` tracks short-term tasks and open work.
- `doc/requirements.md` captures the functional and technical requirements as they become clearer.
- `doc/user_stories.md` records intended user-facing behavior and usage scenarios.
- `doc/decisions.md` documents implementation and architecture decisions, along with the reasoning behind them.
- `AGENTS.md` defines working rules and engineering expectations for AI-assisted collaboration in this repository.
