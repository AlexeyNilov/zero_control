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

> **When** the bot starts, **the system shall** subscribe to the MQTT topic `zero-control/notify` on the broker identified by `MQTT_BROKER_URL`.

> **When** `MQTT_BROKER_URL` is not configured, **the system shall** use `tcp://localhost:1883` as the MQTT broker URL.

> **While** the bot is running, **when** a non-empty MQTT message is received on `zero-control/notify`, **the system shall** send the message payload as a Telegram notification to the chat identified by `MAIN_CHAT_ID`.

> **While** the bot is running, **when** a blank MQTT message is received on `zero-control/notify`, **the system shall** ignore the message and write a log entry.

> **When** the bot receives a message in a Telegram chat from a Telegram user ID listed in `AUTHORIZED_IDS`, **the system shall** write a log entry indicating message receipt without recording the message text.

> **When** the bot receives a message in a Telegram chat from a Telegram user ID that is not listed in `AUTHORIZED_IDS`, **the system shall** ignore the message.

> **While** the bot is running, **when** an operational event requires developer attention, **the system shall** post a notification to the Telegram chat identified by the `DEVELOPER_CHAT_ID` environment variable.

> **When** the bot successfully connects to the Telegram API and passes authentication during startup, **the system shall** post a startup notification to the Telegram chat identified by the `DEVELOPER_CHAT_ID` environment variable.

> **When** the bot encounters a critical runtime failure that prevents normal operation or requires manual intervention, **the system shall** post a failure notification to the Telegram chat identified by the `DEVELOPER_CHAT_ID` environment variable.

> **When** the bot receives the `/status` command in a Telegram chat from a Telegram user ID listed in `AUTHORIZED_IDS`, **the system shall** reply in that chat with `zero_control is online` and the current LAN IP address when available.
