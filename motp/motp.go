package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"strconv"
	"time"

	speakeasy "github.com/bgentry/speakeasy"
)

func main() {
	var (
		currentUser *user.User
		outFile     string
		output      string
		vpnUser     string
		hash        string
		otp         string
		secret      string
		pin         string
		err         error
		nowStr      string
	)

	currentUser, err = user.Current()

	flag.StringVar(&vpnUser, "u", currentUser.Name, "OpenVPN user")
	flag.StringVar(&outFile, "o", "", "Output file path")
	flag.Parse()

	if outFile == "" {
		fmt.Println("Must provide output file path to write OTP to.")
		os.Exit(1)
	}

	if secret, err = speakeasy.Ask("Secret (not echoed): "); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if pin, err = speakeasy.Ask("PIN (not echoed): "); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	nowStr = strconv.FormatInt(time.Now().Unix()/10, 10)
	hash = fmt.Sprintf("%x", md5.Sum([]byte(nowStr+secret+pin)))
	otp = hash[:6]
	output = fmt.Sprintf("%s\n%s", vpnUser, otp)
	ioutil.WriteFile(outFile, []byte(output), 0644)
}
