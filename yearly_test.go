package recurrence

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestYearlyPattern(t *testing.T) {
	Convey("With a yearly pattern on 29th february", t, func() {
		r := Recurrence{
			Frequence: Yearly,
			Interval:  1,
			Start:     time.Date(2016, 2, 29, 12, 0, 0, 0, time.UTC),
		}
		Convey("There should be an event 2016", func() {
			event := r.GetNextDate(time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC))
			So(event, ShouldHappenOn, time.Date(2016, 2, 29, 12, 0, 0, 0, time.UTC))
		})
		Convey("The next event should be 2020", func() {
			event := r.GetNextDate(time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC))
			So(event, ShouldHappenOn, time.Date(2020, 2, 29, 12, 0, 0, 0, time.UTC))
		})
		Convey("If there is an enddate 2018", func() {
			r.End = time.Date(2018, 12, 31, 0, 0, 0, 0, time.UTC)
			Convey("There should be only one event", func() {
				event := r.GetNextDate(time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC))
				So(event, ShouldHappenOn, time.Date(2016, 2, 29, 12, 0, 0, 0, time.UTC))
				event = r.GetNextDate(event)
				So(event, ShouldNotHappen)
			})
		})
	})
	Convey("With an event on 1st march every 3 years", t, func() {
		r := Recurrence{
			Frequence: Yearly,
			Interval:  3,
			Start:     time.Date(2016, 3, 1, 12, 0, 0, 0, time.UTC),
		}
		Convey("The first few events should Happen on", func() {
			for year := 2016; year < 2040; year += 3 {
				Convey(fmt.Sprintf("%d", year), func() {
					event := r.GetNextDate(time.Date(year-3, 3, 2, 0, 0, 0, 0, time.UTC))
					So(event, ShouldHappenOn, time.Date(year, 3, 1, 12, 0, 0, 0, time.UTC))
				})
			}
		})
		Convey("If there is an end 2020", func() {
			r.End = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
			Convey("There Should be only 2 events", func() {
				first := r.GetNextDate(time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC))
				So(first, ShouldHappenOn, time.Date(2016, 3, 1, 12, 0, 0, 0, time.UTC))
				second := r.GetNextDate(first)
				So(second, ShouldHappenOn, time.Date(2019, 3, 1, 12, 0, 0, 0, time.UTC))
				event := r.GetNextDate(second)
				So(event, ShouldNotHappen)
			})
		})
	})
}
