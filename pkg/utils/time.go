package utils

import (
	"time"
)


func GetTimeInDateFormat() time.Time {
	// Get current time
	currentTime := time.Now()

	// Return current time as time.Date
	return time.Date(
		currentTime.Year(),
		currentTime.Month(),
		currentTime.Day(),
		currentTime.Hour(),
		currentTime.Minute(),
		currentTime.Second(),
		currentTime.Nanosecond(),
		currentTime.Location(),
	)
}


func SetTimeInDateFormat(hour int, minute int) time.Time {
	// Get current time
	currentTime := time.Now()

	// Return current time as time.Date
	return time.Date(
		currentTime.Year(),
		currentTime.Month(),
		currentTime.Day(),
		currentTime.Hour(),
		currentTime.Minute(),
		currentTime.Second(),
		currentTime.Nanosecond(),
		currentTime.Location(),
	)
}
