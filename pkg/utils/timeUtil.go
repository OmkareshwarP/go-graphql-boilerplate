package utils

import (
	"time"
)

func GetCurrentTime() time.Time {
	return time.Now().UTC() // Ensure the time is in UTC
}

func GetFromDate(numberOfDays int) (fromDate, toDate, previousWindowFromDate, previousWindowToDate string) {
	// Get the current time in UTC
	currentTime := time.Now().UTC()

	// Format the toDate (current date and time)
	toDate = currentTime.Format("2006-01-02 15:04:05")

	// Calculate the fromDate by subtracting numberOfDays from current date and time
	fromDateTime := currentTime.Add(-time.Duration(numberOfDays) * 24 * time.Hour)

	// Format the fromDate as per the desired format
	fromDate = fromDateTime.Format("2006-01-02 15:04:05")

	// Calculate the toDate and fromDate for the previous window (2 * numberOfDays days ago to 7 days ago)
	previousWindowFromDateTime := fromDateTime.Add(-time.Duration(numberOfDays) * 24 * time.Hour)

	// Format the previous window dates as per the desired format
	previousWindowFromDate = previousWindowFromDateTime.Format("2006-01-02 15:04:05")
	previousWindowToDate = fromDate

	return fromDate, toDate, previousWindowFromDate, previousWindowToDate
}

func GetCurrentEpochTimestampInMS() int64 {
	return time.Now().UnixMilli()
}
