package recurrence

type Type int

const (
	// Repeat every X days
	Daily Type = iota
	// Repeat every X weeks on some of the Week-Days
	Weekly
	// Repeat every X months on the n-th day of the month
	MonthlyXth
	// Repeat every X months on the n-th weekday of the month
	Monthly
	// Repeat every X years
	Yearly
)
