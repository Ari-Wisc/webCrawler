package main

import (
    "fmt"
    "net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
    // Check page limit with mutex lock
    cfg.mu.Lock()
    if len(cfg.pages) >= cfg.maxPages {
        cfg.mu.Unlock()
        return
    }
    cfg.mu.Unlock()

    cfg.concurrencyControl <- struct{}{}
    defer func() {
        <-cfg.concurrencyControl
        cfg.wg.Done()
    }()

    currentURL, err := url.Parse(rawCurrentURL)
    if err != nil {
        fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
        return
    }

    // Skip other websites
    if currentURL.Hostname() != cfg.baseURL.Hostname() {
        return
    }

    normalizedURL, err := normalizeURL(rawCurrentURL)
    if err != nil {
        fmt.Printf("Error - normalizeURL: %v\n", err)
        return
    }

    isFirst := cfg.addPageVisit(normalizedURL)
    if !isFirst {
        return
    }

    // Check again after adding the page
    cfg.mu.Lock()
    if len(cfg.pages) > cfg.maxPages {
        cfg.mu.Unlock()
        return
    }
    cfg.mu.Unlock()

    fmt.Printf("crawling %s\n", rawCurrentURL)

    htmlBody, err := getHTML(rawCurrentURL)
    if err != nil {
        fmt.Printf("Error - getHTML: %v\n", err)
        return
    }

    nextURLs, err := getURLsFromHTML(htmlBody, cfg.baseURL)
    if err != nil {
        fmt.Printf("Error - getURLsFromHTML: %v\n", err)
        return
    }

    // Only spawn new goroutines if we haven't reached max pages
    cfg.mu.Lock()
    canSpawnMore := len(cfg.pages) < cfg.maxPages
    cfg.mu.Unlock()

    if canSpawnMore {
        for _, nextURL := range nextURLs {
            cfg.wg.Add(1)
            go cfg.crawlPage(nextURL)
        }
    }
}