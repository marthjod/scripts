package main

import (
	"flag"
	"fmt"
	"github.com/dchest/uniuri"
	"net/http"
	"strconv"
)

func main() {

	var (
		err        error
		port       int
		dir        string
		randPrefix string
	)

	flag.IntVar(&port, "port", 4242, "Local port to listen on")
	flag.StringVar(&dir, "dir", "", "Directory with files to serve")
	flag.Parse()

	randPrefix = "/" + uniuri.NewLen(32) + "/"

	http.Handle(randPrefix, http.StripPrefix(randPrefix, http.FileServer(http.Dir(dir))))

	fmt.Printf("Serving on :%d%s\n", port, randPrefix)
	err = http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
