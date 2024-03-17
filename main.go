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
			if config.ProtectedMetrics {
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
	token, err := api.GetAuthToken()
	if err != nil {
		log.Println("Error getting auth token:", err)
		return
	}

	s := gocron.NewScheduler(time.Local)

	s.Every(config.UpdateInterval).Seconds().Do(func() {
		go api.FetchNodesStatus(token)
	})
	s.Every(config.UpdateInterval).Seconds().Do(func() {
		go api.FetchNodesUsage(token)
	})
	s.Every(config.UpdateInterval).Seconds().Do(func() {
		go api.FetchSystemStats(token)
	})
	s.Every(config.UpdateInterval).Seconds().Do(func() {
		go api.FetchUsersStats(token)
	})

	s.Every(config.UpdateInterval).Seconds().Do(func() {
		go api.FetchCoreStatus(token)
	})

	go s.StartAsync()

	http.Handle("/metrics", BasicAuthMiddleware(config.MetricsUsername,
		config.MetricsPassword)(promhttp.Handler()))
	log.Printf("Starting server on :%s", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
