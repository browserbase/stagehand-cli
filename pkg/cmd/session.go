// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/browserbase/stagehand-cli/internal/apiquery"
	"github.com/browserbase/stagehand-cli/internal/requestflag"
	"github.com/browserbase/stagehand-go"
	"github.com/browserbase/stagehand-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var sessionsAct = cli.Command{
	Name:  "act",
	Usage: "Executes a browser action using natural language instructions or a predefined\nAction object.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:  "id",
			Usage: "Unique session identifier",
		},
		&requestflag.Flag[any]{
			Name:     "input",
			Usage:    "Natural language instruction or Action object",
			BodyPath: "input",
		},
		&requestflag.Flag[string]{
			Name:     "frame-id",
			Usage:    "Target frame ID for the action",
			BodyPath: "frameId",
		},
		&requestflag.Flag[any]{
			Name:     "options",
			BodyPath: "options",
		},
		&requestflag.Flag[bool]{
			Name:     "stream-response",
			Usage:    "Whether to stream the response via SSE",
			BodyPath: "streamResponse",
		},
		&requestflag.Flag[string]{
			Name:       "x-language",
			Usage:      "Client SDK language",
			HeaderPath: "x-language",
		},
		&requestflag.Flag[string]{
			Name:       "x-sdk-version",
			Usage:      "Version of the Stagehand SDK",
			HeaderPath: "x-sdk-version",
		},
		&requestflag.Flag[requestflag.DateTimeValue]{
			Name:       "x-sent-at",
			Usage:      "ISO timestamp when request was sent",
			HeaderPath: "x-sent-at",
		},
		&requestflag.Flag[string]{
			Name:       "x-stream-response",
			Usage:      "Whether to stream the response via SSE",
			HeaderPath: "x-stream-response",
		},
	},
	Action:          handleSessionsAct,
	HideHelpCommand: true,
}

var sessionsEnd = cli.Command{
	Name:  "end",
	Usage: "Terminates the browser session and releases all associated resources.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:  "id",
			Usage: "Unique session identifier",
		},
		&requestflag.Flag[string]{
			Name:       "x-language",
			Usage:      "Client SDK language",
			HeaderPath: "x-language",
		},
		&requestflag.Flag[string]{
			Name:       "x-sdk-version",
			Usage:      "Version of the Stagehand SDK",
			HeaderPath: "x-sdk-version",
		},
		&requestflag.Flag[requestflag.DateTimeValue]{
			Name:       "x-sent-at",
			Usage:      "ISO timestamp when request was sent",
			HeaderPath: "x-sent-at",
		},
		&requestflag.Flag[string]{
			Name:       "x-stream-response",
			Usage:      "Whether to stream the response via SSE",
			HeaderPath: "x-stream-response",
		},
	},
	Action:          handleSessionsEnd,
	HideHelpCommand: true,
}

var sessionsExecute = cli.Command{
	Name:  "execute",
	Usage: "Runs an autonomous AI agent that can perform complex multi-step browser tasks.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:  "id",
			Usage: "Unique session identifier",
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
			Usage:    "Target frame ID for the agent",
			BodyPath: "frameId",
		},
		&requestflag.Flag[bool]{
			Name:     "stream-response",
			Usage:    "Whether to stream the response via SSE",
			BodyPath: "streamResponse",
		},
		&requestflag.Flag[string]{
			Name:       "x-language",
			Usage:      "Client SDK language",
			HeaderPath: "x-language",
		},
		&requestflag.Flag[string]{
			Name:       "x-sdk-version",
			Usage:      "Version of the Stagehand SDK",
			HeaderPath: "x-sdk-version",
		},
		&requestflag.Flag[requestflag.DateTimeValue]{
			Name:       "x-sent-at",
			Usage:      "ISO timestamp when request was sent",
			HeaderPath: "x-sent-at",
		},
		&requestflag.Flag[string]{
			Name:       "x-stream-response",
			Usage:      "Whether to stream the response via SSE",
			HeaderPath: "x-stream-response",
		},
	},
	Action:          handleSessionsExecute,
	HideHelpCommand: true,
}

