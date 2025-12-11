# Stagehand CLI

The official CLI for the [Stagehand REST API](https://browserbase.com).

It is generated with [Stainless](https://www.stainless.com/).

## Installation

### Installing with Go

```sh
go install 'github.com/stainless-sdks/stagehand-cli/cmd/stagehand@latest'
```

### Running Locally

```sh
go run cmd/stagehand/main.go
```

## Usage

The CLI follows a resource-based command structure:

```sh
stagehand [resource] [command] [flags]
```

```sh
stagehand sessions start \
  --env LOCAL
```

For details about specific commands, use the `--help` flag.

## Global Flags

- `--debug` - Enable debug logging (includes HTTP request/response details)
- `--version`, `-v` - Show the CLI version
