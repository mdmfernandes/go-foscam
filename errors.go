package foscam

import (
	"errors"
	"fmt"
)

// ErrCameraInvalidModel is the error thrown for an invalid camera model.
var ErrCameraInvalidModel = errors.New("invalid model")

// CameraError represents a camera generic error.
type CameraError struct {
	Msg string
}

func (c *CameraError) Error() string {
	return fmt.Sprintf("Camera error: %s", c.Msg)
}

// BadStatusError represents an error in the status code returned by the camera.
type BadStatusError struct {
	URL      string
	Status   int
	Expected int
}

func (b *BadStatusError) Error() string {
	return fmt.Sprintf("Bad status code %d from '%s'. Expected: %d", b.Status, b.URL, b.Expected)
}

// BadResponseError represents a bad response from the camera.
type BadResponseError struct {
	Want any
	Got  any
}

func (b *BadResponseError) Error() string {
	return fmt.Sprintf("Unexpected response. Want: %#v, Got: %#v", b.Want, b.Got)
}
