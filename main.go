package main

import (
	"log"
	"marzban-exporter/api"
	"marzban-exporter/config"
	"marzban-exporter/metrics"
	"net/http"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	prometheus.MustRegister(
		// Node-related metrics
		metrics.NodesStatus,
		metrics.NodesUplink,
		metrics.NodesDownlink,

		// User-related metrics
		metrics.UserDataLimit,
		metrics.UserUsedTraffic,
		metrics.UserLifetimeUsedTraffic,
		metrics.UserExpirationDate,
		metrics.UserOnline,

		// System-related metrics
		metrics.CoreStarted,
		metrics.MemTotal,
		metrics.MemUsed,
		metrics.CpuCores,
		metrics.CpuUsage,
		metrics.TotalUser,
		metrics.UsersActive,
		metrics.IncomingBandwidth,
		metrics.OutgoingBandwidth,
		metrics.IncomingBandwidthSpeed,
		metrics.OutgoingBandwidthSpeed,
	)
}

func BasicAuthMiddleware(username, password string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if config.CLIConfig.ProtectedMetrics {
				user, pass, ok := r.BasicAuth()
				if !ok || user != username || pass != password {
					w.Header().Set("WWW-Authenticate", `Basic realm="metrics"`)
					http.Error(w, "Unauthorized.", http.StatusUnauthorized)
					return
				}
			}
			next.ServeHTTP(w, r)
		})
	}
}

func main() {
	config.Parse()

	s := gocron.NewScheduler(time.Local)

	s.Every(config.CLIConfig.UpdateInterval).Seconds().Do(func() {
		token, err := api.GetAuthToken()
		if err != nil {
			log.Println("Error getting auth token:", err)
			return
		}

		log.Print("Starting to collect metrics")

		log.Print("Collecting NodesStatus metrics")
		api.FetchNodesStatus(token)
		log.Print("Finished collecting NodesStatus metrics")

		log.Print("Collecting NodesUsage metrics")
		api.FetchNodesUsage(token)
		log.Print("Finished collecting NodesUsage metrics")

		log.Print("Collecting SystemStats metrics")
		api.FetchSystemStats(token)
		log.Print("Finished collecting SystemStats metrics")

		log.Print("Collecting UsersStats metrics")
		api.FetchUsersStats(token)
		log.Print("Finished collecting UsersStats metrics")

		log.Print("Collecting CoreStatus metrics")
		api.FetchCoreStatus(token)
		log.Print("Finished collecting CoreStatus metrics")

		log.Print("Finished all metric collection")
	})

	go s.StartAsync()

	http.Handle("/metrics", BasicAuthMiddleware(config.CLIConfig.MetricsUsername,
		config.CLIConfig.MetricsPassword)(promhttp.Handler()))
	log.Printf("Starting server on :%s", config.CLIConfig.Port)
	log.Fatal(http.ListenAndServe(":"+config.CLIConfig.Port, nil))
}
