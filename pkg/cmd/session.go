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

var sessionsAct = requestflag.WithInnerFlags(cli.Command{
	Name:    "act",
	Usage:   "Executes a browser action using natural language instructions or a predefined\nAction object.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Usage:    "Unique session identifier",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "input",
			Usage:    "Natural language instruction or Action object",
			Required: true,
			BodyPath: "input",
		},
		&requestflag.Flag[string]{
			Name:     "frame-id",
			Usage:    "Target frame ID for the action",
			BodyPath: "frameId",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "options",
			BodyPath: "options",
		},
		&requestflag.Flag[bool]{
			Name:     "stream-response",
			Usage:    "Whether to stream the response via SSE",
			BodyPath: "streamResponse",
		},
		&requestflag.Flag[any]{
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
}, map[string][]requestflag.HasOuterFlag{
	"options": {
		&requestflag.InnerFlag[any]{
			Name:       "options.model",
			Usage:      "Model name string with provider prefix (e.g., 'openai/gpt-5-nano', 'anthropic/claude-4.5-opus')",
			InnerField: "model",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "options.timeout",
			Usage:      "Timeout in ms for the action",
			InnerField: "timeout",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "options.variables",
			Usage:      "Variables to substitute in the action instruction",
			InnerField: "variables",
		},
	},
})

