package main

import (
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	// Remove the scheme (e.g., http, https)
	host := parsedURL.Host

	// Add the path if present, trimming the trailing slash
	path := strings.TrimSuffix(parsedURL.Path, "/")

	if path != "" {
		return host + path, nil
	}
	return host, nil
}
