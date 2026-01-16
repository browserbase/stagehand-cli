# Stagehand CLI

The official CLI for the [Stagehand REST API](https://docs.stagehand.dev).

It is generated with [Stainless](https://www.stainless.com/).

<!-- x-release-please-start-version -->

## Installation

### Installing with Go

```sh
go install 'github.com/browserbase/stagehand-cli/cmd/stagehand@latest'
```

<!-- x-release-please-end -->

### Running Locally

```sh
./scripts/run args...
```

## Usage

The CLI follows a resource-based command structure:

```sh
stagehand [resource] [command] [flags]
```

```sh
stagehand sessions act \
  --id 00000000-your-session-id-000000000000 \
  --input 'click the first link on the page' \
  --frame-id frameId \
  --options '{model: openai/gpt-4o, timeout: 30000, variables: {username: john_doe}}' \
  --stream-response false \
  --x-stream-response true
```

For details about specific commands, use the `--help` flag.

## Global Flags

- `--help` - Show command line usage
- `--debug` - Enable debug logging (includes HTTP request/response details)
- `--version`, `-v` - Show the CLI version
- `--base-url` - Use a custom API backend URL
- `--format` - Change the output format (`auto`, `explore`, `json`, `jsonl`, `pretty`, `raw`, `yaml`)
- `--format-error` - Change the output format for errors (`auto`, `explore`, `json`, `jsonl`, `pretty`, `raw`, `yaml`)
- `--transform` - Transform the data output using [GJSON syntax](https://github.com/tidwall/gjson/blob/master/SYNTAX.md)
- `--transform-error` - Transform the error output using [GJSON syntax](https://github.com/tidwall/gjson/blob/master/SYNTAX.md)
