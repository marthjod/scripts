package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {

	var (
		from  string
		to    string
		slots int64
		err   error
		t1    time.Time
		t2    time.Time
		d     time.Duration
		slot  int64
		t     int64
	)

	// -from "$(date -d"last Monday 07:30" +%FT%TZ)" -to "$(date -d"today 08:30" +%FT%TZ )"
	flag.StringVar(&from, "from", "", "Start time (RFC3339)")
	flag.StringVar(&to, "to", "", "End time (RFC3339)")
	flag.Int64Var(&slots, "slots", 2, "Number of slots")
	flag.Parse()

	if from == "" || to == "" {
		flag.Usage()
		os.Exit(1)
	}

	t1, err = time.Parse(time.RFC3339, from)
	if err != nil {
		panic(err.Error())
	}
	t2, err = time.Parse(time.RFC3339, to)
	if err != nil {
		panic(err.Error())
	}

	d = t2.Sub(t1)
	slot = int64(d) / slots

	fmt.Printf("%v/%v = %v\n", d, slots, slot)
	for t = int64(t1.UnixNano()); t < int64(t2.UnixNano()); t += slot {
		fmt.Printf("%v\n", t)
	}

}
