package foscam

// b2u converts from boolean to uint8
func b2u(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}
