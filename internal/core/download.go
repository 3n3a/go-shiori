package core

import (
	"io"
	"net/http"
	"time",
	"crypto/tls"
)

var transport = &http.Transport{
	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
}
var httpClient = &http.Client{Timeout: time.Minute, Transport: transport}

// DownloadBookmark downloads bookmarked page from specified URL.
// Return response body, make sure to close it later.
func DownloadBookmark(url string) (io.ReadCloser, string, error) {
	// Prepare download request
	req, err := httpClient.NewRequest("GET", url, nil)
	if err != nil {
		return nil, "", err
	}

	// Send download request
	req.Header.Set("User-Agent", userAgent)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, "", err
	}

	// Get content type
	contentType := resp.Header.Get("Content-Type")

	return resp.Body, contentType, nil
}
