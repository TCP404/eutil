package eutil

import (
	"slices"
	"testing"
	"time"
)

func Test_DateRanger(t *testing.T) {
	var (
		startTime = time.Date(2024, time.November, 1, 0, 0, 0, 0, time.UTC)
		endTime   = time.Date(2024, time.November, 5, 0, 0, 0, 0, time.UTC)
		step      = 24 * time.Hour
		except    = []time.Time{
			time.Date(2024, time.November, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.November, 2, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.November, 3, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.November, 4, 0, 0, 0, 0, time.UTC),
		}
		got = make([]time.Time, 0, 4)
	)

	for d := range DateRanger(startTime, endTime, step) {
		got = append(got, d)
	}

	if gotLen, exceptLen := len(got), len(except); gotLen != exceptLen {
		t.Fatalf("except length: %v, but got length: %v", exceptLen, gotLen)
	}
	if !slices.EqualFunc(got, except, func(a, b time.Time) bool { return a.Equal(b) }) {
		t.Fatalf("except: %v, but got: %v", except, got)
	}
}

func Test_TimeRanger(t *testing.T) {
	var (
		startTime = time.Date(2024, time.November, 1, 0, 0, 0, 0, time.UTC)
		endTime   = time.Date(2024, time.November, 3, 0, 0, 0, 0, time.UTC)
		step      = 12 * time.Hour
		got       = make([][2]time.Time, 0, 4)
		except    = [][2]time.Time{
			{
				time.Date(2024, time.November, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2024, time.November, 1, 12, 0, 0, 0, time.UTC),
			},
			{
				time.Date(2024, time.November, 1, 12, 0, 0, 0, time.UTC),
				time.Date(2024, time.November, 2, 0, 0, 0, 0, time.UTC),
			},
			{
				time.Date(2024, time.November, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2024, time.November, 2, 12, 0, 0, 0, time.UTC),
			},
			{
				time.Date(2024, time.November, 2, 12, 0, 0, 0, time.UTC),
				time.Date(2024, time.November, 3, 0, 0, 0, 0, time.UTC),
			},
		}
	)

	for s, e := range TimeRanger(startTime, endTime, step) {
		got = append(got, [2]time.Time{s, e})
	}
	if !slices.EqualFunc(got, except, func(a, b [2]time.Time) bool {
		return a[0].Equal(b[0]) && a[1].Equal(b[1])
	}) {
		t.Fatalf("except: %v, but got: %v", except, got)
	}
}
