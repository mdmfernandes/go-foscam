package foscam

import (
	"net/http"
)

// Model is the camera model
type Model uint8

// Allowed models
const (
	FI9800P Model = iota
	FI8919W
)

// modelNames maps each camera model constant to its string representation.
var modelNames = map[Model]string{
	FI9800P: "FI9800P",
	FI8919W: "FI8919W",
}

// String returns the string representation of the camera model.
func (m Model) String() string {
	if name, ok := modelNames[m]; ok {
		return name
	}
	// Fallback for any undefined models
	return "invalid model"
}

// Camera is the common interface implemented by all camera models.
type Camera interface {
	// Change camera motion status to the provided value.
	// true: enable, false: disable.
	ChangeMotionStatus(enable bool) error
	// Snap a picture from the camera.
	SnapPicture() ([]byte, error)
	// GetMotionStatus() error
}

// HTTPClient interface.
type HTTPClient interface {
	Get(url string) (*http.Response, error)
}

// Config is the camera configuration.
type Config struct {
	URL      string
	User     string
	Password string
}

// NewCamera is a camera interface factory.
// Creates a camera by providing its model and configuration.
// HTTPCLient is the client used to make requests to the cameras. Default is `http.Client`.
func NewCamera(m Model, cfg Config, client ...HTTPClient) (cam Camera, err error) {
	var c HTTPClient

	if len(client) == 0 {
		c = &http.Client{}
	} else {
		c = client[0]
	}

	// Initialize the camera
	switch m.String() {
	case "FI9800P":
		cam = &fi9800p{
			Client:   c,
			URL:      cfg.URL,
			User:     cfg.User,
			Password: cfg.Password,
		}
	case "FI8919W":
		cam = &fi8919w{
			Client:   c,
			URL:      cfg.URL,
			User:     cfg.User,
			Password: cfg.Password,
		}
	default:
		err = ErrCameraInvalidModel
	}
	return
}
