package recurrence

import (
	"testing"
	"time"
)

func Test_WeeklyPatternToInt(t *testing.T) {
	if r := WeeklyPatternToInt(time.Sunday, time.Sunday); r != 1 {
		t.Errorf("Sunday failed. Got %v", r)
	}
	if r := WeeklyPatternToInt(time.Sunday, time.Monday); r != 2 {
		t.Errorf("Monday failed. Got %v", r)
	}
	if r := WeeklyPatternToInt(time.Sunday, time.Sunday, time.Monday); r != 3 {
		t.Errorf("Sunday,Monday failed. Got %v", r)
	}
	if r := WeeklyPatternToInt(time.Sunday, time.Sunday, time.Sunday); r != 1 {
		t.Errorf("Sunday,Sunday failed. Got %v", r)
	}
	if r := WeeklyPatternToInt(time.Sunday, time.Sunday, time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday); r != 127 {
		t.Errorf("All-Days failed. Got %v", r)
	}

	if r := WeeklyPatternToInt(time.Monday, time.Sunday); r != 257 {
		t.Error("FirstDayOfWeek Monday encoding failed")
	}
	if r := WeeklyPatternToInt(time.Tuesday, time.Sunday); r != 513 {
		t.Error("FirstDayOfWeek Tuesday encoding failed")
	}
}

func Test_WeeklyPattern(t *testing.T) {
	local, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		t.Errorf("Failed to load local: %s", err)
	}
	r := Recurrence{
		Type:      Weekly,
		Location:  local,
		Frequence: 2, // Every 2 weeks
		Pattern:   WeeklyPatternToInt(time.Monday, time.Monday, time.Saturday),
		Start:     time.Date(2016, 1, 1, 12, 0, 0, 0, local),
		End:       time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	tests := map[time.Time]time.Time{
		time.Date(2017, 1, 3, 0, 0, 0, 0, time.UTC):   time.Time{},
		time.Date(2015, 12, 12, 0, 0, 0, 0, time.UTC): time.Date(2016, 1, 2, 12, 0, 0, 0, local),
		time.Date(2016, 1, 3, 0, 0, 0, 0, time.UTC):   time.Date(2016, 1, 11, 12, 0, 0, 0, local),
		time.Date(2016, 2, 7, 0, 0, 0, 0, time.UTC):   time.Date(2016, 2, 8, 12, 0, 0, 0, local),
	}

	for d, expected := range tests {
		result := r.GetNextDate(d)
		if result != expected {
			t.Errorf("Failed for %v. Got %v expected %v", d, result, expected)
		}
	}

	r.End = time.Time{}
	if dat := r.GetNextDate(time.Date(2017, 1, 1, 1, 0, 0, 0, time.UTC)); dat != time.Date(2017, 1, 9, 12, 0, 0, 0, local) {
		t.Error("no-end-date failed. Got", dat)
	}
}
