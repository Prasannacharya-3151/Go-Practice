package intermediate

import (
	"fmt"
	"time"
)

func main() {
	//1. Current time 
	now := time.Now()
	fmt.Println("Current time:", now)

	//2. Formate time
	fmt.Println(
		"Formatted Time:",
		now.Format("02-02-2006 25:04:05"),
	)

	//3. Parse String to time
	dateStr:= "2026-06-06"
	parsedTime, _ := time.Parse(
		"2006-01-02",
		dateStr,
	)
	fmt.Println("Parsed Time:", parsedTime)

	//4. Add time
	futureTime := now.Add(24 * time.Hour)
	fmt.Println("Tommorrow:", futureTime)

	// 5. Time Difference
	diff := futureTime.Sub(now)
	fmt.Println("Difference:", diff)

	// 6. Sleep
	fmt.Println("Waiting 3 seconds...")
	time.Sleep(3 * time.Second)
	fmt.Println("Done!")
}