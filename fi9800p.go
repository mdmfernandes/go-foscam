package foscam

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
)

// fi9800p is a camera Foscam FI9800P.
// We don't need to export this struct since we are using an interface factory.
// see the file foscam.go for more details
type fi9800p struct {
	Client HTTPClient
	Config
}

type fi9800pMotion struct {
	XMLName          xml.Name `xml:"CGI_Result" url:"-"`
	IsEnable         uint8    `xml:"isEnable" url:"isEnable"`
	Linkage          uint8    `xml:"linkage" url:"linkage"`
	SnapInterval     uint8    `xml:"snapInterval" url:"snapInterval"`
	Sensitivity      uint8    `xml:"sensitivity" url:"sensitivity"`
	TriggerInterval  uint8    `xml:"triggerInterval" url:"triggerInterval"`
	IsMovAlarmEnable uint8    `xml:"isMovAlarmEnable" url:"isMovAlarmEnable"`
	IsPirAlarmEnable uint8    `xml:"isPirAlarmEnable" url:"isPirAlarmEnable"`
	Schedule0        uint64   `xml:"schedule0" url:"schedule0"`
	Schedule1        uint64   `xml:"schedule1" url:"schedule1"`
	Schedule2        uint64   `xml:"schedule2" url:"schedule2"`
	Schedule3        uint64   `xml:"schedule3" url:"schedule3"`
	Schedule4        uint64   `xml:"schedule4" url:"schedule4"`
	Schedule5        uint64   `xml:"schedule5" url:"schedule5"`
	Schedule6        uint64   `xml:"schedule6" url:"schedule6"`
	Area0            uint16   `xml:"area0" url:"area0"`
	Area1            uint16   `xml:"area1" url:"area1"`
	Area2            uint16   `xml:"area2" url:"area2"`
	Area3            uint16   `xml:"area3" url:"area3"`
	Area4            uint16   `xml:"area4" url:"area4"`
	Area5            uint16   `xml:"area5" url:"area5"`
	Area6            uint16   `xml:"area6" url:"area6"`
	Area7            uint16   `xml:"area7" url:"area7"`
	Area8            uint16   `xml:"area8" url:"area8"`
	Area9            uint16   `xml:"area9" url:"area9"`
}

type fi9800pResponse struct {
	XMLName xml.Name `xml:"CGI_Result"`
	Result  int      `xml:"result"`
}

// updateMotionDetect updates the motion detection configuration.
func (c *fi9800p) updateMotionDetect(mc fi9800pMotion) error {
	// Construct the URL
	qc, _ := query.Values(c)  // Credentials
	qm, _ := query.Values(mc) // Motion Config
	url := fmt.Sprintf("%s/cgi-bin/CGIProxy.fcgi?cmd=setMotionDetectConfig&%s&%s", c.URL, qm.Encode(), qc.Encode())

	res, err := c.Client.Get(url)
	if err != nil {
		return &CameraError{err.Error()}
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return &BadStatusError{URL: c.URL, Status: res.StatusCode, Expected: http.StatusOK}
	}

	b, _ := io.ReadAll(res.Body)

	var mr fi9800pResponse

	if err := xml.Unmarshal(b, &mr); err != nil {
		return err
	}

	if mr.Result != 0 {
		return &BadResponseError{Want: 0, Got: mr.Result}
	}

	return nil
}

// GetMotionDetect retrieves the motion detection configuration.
func (c *fi9800p) GetMotionDetect() (fi9800pMotion, error) {
	var mc fi9800pMotion

	// Construct the URL
	q, _ := query.Values(c)
	url := fmt.Sprintf("%s/cgi-bin/CGIProxy.fcgi?cmd=getMotionDetectConfig&%s", c.URL, q.Encode())

	res, err := c.Client.Get(url)
	if err != nil {
		return mc, &CameraError{err.Error()}
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return mc, &BadStatusError{URL: c.URL, Status: res.StatusCode, Expected: http.StatusOK}
	}

	b, _ := io.ReadAll(res.Body)
	if err = xml.Unmarshal(b, &mc); err != nil {
		return mc, err
	}

	return mc, nil
}

// ChangeMotionStatus enables/disables the camera motion detection.
func (c *fi9800p) ChangeMotionStatus(enable bool) error {
	mc, err := c.GetMotionDetect()
	if err != nil {
		return err
	}

	// If the camera status is already the desired, don't send the request
	if mc.IsEnable == b2u(enable) {
		return nil
	}
	mc.IsEnable = b2u(enable)

	return c.updateMotionDetect(mc)
}
