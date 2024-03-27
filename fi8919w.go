package foscam

import (
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
)

// fi8919w is a camera Foscam FI8919W
// We don't need to export this struct since we are using an interface fictory
// see the file foscam.go for more details
type fi8919w struct {
	Client HTTPClient
	Config
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
