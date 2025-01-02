package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
    if len(os.Args) < 3 {
        fmt.Println("Usage: ./crawler <URL> <maxConcurrency> <maxPages>")
        return
    }

    rawBaseURL := os.Args[1]
    maxConcurrency, err := strconv.Atoi(os.Args[2])
    if err != nil {
        fmt.Println("Invalid maxConcurrency value")
        return
    }
    maxPages, err := strconv.Atoi(os.Args[3])
    if err != nil {
        fmt.Println("Invalid maxPages value")
        return
    }

    cfg, err := configure(rawBaseURL, maxConcurrency)
    if err != nil {
        fmt.Printf("Error - configure: %v", err)
        return
    }
    cfg.maxPages = maxPages // Set maxPages

    fmt.Printf("starting crawl of: %s...\n", rawBaseURL)

    cfg.wg.Add(1)
    go cfg.crawlPage(rawBaseURL)
    cfg.wg.Wait()

    for normalizedURL, count := range cfg.pages {
        fmt.Printf("%d - %s\n", count, normalizedURL)
    }
}
