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
)
