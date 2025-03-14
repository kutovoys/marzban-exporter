package models

import (
	"encoding/json"
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

type LoginResponse struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}

type ObjectResponse struct {
	Success bool            `json:"success"`
	Msg     string          `json:"msg"`
	Obj     json.RawMessage `json:"obj"`
}

type ProcessState string

const (
	Running ProcessState = "running"
	Stop    ProcessState = "stop"
	Error   ProcessState = "error"
)

type ServerStatusResponse struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Obj     *struct {
		Cpu         float64 `json:"cpu"`
		CpuCores    int     `json:"cpuCores"`
		CpuSpeedMhz float64 `json:"cpuSpeedMhz"`
		Mem         struct {
			Current uint64 `json:"current"`
			Total   uint64 `json:"total"`
		} `json:"mem"`
		Swap struct {
			Current uint64 `json:"current"`
			Total   uint64 `json:"total"`
		} `json:"swap"`
		Disk struct {
			Current uint64 `json:"current"`
			Total   uint64 `json:"total"`
		} `json:"disk"`
		Xray struct {
			State    ProcessState `json:"state"`
			ErrorMsg string       `json:"errorMsg"`
			Version  string       `json:"version"`
		} `json:"xray"`
		Uptime   uint64    `json:"uptime"`
		Loads    []float64 `json:"loads"`
		TcpCount int       `json:"tcpCount"`
		UdpCount int       `json:"udpCount"`
		NetIO    struct {
			Up   uint64 `json:"up"`
			Down uint64 `json:"down"`
		} `json:"netIO"`
		NetTraffic struct {
			Sent uint64 `json:"sent"`
			Recv uint64 `json:"recv"`
		} `json:"netTraffic"`
		PublicIP struct {
			IPv4 string `json:"ipv4"`
			IPv6 string `json:"ipv6"`
		} `json:"publicIP"`
		AppStats struct {
			Threads uint32 `json:"threads"`
			Mem     uint64 `json:"mem"`
			Uptime  uint64 `json:"uptime"`
		} `json:"appStats"`
	} `json:"obj"`
}
