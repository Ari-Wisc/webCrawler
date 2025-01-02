package main

import (
    "fmt"
    "sort"
)

// pageResult represents a single page and its link count
type pageResult struct {
    url   string
    count int
}

// sortPages converts the pages map into a sorted slice of pageResult
func sortPages(pages map[string]int) []pageResult {
    // Create slice of pageResult structs
    results := make([]pageResult, 0, len(pages))
    for url, count := range pages {
        results = append(results, pageResult{url: url, count: count})
    }

    // Sort by count (descending) and URL (ascending)
    sort.Slice(results, func(i, j int) bool {
        // If counts are different, sort by count descending
        if results[i].count != results[j].count {
            return results[i].count > results[j].count
        }
        // If counts are equal, sort by URL ascending
        return results[i].url < results[j].url
    })

    return results
}

func printReport(pages map[string]int, baseURL string) {
    // Print header
    fmt.Printf("=============================\n")
    fmt.Printf("  REPORT for %s\n", baseURL)
    fmt.Printf("=============================\n\n")

    // Sort the pages
    sortedPages := sortPages(pages)

    // Print each result
    for _, page := range sortedPages {
        fmt.Printf("Found %d internal links to %s\n", page.count, page.url)
    }
}