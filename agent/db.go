package main

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
)

var db *sql.DB

func InitDB() {
    connStr := "host=postgres user=monitor password=monitor dbname=monitor sslmode=disable"
    var err error
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
}

func SavePing(result PingResult) {
    _, err := db.Exec(`INSERT INTO ping_results (host, packet_loss, avg_latency, created_at) VALUES ($1, $2, $3, NOW())`,
        result.Host, result.PacketLoss, result.AvgLatency)
    if err != nil {
        log.Println("SavePing:", err)
    }
}

func SaveHTTP(result HTTPResult) {
    _, err := db.Exec(`INSERT INTO http_results (url, status_code, duration_ms, created_at) VALUES ($1, $2, $3, NOW())`,
        result.URL, result.StatusCode, result.DurationMs)
    if err != nil {
        log.Println("SaveHTTP:", err)
    }
}
