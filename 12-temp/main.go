package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("schedule", Schedule("7/25/2019 13:45:00"))
	fmt.Println("has passed", HasPassed("December 9, 2112 11:59:59"))
	fmt.Println("is afternoon appointment", IsAfternoonAppointment("Friday, September 9, 2112 11:59:59"))
	fmt.Println("description", Description("7/25/2019 13:45:00"))
	fmt.Println(AnniversaryDate())
}

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/02/2006 15:04:05"

	formatedTime, err := time.Parse(layout, date)

	if nil != err {
		fmt.Println(err)
	}

	return formatedTime
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	// expected input layout HasPassed("July 25, 2019 13:45:00")
	layout := "January 2, 2006 15:04:05"
	parsedDate, err := time.Parse(layout, date)
	if nil != err {
		fmt.Println(err)
	}
	now := time.Now()

	return parsedDate.Before(now)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	// example call IsAfternoonAppointment("Thursday, July 25, 2019 13:45:00")
	layout := "Monday, January 2, 2006 15:04:05"
	parsedDate, err := time.Parse(layout, date)

	if nil != err {
		fmt.Println(err)
	}

	appointmentHour := parsedDate.Hour()

	return appointmentHour >= 12 && appointmentHour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	// example call Description("7/25/2019 13:45:00")
	layout := "1/2/2006 15:04:05"
	t, err := time.Parse(layout, date)
	if nil != err {
		fmt.Println(err)
	}
	description := fmt.Sprintf("You have an appointment on %s, %s %v, %v, at %v:%v.", t.Weekday(), t.Month(), t.Day(), t.Year(), t.Hour(), t.Minute())
	return description
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	// assuming current year is example call AnniversaryDate() => // => 2020-09-15 00:00:00 +0000 UTC
	currentYear := time.Now().Year()
	anniversaryThisYear := time.Date(currentYear, time.September, 15, 0, 0, 0, 0, time.UTC)
	return anniversaryThisYear
}
