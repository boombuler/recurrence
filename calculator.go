package recurrence

import (
	"time"
)

const (
	day = 24 * time.Hour
)

func (p Recurrence) GetNextDate(d time.Time) time.Time {
	if p.Frequence == 0 {
		p.Frequence = 1
	}
	switch p.Type {
	case Daily:
		return p.ndDaily(d)
	case Weekly:
		return p.ndWeekly(d)
	}
	return time.Time{}
}

func (p Recurrence) dateOf(t time.Time) time.Time {
	y, m, d := t.In(p.Location).Date()
	return time.Date(y, m, d, 0, 0, 0, 0, p.Location)
}

func (p Recurrence) ndDaily(d time.Time) time.Time {
	start := p.Start.In(p.Location)
	end := p.End.In(p.Location)
	if end.After(start) && d.After(end) {
		return time.Time{}
	}
	if d.Before(start) {
		return start
	}

	startDate := p.dateOf(start)
	timeOfDay := start.Sub(startDate)
	d = d.In(p.Location)

	dateOfD := p.dateOf(d)

	daysBetween := int(dateOfD.Sub(startDate) / day)
	freq := int(p.Frequence)
	daysToAdd := (freq - (daysBetween % freq)) % freq

	res := dateOfD.Add(time.Duration(daysToAdd)*day + timeOfDay)

	if !res.After(d) {
		res = res.Add(time.Duration(p.Frequence) * day)
	}
	if end.After(start) && res.After(end) {
		return time.Time{}
	}
	return res
}

func (p Recurrence) ndWeekly(d time.Time) time.Time {
	start := p.Start.In(p.Location)
	end := p.End.In(p.Location)
	if end.After(start) && !end.After(d) {
		return time.Time{}
	}
	d = d.In(p.Location)

	startDate := p.dateOf(start)
	timeOfDay := start.Sub(startDate)

	startOfWeek, _ := IntToWeeklyPattern(p.Pattern)
	days := p.Pattern & 255

	weekStart := startDate.Add(time.Duration(-(7+int(start.Weekday()-startOfWeek))%7) * day)
	if d.Before(weekStart) {
		d = weekStart
	}
	cycleLength := time.Duration(p.Frequence*7) * day

	// Skip already passed cycles.
	weekStart = weekStart.Add(time.Duration(int(d.Sub(weekStart)/cycleLength)) * cycleLength)
	dayOfD := p.dateOf(d)

	for ws := weekStart; end.Before(start) || !end.Before(ws); ws = ws.Add(cycleLength) {
		for i := 0; i < 7; i++ {
			dat := ws.Add(time.Duration(i) * day)
			if dat.Before(dayOfD) || dat.Before(startDate) {
				continue
			}

			wd := int(1 << uint(dat.Weekday()))
			if (days & wd) != wd {
				continue
			}
			dat = dat.Add(timeOfDay)
			if dat.Before(d) {
				continue
			}
			if end.After(start) && dat.After(end) {
				return time.Time{}
			}
			return dat
		}
	}
	// This should not happen...
	return time.Time{}
}
