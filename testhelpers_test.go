package recurrence

import (
	"fmt"
	"time"
)

func ShouldHappenOn(actual interface{}, expected ...interface{}) string {
	if len(expected) != 1 {
		return "Invalid parameters for ShouldHappenOn"
	}
	actualTime, firstOk := actual.(time.Time)
	expectedTime, secondOk := expected[0].(time.Time)

	if !firstOk || !secondOk {
		return "ShouldHappenOn should be used on time values"
	}

	if actualTime.Before(expectedTime) || actualTime.After(expectedTime) {
		return fmt.Sprintf("Expected to happen on %v but happened on %v", expectedTime, actualTime)
	}

	return ""
}

func ShouldNotHappen(actual interface{}, expected ...interface{}) string {
	actualTime, ok := actual.(time.Time)

	if !ok {
		return fmt.Sprintf("Expected %v to be a time value", actual)
	}

	if !actualTime.IsZero() {
		return fmt.Sprintf("Time value should be zero but is %v", actualTime)
	}

	return ""
}
