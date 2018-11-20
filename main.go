package main

import (
        "flag"
        "os"
        "log"
        "time"
        "net/http"
)

var (
        flag_url        = flag.String("url", "", "Url to test")
        flag_timeout    = flag.Int("timeout", 5, "Timeout seconds")
)

func main() {
        flag.Parse()

        if *flag_url == "" {
                flag.Usage()
                os.Exit(1)
        }
        log.Printf("Connecting to %s...", *flag_url)

        httpClient := http.Client { Timeout:  time.Duration(*flag_timeout) * time.Second }

        resp, err:= httpClient.Get(*flag_url)
        if err != nil {
                log.Fatalf("Error occurred during connection %s ",err)
        }
        defer resp.Body.Close()
        if resp.StatusCode != 200 {
                log.Fatalf("HTTP response %d ",resp.StatusCode)
        }
        log.Printf("connection succeded")
}

