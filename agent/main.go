package main

import (
    "time"
)

func main() {
    InitDB()

    for {
        hosts := []string{"8.8.8.8", "google.com"}
        for _, host := range hosts {
            res, err := RunPing(host)
            if err == nil {
                SavePing(res)
            }
        }

        urls := []string{"https://google.com", "https://youtube.com", "https://rnp.br"}
        for _, url := range urls {
            res, err := CheckHTTP(url)
            if err == nil {
                SaveHTTP(res)
            }
        }

        time.Sleep(300 * time.Second)
    }
}
