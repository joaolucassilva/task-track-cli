# Task Track CLI

Task Tracker CLI is a simple command-line interface (CLI) application designed to help you track and manage your tasks.
This project allows you to add, update, delete, and list tasks, as well as mark them as in progress or done.
All tasks are stored in a JSON file in the current directory.

Challenge from [roadmap.sh](https://roadmap.sh/projects/task-tracker).

## Features

- **Add tasks:** Create new tasks with a description.
- **Update tasks:** Modify the description of existing tasks.
- **Delete tasks:** Remove tasks that are no longer needed.
- **Mark as in progress:** Set a task's status to "in-progress".
- **Mark as done:** Set a task's status to "done".
- **List tasks:** View all tasks, of filter by status, (done, todo, in-progress)

## Requirements

- **Go:** This project requires [GO](https://golang.org/dl/) to be installed on your machine.

## Installation

1. **Clone the repository:**
   ```bash
    git clone https://github.com/joaolucassilva/task-track-cli.git
    cd task-track-cli
    ```
2. **Install dependencies:**
    ```bash
    go mod tidy   
    ```
3. **Build the application:**
    ```bash
    go build -o task-track-cli
    ```
4. **Run the application:**
   ```bash
   ./task-track-cli <command> [arguments]
   ```

## Usage

To use the Task Tracker CLI, run the application from the command line with the appropriate arguments.

### Commands

- **Add a task:**
- **Update a task:**
- **Delete a task:**
- **Mark a task as in progress:**
- **Mark a task as done:**
- **List all tasks:**
- **List tasks that are todo:**
- **List tasks that are done:**
- **List tasks that are in-progress:**

### Example:

```bash
# Add a new task
./task-track-cli add "Buy groceries"

# Update a task
./task-track-cli update 1 "Buy groceries and cook dinner"

# Deleting a task
./task-track-cli delete 1

# Marking a tasks as in progress
./task-track-cli mark-in-progress 1

# Marking a tasks as in done
./task-track-cli mark-done 1

# List all tasks
./task-track-cli list

# List tasks by status
./task-track-cli list done
./task-track-cli list todo
./task-track-cli list in-progress
```
