package main

import (
	"fmt"
	"time"
)

func dayOfTheWeek(day int, month int, year int) string {
	// 1) find out what 1 Jan 1971 was
	// Friday - https://takemeback.to/01-January-1971
	// 2) find out leap years between 1971 - 2100
	// hmm was on wrong track. Can just use built-in package time
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local).Weekday().String()
}

// initially made this as ancillary function but then I don't actually need it. But good practice with implementing more functions
func isLeapYear(year int) bool {
	// In the Gregorian calendar, three criteria must be taken into account to identify leap years:
	// a better way than storing an array of hardcoded known leap years between the date range
	// use math: https://www.timeanddate.com/date/leapyear.html - basically 3 criteria
	isLeapYear := false
	criteria1 := year%4 == 0   // general rule if true then == leap year
	criteria2 := year%100 == 0 // general rule every century is not a leap year
	criteria3 := year%400 == 0 // general rule if divisible by 400 is a leap year

	if criteria1 {
		isLeapYear = true
	}

	if criteria2 {
		isLeapYear = false
	}

	if criteria3 {
		isLeapYear = true
	}

	return isLeapYear
}

func main() {
	day := 31
	month := 8
	year := 2019
	fmt.Println(dayOfTheWeek(day, month, year))

}

/* -------- NOTES -----------

Was initially confused about how to represent the logic here in a boolean test
https://www.mathsisfun.com/leap-years.html

But when I put it this way it makes sense:

	if criteria1 {
		isLeapYear == true
	}

	if criteria2 {
		isLeapYear == false
	}

	if criteria3 {
		isLeapYear == true
	}

But then I soon realised these are the same (i.e. the comparison operations not the return statement)?

	if criteria1 && !criteria2 && criteria3 {
		return true
	} else {
		return false
	}
*/
