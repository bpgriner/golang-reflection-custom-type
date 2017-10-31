## Golang Reflection Example with Custom Types and Nested Slices

Get a custom struct type's field by name
```go
r := reflect.ValueOf(&yourStructType)
f := reflect.Indirect(r).FieldByName("fieldNameInYourStructType")
```

Convert that return Value into it's proper type
```go
yourStructTypes, ok := f.Interface().([]YourStructType)
if !ok {
    panic("f is not a []YourStructType")
}
```

### Alternative Approach
In the example found in main.go of this repo, reflection is being used. However, in this example, the same problem can be solved with the the use of a map:
```go
package main

import (
	"fmt"
)

type DayOfWeek string

type LocalTime string

type ScheduleDetail struct {
	Weekly map[DayOfWeek][]LocalTime
}

func main() {
	scheduleDetail := ScheduleDetail{
		Weekly: map[DayOfWeek][]LocalTime{
			"Sunday":    {"0900", "1100", "1300"},
			"Monday":    {"0900", "1100", "1300"},
			"Tuesday":   {"0900", "1100", "1300"},
			"Wednesday": {"0900", "1100", "1300", "1400"},
			"Thursday":  {"0900", "1100", "1300"},
			"Friday":    {"0900", "1100", "1300"},
			"Saturday":  {"0900", "1100", "1300"},
		},
	}

	dayOfWeek := "Wednesday"

	startTimes := scheduleDetail.Weekly[DayOfWeek(dayOfWeek)]
	fmt.Println(startTimes)
}
```
