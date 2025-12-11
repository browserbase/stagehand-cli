// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/stainless-sdks/stagehand-cli/internal/apiquery"
	"github.com/stainless-sdks/stagehand-cli/internal/requestflag"
	"github.com/stainless-sdks/stagehand-go"
	"github.com/stainless-sdks/stagehand-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var sessionsAct = cli.Command{
	Name:  "act",
	Usage: "Performs a browser action based on natural language instruction or a specific\naction object returned by observe().",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "session-id",
		},
		&requestflag.Flag[any]{
			Name:     "input",
			Usage:    "Natural language instruction",
			BodyPath: "input",
		},
		&requestflag.Flag[string]{
			Name:     "frame-id",
			Usage:    "Frame ID to act on (optional)",
			BodyPath: "frameId",
		},
		&requestflag.Flag[any]{
			Name:     "options",
			BodyPath: "options",
		},
		&requestflag.Flag[string]{
			Name:       "x-stream-response",
			Default:    "true",
			HeaderPath: "x-stream-response",
		},
	},
	Action:          handleSessionsAct,
	HideHelpCommand: true,
}

var sessionsEnd = cli.Command{
	Name:  "end",
	Usage: "Closes the browser and cleans up all resources associated with the session.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "session-id",
		},
	},
	Action:          handleSessionsEnd,
	HideHelpCommand: true,
}

var sessionsExecuteAgent = cli.Command{
	Name:  "execute-agent",
	Usage: "Runs an autonomous agent that can perform multiple actions to complete a complex\ntask.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "session-id",
		},
		&requestflag.Flag[any]{
			Name:     "agent-config",
			BodyPath: "agentConfig",
		},
		&requestflag.Flag[any]{
			Name:     "execute-options",
			BodyPath: "executeOptions",
		},
		&requestflag.Flag[string]{
			Name:     "frame-id",
			BodyPath: "frameId",
		},
		&requestflag.Flag[string]{
			Name:       "x-stream-response",
			Default:    "true",
			HeaderPath: "x-stream-response",
		},
	},
	Action:          handleSessionsExecuteAgent,
	HideHelpCommand: true,
}

var sessionsExtract = cli.Command{
	Name:  "extract",
	Usage: "Extracts data from the current page using natural language instructions and\noptional JSON schema for structured output.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "session-id",
		},
		&requestflag.Flag[string]{
			Name:     "frame-id",
			Usage:    "Frame ID to extract from",
			BodyPath: "frameId",
		},
		&requestflag.Flag[string]{
			Name:     "instruction",
			Usage:    "Natural language instruction for extraction",
			BodyPath: "instruction",
		},
		&requestflag.Flag[any]{
			Name:     "options",
			BodyPath: "options",
		},
		&requestflag.Flag[any]{
			Name:     "schema",
			Usage:    "JSON Schema for structured output",
			BodyPath: "schema",
		},
		&requestflag.Flag[string]{
			Name:       "x-stream-response",
			Default:    "true",
			HeaderPath: "x-stream-response",
		},
	},
	Action:          handleSessionsExtract,
	HideHelpCommand: true,
}

var sessionsNavigate = cli.Command{
	Name:  "navigate",
	Usage: "Navigates the browser to the specified URL and waits for page load.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "session-id",
		},
		&requestflag.Flag[string]{
			Name:     "url",
			Usage:    "URL to navigate to",
			BodyPath: "url",
		},
		&requestflag.Flag[string]{
			Name:     "frame-id",
			BodyPath: "frameId",
		},
		&requestflag.Flag[any]{
			Name:     "options",
			BodyPath: "options",
		},
		&requestflag.Flag[string]{
			Name:       "x-stream-response",
			Default:    "true",
			HeaderPath: "x-stream-response",
		},
	},
	Action:          handleSessionsNavigate,
	HideHelpCommand: true,
}

var sessionsObserve = cli.Command{
	Name:  "observe",
	Usage: "Returns a list of candidate actions that can be performed on the page,\noptionally filtered by natural language instruction.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "session-id",
		},
		&requestflag.Flag[string]{
			Name:     "frame-id",
			Usage:    "Frame ID to observe",
			BodyPath: "frameId",
		},
		&requestflag.Flag[string]{
			Name:     "instruction",
			Usage:    "Natural language instruction to filter actions",
			BodyPath: "instruction",
		},
		&requestflag.Flag[any]{
			Name:     "options",
			BodyPath: "options",
		},
		&requestflag.Flag[string]{
			Name:       "x-stream-response",
			Default:    "true",
			HeaderPath: "x-stream-response",
		},
	},
	Action:          handleSessionsObserve,
	HideHelpCommand: true,
}

