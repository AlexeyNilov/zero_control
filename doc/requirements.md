# Requirements

## EARS (Easy Approach to Requirements Syntax)

Use the EARS structure for precise requirements:

> **While** `<optional precondition>`, **when** `<optional trigger>`, **the system shall** `<system response>`.

This helps ensure requirements are:

* Context-aware
* Trigger-based
* Action-specific

## Actual requirements

> **When** the bot successfully connects to the Telegram API and passes authentication during startup, **the system shall** write a log entry indicating successful startup.

> **When** the bot starts, **the system shall** read the `AUTHORIZED_IDS` environment variable as the configured allowlist of Telegram user IDs that may interact with the bot.

> **When** the bot receives a message in a Telegram chat from a Telegram user ID listed in `AUTHORIZED_IDS`, **the system shall** write a log entry indicating message receipt without recording the message text.

> **When** the bot receives a message in a Telegram chat from a Telegram user ID that is not listed in `AUTHORIZED_IDS`, **the system shall** ignore the message.

> **While** the bot is running, **when** an operational event requires developer attention, **the system shall** post a notification to the Telegram chat identified by the `DEVELOPER_CHAT_ID` environment variable.

> **When** the bot successfully connects to the Telegram API and passes authentication during startup, **the system shall** post a startup notification to the Telegram chat identified by the `DEVELOPER_CHAT_ID` environment variable.

> **When** the bot encounters a critical runtime failure that prevents normal operation or requires manual intervention, **the system shall** post a failure notification to the Telegram chat identified by the `DEVELOPER_CHAT_ID` environment variable.

> **When** the bot receives the `/status` command in a Telegram chat from a Telegram user ID listed in `AUTHORIZED_IDS`, **the system shall** reply in that chat with `zero_control is online`.
