package foscam

import (
	"io"
	"net/http"
)

// Send a GET request to URL via the provided client.
func getRequest(client HTTPClient, url string) ([]byte, error) {
	res, err := client.Get(url)
	if err != nil {
		return nil, &CameraError{err.Error()}
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, &BadStatusError{URL: url, Status: res.StatusCode, Expected: http.StatusOK}
	}

	b, _ := io.ReadAll(res.Body)

	return b, nil
}

// Get a snapshot from the URL, via the provided client.
func getSnap(client HTTPClient, url string) ([]byte, error) {
	b, err := getRequest(client, url)
	if err != nil {
		return nil, err
	}

	// Check that camera returns a JPEG image
	if err = isJpeg(b); err != nil {
		return nil, err
	}

	return b, nil
}
