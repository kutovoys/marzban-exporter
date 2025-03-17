package models

import (
	"fmt"

	"github.com/alecthomas/kong"
)

type CLI struct {
	Port             string      `name:"metrics-port" help:"Port to listen on" default:"9090" env:"METRICS_PORT"`
	ProtectedMetrics bool        `name:"metrics-protected" help:"Whether metrics are protected by basic auth" default:"false" env:"METRICS_PROTECTED"`
	MetricsUsername  string      `name:"metrics-username" help:"Username for metrics if protected by basic auth" default:"metricsUser" env:"METRICS_USERNAME"`
	MetricsPassword  string      `name:"metrics-password" help:"Password for metrics if protected by basic auth" default:"MetricsVeryHardPassword" env:"METRICS_PASSWORD"`
	UpdateInterval   int         `name:"update-interval" help:"Interval for metrics update in seconds" default:"60" env:"UPDATE_INTERVAL"`
	TimeZone         string      `name:"timezone" help:"Timezone used in the application" default:"UTC" env:"TIMEZONE"`
	BaseURL          string      `name:"panel-base-url" help:"Panel base URL" env:"PANEL_BASE_URL"`
	ApiUsername      string      `name:"panel-username" help:"Panel username" env:"PANEL_USERNAME" required:""`
	ApiPassword      string      `name:"panel-password" help:"Panel password" env:"PANEL_PASSWORD" required:""`
	Version          VersionFlag `name:"version" help:"Print version information and quit"`
}

type VersionFlag string

func (v VersionFlag) Decode(ctx *kong.DecodeContext) error { return nil }
func (v VersionFlag) IsBool() bool                         { return true }
func (v VersionFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error {
	fmt.Println("3X-UI Exporter (Fork)")
	fmt.Printf("Version:\t %s\n", vars["version"])
	fmt.Printf("Commit:\t %s\n", vars["commit"])
	fmt.Printf("Github (Marzban): https://github.com/kutovoys/marzban-exporter\n")
	fmt.Printf("GitHub (3X-UI Fork): https://github.com/hteppl/3x-ui-exporter\n")
	app.Exit(0)
	return nil
}
