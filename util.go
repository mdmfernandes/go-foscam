package foscam

import (
	"net/http"
)

const jpegMime string = "image/jpeg"

// Check if the provided data is a JPEG image, using mimesnif.
func isJpeg(data []byte) error {
	if mime := http.DetectContentType(data); mime != jpegMime {
		return &InvalidMIMETypeError{Want: jpegMime, Got: mime}
	}
	return nil
}

// b2u converts from boolean to uint8.
func b2u(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}
