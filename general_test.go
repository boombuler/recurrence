package recurrence

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestInvalid(t *testing.T) {
	Convey("With an invalid recurrence type", t, func() {
		r := Recurrence{
			Frequence: Frequence(999),
			Interval:  5,
			Location:  time.UTC,
			Start:     time.Date(2016, 1, 1, 12, 0, 0, 0, time.UTC),
		}
		Convey("the there should be no event", func() {
			nextEvent := r.GetNextDate(time.Now())
			So(nextEvent, ShouldNotHappen)
		})
	})
	Convey("A daily recurrence without an interval should act like every day", t, func() {
		r := Recurrence{
			Frequence: Daily,
			Interval:  0,
			Location:  time.UTC,
			Start:     time.Date(2016, 1, 1, 12, 0, 0, 0, time.UTC),
		}
		event := r.GetNextDate(time.Date(2016, 1, 3, 12, 0, 0, 0, time.UTC))
		So(event, ShouldHappenOn, time.Date(2016, 1, 4, 12, 0, 0, 0, time.UTC))
	})
}

func TestNonRepeating(t *testing.T) {
	Convey("Without a repeating frequence", t, func() {
		start := time.Date(2016, 1, 1, 12, 0, 0, 0, time.UTC)
		r := Recurrence{
			Start: start,
		}
		Convey("It should return the start date", func() {
			event := r.GetNextDate(time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC))
			So(event, ShouldHappenOn, start)
		})
		Convey("There should be no second event", func() {
			event := r.GetNextDate(start)
			So(event, ShouldNotHappen)
		})
	})
}
