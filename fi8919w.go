package foscam

import (
	"fmt"
	"io"

	"github.com/google/go-querystring/query"
)

// fi8919w is a camera Foscam FI8919W
// We don't need to export this struct since we are using an interface fictory
// see the file foscam.go for more details
type fi8919w struct {
	Client HTTPClient
	Config
}

func (c *fi8919w) ChangeMotionStatus(enable bool) error {
	q, _ := query.Values(c)
	url := fmt.Sprintf("%s/set_alarm.cgi?%s&motion_armed=%d",
		c.URL,
		q.Encode(),
		b2u(enable))

	res, err := c.Client.Get(url)
	if err != nil {
		return err
	}

	b, _ := io.ReadAll(res.Body)

	if res.StatusCode != 200 {
		err = fmt.Errorf("unexpected response status from camera: %d", res.StatusCode)
	} else if string(b) != "ok.\n" {
		err = fmt.Errorf("unexpected response text from camera: %s", string(b))
	}

	return err
}
