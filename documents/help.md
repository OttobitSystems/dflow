# dflow

**dflow** is a CLI-based tool designed for managing "flows"—sequences of work or actions—and tracking sessions and logs associated with those flows. It provides commands for creating, listing, and managing flows, as well as logging activities and generating operational recaps.

---

## Features

- **Flow Management:** Create and list flows, each representing a distinct workflow or process.
- **Session Tracking:** Start and end sessions within flows to monitor work periods.
- **Logging:** Store and retrieve logs for each session, supporting detailed tracking and auditing.
- **Recaps:** Generate summaries of flow usage, including time spent and last entry information.
- **CLI Interface:** All features are accessible via an extensible command-line interface (CLI), built using [cobra](https://github.com/spf13/cobra).

---

## Architecture Overview

- **Main Entry Point:**  
  The application starts from `dflow.go`, which initializes the database and registers CLI commands.

- **Commands:**  
  - Located in `internal/commands/`
  - Examples:
    - `list`: Lists all available flows.
    - `logs`: Shows logs for a flow.
    - `recap`: Displays a summary table of flow activity.
    - `config` & `space`: Manage configuration and flow spaces.

- **Persistency Layer:**  
  - Located in `internal/persistency/`
  - Uses [GORM](https://gorm.io/) with SQLite for data storage.
  - Models:
    - **Flow:** Main workflow entity.
    - **Session:** Instance of time spent in a flow.
    - **Log:** Messages/events tied to flows and sessions.
  - Functions:
    - Initialize the database (`InitDatabase`)
    - Create flows and sessions
    - Store and retrieve logs

- **Recap Logic:**  
  - Located in `internal/recap/`
  - Aggregates data on flows and sessions to provide operational summaries.

---

## Data Model

- **Flow**
  - `Name` (primary key)
  - `CreatedAt`
  - Relations: Sessions, Logs

- **Session**
  - `ID` (primary key)
  - `FlowID` (foreign key)
  - `StartedAt`, `CompletedAt`

- **Log**
  - `ID` (primary key)
  - `FlowID` (foreign key)
  - `SessionID` (foreign key)
  - `TimeStamp`
  - `Log` (message)

---

## Usage (CLI)

After building and installing `dflow`, you can use the following commands:

```sh
dflow list
# Lists all flows

dflow logs <flow-name>
# Shows logs for the specified flow

dflow recap
# Shows a summary table for all flows

dflow config
# Manage configuration

dflow space
# Manage flow space
```

---

## Example: Creating and Using Flows

1. **Initialize the database** (happens automatically on start).
2. **Create a new flow:**  
   `dflow config set create-flow <flow-name>`
3. **Enter a flow session:**  
   `dflow enter <flow-name>`
4. **Log activities:**  
   `dflow logs <flow-name>`
5. **Review usage:**  
   `dflow recap`

---

## Development & Testing

- Unit tests are located in `tests/`, e.g., `repository_test.go` tests database initialization and flow creation.
- Main dependencies: `cobra`, `gorm`, `sqlite`, `bubbletea` (for TUI elements).

---

## Extending

- Add new commands in `internal/commands/`.
- Extend data models in `internal/persistency/models/`.
- Integrate UI features via `internal/tui/` using `bubbletea`.

---

## License

MIT (see LICENSE file)

---

## Authors & Contributors

- Maintained by OttobitSystems and contributors.

---

## References

- [cobra](https://github.com/spf13/cobra) - CLI library
- [gorm](https://gorm.io/) - ORM for Go
- [bubbletea](https://github.com/charmbracelet/bubbletea) - TUI library

