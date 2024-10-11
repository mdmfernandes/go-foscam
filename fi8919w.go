package foscam

import (
	"fmt"

	"github.com/google/go-querystring/query"
)

// fi8919w is a camera Foscam FI8919W
// We don't need to export this struct since we are using an interface factory.
// See the file foscam.go for more details
type fi8919w struct {
	Client HTTPClient `url:"-"`
	// The go-querystring values may change from camera to camera, so we can't
	// use Config (from foscam.go) directly here.
	URL      string `url:"-"`
	User     string `url:"user"`
	Password string `url:"pwd"`
}

// ChangeMotionStatus enables/disables the camera motion detection.
func (c *fi8919w) ChangeMotionStatus(enable bool) error {
	q, _ := query.Values(c)
	url := fmt.Sprintf("%s/set_alarm.cgi?%s&motion_armed=%d",
		c.URL,
		q.Encode(),
		b2u(enable))

	b, err := getRequest(c.Client, url)
	if err != nil {
		return err
	}

	got := string(b)
	want := "ok.\n"
	if got != want {
		return &BadResponseError{want, got}
	}

	return nil
}

// SnapPicture takes a snapshot and returns the picture in a byte slice.
func (c *fi8919w) SnapPicture() ([]byte, error) {
	q, _ := query.Values(c)
	url := fmt.Sprintf("%s/snapshot.cgi?%s", c.URL, q.Encode())

	return getSnap(c.Client, url)
}
