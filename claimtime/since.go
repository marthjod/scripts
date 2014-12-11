package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	var (
		sinceUTC   string
		now        time.Time
		todayStamp []string
		err        error
	)

	flag.StringVar(&sinceUTC, "since-utc", "", "Start time (delta = now - start)")
	flag.Parse()

	if len(sinceUTC) == 0 {
		fmt.Println("Need start time as argument")
		os.Exit(1)
	}

	now = time.Now().UTC()
	fmt.Printf("Now: %v\n", now)

	todayStamp = []string{
		now.Weekday().String()[:3],
		now.Month().String()[:3],
		strconv.Itoa(now.Day()),
		sinceUTC,
		strconv.Itoa(now.Year()),
	}

	start, err := time.Parse(time.ANSIC, strings.Join(todayStamp, " "))
	if err != nil {
		fmt.Printf("%v\n", err.Error())
	} else {
		start = start.UTC()
		fmt.Printf("Delta: %v\n", time.Since(start))
	}
}
