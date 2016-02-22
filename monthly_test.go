package recurrence

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestMonthly(t *testing.T) {
	Convey("With a monthly recurrence on the 1st day every 3 months", t, func() {
		r := Recurrence{
			Frequence: MonthlyXth,
			Location:  time.UTC,
			Interval:  3,
			Start:     time.Date(2016, 1, 1, 12, 0, 0, 0, time.UTC),
		}
		event := time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)
		Convey("The first event should be on first january", func() {
			event := r.GetNextDate(event)
			So(event, ShouldHappenOn, time.Date(2016, 1, 1, 12, 0, 0, 0, time.UTC))
			Convey("and the second event should be on first april", func() {
				event := r.GetNextDate(event)
				So(event, ShouldHappenOn, time.Date(2016, 4, 1, 12, 0, 0, 0, time.UTC))
			})
			Convey("With the recurrence end in february there should be no second event", func() {
				r.End = time.Date(2016, 2, 1, 0, 0, 0, 0, time.UTC)
				event := r.GetNextDate(event)
				So(event, ShouldNotHappen)
			})
		})
	})
	Convey("With a monthly recurrence on the 31th of each month", t, func() {
		r := Recurrence{
			Frequence: MonthlyXth,
			Location:  time.UTC,
			Interval:  1,
			Start:     time.Date(2016, 1, 31, 12, 0, 0, 0, time.UTC),
		}
		Convey("It should happen in january", func() {
			event := r.GetNextDate(time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC))
			So(event, ShouldHappenOn, time.Date(2016, 1, 31, 12, 0, 0, 0, time.UTC))
		})
		Convey("But not on february", func() {
			event := r.GetNextDate(time.Date(2016, 2, 1, 0, 0, 0, 0, time.UTC))
			So(event, ShouldHappenOn, time.Date(2016, 3, 31, 12, 0, 0, 0, time.UTC))
		})
	})
	Convey("With a monthly recurrence on the 1st every 3 months starting in november", t, func() {
		start := time.Date(2016, 11, 1, 12, 0, 0, 0, time.UTC)
		r := Recurrence{
			Frequence: MonthlyXth,
			Location:  time.UTC,
			Interval:  3,
			Start:     start,
		}
		Convey("There should be an occurrence on every 1st february", func() {
			for year := 2017; year < 2027; year++ {
				startOfYear := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
				event := r.GetNextDate(startOfYear)
				So(event, ShouldHappenOn, time.Date(year, 2, 1, 12, 0, 0, 0, time.UTC))
			}
		})
	})
}
