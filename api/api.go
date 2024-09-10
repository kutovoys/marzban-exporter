package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"marzban-exporter/config"
	"marzban-exporter/metrics"
	"marzban-exporter/models"
	"net"
	"net/http"
	"sync"
	"time"
)

var tokenCache struct {
	Token     string
	ExpiresAt time.Time
	sync.Mutex
}

func GetAuthToken() (string, error) {
	tokenCache.Lock()
	defer tokenCache.Unlock()

	if tokenCache.Token != "" && time.Now().Before(tokenCache.ExpiresAt) {
		return tokenCache.Token, nil
	}

	path := "/api/admin/token"
	data := []byte(fmt.Sprintf("username=%s&password=%s", config.CLIConfig.ApiUsername, config.CLIConfig.ApiPassword))
	httpClient := createHttpClient()
	req, err := createRequest("POST", path, "")
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Body = io.NopCloser(bytes.NewBuffer(data))

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making auth request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading auth response body: %v", err)
	}

	var tokenResponse struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		return "", fmt.Errorf("error unmarshaling auth token response: %v", err)
	}

	tokenCache.Token = tokenResponse.AccessToken
	tokenCache.ExpiresAt = time.Now().Add(12 * time.Hour)

	return tokenCache.Token, nil
}

func FetchNodesStatus(token string) {
	path := "/api/nodes"
	body, err := sendRequest(path, token)
	if err != nil {
		log.Println("Error making request for nodes status:", err)
		return
	}

	var nodes []models.Node
	if err := json.Unmarshal(body, &nodes); err != nil {
		log.Println("Error unmarshaling response:", err)
		return
	}

	for _, node := range nodes {
		status := 0.0
		if node.Status == "connected" {
			status = 1.0
		}
		metrics.NodesStatus.WithLabelValues(fmt.Sprintf("%d", node.ID), node.Name).Set(status)
	}
}

func FetchNodesUsage(token string) {
	path := "/api/nodes/usage"
	body, err := sendRequest(path, token)
	if err != nil {
		log.Println("Error making request for nodes usage:", err)
		return
	}

	var usageResponse models.UsageResponse
	if err := json.Unmarshal(body, &usageResponse); err != nil {
		log.Println("Error unmarshaling response:", err)
		return
	}

	for _, usage := range usageResponse.Usages {
		id := "0"
		if usage.NodeID != nil {
			id = fmt.Sprintf("%d", *usage.NodeID)
		}
		metrics.NodesUplink.WithLabelValues(id, usage.NodeName).Set(float64(usage.Uplink))
		metrics.NodesDownlink.WithLabelValues(id, usage.NodeName).Set(float64(usage.Downlink))
	}
}

func FetchSystemStats(token string) {
	path := "/api/system"
	body, err := sendRequest(path, token)
	if err != nil {
		log.Println("Error making request for system stats:", err)
		return
	}

	var stats models.SystemStats
	if err := json.Unmarshal(body, &stats); err != nil {
		log.Println("Error unmarshaling response:", err)
		return
	}

	metrics.MemTotal.Set(stats.MemTotal)
	metrics.MemUsed.Set(stats.MemUsed)
	metrics.CpuCores.Set(stats.CpuCores)
	metrics.CpuUsage.Set(stats.CpuUsage)
	metrics.TotalUser.Set(stats.TotalUser)
	metrics.UsersActive.Set(stats.UsersActive)
	metrics.IncomingBandwidth.Set(stats.IncomingBandwidth)
	metrics.OutgoingBandwidth.Set(stats.OutgoingBandwidth)
	metrics.IncomingBandwidthSpeed.Set(stats.IncomingBandwidthSpeed)
	metrics.OutgoingBandwidthSpeed.Set(stats.OutgoingBandwidthSpeed)
}

func FetchCoreStatus(token string) {
	path := "/api/core"
	body, err := sendRequest(path, token)
	if err != nil {
		log.Println("Error making request for core status:", err)
		return
	}

	var coreResponse struct {
		Version       string `json:"version"`
		Started       bool   `json:"started"`
		LogsWebsocket string `json:"logs_websocket"`
	}
	if err := json.Unmarshal(body, &coreResponse); err != nil {
		log.Println("Error unmarshaling core status response:", err)
		return
	}

	var startedValue float64
	if coreResponse.Started {
		startedValue = 1.0
	} else {
		startedValue = 0.0
	}

	metrics.CoreStarted.WithLabelValues(coreResponse.Version).Set(startedValue)
}

func FetchUsersStats(token string) {
	pageSize := 250
	offset := 0
	totalUsersFetched := 0

	location, err := time.LoadLocation(config.CLIConfig.TimeZone)
	if err != nil {
		log.Println("Error setting timezone:", err)
		return
	}

	for {
		path := fmt.Sprintf("/api/users?limit=%d&offset=%d", pageSize, offset)
		body, err := sendRequest(path, token)
		if err != nil {
			log.Println("Error making request for user stats:", err)
			return
		}

		var usersResponse models.UsersResponse
		if err := json.Unmarshal(body, &usersResponse); err != nil {
			log.Println("Error unmarshaling response:", err)
			return
		}

		now := time.Now().In(location)
		for _, user := range usersResponse.Users {
			var onlineValue float64 = 0
			if user.OnlineAt != nil {
				onlineAt, err := time.Parse("2006-01-02T15:04:05", *user.OnlineAt)
				if err != nil {
					continue
				}
				onlineAt = onlineAt.In(location)
				if now.Sub(onlineAt) <= time.Duration(config.CLIConfig.InactivityTime)*time.Minute {
					onlineValue = 1
				}
			}

			metrics.UserOnline.WithLabelValues(user.Username).Set(onlineValue)
			metrics.UserDataLimit.WithLabelValues(user.Username).Set(user.DataLimit)
			metrics.UserUsedTraffic.WithLabelValues(user.Username).Set(user.UsedTraffic)
			metrics.UserLifetimeUsedTraffic.WithLabelValues(user.Username).Set(user.LifetimeUsedTraffic)
			metrics.UserExpirationDate.WithLabelValues(user.Username).Set(user.Expire)
		}

		fetchedUsersCount := len(usersResponse.Users)
		totalUsersFetched += fetchedUsersCount
		if fetchedUsersCount < pageSize {
			log.Printf("Fetched all users, total: %d", totalUsersFetched)
			break
		}
		offset += pageSize
	}
}

func createHttpClient() *http.Client {
	if config.CLIConfig.SocketPath != "" {
		return &http.Client{
			Transport: &http.Transport{
				DialContext: func(ctx context.Context, _, _ string) (net.Conn, error) {
					return net.Dial("unix", config.CLIConfig.SocketPath)
				},
			},
			Timeout: 30 * time.Second,
		}
	}
	return &http.Client{Timeout: 30 * time.Second}
}

func createRequest(method, path, token string) (*http.Request, error) {
	url := fmt.Sprintf("%s%s", config.CLIConfig.BaseURL, path)
	if config.CLIConfig.SocketPath != "" {
		url = "http://unix" + path
	}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	return req, nil
}

func sendRequest(path, token string) ([]byte, error) {
	httpClient := createHttpClient()
	req, err := createRequest("GET", path, token)
	if err != nil {
		return nil, err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
