# pls

`pls` is a small command-line tool for defining named groups of shell commands in a file and running them as jobs. It is useful for keeping common development tasks such as building, testing, formatting, or starting an application in one project-local file.

## How It Works

`pls` reads a file named `plsfile` by default. Each job starts with a line beginning with `#:` followed by the job name. Every non-empty line after that is treated as a shell command belonging to the job, until the next job declaration.

When a job is run, its commands are joined into a shell script and executed with `sh -c`. The script uses `set -e`, so execution stops when a command fails.

Lines beginning with `//`, and inline text after `//`, are treated as comments. Blank lines are ignored.

## Example `plsfile`

```text
// Build the application.
#: build
go build ./...

// Run the test suite.
#: test
go test ./...

// Format the Go source files.
#: format
gofmt -w .

// Start the application locally.
#: dev
go run .
```

The example defines four jobs: `build`, `test`, `format`, and `dev`. Commands are run in the order they appear in the file.

## Installation

Build the executable from the project directory:

```sh
go build -o pls .
```

Alternatively, run it without creating a binary:

```sh
go run . build
```

The project requires Go 1.27 or newer.

## Usage

Run a job from the default `plsfile` in the current directory:

```sh
./pls build
```

Run a job from another file:

```sh
./pls --file path/to/plsfile test
```

The short option `-f` can be used instead:

```sh
./pls -f path/to/plsfile test
```

If no job name is provided, `pls` prints the available jobs:

```sh
./pls
```

See the command-line help for the available options:

```sh
./pls --help
```

### Options

| Option | Description | Default |
| --- | --- | --- |
| `-f`, `--file` | Path to the job file | `plsfile` |
| `-h`, `--help` | Show command help | |

## File Syntax

### Job declarations

Declare a job with `#:` and a name:

```text
#: clean
rm -rf ./build
```

The job name may contain spaces, although invoking jobs with simple names is recommended. A job declaration without a name is invalid.

### Commands

All lines after a job declaration are passed to the shell as commands:

```text
#: check
go vet ./...
go test ./...
```

Because commands run through `sh`, shell features such as pipes, environment variables, command substitution, and command chaining can be used where supported by the system shell.

### Comments

Use `//` for full-line or inline comments:

```text
// This is a comment.
#: greet
echo "hello" // This part is also ignored
```

Currently, `//` starts a comment wherever it appears on a line. It cannot be used as literal command text without being treated as a comment marker.

## Behavior and Errors

- Commands inherit the current working directory, standard input, standard output, and standard error.
- Commands run sequentially in the order they appear.
- If a command exits unsuccessfully, remaining commands in that job are not run.
- If the selected file does not exist, `pls` reports that the file was not found.
- If the requested job does not exist, `pls` prints the available job names.
- Only the first positional argument is used as the job name.

## Development

Run the test suite with:

```sh
go test ./...
```

The parser implementation is in `internal/plsfile`, and the CLI entry point is in `cmd`.
