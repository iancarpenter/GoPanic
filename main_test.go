// provide a working test for the main.go file and test the timeIn function
// without using external libraries
package main

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeIn(t *testing.T) {
	tz := "Europe/London"

	t1, err := timeIn(tz)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	t2, err := time.LoadLocation(tz)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// output this to the test results console
	t.Log("Location:", t1.Location().String())
	t.Log("Location:", t2.String())

	if t1.Location().String() != t2.String() {
		t.Fatalf("Expected location %v, got %v", t2.String(), t1.Location().String())
	}

	fmt.Println("Current time in", tz, "is", t1)
}
