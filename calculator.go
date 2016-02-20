package recurrence

import (
	"time"
)

const (
	day = 24 * time.Hour
)

func (p Recurrence) GetNextDate(d time.Time) time.Time {
	switch p.Type {
	case Daily:
		return p.ndDaily(d)
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

	daysToAdd := (p.Frequence - (daysBetween % p.Frequence)) % p.Frequence

	res := dateOfD.Add(time.Duration(daysToAdd)*day + timeOfDay)

	if !res.After(d) {
		res = res.Add(time.Duration(p.Frequence) * day)
	}
	if end.After(start) && res.After(end) {
		return time.Time{}
	}
	return res
}
