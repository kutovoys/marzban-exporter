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
	token, err := api.GetAuthToken()
	if err != nil {
		log.Println("Error getting auth token:", err)
		return
	}

	s := gocron.NewScheduler(time.Local)

	s.Every(config.CLIConfig.UpdateInterval).Seconds().Do(func() {
		log.Print("Starting to collect NodesStatus metrics")
		go api.FetchNodesStatus(token)
		log.Print("Finished collecting NodesStatus metrics")
	})
	s.Every(config.CLIConfig.UpdateInterval).Seconds().Do(func() {
		log.Print("Starting to collect NodesUsage metrics")
		go api.FetchNodesUsage(token)
		log.Print("Finished collecting NodesUsage metrics")
	})
	s.Every(config.CLIConfig.UpdateInterval).Seconds().Do(func() {
		log.Print("Starting to collect SystemStats metrics")
		go api.FetchSystemStats(token)
		log.Print("Finished collecting SystemStats metrics")
	})
	s.Every(config.CLIConfig.UpdateInterval).Seconds().Do(func() {
		log.Print("Starting to collect UsersStats metrics")
		go api.FetchUsersStats(token)
		log.Print("Finished collecting UsersStats metrics")
	})
	s.Every(config.CLIConfig.UpdateInterval).Seconds().Do(func() {
		log.Print("Starting to collect CoreStatus metrics")
		go api.FetchCoreStatus(token)
		log.Print("Finished collecting CoreStatus metrics")
	})

	go s.StartAsync()

	http.Handle("/metrics", BasicAuthMiddleware(config.CLIConfig.MetricsUsername,
		config.CLIConfig.MetricsPassword)(promhttp.Handler()))
	log.Printf("Starting server on :%s", config.CLIConfig.Port)
	log.Fatal(http.ListenAndServe(":"+config.CLIConfig.Port, nil))
}
