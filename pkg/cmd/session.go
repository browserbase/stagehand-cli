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
		&requestflag.Flag[any]{
			Name: "id",
		},
		&requestflag.Flag[any]{
			Name:     "body",
			BodyRoot: true,
		},
		&requestflag.Flag[any]{
			Name:       "x-language",
			HeaderPath: "x-language",
		},
		&requestflag.Flag[any]{
			Name:       "x-sdk-version",
			HeaderPath: "x-sdk-version",
		},
		&requestflag.Flag[any]{
			Name:       "x-sent-at",
			HeaderPath: "x-sent-at",
		},
		&requestflag.Flag[any]{
			Name:       "x-stream-response",
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
		&requestflag.Flag[any]{
			Name: "id",
		},
		&requestflag.Flag[any]{
			Name:       "x-language",
			HeaderPath: "x-language",
		},
		&requestflag.Flag[any]{
			Name:       "x-sdk-version",
			HeaderPath: "x-sdk-version",
		},
		&requestflag.Flag[any]{
			Name:       "x-sent-at",
			HeaderPath: "x-sent-at",
		},
		&requestflag.Flag[any]{
			Name:       "x-stream-response",
			HeaderPath: "x-stream-response",
		},
	},
	Action:          handleSessionsEnd,
	HideHelpCommand: true,
}

var sessionsExecuteAgent = cli.Command{
	Name:  "execute-agent",
	Usage: "Runs an autonomous AI agent that can perform complex multi-step browser tasks.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name: "id",
		},
		&requestflag.Flag[any]{
			Name:     "body",
			BodyRoot: true,
		},
		&requestflag.Flag[any]{
			Name:       "x-language",
			HeaderPath: "x-language",
		},
		&requestflag.Flag[any]{
			Name:       "x-sdk-version",
			HeaderPath: "x-sdk-version",
		},
		&requestflag.Flag[any]{
			Name:       "x-sent-at",
			HeaderPath: "x-sent-at",
		},
		&requestflag.Flag[any]{
			Name:       "x-stream-response",
			HeaderPath: "x-stream-response",
		},
	},
	Action:          handleSessionsExecuteAgent,
	HideHelpCommand: true,
}

var sessionsExtract = cli.Command{
	Name:  "extract",
	Usage: "Extracts structured data from the current page using AI-powered analysis.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name: "id",
		},
		&requestflag.Flag[any]{
			Name:     "body",
			BodyRoot: true,
		},
		&requestflag.Flag[any]{
			Name:       "x-language",
			HeaderPath: "x-language",
		},
		&requestflag.Flag[any]{
			Name:       "x-sdk-version",
			HeaderPath: "x-sdk-version",
		},
		&requestflag.Flag[any]{
			Name:       "x-sent-at",
			HeaderPath: "x-sent-at",
		},
		&requestflag.Flag[any]{
			Name:       "x-stream-response",
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
		&requestflag.Flag[any]{
			Name: "id",
		},
		&requestflag.Flag[any]{
			Name:     "body",
			BodyRoot: true,
		},
		&requestflag.Flag[any]{
			Name:       "x-language",
			HeaderPath: "x-language",
		},
		&requestflag.Flag[any]{
			Name:       "x-sdk-version",
			HeaderPath: "x-sdk-version",
		},
		&requestflag.Flag[any]{
			Name:       "x-sent-at",
			HeaderPath: "x-sent-at",
		},
		&requestflag.Flag[any]{
			Name:       "x-stream-response",
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
		&requestflag.Flag[any]{
			Name: "id",
		},
		&requestflag.Flag[any]{
			Name:     "body",
			BodyRoot: true,
		},
		&requestflag.Flag[any]{
			Name:       "x-language",
			HeaderPath: "x-language",
		},
		&requestflag.Flag[any]{
			Name:       "x-sdk-version",
			HeaderPath: "x-sdk-version",
		},
		&requestflag.Flag[any]{
			Name:       "x-sent-at",
			HeaderPath: "x-sent-at",
		},
		&requestflag.Flag[any]{
			Name:       "x-stream-response",
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
		&requestflag.Flag[any]{
			Name:     "body",
			BodyRoot: true,
		},
		&requestflag.Flag[any]{
			Name:       "x-language",
			HeaderPath: "x-language",
		},
		&requestflag.Flag[any]{
			Name:       "x-sdk-version",
			HeaderPath: "x-sdk-version",
		},
		&requestflag.Flag[any]{
			Name:       "x-sent-at",
			HeaderPath: "x-sent-at",
		},
		&requestflag.Flag[any]{
			Name:       "x-stream-response",
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
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Sessions.Act(
		ctx,
		cmd.Value("id").(any),
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
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Sessions.End(
		ctx,
		cmd.Value("id").(any),
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

func handleSessionsExecuteAgent(ctx context.Context, cmd *cli.Command) error {
	client := stagehand.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
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
		cmd.Value("id").(any),
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
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Sessions.Extract(
		ctx,
		cmd.Value("id").(any),
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
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Sessions.Navigate(
		ctx,
		cmd.Value("id").(any),
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
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Sessions.Observe(
		ctx,
		cmd.Value("id").(any),
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