var sessionsStart = cli.Command{
	Name:  "start",
	Usage: "Initializes a new Stagehand session with a browser instance. Returns a session\nID that must be used for all subsequent requests.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "env",
			Usage:    "Environment to run the browser in",
			BodyPath: "env",
		},
		&requestflag.Flag[string]{
			Name:     "api-key",
			Usage:    "API key for Browserbase (required when env=BROWSERBASE)",
			BodyPath: "apiKey",
		},
		&requestflag.Flag[int64]{
			Name:     "dom-settle-timeout",
			Usage:    "Timeout in ms to wait for DOM to settle",
			BodyPath: "domSettleTimeout",
		},
		&requestflag.Flag[any]{
			Name:     "local-browser-launch-options",
			Usage:    "Options for local browser launch",
			BodyPath: "localBrowserLaunchOptions",
		},
		&requestflag.Flag[string]{
			Name:     "model",
			Usage:    "AI model to use for actions",
			BodyPath: "model",
		},
		&requestflag.Flag[string]{
			Name:     "project-id",
			Usage:    "Project ID for Browserbase (required when env=BROWSERBASE)",
			BodyPath: "projectId",
		},
		&requestflag.Flag[bool]{
			Name:     "self-heal",
			Usage:    "Enable self-healing for failed actions",
			BodyPath: "selfHeal",
		},
		&requestflag.Flag[string]{
			Name:     "system-prompt",
			Usage:    "Custom system prompt for AI actions",
			BodyPath: "systemPrompt",
		},
		&requestflag.Flag[int64]{
			Name:     "verbose",
			Usage:    "Logging verbosity level",
			BodyPath: "verbose",
		},
	},
	Action:          handleSessionsStart,
	HideHelpCommand: true,
}

func handleSessionsAct(ctx context.Context, cmd *cli.Command) error {
	client := stagehand.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("session-id") && len(unusedArgs) > 0 {
		cmd.Set("session-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := stagehand.SessionActParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Sessions.Act(
		ctx,
		cmd.Value("session-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "sessions act", obj, format, transform)
}

func handleSessionsEnd(ctx context.Context, cmd *cli.Command) error {
	client := stagehand.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("session-id") && len(unusedArgs) > 0 {
		cmd.Set("session-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Sessions.End(ctx, cmd.Value("session-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "sessions end", obj, format, transform)
}

func handleSessionsExecuteAgent(ctx context.Context, cmd *cli.Command) error {
	client := stagehand.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("session-id") && len(unusedArgs) > 0 {
		cmd.Set("session-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := stagehand.SessionExecuteAgentParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Sessions.ExecuteAgent(
		ctx,
		cmd.Value("session-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "sessions execute-agent", obj, format, transform)
}

func handleSessionsExtract(ctx context.Context, cmd *cli.Command) error {
	client := stagehand.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("session-id") && len(unusedArgs) > 0 {
		cmd.Set("session-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := stagehand.SessionExtractParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Sessions.Extract(
		ctx,
		cmd.Value("session-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "sessions extract", obj, format, transform)
}

func handleSessionsNavigate(ctx context.Context, cmd *cli.Command) error {
	client := stagehand.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("session-id") && len(unusedArgs) > 0 {
		cmd.Set("session-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := stagehand.SessionNavigateParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Sessions.Navigate(
		ctx,
		cmd.Value("session-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "sessions navigate", obj, format, transform)
}

func handleSessionsObserve(ctx context.Context, cmd *cli.Command) error {
	client := stagehand.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("session-id") && len(unusedArgs) > 0 {
		cmd.Set("session-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := stagehand.SessionObserveParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Sessions.Observe(
		ctx,
		cmd.Value("session-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "sessions observe", obj, format, transform)
}

func handleSessionsStart(ctx context.Context, cmd *cli.Command) error {
	client := stagehand.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := stagehand.SessionStartParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Sessions.Start(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "sessions start", obj, format, transform)
}
