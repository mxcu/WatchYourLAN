package clickhouse

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	"github.com/aceberg/WatchYourLAN/internal/models"
)

// Global HTTP Client to enable Keep-Alive Connection-Reuse
var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

// Add - write data to ClickHouse
func Add(appConfig models.Conf, oneHist models.Host) {
	if appConfig.ClickhouseAddr == "" {
		return
	}

	db := appConfig.ClickhouseDB
	if db == "" {
		db = "default"
	}

	table := appConfig.ClickhouseTable
	if table == "" {
		table = "watch_your_lan"
	}

	insertSQL := fmt.Sprintf("INSERT INTO %s.%s (ts, known, mac, iface, ip, name) FORMAT JSONEachRow", db, table)
	jsonRow := map[string]any{
		"ts":    time.Now().UTC().Format("2006-01-02 15:04:05"),
		"known": oneHist.Known,
		"mac":   oneHist.Mac,
		"iface": oneHist.Iface,
		"ip":    oneHist.IP,
		"name":  oneHist.Name,
	}
	jsonData, err := json.Marshal(jsonRow)
	if err != nil {
		slog.Error("ClickHouse JSON error", "err", err)
		return
	}

	reqURL := fmt.Sprintf("%s/?query=%s", appConfig.ClickhouseAddr, url.QueryEscape(insertSQL))
	req, err := http.NewRequest(http.MethodPost, reqURL, bytes.NewBuffer(jsonData))
	if err != nil {
		slog.Error("ClickHouse Request error", "err", err)
		return
	}

	req.Header.Set("X-ClickHouse-User", appConfig.ClickhouseUser)
	req.Header.Set("X-ClickHouse-Key", appConfig.ClickhousePassword)

	res, err := httpClient.Do(req)
	if err != nil {
		slog.Error("ClickHouse network error", "err", err)
		return
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	if res.StatusCode != http.StatusOK {
		slog.Error("ClickHouse HTTP error", "status", res.StatusCode, "body", string(body))
	}
}
