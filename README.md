# dflow - User Guide

Welcome to **dflow**! This guide will help you get started with installation, basic commands, and understanding how to use this tool to manage your workflows.

---

## What is dflow?

**dflow** is a command-line application for managing *flows* (work processes), tracking sessions, and logging activities. If you work on multiple projects or tasks and want a simple way to keep track of time and details for each, `dflow` is for you.

---

## Installation

### Prerequisites

- **Go** (Golang) installed ([download here](https://golang.org/dl/))
- SQLite (installed automatically by Go dependencies)

### Steps

1. **Clone the repository**:
   ```sh
   git clone https://github.com/OttobitSystems/dflow.git
   cd dflow
   ```

2. **Build the application**:
   ```sh
   go build -o dflow
   ```

3. **Run dflow**:
   ```sh
   ./dflow
   ```

---

## Basic Commands

Here are the most common commands you'll use in `dflow`:

### 1. List all flows

```sh
dflow list
```
*Shows all your available flows (work processes).*

---

### 2. Create a new flow

```sh
dflow config set create-flow <flow-name>
```
*Creates a new flow named `<flow-name>`. Use this for new projects or tasks.*

---

### 3. Enter a flow session

```sh
dflow enter <flow-name>
```
*Starts a new session in the flow named `<flow-name>`. Use this when you begin working on a flow.*

---

### 4. Log an activity

```sh
dflow logs <flow-name>
```
*Shows logs for the specified flow. This is useful for reviewing what happened in previous sessions.*

---

### 5. Get a recap

```sh
dflow recap
```
*Displays a summary table showing each flow, last entry, and total time spent.*

---

### 6. Manage configuration

```sh
dflow config
```
*Access configuration options for dflow.*

---

### 7. Manage flow space

```sh
dflow space
```
*Manage your flow spaces (advanced usage).*

---

## Example Workflow

1. **Create a flow** for your project:
   ```sh
   dflow config set create-flow ProjectX
   ```

2. **Start working** on ProjectX:
   ```sh
   dflow enter ProjectX
   ```

3. **Check logs** for ProjectX:
   ```sh
   dflow logs ProjectX
   ```

4. **Review your time** spent:
   ```sh
   dflow recap
   ```

---

## FAQ

**Where is my data stored?**  
All data is stored locally in an SQLite database file (`dflow.db`).

**Can I use dflow with a team?**  
dflow is primarily designed for individual use, but flows and logs can be exported or shared.

**How do I reset or delete a flow?**  
Currently, you’ll need to use database tools to manually remove flows.

---

## Need Help?

- Check the repository's [Issues](https://github.com/OttobitSystems/dflow/issues) for help or to report a bug.
- Contact the maintainers via GitHub.

---

## License

MIT

---

Happy tracking!
