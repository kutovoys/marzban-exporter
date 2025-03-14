package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	// User-related metrics
	OnlineUsersCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "x_ui_total_online_users",
			Help: "Total number of online users",
		},
	)
	// Inbound-related metrics
	InboundUp = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "x_ui_inbound_up_bytes",
			Help: "Total uploaded bytes per inbound",
		}, []string{"id", "remark"},
	)
	InboundDown = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "x_ui_inbound_down_bytes",
			Help: "Total downloaded bytes per inbound",
		}, []string{"id", "remark"},
	)
	// Client-related metrics
	ClientUp = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "x_ui_client_up_bytes",
			Help: "Total uploaded bytes per client",
		}, []string{"id", "email"},
	)
	ClientDown = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "x_ui_client_down_bytes",
			Help: "Total downloaded bytes per client",
		}, []string{"id", "email"},
	)
	// System-related metrics
	XrayVersion = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "x_ui_xray_version",
			Help: "XRay version used by 3X-UI",
		},
		[]string{"version"},
	)
	PanelThreads = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "x_ui_panel_threads",
			Help: "3X-UI panel threads",
		},
	)
	PanelMemory = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "x_ui_panel_memory",
			Help: "3X-UI panel memory usage",
		},
	)
	PanelUptime = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "x_ui_panel_uptime",
			Help: "3X-UI panel uptime",
		},
	)
)
