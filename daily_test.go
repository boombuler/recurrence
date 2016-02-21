package recurrence

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestDailyEvery5Days(t *testing.T) {
	Convey("With a daily event every 5 days", t, func() {
		r := Recurrence{
			Type:      Daily,
			Frequence: 5,
			Location:  time.UTC,
			Start:     time.Date(2016, 1, 1, 12, 0, 0, 0, time.UTC),
		}

		Convey("which ends 2017", func() {
			r.End = time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC)
			Convey("there should be no event 2017", func() {
				nextEvent := r.GetNextDate(time.Date(2017, 1, 1, 1, 0, 0, 0, time.UTC))
				So(nextEvent, ShouldNotHappen)
			})
			Convey("the first event should be on 1st january", func() {
				nextEvent := r.GetNextDate(time.Date(2015, 12, 12, 0, 0, 0, 0, time.UTC))
				So(nextEvent, ShouldHappenOn, time.Date(2016, 1, 1, 12, 0, 0, 0, time.UTC))
			})
			Convey("the second event should be on 6th january", func() {
				nextEvent := r.GetNextDate(time.Date(2016, 1, 1, 12, 0, 0, 0, time.UTC))
				So(nextEvent, ShouldHappenOn, time.Date(2016, 1, 6, 12, 0, 0, 0, time.UTC))
			})
			Convey("the last event should be on 31th december", func() {
				lastEvent := r.GetNextDate(time.Date(2016, 12, 27, 0, 0, 0, 0, time.UTC))
				So(lastEvent, ShouldHappenOn, time.Date(2016, 12, 31, 12, 0, 0, 0, time.UTC))

				nextEvent := r.GetNextDate(lastEvent)
				So(nextEvent, ShouldNotHappen)
			})
		})
		Convey("which doesn't end", func() {
			r.End = time.Time{}
			Convey("there should be an event on 5th january 2017", func() {
				nextEvent := r.GetNextDate(time.Date(2017, 1, 1, 1, 0, 0, 0, time.UTC))
				So(nextEvent, ShouldHappenOn, time.Date(2017, 1, 5, 12, 0, 0, 0, time.UTC))
			})
		})
	})
}
