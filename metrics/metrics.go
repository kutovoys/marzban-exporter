package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	// Node-related metrics
	NodesStatus = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "marzban_nodes_status",
			Help: "Status of Marzban nodes",
		},
		[]string{"name", "address", "id", "usage_coefficient", "xray_version", "status"},
	)
	NodesUplink = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "marzban_nodes_uplink",
			Help: "Uplink traffic of Marzban nodes",
		},
		[]string{"id", "name"},
	)
	NodesDownlink = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "marzban_nodes_downlink",
			Help: "Downlink traffic of Marzban nodes",
		},
		[]string{"id", "name"},
	)

	// User-related metrics
	UserDataLimit = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "marzban_user_data_limit",
			Help: "Data limit of the user",
		},
		[]string{"data_limit_reset_strategy", "note", "username", "status", "last_user_agent"},
	)
	UserUsedTraffic = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "marzban_user_used_traffic",
			Help: "Used traffic of the user",
		},
		[]string{"data_limit_reset_strategy", "note", "username", "status", "last_user_agent"},
	)
	UserLifetimeUsedTraffic = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "marzban_user_lifetime_used_traffic",
			Help: "Lifetime used traffic of the user",
		},
		[]string{"data_limit_reset_strategy", "note", "username", "status", "last_user_agent"},
	)
	UserExpirationDate = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "marzban_user_expiration_date",
			Help: "User's subscription expiration date",
		},
		[]string{"data_limit_reset_strategy", "note", "username", "status", "last_user_agent"},
	)
	UserOnline = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "marzban_user_online",
			Help: "Whether a user is online within the last 2 minutes",
		},
		[]string{"note", "username", "status", "last_user_agent"},
	)

	// System-related metrics
	CoreStarted = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "marzban_core_started",
			Help: "Indicates if Marzban core is started",
		},
		[]string{"version"},
	)
	MemTotal = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "marzban_panel_mem_total",
		Help: "Panel Total memory",
	})
	MemUsed = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "marzban_panel_mem_used",
		Help: "Panel Used memory",
	})
	CpuCores = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "marzban_panel_cpu_cores",
		Help: "Panel Number of CPU cores",
	})
	CpuUsage = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "marzban_panel_cpu_usage",
		Help: "Panel CPU usage",
	})
	TotalUser = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "marzban_panel_total_users",
		Help: "Total number of users",
	})
	UsersActive = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "marzban_panel_users_active",
		Help: "Number of active users",
	})
	IncomingBandwidth = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "marzban_all_incoming_bandwidth",
		Help: "Incoming bandwidth with all nodes",
	})
	OutgoingBandwidth = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "marzban_all_outgoing_bandwidth",
		Help: "Outgoing bandwidth with all nodes",
	})
	IncomingBandwidthSpeed = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "marzban_panel_incoming_bandwidth_speed",
		Help: "Panel incoming bandwidth speed",
	})
	OutgoingBandwidthSpeed = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "marzban_panel_outgoing_bandwidth_speed",
		Help: "Panel outgoing bandwidth speed",
	})
)
