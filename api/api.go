package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"marzban-exporter/config"
	"marzban-exporter/metrics"
	"marzban-exporter/models"
	"net/http"
	"regexp"
	"time"
)

var httpClient = &http.Client{
	Timeout: time.Second * 10,
}

var firstWordRegexp = regexp.MustCompile(`^[a-zA-Z]+`)

func sendRequest(url, token string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func FetchNodesStatus(token string) {
	url := fmt.Sprintf("%s/api/nodes", config.CLIConfig.BaseURL)
	body, err := sendRequest(url, token)
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
		metrics.NodesStatus.WithLabelValues(node.Name, node.Address, fmt.Sprintf("%d", node.ID), fmt.Sprintf("%f", node.UsageCoef), node.XrayVer, node.Status).Set(status)
	}
}

func FetchNodesUsage(token string) {
	url := fmt.Sprintf("%s/api/nodes/usage", config.CLIConfig.BaseURL)
	body, err := sendRequest(url, token)
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
	url := fmt.Sprintf("%s/api/system", config.CLIConfig.BaseURL)
	body, err := sendRequest(url, token)
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
	url := fmt.Sprintf("%s/api/core", config.CLIConfig.BaseURL)
	body, err := sendRequest(url, token)
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
	url := fmt.Sprintf("%s/api/users", config.CLIConfig.BaseURL)
	body, err := sendRequest(url, token)
	if err != nil {
		log.Println("Error making request for user stats:", err)
		return
	}

	var usersResponse models.UsersResponse
	if err := json.Unmarshal(body, &usersResponse); err != nil {
		log.Println("Error unmarshaling response:", err)
		return
	}

	location, err := time.LoadLocation(config.CLIConfig.TimeZone)
	if err != nil {
		log.Println("Error setting timezone:", err)
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
		firstWord := firstWordRegexp.FindString(user.SubLastUserAgent)

		metrics.UserOnline.WithLabelValues(user.Note, user.Username, user.Status, firstWord).Set(onlineValue)
		metrics.UserDataLimit.WithLabelValues(user.DataLimitResetStrategy, user.Note, user.Username, user.Status, firstWord).Set(user.DataLimit)
		metrics.UserUsedTraffic.WithLabelValues(user.DataLimitResetStrategy, user.Note, user.Username, user.Status, firstWord).Set(user.UsedTraffic)
		metrics.UserLifetimeUsedTraffic.WithLabelValues(user.DataLimitResetStrategy, user.Note, user.Username, user.Status, firstWord).Set(user.LifetimeUsedTraffic)
	}
}

func GetAuthToken() (string, error) {
	url := fmt.Sprintf("%s/api/admin/token", config.CLIConfig.BaseURL)
	data := []byte(fmt.Sprintf("username=%s&password=%s", config.CLIConfig.ApiUsername, config.CLIConfig.ApiPassword))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var tokenResponse models.AuthTokenResponse
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		return "", err
	}

	return tokenResponse.AccessToken, nil
}
