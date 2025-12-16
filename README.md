# Stagehand CLI

The official CLI for the [Stagehand REST API](https://docs.stagehand.dev).

It is generated with [Stainless](https://www.stainless.com/).

## Installation

### Installing with Go

<!-- x-release-please-start-version -->

```sh
go install 'github.com/browserbase/stagehand-cli/cmd/stagehand@latest'
```

### Running Locally

<!-- x-release-please-start-version -->

```sh
go run cmd/stagehand/main.go
```

<!-- x-release-please-end -->

## Usage

The CLI follows a resource-based command structure:

```sh
stagehand [resource] [command] [flags]
```

```sh
stagehand sessions act \
  --session-id 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e \
  --input 'click the first link on the page'
```

For details about specific commands, use the `--help` flag.

## Global Flags

- `--debug` - Enable debug logging (includes HTTP request/response details)
- `--version`, `-v` - Show the CLI version
