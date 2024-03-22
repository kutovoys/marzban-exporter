package config

import (
	"errors"
	"fmt"
	"marzban-exporter/models"
	"os"

	"github.com/alecthomas/kong"
)

var CLIConfig models.CLI

func Parse() {
	kong.Parse(&CLIConfig,
		kong.Name("marzban-exporter"),
		kong.Description("A command-line application for exporting Marzban metrics."),
	)
	if err := validate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}

func validate() error {
	if CLIConfig.BaseURL == "" && CLIConfig.SocketPath == "" {
		return errors.New("marzban-exporter: error: either --marzban-base-url or --marzban-socket must be provided")
	}
	return nil
}
