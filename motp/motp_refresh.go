package main

import (
	"crypto/md5"
	"fmt"
	//clipboard "github.com/atotto/clipboard"
	speakeasy "github.com/bgentry/speakeasy"
	terminal "github.com/wsxiaoys/terminal"
	color "github.com/wsxiaoys/terminal/color"
	"os"
	"os/signal"
	"strconv"
	"time"
)

func refresh(prefix, otp string) {
	terminal.Stdout.ClearLine()
	terminal.Stdout.Left(len(prefix + " " + otp))
	color.Printf("%s @{y}%s", prefix, otp)
	//clipboard.WriteAll(otp)
}

func main() {
	var (
		secret string
		pin    string
		err    error
		nowStr string
		hash   string
	)

	secret, err = speakeasy.Ask("Secret (not echoed): ")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	pin, err = speakeasy.Ask("PIN (not echoed): ")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for _ = range c {
			terminal.Stdout.Reset()
			os.Exit(3)
		}
	}()

	for {
		// 10-sec granularity
		nowStr = strconv.FormatInt(time.Now().Unix()/10, 10)

		hash = fmt.Sprintf("%x\n", md5.Sum([]byte(nowStr+secret+pin)))

		//fmt.Printf("OTP: %s\n", hash[:6])
		refresh("OTP", hash[:6])

		time.Sleep(10 * time.Second)
	}
}
