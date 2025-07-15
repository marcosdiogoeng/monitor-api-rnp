package main

import (
    "os/exec"
    "regexp"
    "strconv"
)

type PingResult struct {
    Host       string
    PacketLoss float64
    AvgLatency float64
}

func RunPing(host string) (PingResult, error) {
    out, err := exec.Command("ping", "-c", "4", host).Output()
    if err != nil {
        return PingResult{}, err
    }

    output := string(out)
    lossRegex := regexp.MustCompile(`, ([\d.]+)% packet loss`)
    latencyRegex := regexp.MustCompile(`rtt .* = .*?/([\d.]+)/`)

    lossMatch := lossRegex.FindStringSubmatch(output)
    latencyMatch := latencyRegex.FindStringSubmatch(output)

    loss, _ := strconv.ParseFloat(lossMatch[1], 64)
    latency, _ := strconv.ParseFloat(latencyMatch[1], 64)

    return PingResult{Host: host, PacketLoss: loss, AvgLatency: latency}, nil
}
