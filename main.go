package main

import (
	"fmt"
	"time"

	"github.com/thlib/go-timezone-local/tzlocal"
)

func main() {
	for {
		fmt.Println(LazyInstallTime())
		time.Sleep(5 * time.Second)
	}
	// 2021-10-31 20:00:00 +0100 CET
}

func TruncateToDay(t time.Time) time.Time {
	tzname, err := tzlocal.RuntimeTZ()
	fmt.Println(tzname, err)

	// example:
	// tzname = "Europe/Berlin"

	// now you can use tzname to properly set up a location:
	loc, _ := time.LoadLocation(tzname)

	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, loc)
}

func LazyInstallTime() time.Time {
	return TruncateToDay(time.Now()).Add(24 * time.Hour)
}