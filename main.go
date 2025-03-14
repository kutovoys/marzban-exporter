package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"x-ui-exporter/api"
	"x-ui-exporter/config"
	"x-ui-exporter/metrics"

	"github.com/go-co-op/gocron"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	version = "unknown"
	commit  = "unknown"
)

func init() { //
	prometheus.MustRegister(
		// User-related metrics
		metrics.OnlineUsersCount,

		// System-related metrics
		metrics.XrayVersion,
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
	config.Parse(version, commit)

	fmt.Println("3X-UI Exporter (Fork)", version)

	s := gocron.NewScheduler(time.Local)

	s.Every(config.CLIConfig.UpdateInterval).Seconds().Do(func() {
		token, err := api.GetAuthToken()
		if err != nil {
			log.Println("Error getting auth token:", err)
			return
		}

		log.Print("Starting to collect metrics")

		log.Print("Collecting UsersStats metrics")
		api.FetchOnlineUsersCount(token)
		log.Print("Finished collecting UsersStats metrics")

		log.Print("Collecting ServerStatus metrics")
		api.FetchServerStatus(token)
		log.Print("Finished collecting ServerStatus metrics")

		log.Print("Finished all metric collection")
	})

	go s.StartAsync()

	http.Handle("/metrics", BasicAuthMiddleware(config.CLIConfig.MetricsUsername,
		config.CLIConfig.MetricsPassword)(promhttp.Handler()))
	log.Printf("Starting server on :%s", config.CLIConfig.Port)
	log.Fatal(http.ListenAndServe(":"+config.CLIConfig.Port, nil))
}
