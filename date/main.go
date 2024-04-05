package main

import (
	"fmt"
	"time"
)

func main() {

	// Time package
	now := time.Now()
	fmt.Printf("now: %v, type: %T\n", now, now)

	year := now.Year()
	// Get the number of the month: month := (int)(now.Month())
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("year: %v, month: %v, day: %v, hour: %v, minute: %v, second: %v\n", year, month, day, hour, minute, second)
	// Format the time
	fmt.Println(now.Format("2006-01-02 15:04:05"))

	dateStr := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, minute, second)
	fmt.Println(dateStr)

	// Parse time
	timeStr := "2021-01-01 12:00:00"
	parseTime, _ := time.Parse("2006-01-02 15:04:05", timeStr)
	fmt.Println(parseTime)
}
