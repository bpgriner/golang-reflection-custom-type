package main

import (
	"fmt"
	"reflect"
)

type LocalTime string

type ScheduleTypeWeekly struct {
	Sunday    []LocalTime
	Monday    []LocalTime
	Tuesday   []LocalTime
	Wednesday []LocalTime
	Thursday  []LocalTime
	Friday    []LocalTime
	Saturday  []LocalTime
}

func main() {
	dayOfWeek := "Wednesday"

	scheduleTypeWeekly := ScheduleTypeWeekly{
		Sunday:    []LocalTime{"0900", "1000", "1200", "1400"},
		Monday:    []LocalTime{"1200", "1600"},
		Tuesday:   []LocalTime{"1200", "1600"},
		Wednesday: []LocalTime{"1200", "1600"},
		Thursday:  []LocalTime{"1200", "1600"},
		Friday:    []LocalTime{"1200", "1600"},
		Saturday:  []LocalTime{"0900", "1000", "1200", "1400"},
	}

	r := reflect.ValueOf(&scheduleTypeWeekly)
	f := reflect.Indirect(r).FieldByName(dayOfWeek)

	fmt.Println(f)
	fmt.Println(f.Kind())
	fmt.Println(f.Slice(0, 2))

	startTimes, ok := f.Interface().([]LocalTime)
	if !ok {
		panic("f is not a []LocalTime")
	}

	fmt.Print("\nstartTimes: ")
	fmt.Println(startTimes)
}
