package recurrence

import (
	"testing"
	"time"
)

func TestDailyEvery5Days(t *testing.T) {
	r := Recurrence{
		Type:      Daily,
		Frequence: 5,
		Location:  time.UTC,
		Start:     time.Date(2016, 1, 1, 12, 0, 0, 0, time.UTC),
		End:       time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	tests := map[time.Time]time.Time{
		time.Date(2015, 12, 12, 0, 0, 0, 0, time.UTC): time.Date(2016, 1, 1, 12, 0, 0, 0, time.UTC),
		time.Date(2016, 1, 2, 0, 0, 0, 0, time.UTC):   time.Date(2016, 1, 6, 12, 0, 0, 0, time.UTC),
		time.Date(2017, 1, 1, 1, 0, 0, 0, time.UTC):   time.Time{},
		time.Date(2016, 1, 6, 12, 0, 0, 0, time.UTC):  time.Date(2016, 1, 11, 12, 0, 0, 0, time.UTC),
		time.Date(2016, 1, 6, 11, 0, 0, 0, time.UTC):  time.Date(2016, 1, 6, 12, 0, 0, 0, time.UTC),
		time.Date(2016, 1, 30, 0, 0, 0, 0, time.UTC):  time.Date(2016, 1, 31, 12, 0, 0, 0, time.UTC),
		time.Date(2016, 3, 1, 0, 0, 0, 0, time.UTC):   time.Date(2016, 3, 1, 12, 0, 0, 0, time.UTC),
	}

	for d, expected := range tests {
		result := r.GetNextDate(d)
		if result != expected {
			t.Errorf("Failed for %v. Got %v expected %v", d, result, expected)
		}
	}

	r.End = time.Time{}
	if dat := r.GetNextDate(time.Date(2017, 1, 1, 1, 0, 0, 0, time.UTC)); dat != time.Date(2017, 1, 5, 12, 0, 0, 0, time.UTC) {
		t.Error("no-end-date failed")
	}
}