var sessionsEnd = cli.Command{
	Name:    "end",
	Usage:   "Terminates the browser session and releases all associated resources.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Usage:    "Unique session identifier",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "-force-body",
			BodyPath: "_forceBody",
		},
		&requestflag.Flag[any]{
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

var sessionsExecute = requestflag.WithInnerFlags(cli.Command{
	Name:    "execute",
	Usage:   "Runs an autonomous AI agent that can perform complex multi-step browser tasks.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Usage:    "Unique session identifier",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "agent-config",
			Required: true,
			BodyPath: "agentConfig",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "execute-options",
			Required: true,
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
		&requestflag.Flag[any]{
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
}, map[string][]requestflag.HasOuterFlag{
	"agent-config": {
		&requestflag.InnerFlag[bool]{
			Name:       "agent-config.cua",
			Usage:      "Enable Computer Use Agent mode",
			InnerField: "cua",
		},
		&requestflag.InnerFlag[any]{
			Name:       "agent-config.model",
			Usage:      "Model name string with provider prefix (e.g., 'openai/gpt-5-nano', 'anthropic/claude-4.5-opus')",
			InnerField: "model",
		},
		&requestflag.InnerFlag[string]{
			Name:       "agent-config.provider",
			Usage:      "AI provider for the agent (legacy, use model: openai/gpt-5-nano instead)",
			InnerField: "provider",
		},
		&requestflag.InnerFlag[string]{
			Name:       "agent-config.system-prompt",
			Usage:      "Custom system prompt for the agent",
			InnerField: "systemPrompt",
		},
	},
	"execute-options": {
		&requestflag.InnerFlag[string]{
			Name:       "execute-options.instruction",
			Usage:      "Natural language instruction for the agent",
			InnerField: "instruction",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "execute-options.highlight-cursor",
			Usage:      "Whether to visually highlight the cursor during execution",
			InnerField: "highlightCursor",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "execute-options.max-steps",
			Usage:      "Maximum number of steps the agent can take",
			InnerField: "maxSteps",
		},
	},
})

var sessionsExtract = requestflag.WithInnerFlags(cli.Command{
	Name:    "extract",
	Usage:   "Extracts structured data from the current page using AI-powered analysis.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Usage:    "Unique session identifier",
			Required: true,
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
		&requestflag.Flag[map[string]any]{
			Name:     "options",
			BodyPath: "options",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "schema",
			Usage:    "JSON Schema defining the structure of data to extract",
			BodyPath: "schema",
		},
		&requestflag.Flag[bool]{
			Name:     "stream-response",
			Usage:    "Whether to stream the response via SSE",
			BodyPath: "streamResponse",
		},
		&requestflag.Flag[any]{
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
}, map[string][]requestflag.HasOuterFlag{
	"options": {
		&requestflag.InnerFlag[any]{
			Name:       "options.model",
			Usage:      "Model name string with provider prefix (e.g., 'openai/gpt-5-nano', 'anthropic/claude-4.5-opus')",
			InnerField: "model",
		},
		&requestflag.InnerFlag[string]{
			Name:       "options.selector",
			Usage:      "CSS selector to scope extraction to a specific element",
			InnerField: "selector",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "options.timeout",
			Usage:      "Timeout in ms for the extraction",
			InnerField: "timeout",
		},
	},
})

var sessionsNavigate = requestflag.WithInnerFlags(cli.Command{
	Name:    "navigate",
	Usage:   "Navigates the browser to the specified URL.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Usage:    "Unique session identifier",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "url",
			Usage:    "URL to navigate to",
			Required: true,
			BodyPath: "url",
		},
		&requestflag.Flag[string]{
			Name:     "frame-id",
			Usage:    "Target frame ID for the navigation",
			BodyPath: "frameId",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "options",
			BodyPath: "options",
		},
		&requestflag.Flag[bool]{
			Name:     "stream-response",
			Usage:    "Whether to stream the response via SSE",
			BodyPath: "streamResponse",
		},
		&requestflag.Flag[any]{
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
}, map[string][]requestflag.HasOuterFlag{
	"options": {
		&requestflag.InnerFlag[string]{
			Name:       "options.referer",
			Usage:      "Referer header to send with the request",
			InnerField: "referer",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "options.timeout",
			Usage:      "Timeout in ms for the navigation",
			InnerField: "timeout",
		},
		&requestflag.InnerFlag[string]{
			Name:       "options.wait-until",
			Usage:      "When to consider navigation complete",
			InnerField: "waitUntil",
		},
	},
})

var sessionsObserve = requestflag.WithInnerFlags(cli.Command{
	Name:    "observe",
	Usage:   "Identifies and returns available actions on the current page that match the\ngiven instruction.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Usage:    "Unique session identifier",
			Required: true,
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
		&requestflag.Flag[map[string]any]{
			Name:     "options",
			BodyPath: "options",
		},
		&requestflag.Flag[bool]{
			Name:     "stream-response",
			Usage:    "Whether to stream the response via SSE",
			BodyPath: "streamResponse",
		},
		&requestflag.Flag[any]{
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
}, map[string][]requestflag.HasOuterFlag{
	"options": {
		&requestflag.InnerFlag[any]{
			Name:       "options.model",
			Usage:      "Model name string with provider prefix (e.g., 'openai/gpt-5-nano', 'anthropic/claude-4.5-opus')",
			InnerField: "model",
		},
		&requestflag.InnerFlag[string]{
			Name:       "options.selector",
			Usage:      "CSS selector to scope observation to a specific element",
			InnerField: "selector",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "options.timeout",
			Usage:      "Timeout in ms for the observation",
			InnerField: "timeout",
		},
	},
})

var sessionsStart = requestflag.WithInnerFlags(cli.Command{
	Name:    "start",
	Usage:   "Creates a new browser session with the specified configuration. Returns a\nsession ID used for all subsequent operations.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "model-name",
			Usage:    "Model name to use for AI operations",
			Required: true,
			BodyPath: "modelName",
		},
		&requestflag.Flag[float64]{
			Name:     "act-timeout-ms",
			Usage:    "Timeout in ms for act operations (deprecated, v2 only)",
			BodyPath: "actTimeoutMs",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "browser",
			BodyPath: "browser",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "browserbase-session-create-params",
			BodyPath: "browserbaseSessionCreateParams",
		},
		&requestflag.Flag[string]{
			Name:     "browserbase-session-id",
			Usage:    "Existing Browserbase session ID to resume",
			BodyPath: "browserbaseSessionID",
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
		&requestflag.Flag[float64]{
			Name:     "verbose",
			Usage:    "Logging verbosity level (0=quiet, 1=normal, 2=debug)",
			BodyPath: "verbose",
		},
		&requestflag.Flag[bool]{
			Name:     "wait-for-captcha-solves",
			Usage:    "Wait for captcha solves (deprecated, v2 only)",
			BodyPath: "waitForCaptchaSolves",
		},
		&requestflag.Flag[any]{
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
}, map[string][]requestflag.HasOuterFlag{
	"browser": {
		&requestflag.InnerFlag[string]{
			Name:       "browser.cdp-url",
			Usage:      "Chrome DevTools Protocol URL for connecting to existing browser",
			InnerField: "cdpUrl",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "browser.launch-options",
			InnerField: "launchOptions",
		},
		&requestflag.InnerFlag[string]{
			Name:       "browser.type",
			Usage:      "Browser type to use",
			InnerField: "type",
		},
	},
	"browserbase-session-create-params": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "browserbase-session-create-params.browser-settings",
			InnerField: "browserSettings",
		},
		&requestflag.InnerFlag[string]{
			Name:       "browserbase-session-create-params.extension-id",
			InnerField: "extensionId",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "browserbase-session-create-params.keep-alive",
			InnerField: "keepAlive",
		},
		&requestflag.InnerFlag[string]{
			Name:       "browserbase-session-create-params.project-id",
			InnerField: "projectId",
		},
		&requestflag.InnerFlag[any]{
			Name:       "browserbase-session-create-params.proxies",
			InnerField: "proxies",
		},
		&requestflag.InnerFlag[string]{
			Name:       "browserbase-session-create-params.region",
			InnerField: "region",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "browserbase-session-create-params.timeout",
			InnerField: "timeout",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "browserbase-session-create-params.user-metadata",
			InnerField: "userMetadata",
		},
	},
})

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

	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	stream := client.Sessions.ActStreaming(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	return ShowJSONIterator(os.Stdout, "sessions act", stream, format, transform)
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
		ApplicationJSON,
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

	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	stream := client.Sessions.ExecuteStreaming(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	return ShowJSONIterator(os.Stdout, "sessions execute", stream, format, transform)
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

	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	stream := client.Sessions.ExtractStreaming(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	return ShowJSONIterator(os.Stdout, "sessions extract", stream, format, transform)
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

	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	stream := client.Sessions.ObserveStreaming(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	return ShowJSONIterator(os.Stdout, "sessions observe", stream, format, transform)
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
