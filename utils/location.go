package utils

import "math"

// Hsin ...
func Hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

// IsInvalidLatLng ...
func IsInvalidLatLng(lat, lng float64) bool {
	isEmpty := lat == 0 && lng == 0
	isInvalidLat := lat < -90 && lat > 90
	isInvalidLng := lng < -180 && lng > 180
	return isEmpty || isInvalidLat || isInvalidLng
}
