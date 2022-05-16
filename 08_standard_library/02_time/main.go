package main

import "time"

func SpanYear(t time.Time) (time.Time, time.Time)        {}
func SpanMonth(t time.Time) (time.Time, time.Time)       {}
func SpanWeek(t time.Time) (time.Time, time.Time)        {}
func SpanWorkingWeek(t time.Time) (time.Time, time.Time) {}
func SpanDay(t time.Time) (time.Time, time.Time)         {}
func SpanHour(t time.Time) (time.Time, time.Time)        {}

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
}
