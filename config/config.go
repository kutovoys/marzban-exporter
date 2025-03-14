package config

import (
	"errors"
	"fmt"
	"marzban-exporter/models"
	"os"

	"github.com/alecthomas/kong"
)

var CLIConfig models.CLI

func Parse(version, commit string) {
	ctx := kong.Parse(&CLIConfig,
		kong.Name("x-ui-exporter"),
		kong.Description("A command-line application for exporting 3X-UI metrics."),
		kong.Vars{
			"version": version,
			"commit":  commit,
		},
	)

	if err := validate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		ctx.Exit(2)
	}
}

func validate() error {
	if CLIConfig.BaseURL == "" {
		return errors.New("x-ui-exporter: error: --panel-base-url must be provided")
	}
	if CLIConfig.ApiUsername == "" {
		return errors.New("x-ui-exporter: error: --panel-username must be provided")
	}
	if CLIConfig.ApiPassword == "" {
		return errors.New("x-ui-exporter: error: --panel-password must be provided")
	}
	return nil
}
