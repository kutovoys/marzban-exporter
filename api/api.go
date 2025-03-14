package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"marzban-exporter/config"
	"marzban-exporter/metrics"
	"marzban-exporter/models"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

var cookieCache struct {
	Cookie    http.Cookie
	ExpiresAt time.Time
	sync.Mutex
}

func createHttpClient() *http.Client {
	return &http.Client{Timeout: 30 * time.Second}
}

// API logic partially was taken from the client3xui module
// https://github.com/digilolnet/client3xui

func GetAuthToken() (*http.Cookie, error) {
	cookieCache.Lock()
	defer cookieCache.Unlock()

	if cookieCache.Cookie.Name != "" && time.Now().Before(cookieCache.ExpiresAt) {
		return &cookieCache.Cookie, nil
	}

	path := config.CLIConfig.BaseURL + "/login"
	data := url.Values{
		"username": {config.CLIConfig.ApiUsername},
		"password": {config.CLIConfig.ApiPassword},
	}

	req, err := http.NewRequest("POST", path, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := createHttpClient()
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var loginResp struct {
		Success bool   `json:"success"`
		Msg     string `json:"msg"`
	}
	if err := json.Unmarshal(body, &loginResp); err != nil {
		return nil, err
	}

	if !loginResp.Success {
		return nil, errors.New(loginResp.Msg)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("authentication failed")
	}

	for _, cookie := range resp.Cookies() {
		if cookie.Name == "3x-ui" {
			cookieCache.Cookie = *cookie
			cookieCache.ExpiresAt = cookie.Expires.Add(-6 * time.Hour)
		}
	}

	if cookieCache.Cookie.Name == "" {
		return nil, errors.New("no cookies found in auth response")
	}

	return &cookieCache.Cookie, nil
}

func FetchOnlineUsersCount(cookie *http.Cookie) {
	body, err := sendRequest("/panel/inbound/onlines", "POST", cookie)
	if err != nil {
		log.Println("Error making request for system stats:", err)
		return
	}

	var response models.ObjectResponse
	if err := json.Unmarshal(body, &response); err != nil {
		log.Println("Error unmarshaling response:", err)
		return
	}

	var arr []json.RawMessage
	err = json.Unmarshal(response.Obj, &arr)
	if err != nil {
		log.Println("Error converting Obj as array:", err)
		return
	}

	metrics.OnlineUsersCount.Set(float64(len(arr)))
}

func FetchServerStatus(cookie *http.Cookie) {
	body, err := sendRequest("/server/status", "POST", cookie)
	if err != nil {
		log.Println("Error making request for system stats:", err)
		return
	}

	var response models.ServerStatusResponse
	if err := json.Unmarshal(body, &response); err != nil {
		log.Println("Error unmarshaling response:", err)
		return
	}

	metrics.XrayVersion.WithLabelValues(response.Obj.Xray.Version).Set(1)
}

func createRequest(method, path string, cookie *http.Cookie) (*http.Request, error) {
	url := fmt.Sprintf("%s%s", config.CLIConfig.BaseURL, path)

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.AddCookie(cookie)
	return req, nil
}

func sendRequest(path, method string, cookie *http.Cookie) ([]byte, error) {
	req, err := createRequest(method, path, cookie)
	if err != nil {
		return nil, err
	}

	resp, err := createHttpClient().Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
