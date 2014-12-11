// based on example found on https://github.com/howeyc/fsnotify
package main

import (
	"flag"
	"fmt"
	"github.com/howeyc/fsnotify"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
)

type modifyHandler func(string, string)

func main() {

	var (
		err            error
		watcher        *fsnotify.Watcher
		done           chan bool
		includePattern string
		include        *regexp.Regexp
		watchedFile    string
	)

	flag.StringVar(&watchedFile, "watch", "none", `Directory to watch for modification events`)
	flag.StringVar(&includePattern, "include", "", `Filename pattern to include (regex)`)
	flag.Parse()

	include = nil
	if includePattern != "" {
		include = regexp.MustCompile(includePattern)
	}

	watcher, err = fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	done = make(chan bool)

	go func(include *regexp.Regexp, hdlr modifyHandler) {
		for {
			select {
			case ev := <-watcher.Event:
				// log.Println("event:", ev)
				if include == nil || include.MatchString(ev.Name) {
					if ev.IsModify() || ev.IsCreate() {
						log.Printf("Calling handler\n")
						hdlr(ev.Name, "make")
					}
				}
			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}
	}(include, triggerCmd)

	err = watcher.Watch(watchedFile)
	if err != nil {
		fmt.Printf("%s: %v\n", watchedFile, err)
		os.Exit(1)
	}

	<-done

	watcher.Close()
}

func triggerCmd(fname string, command string) {
	var (
		err    error
		cmd    *exec.Cmd
		cmdOut string
		cmdErr string
		b      []byte
		stdout io.ReadCloser
		stderr io.ReadCloser
	)

	cmd = exec.Command(command)

	stdout, err = cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("%s StdoutPipe(): %v\n", cmd.Path, err)
	}

	stderr, err = cmd.StderrPipe()
	if err != nil {
		log.Fatalf("%s StderrPipe(): %v\n", cmd.Path, err)
	}

	if err = cmd.Start(); err != nil {
		log.Fatalf("%s Start(): %v\n", cmd.Path, err)
	}

	if b, err = ioutil.ReadAll(stdout); err == nil {
		cmdOut = string(b)
	}

	if b, err = ioutil.ReadAll(stderr); err == nil {
		cmdErr = string(b)
	}

	if cmdOut != "" {
		fmt.Print(cmdOut)
	}
	if cmdErr != "" {
		fmt.Print(cmdErr)
	}
}
