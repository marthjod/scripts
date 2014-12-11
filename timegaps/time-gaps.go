package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"time"
)

func readLinesFromFile(path string) []string {

	var (
		file    *os.File
		scanner *bufio.Scanner
		list    []string
		err     error
	)

	list = []string{}

	file, err = os.Open(path)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}

	return list
}

func formatTimes(timeStr string) string {
	var (
		re    *regexp.Regexp
		err   error
		times string
	)

	times = ""

	if re, err = regexp.Compile("[0-9]{4}-[0-9]{2}-[0-9]{2}T"); err == nil {
		times = re.ReplaceAllString(timeStr, "")

		if re, err = regexp.Compile("Z"); err == nil {
			times = re.ReplaceAllString(times, "")
		}
	}

	return times
}

func main() {

	var (
		lines      []string
		path       string
		prefix     string
		stamp      string
		i          int
		re         *regexp.Regexp
		err        error
		current    time.Time
		last       time.Time
		diff       time.Duration
		minDiffStr string
		minDiff    time.Duration
		diffs      map[string]time.Duration
	)

	flag.StringVar(&path, "f", "none", "File to parse")
	flag.StringVar(&prefix, "t", "", "Timestamp pattern	in file (regex)")
	flag.StringVar(&minDiffStr, "min-duration", "1s", "Minimum duration to consider in output")
	flag.Parse()

	re, err = regexp.Compile(prefix)
	if err != nil {
		fmt.Printf("Cannot compile regex: %v\n", err.Error())
		os.Exit(1)
	}

	minDiff, err = time.ParseDuration(minDiffStr)
	if err != nil {
		fmt.Printf("No valid duration format: %v\n", err.Error())
		os.Exit(1)
	}

	diffs = make(map[string]time.Duration, 10)

	lines = readLinesFromFile(path)
	for i = 0; i < len(lines); i++ {
		if re.MatchString(lines[i]) {
			stamp = re.FindStringSubmatch(lines[i])[0]
			t, err := time.Parse(time.ANSIC, "Mon Jan 01 "+stamp+" 2104")
			if err != nil {
				fmt.Printf("%v\n", err.Error())
			} else {
				current = t
				diff = current.Sub(last)
				if diff >= minDiff {
					diffs[formatTimes(last.Format(time.RFC3339))+
						"-"+
						formatTimes(current.Format(time.RFC3339))] = diff
				}
				last = current
			}
		}
	}

	// print durations
	for times, diff := range diffs {
		fmt.Printf("%v\t\t(%v)\n", diff, times)
	}

}
