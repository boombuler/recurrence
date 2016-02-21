package recurrence

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestWeeklyPatternToInt(t *testing.T) {
	Convey("With sunday as first day of week", t, func() {
		startOfWeek := time.Sunday
		Convey("sunday should encode to 1", func() {
			So(WeeklyPatternToInt(startOfWeek, time.Sunday), ShouldEqual, 1)
		})
		Convey("monday should encode to 2", func() {
			So(WeeklyPatternToInt(startOfWeek, time.Monday), ShouldEqual, 2)
		})
		Convey("sunday and mondy should encode to 3", func() {
			So(WeeklyPatternToInt(startOfWeek, time.Sunday, time.Monday), ShouldEqual, 3)
		})
		Convey("passing the same day twice should not change the result", func() {
			So(WeeklyPatternToInt(startOfWeek, time.Sunday, time.Sunday), ShouldEqual, 1)
		})
		Convey("passing all days should encode to 127", func() {
			So(WeeklyPatternToInt(time.Sunday, time.Sunday, time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday), ShouldEqual, 127)
		})
	})
	Convey("With monday as first day of week", t, func() {
		Convey("sunday should encode to 257", func() {
			So(WeeklyPatternToInt(time.Monday, time.Sunday), ShouldEqual, 257)
		})
	})
	Convey("With tuesday as first day of week", t, func() {
		Convey("sunday should encode to 513", func() {
			So(WeeklyPatternToInt(time.Tuesday, time.Sunday), ShouldEqual, 513)
		})
	})
}

func TestWeeklyPattern(t *testing.T) {
	Convey("In berlin", t, func() {
		local, err := time.LoadLocation("Europe/Berlin")
		if err != nil {
			t.Errorf("Failed to load local: %s", err)
		}

		Convey("With a weekly recurrence that happens every two weeks on monday and saturday", func() {
			r := Recurrence{
				Type:      Weekly,
				Location:  local,
				Frequence: 2, // Every 2 weeks
				Pattern:   WeeklyPatternToInt(time.Monday, time.Monday, time.Saturday),
				Start:     time.Date(2016, 1, 1, 12, 0, 0, 0, local),
			}
			Convey("and an end in early 2017", func() {
				r.End = time.Date(2017, 1, 12, 23, 59, 0, 0, time.UTC)
				Convey("there should be no event after the end", func() {
					nextEvent := r.GetNextDate(r.End)
					So(nextEvent, ShouldNotHappen)
				})
				Convey("there should be no after 10th january 2017", func() {
					nextEvent := r.GetNextDate(time.Date(2017, 1, 10, 0, 0, 0, 0, time.UTC))
					So(nextEvent, ShouldNotHappen)
				})
				Convey("the first event should be on 2nd january", func() {
					nextEvent := r.GetNextDate(time.Date(2015, 12, 1, 0, 0, 0, 0, time.UTC))
					So(nextEvent, ShouldHappenOn, time.Date(2016, 1, 2, 12, 0, 0, 0, local))
				})
				Convey("the second event 11th january", func() {
					nextEvent := r.GetNextDate(time.Date(2016, 1, 2, 13, 0, 0, 0, time.UTC))
					So(nextEvent, ShouldHappenOn, time.Date(2016, 1, 11, 12, 0, 0, 0, local))
				})
				Convey("there should be another event on 12th feburary", func() {
					nextEvent := r.GetNextDate(time.Date(2016, 2, 7, 0, 0, 0, 0, time.UTC))
					So(nextEvent, ShouldHappenOn, time.Date(2016, 2, 8, 12, 0, 0, 0, local))
				})
			})

			Convey("without an enddate", func() {
				r.End = time.Time{}
				Convey("there should be an event in 2017", func() {
					nextEvent := r.GetNextDate(time.Date(2017, 1, 1, 1, 0, 0, 0, time.UTC))
					So(nextEvent, ShouldHappenOn, time.Date(2017, 1, 9, 12, 0, 0, 0, local))
				})
			})
		})
	})
}
