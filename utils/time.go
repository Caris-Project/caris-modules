package utils

import (
	"time"
)

//
// NOTE: due to unique timezone in server's code, all time using will be convert to HCM timezone (UTC +7)
// All function generate time, must be call util function is here
// WARNING: don't accept call time.Now() directly
//

const timezoneHCM = "Asia/Ho_Chi_Minh"

// getHCMLocation ...
func getHCMLocation() *time.Location {
	l, _ := time.LoadLocation(timezoneHCM)
	return l
}

// TimeNow ...
func TimeNow() time.Time {
	loc := getHCMLocation()
	return time.Now().In(loc)
}

// TimeWithHoursAgo ...
func TimeWithHoursAgo(hoursAgo int) time.Time {
	loc := getHCMLocation()
	t := time.Now().Add(time.Hour * -1 * time.Duration(hoursAgo)).In(loc)
	return t
}
