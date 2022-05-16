package main

import (
	"fmt"
	"time"
)

func SpanYear(t time.Time) (time.Time, time.Time) {
	startYear := time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
	endYear := startYear.AddDate(1, 0, 0).Add(-time.Nanosecond)
	return startYear, endYear
}

func SpanMonth(t time.Time) (time.Time, time.Time) {
	startMonth := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	endMonth := startMonth.AddDate(0, 1, 0).Add(-time.Nanosecond)
	return startMonth, endMonth
}

func SpanWeek(t time.Time) (time.Time, time.Time) {
	t.Weekday()
	startWeek := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	if t.Weekday() == 0 {
		startWeek = startWeek.AddDate(0, 0, 1)
	} else {
		startWeek = startWeek.AddDate(0, 0, int(t.Weekday())-1)
	}

	endWeek := startWeek.AddDate(0, 0, 7).Add(-time.Nanosecond)
	return startWeek, endWeek
}

func SpanWorkingWeek(t time.Time) (time.Time, time.Time) {
	startWeek, endWeek := SpanWeek(t)
	return startWeek, endWeek.AddDate(0, 0, -2)
}
func SpanDay(t time.Time) (time.Time, time.Time) {
	startDay := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	endDay := startDay.Add(time.Hour*24 - time.Nanosecond)
	return startDay, endDay
}

func SpanHour(t time.Time) (time.Time, time.Time) {
	startHour := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.Location())
	endHour := startHour.Add(time.Hour - time.Nanosecond)
	return startHour, endHour
}

/*
	Task:
		Implement the empy functions.
		You receive a time.Time value and should return two time.Time values dependent on Span.

		For example your t is 2022-04-12T12:02:32 then your SpanYear should return:
			2022-01-01T00:00:00, 2022-12-31T23:59:59

		For example your t is 2022-04-12T12:02:32 then your SpanWeek should return:
			2022-04-11T00:00:00, 2022-04-17T23:59:59

*/
func main() {
	now := time.Now()
	fmt.Println(SpanYear(now))
	fmt.Println(SpanMonth(now))
	fmt.Println(SpanWeek(now))
	fmt.Println(SpanWorkingWeek(now))
	fmt.Println(SpanDay(now))
	fmt.Println(SpanHour(now))
}
