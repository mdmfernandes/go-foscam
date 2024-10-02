package foscam

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/google/go-querystring/query"
)

// fi8919w is a camera Foscam FI8919W
// We don't need to export this struct since we are using an interface fictory
// see the file foscam.go for more details
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

	res, err := c.Client.Get(url)
	if err != nil {
		return &CameraError{err.Error()}
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return &BadStatusError{URL: c.URL, Status: res.StatusCode, Expected: http.StatusOK}
	}

	b, _ := io.ReadAll(res.Body)
	got := string(b)
	want := "ok.\n"
	if got != want {
		return &BadResponseError{want, got}
	}

	return nil
}

func (c *fi8919w) SnapPicture() ([]byte, error) {
	q, _ := query.Values(c)
	url := fmt.Sprintf("%s/snapshot.cgi?%s", c.URL, q.Encode())

	res, err := c.Client.Get(url)
	if err != nil {
		return nil, &CameraError{err.Error()}
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, &BadStatusError{URL: c.URL, Status: res.StatusCode, Expected: http.StatusOK}
	}

	b, _ := io.ReadAll(res.Body)

	// Check that camera returns a JPEG image
	if mime := http.DetectContentType(b); mime != jpegMime {
		// If camera returns plain text, show it in the error message
		if strings.Contains(mime, "text/plain") {
			mime = string(b)
		}
		return nil, &BadResponseError{Want: jpegMime, Got: mime}
	}

	return b, nil
}
