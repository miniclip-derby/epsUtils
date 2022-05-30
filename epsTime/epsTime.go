package epsTime

import (
	"fmt"
	"strconv"
	"time"
)

// func encodeDateTime(dt time.Time) string {
// 	if dt.IsZero() {
// 		return fmt.Sprintf("Empty")
// 	}
// 	return fmt.Sprintf("%#v", dt.UTC().UnixNano())
// }

// func decodeDate(date string) (time.Time, error) {
// 	return time.Parse("2006-01-02", date)
// }

func ToString(d time.Time) string {
	return d.Format("Monday January 2 2006")
}

func SecondsToMinSec(seconds int) (int, int) {
	var m, s int

	if seconds > 60 {
		m = seconds / (60)
		s = seconds - (m * 60)
	}

	return m, s
}

func SecondsToMinSecString(seconds int) string {
	m, s := SecondsToMinSec(seconds)
	return fmt.Sprintf("%2v:%.2v", m, s)
}

func SecondsToMinSecStringLeftAlign(seconds int) string {
	m, s := SecondsToMinSec(seconds)
	return fmt.Sprintf("%v:%.2v", m, s)
}

/*
func decodeDateAndTime(date string) (time.Time, error) {
	return time.Parse("2006-01-02 03:04 PM", date)
}
*/
// func decodeDateTime(date string, timeInt int) (time.Time, error) {
// 	goDateTime, err := time.Parse("2006-01-02", date)
// 	if err != nil {
// 		return getServerTime(), err
// 	}
//
// 	goSeconds := time.Second * time.Duration(timeInt)
// 	goDateTime = goDateTime.Add(goSeconds)
//
// 	return goDateTime, nil
// }

// func decodeUnixNano(nano string) (time.Time, error) {
// 	nsec, err := strconv.ParseInt(nano, 10, 64)
// 	if err != nil {
// 		nsec = 0
// 	}
// 	return time.Unix(0, nsec), err
// }

func FromUnixSeconds(seconds string) time.Time {
	sec, err := strconv.ParseInt(seconds, 10, 64)
	if err != nil {
		sec = 0
	}
	return time.Unix(sec, 0)
}

// func dateQuantToDay(t time.Time) time.Time {
// 	year, month, day := t.Date()
// 	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
// }

// func timeDurationSeconds(seconds int) time.Duration {
// 	return time.Second * time.Duration(seconds)
// }
