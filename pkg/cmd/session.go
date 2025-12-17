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
