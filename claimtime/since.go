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
		since      string
		now        time.Time
		start      time.Time
		location   *time.Location
		todayStamp []string
		err        error
	)

	flag.StringVar(&since, "since", "", "Start time (delta := now - start)")
	flag.Parse()

	if len(since) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	location, err = time.LoadLocation("Local")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	} else {
		fmt.Printf("Location: %v\n", location)
	}

	now = time.Now().Local()
	fmt.Printf("Now:\t%v\n", now)

	todayStamp = []string{
		now.Weekday().String()[:3],
		now.Month().String()[:3],
		strconv.Itoa(now.Day()),
		since,
		strconv.Itoa(now.Year()),
	}

	start, err = time.ParseInLocation(time.ANSIC, strings.Join(todayStamp, " "), location)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	} else {
		fmt.Printf("Start:\t%v\n", start)
		fmt.Printf("Delta:\t%v\n", time.Since(start))
	}
}
