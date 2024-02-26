package main

import (
	"os"

	"github.com/mdmfernandes/go-foscam"
)

type config struct {
	url      string
	user     string
	password string
}

func main() {
	// Generate config
	c := &foscam.Config{
		URL:      os.Getenv("FOSCAM_URL"),
		User:     os.Getenv("FOSCAM_USER"),
		Password: os.Getenv("FOSCAM_PASSWORD"),
	}

	// Initialize a camera
	cam, err := foscam.New(foscam.FI9800P, *c)
	if err != nil {
		panic(err)
	}

	// Operate camera
	if err = cam.ChangeMotionStatus(false); err != nil {
		panic(err)
	}
}
