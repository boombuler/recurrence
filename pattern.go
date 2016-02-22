package recurrence

import (
	"time"
)

type Frequence int

const (
	NotRepeating Frequence = iota
	// Repeat every X days
	Daily
	// Repeat every X weeks on some of the Week-Days
	Weekly
	// Repeat every X months on the n-th day of the month
	MonthlyXth
	// Repeat every X months on the n-th weekday of the month
	Monthly
	// Repeat every X years
	Yearly
)

func WeeklyPatternToInt(firstDayOfWeek time.Weekday, days ...time.Weekday) int {
	result := 0
	for _, d := range days {
		result = result | (1 << uint(d))
	}
	return result | int(firstDayOfWeek<<8)
}

func IntToWeeklyPattern(value int) (time.Weekday, []time.Weekday) {
	result := make([]time.Weekday, 0, 7)
	for i := time.Sunday; i <= time.Saturday; i++ {
		test := 1 << uint(i)
		if value&test == test {
			result = append(result, i)
		}
	}
	firstDay := time.Weekday(value >> 8)
	return firstDay, result
}
