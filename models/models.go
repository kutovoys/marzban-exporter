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
	InactivityTime   int         `name:"inactivity-time" help:"Time in minutes after which a user is considered inactive" default:"2" env:"INACTIVITY_TIME"`
	BaseURL          string      `name:"marzban-base-url" help:"Marzban panel base URL" env:"MARZBAN_BASE_URL"`
	ApiUsername      string      `name:"marzban-username" help:"Marzban panel username" env:"MARZBAN_USERNAME" required:""`
	ApiPassword      string      `name:"marzban-password" help:"Marzban panel password" env:"MARZBAN_PASSWORD" required:""`
	SocketPath       string      `name:"marzban-socket" help:"Path to Marzban Unix Domain Socket" env:"MARZBAN_SOCKET"`
	Version          VersionFlag `name:"version" help:"Print version information and quit"`
}

type VersionFlag string

func (v VersionFlag) Decode(ctx *kong.DecodeContext) error { return nil }
func (v VersionFlag) IsBool() bool                         { return true }
func (v VersionFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error {
	fmt.Println("Marzban Exporter")
	fmt.Printf("Version:\t %s\n", vars["version"])
	fmt.Printf("Commit:\t %s\n", vars["commit"])
	fmt.Printf("GitHub: https://github.com/kutovoys/marzban-exporter\n")
	app.Exit(0)
	return nil
}

type AuthTokenResponse struct {
	AccessToken string `json:"access_token"`
}

type Node struct {
	Name      string  `json:"name"`
	Address   string  `json:"address"`
	ID        int     `json:"id"`
	Status    string  `json:"status"`
	UsageCoef float64 `json:"usage_coefficient"`
	XrayVer   string  `json:"xray_version"`
}

type NodeUsage struct {
	NodeID   *int   `json:"node_id"`
	NodeName string `json:"node_name"`
	Uplink   int64  `json:"uplink"`
	Downlink int64  `json:"downlink"`
}

type SystemStats struct {
	Version                string  `json:"version"`
	MemTotal               float64 `json:"mem_total"`
	MemUsed                float64 `json:"mem_used"`
	CpuCores               float64 `json:"cpu_cores"`
	CpuUsage               float64 `json:"cpu_usage"`
	TotalUser              float64 `json:"total_user"`
	UsersActive            float64 `json:"users_active"`
	IncomingBandwidth      float64 `json:"incoming_bandwidth"`
	OutgoingBandwidth      float64 `json:"outgoing_bandwidth"`
	IncomingBandwidthSpeed float64 `json:"incoming_bandwidth_speed"`
	OutgoingBandwidthSpeed float64 `json:"outgoing_bandwidth_speed"`
}

type User struct {
	Proxies                map[string]interface{} `json:"proxies"`
	Expire                 float64                `json:"expire"`
	DataLimit              float64                `json:"data_limit"`
	DataLimitResetStrategy string                 `json:"data_limit_reset_strategy"`
	Inbounds               map[string][]string    `json:"inbounds"`
	Note                   string                 `json:"note"`
	SubUpdatedAt           string                 `json:"sub_updated_at"`
	SubLastUserAgent       string                 `json:"sub_last_user_agent"`
	OnlineAt               *string                `json:"online_at"`
	OnHoldExpireDuration   *string                `json:"on_hold_expire_duration"`
	OnHoldTimeout          *string                `json:"on_hold_timeout"`
	Username               string                 `json:"username"`
	Status                 string                 `json:"status"`
	UsedTraffic            float64                `json:"used_traffic"`
	LifetimeUsedTraffic    float64                `json:"lifetime_used_traffic"`
	CreatedAt              string                 `json:"created_at"`
	Links                  []string               `json:"links"`
	SubscriptionUrl        string                 `json:"subscription_url"`
	ExcludedInbounds       map[string][]string    `json:"excluded_inbounds"`
}

type UsersResponse struct {
	Users []User `json:"users"`
	Total int    `json:"total"`
}

type UsageResponse struct {
	Usages []NodeUsage `json:"usages"`
}
