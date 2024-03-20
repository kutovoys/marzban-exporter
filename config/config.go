package config

import (
	"marzban-exporter/models"

	"github.com/alecthomas/kong"
)

var CLIConfig models.CLI

func Parse() {
	ctx := kong.Parse(&CLIConfig,
		kong.Name("marzban-exporter"),
		kong.Description("A command-line application for exporting Marzban metrics."),
	)
	// Use ctx if needed
	_ = ctx
}
