package main

import (
    "net/http"
    "time"
)

type HTTPResult struct {
    URL        string
    StatusCode int
    DurationMs int64
}

func CheckHTTP(url string) (HTTPResult, error) {
    start := time.Now()
    resp, err := http.Get(url)
    duration := time.Since(start).Milliseconds()

    if err != nil {
        return HTTPResult{URL: url, StatusCode: 0, DurationMs: duration}, err
    }
    defer resp.Body.Close()

    return HTTPResult{URL: url, StatusCode: resp.StatusCode, DurationMs: duration}, nil
}
