// taken from https://www.alexedwards.net/blog/when-is-it-ok-to-panic-in-go
package main

import (
	"fmt"
	"os"
	"time"
)

// Takes a string parameter `zone` which represents the name of the time zone.
// If the time zone cannot be found, it returns an empty time.Time value and an error.
func timeIn(zone string) (time.Time, error) {
	location, err := time.LoadLocation(zone)
	if err != nil {
		return time.Time{}, err
		//panic(err) // Panic if the location cannot be found
	}
	return time.Now().In(location), nil
}

func main() {
	tz := "Europe/Wonderland"

	t, err := timeIn(tz)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Current time in", tz, "is", t)
}
