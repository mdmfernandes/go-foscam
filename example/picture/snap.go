// Snap a picture from the camera.
package main

import (
	"crypto/tls"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/mdmfernandes/go-foscam"
)

func main() {
	// Generate config
	c := &foscam.Config{
		URL:      os.Getenv("FOSCAM_URL"),
		User:     os.Getenv("FOSCAM_USER"),
		Password: os.Getenv("FOSCAM_PASSWORD"),
	}

	// WARN: don't disable TLS verification. This is just for testing purposes!
	customTransport := http.DefaultTransport.(*http.Transport).Clone()
	customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true} // #nosec G402
	client := &http.Client{Transport: customTransport}

	// Initialize a camera
	cam, err := foscam.NewCamera(foscam.FI9800P, *c, client)
	if err != nil {
		panic(err)
	}

	// Operate camera
	snap, err := cam.SnapPicture()
	if err != nil {
		panic(err)
	}

	// Write image to file
	fp := "/tmp/image.jpeg"
	if err = os.WriteFile(fp, snap, 0o600); err != nil {
		panic(err)
	}

	msg := fmt.Sprintf("Exported received image to %s", fp)
	slog.Info(msg)
}
