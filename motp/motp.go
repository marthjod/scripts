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

	"github.com/atotto/clipboard"
	"github.com/bgentry/speakeasy"
)

func main() {

	currentUser, _ := user.Current()

	var (
		output          string
		hash            string
		otp             string
		secret          string
		pin             string
		err             error
		nowStr          string
		vpnUser         = flag.String("u", currentUser.Name, "OpenVPN user")
		outFile         = flag.String("o", "", "Output file path")
		secretClipboard = flag.Bool("s", false, "Read secret from clipboard")
	)
	flag.Parse()

	if *outFile == "" {
		fmt.Println("Must provide output file path to write OTP to.")
		os.Exit(1)
	}

	if *secretClipboard {
		secret, _ = clipboard.ReadAll()
	}
	if secret == "" {
		if secret, err = speakeasy.Ask("Secret (not echoed): "); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}
	if pin, err = speakeasy.Ask("PIN (not echoed): "); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	nowStr = strconv.FormatInt(time.Now().Unix()/10, 10)
	hash = fmt.Sprintf("%x", md5.Sum([]byte(nowStr+secret+pin)))
	otp = hash[:6]
	output = fmt.Sprintf("%s\n%s", *vpnUser, otp)
	ioutil.WriteFile(*outFile, []byte(output), 0644)
}
