# Requirements

## EARS (Easy Approach to Requirements Syntax)

Use the EARS structure for precise requirements:

> **While** `<optional precondition>`, **when** `<optional trigger>`, **the system shall** `<system response>`.

This helps ensure requirements are:

* Context-aware
* Trigger-based
* Action-specific

## Actual requirements

> **When** the bot receives a message in a Telegram chat, **the system shall** write a log entry indicating message receipt without recording the message text.

> **When** the bot successfully connects to the Telegram API and passes authentication during startup, **the system shall** write a log entry indicating successful startup.

> **While** the bot is running, **when** an operational event requires developer attention, **the system shall** post a notification to the Telegram chat identified by the `DEVELOPER_CHAT_ID` environment variable.

> **When** the bot successfully connects to the Telegram API and passes authentication during startup, **the system shall** post a startup notification to the Telegram chat identified by the `DEVELOPER_CHAT_ID` environment variable.

> **When** the bot receives the `/status` command in a Telegram chat, **the system shall** reply in that chat with `zero_control is online`.

> **When** the bot encounters a critical runtime failure that prevents normal operation or requires manual intervention, **the system shall** post a failure notification to the Telegram chat identified by the `DEVELOPER_CHAT_ID` environment variable.
