package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	// User-related metrics
	OnlineUsersCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "total_online_users",
			Help: "Total number of online users",
		},
	)
	// System-related metrics
	XrayVersion = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "xray_version",
			Help: "XRay core version",
		},
		[]string{"version"},
	)
)
