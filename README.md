# GoCLI - Command Line Tool

A feature-rich command-line tool built with Go and Cobra, offering task management, text-to-speech conversion, and password generation.

## Features

GoCLI provides several useful utilities:

- **Task Management** - Create, list, update, and delete tasks
- **Text-to-Speech** - Convert text to spoken audio 
- **Password Generation** - Create secure random passwords

## Installation

### Prerequisites

- Go 1.23 or higher
- [Cobra CLI library](https://github.com/spf13/cobra)
- For TTS: macOS with `say` command (built-in)

### Building from source

1. Clone the repository:
```bash
git clone https://github.com/Karn-P/gocli.git
cd gocli
```

2. Build the application:
```bash
go build -o gocli
```

3. Install (optional):
```bash
go install
```

## Usage

### Task Management Commands

The `todo` command allows you to manage your task list.

#### Adding Tasks

You can add tasks using either command or flag syntax:

```bash
# Using command syntax
gocli todo add "Buy groceries"

# Using flag syntax
gocli todo -a "Buy groceries"
gocli todo --add "Buy groceries"
```

#### Listing Tasks

List all your tasks:

```bash
# Using command syntax
gocli todo list

# Using flag syntax
gocli todo -l
gocli todo --list

# Default behavior (no arguments)
gocli todo
```

The tasks will be displayed in a table format:

| ID  | Description      | CreatedAt           | Completed |
|-----|------------------|---------------------|-----------|
| 1   | Buy groceries    | 2023-06-15 10:30:45 | No        |
| 2   | Clean the house  | 2023-06-14 09:15:22 | Yes       |

#### Removing Tasks

To remove a task by its ID:

```bash
# Using command syntax
gocli todo remove 1

# Using flag syntax
gocli todo -r 1
gocli todo --remove 1
```

#### Marking Tasks as Complete

To mark a task as complete:

```bash
# Using command syntax
gocli todo complete 1

# Using flag syntax
gocli todo -c 1
gocli todo --complete 1
```

### Text-to-Speech Commands

The `tts` command lets you convert text to speech.

#### Speaking Text

To convert text to speech:

```bash
# Using command syntax
gocli tts speak "Hello, world!"

# Using flag syntax
gocli tts -s "Hello, world!"
gocli tts --speak "Hello, world!"
```

### Password Generation Commands

The `passgen` command allows you to generate random passwords.

#### Generating Passwords

To generate a secure password:

```bash
# Basic usage (12 character password with letters only)
gocli passgen

# Customize length
gocli passgen -l 16
gocli passgen --length 16

# Include digits
gocli passgen -d
gocli passgen --digits

# Include special characters
gocli passgen -s
gocli passgen --special-chars

# Full customization
gocli passgen -l 20 -d -s
```

Options:
- `-l, --length`: Set the password length (default: 12)
- `-d, --digits`: Include digits in the password
- `-s, --special-chars`: Include special characters in the password

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

