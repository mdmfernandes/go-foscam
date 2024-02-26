package foscam

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
)

// Model is the camera model
type Model uint8

// Allowed models
const (
	FI9800P Model = iota
	FI8919W
)

// Camera is the common interface implemented by all camera models.
type Camera interface {
	ChangeMotionStatus(enable bool) error
	// GetMotionStatus() error
}

// HTTPClient interface.
type HTTPClient interface {
	Get(url string) (*http.Response, error)
}

// Config is the camera configuration.
type Config struct {
	URL      string `url:"-"`
	User     string `url:"usr"`
	Password string `url:"pwd"`
}

// String returns a string representation of the camera name.
func (m Model) String() string {
	return []string{"FI9800P", "FI8919W"}[m]
}

// New is a camera interface factory
// We can create a camera by providing its model
// HTTPCLient is the client used to make requests to the cameras. Default is `http.Client`
func New(m Model, cfg Config, client ...HTTPClient) (cam Camera, err error) {
	var c HTTPClient

	if len(client) == 0 {
		// Skip SSL verification
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true}
		c = &http.Client{}
	} else {
		c = client[0]
	}

	// Check if camera is accessible
	res, err := c.Get(cfg.URL)
	if err == nil && res.StatusCode != http.StatusOK {
		err = errors.New("expected HTTP status = 200 OK")
	}
	if err != nil {
		return
	}

	// TODO: Make it generic
	switch m.String() {
	case "FI9800P":
		cam = &fi9800p{
			Client: c,
			Config: cfg,
		}
	case "FI8919W":
		cam = &fi8919w{
			Client: c,
			Config: cfg,
		}
	default:
		// We already do this check at the function beginning, but just to be sure
		err = fmt.Errorf("invalid camera model: %s", m.String())
	}
	return
}