var sessionsExtract = cli.Command{
	Name:  "extract",
	Usage: "Extracts structured data from the current page using AI-powered analysis.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:  "id",
			Usage: "Unique session identifier",
		},
		&requestflag.Flag[string]{
			Name:     "frame-id",
			Usage:    "Target frame ID for the extraction",
			BodyPath: "frameId",
		},
		&requestflag.Flag[string]{
			Name:     "instruction",
			Usage:    "Natural language instruction for what to extract",
			BodyPath: "instruction",
		},
		&requestflag.Flag[any]{
			Name:     "options",
			BodyPath: "options",
		},
		&requestflag.Flag[any]{
			Name:     "schema",
			Usage:    "JSON Schema defining the structure of data to extract",
			BodyPath: "schema",
		},
		&requestflag.Flag[bool]{
			Name:     "stream-response",
			Usage:    "Whether to stream the response via SSE",
			BodyPath: "streamResponse",
		},
		&requestflag.Flag[string]{
			Name:       "x-language",
			Usage:      "Client SDK language",
			HeaderPath: "x-language",
		},
		&requestflag.Flag[string]{
			Name:       "x-sdk-version",
			Usage:      "Version of the Stagehand SDK",
			HeaderPath: "x-sdk-version",
		},
		&requestflag.Flag[requestflag.DateTimeValue]{
			Name:       "x-sent-at",
			Usage:      "ISO timestamp when request was sent",
			HeaderPath: "x-sent-at",
		},
		&requestflag.Flag[string]{
			Name:       "x-stream-response",
			Usage:      "Whether to stream the response via SSE",
			HeaderPath: "x-stream-response",
		},
	},
	Action:          handleSessionsExtract,
	HideHelpCommand: true,
}

var sessionsNavigate = cli.Command{
	Name:  "navigate",
	Usage: "Navigates the browser to the specified URL.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:  "id",
			Usage: "Unique session identifier",
		},
		&requestflag.Flag[string]{
			Name:     "url",
			Usage:    "URL to navigate to",
			BodyPath: "url",
		},
		&requestflag.Flag[string]{
			Name:     "frame-id",
			Usage:    "Target frame ID for the navigation",
			BodyPath: "frameId",
		},
		&requestflag.Flag[any]{
			Name:     "options",
			BodyPath: "options",
		},
		&requestflag.Flag[bool]{
			Name:     "stream-response",
			Usage:    "Whether to stream the response via SSE",
			BodyPath: "streamResponse",
		},
		&requestflag.Flag[string]{
			Name:       "x-language",
			Usage:      "Client SDK language",
			HeaderPath: "x-language",
		},
		&requestflag.Flag[string]{
			Name:       "x-sdk-version",
			Usage:      "Version of the Stagehand SDK",
			HeaderPath: "x-sdk-version",
		},
		&requestflag.Flag[requestflag.DateTimeValue]{
			Name:       "x-sent-at",
			Usage:      "ISO timestamp when request was sent",
			HeaderPath: "x-sent-at",
		},
		&requestflag.Flag[string]{
			Name:       "x-stream-response",
			Usage:      "Whether to stream the response via SSE",
			HeaderPath: "x-stream-response",
		},
	},
	Action:          handleSessionsNavigate,
	HideHelpCommand: true,
}

var sessionsObserve = cli.Command{
	Name:  "observe",
	Usage: "Identifies and returns available actions on the current page that match the\ngiven instruction.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:  "id",
			Usage: "Unique session identifier",
		},
		&requestflag.Flag[string]{
			Name:     "frame-id",
			Usage:    "Target frame ID for the observation",
			BodyPath: "frameId",
		},
		&requestflag.Flag[string]{
			Name:     "instruction",
			Usage:    "Natural language instruction for what actions to find",
			BodyPath: "instruction",
		},
		&requestflag.Flag[any]{
			Name:     "options",
			BodyPath: "options",
		},
		&requestflag.Flag[bool]{
			Name:     "stream-response",
			Usage:    "Whether to stream the response via SSE",
			BodyPath: "streamResponse",
		},
		&requestflag.Flag[string]{
			Name:       "x-language",
			Usage:      "Client SDK language",
			HeaderPath: "x-language",
		},
		&requestflag.Flag[string]{
			Name:       "x-sdk-version",
			Usage:      "Version of the Stagehand SDK",
			HeaderPath: "x-sdk-version",
		},
		&requestflag.Flag[requestflag.DateTimeValue]{
			Name:       "x-sent-at",
			Usage:      "ISO timestamp when request was sent",
			HeaderPath: "x-sent-at",
		},
		&requestflag.Flag[string]{
			Name:       "x-stream-response",
			Usage:      "Whether to stream the response via SSE",
			HeaderPath: "x-stream-response",
		},
	},
	Action:          handleSessionsObserve,
	HideHelpCommand: true,
}

