package helper

// Helper functions for HTTP requests

import (
	"net/http"
	"strconv"
)

// GetFileSizeFromHeader extracts file size from http headers
func GetFileSizeFromHeader(url string) (int64, error) {
	// TODO: Implement logic to get "Content-Length" from header and covert to int64
	resp, err := http.Head(url)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()
	// Getting the file size
	contentLength, err := strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 64)
	if err != nil {
		return 0, err
	}

	return contentLength, nil
}
