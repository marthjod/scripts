#!/bin/bash

set -e

VPN_USER=myuser
OTP_FILE=/home/myuser/otp.tmp

rm -f $OTP_FILE
/home/myuser/motp -u $VPN_USER -o $OTP_FILE

cd /home/myuser/openvpn
sudo openvpn --script-security 2  --config vpn-myuser.ovpn --auth-user-pass $OTP_FILE

