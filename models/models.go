package models

type AuthTokenResponse struct {
	AccessToken string `json:"access_token"`
}

type Node struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	ID      int    `json:"id"`
	Status  string `json:"status"`
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
	Expire                 *string                `json:"expire"`
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