var sessionsStart = cli.Command{
	Name:  "start",
	Usage: "Creates a new browser session with the specified configuration. Returns a\nsession ID used for all subsequent operations.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "model-name",
			Usage:    "Model name to use for AI operations",
			BodyPath: "modelName",
		},
		&requestflag.Flag[float64]{
			Name:     "act-timeout-ms",
			Usage:    "Timeout in ms for act operations",
			BodyPath: "actTimeoutMs",
		},
		&requestflag.Flag[any]{
			Name:     "browser",
			BodyPath: "browser",
		},
		&requestflag.Flag[any]{
			Name:     "browserbase-session-create-params",
			BodyPath: "browserbaseSessionCreateParams",
		},
		&requestflag.Flag[string]{
			Name:     "browserbase-session-id",
			Usage:    "Existing Browserbase session ID to resume",
			BodyPath: "browserbaseSessionID",
		},
		&requestflag.Flag[bool]{
			Name:     "debug-dom",
			BodyPath: "debugDom",
		},
		&requestflag.Flag[float64]{
			Name:     "dom-settle-timeout-ms",
			Usage:    "Timeout in ms to wait for DOM to settle",
			BodyPath: "domSettleTimeoutMs",
		},
		&requestflag.Flag[bool]{
			Name:     "experimental",
			BodyPath: "experimental",
		},
		&requestflag.Flag[bool]{
			Name:     "self-heal",
			Usage:    "Enable self-healing for failed actions",
			BodyPath: "selfHeal",
		},
		&requestflag.Flag[string]{
			Name:     "system-prompt",
			Usage:    "Custom system prompt for AI operations",
			BodyPath: "systemPrompt",
		},
		&requestflag.Flag[int64]{
			Name:     "verbose",
			Usage:    "Logging verbosity level (0=quiet, 1=normal, 2=debug)",
			BodyPath: "verbose",
		},
		&requestflag.Flag[bool]{
			Name:     "wait-for-captcha-solves",
			BodyPath: "waitForCaptchaSolves",
		},
		&requestflag.Flag[string]{
			Name:       "x-language",
			Usage:      "Client SDK language",
			HeaderPath: "x-language",
		},
		&requestflag.Flag[string]{
			Name:       "x-sdk-version",
			Usage:      "Version of the Stagehand SDK",
			HeaderPath: "x-sdk-version",
		},
		&requestflag.Flag[requestflag.DateTimeValue]{
			Name:       "x-sent-at",
			Usage:      "ISO timestamp when request was sent",
			HeaderPath: "x-sent-at",
		},
		&requestflag.Flag[string]{
			Name:       "x-stream-response",
			Usage:      "Whether to stream the response via SSE",
			HeaderPath: "x-stream-response",
		},
	},
	Action:          handleSessionsStart,
	HideHelpCommand: true,
}

func handleSessionsAct(ctx context.Context, cmd *cli.Command) error {
	client := stagehand.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
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
		false,
	)
	if err != nil {
		return err
	}

	stream := client.Sessions.ActStreaming(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	for stream.Next() {
		fmt.Printf("%s\n", stream.Current().RawJSON())
	}
	return stream.Err()
}

func handleSessionsEnd(ctx context.Context, cmd *cli.Command) error {
	client := stagehand.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stagehand.SessionEndParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Sessions.End(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "sessions end", obj, format, transform)
}

func handleSessionsExecute(ctx context.Context, cmd *cli.Command) error {
	client := stagehand.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stagehand.SessionExecuteParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	stream := client.Sessions.ExecuteStreaming(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	for stream.Next() {
		fmt.Printf("%s\n", stream.Current().RawJSON())
	}
	return stream.Err()
}

func handleSessionsExtract(ctx context.Context, cmd *cli.Command) error {
	client := stagehand.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
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
		false,
	)
	if err != nil {
		return err
	}

	stream := client.Sessions.ExtractStreaming(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	for stream.Next() {
		fmt.Printf("%s\n", stream.Current().RawJSON())
	}
	return stream.Err()
}

func handleSessionsNavigate(ctx context.Context, cmd *cli.Command) error {
	client := stagehand.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
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
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Sessions.Navigate(
		ctx,
		cmd.Value("id").(string),
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
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
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
		false,
	)
	if err != nil {
		return err
	}

	stream := client.Sessions.ObserveStreaming(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	for stream.Next() {
		fmt.Printf("%s\n", stream.Current().RawJSON())
	}
	return stream.Err()
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
		false,
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
