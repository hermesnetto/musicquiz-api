package date

import "time"

const(
	apiDateLayout = "2006-01-02T15:04:05Z"
)

// GetNow gets the current time in UTC
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowFormatted gets the current time formatted
func GetNowFormatted() string {
	return GetNow().Format(apiDateLayout)
}
