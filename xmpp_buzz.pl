#!/usr/bin/perl

##
# xmpp_buzz.pl
# Buzz an XMPP buddy via command line.
#
# marthjod@gmail.com 
# lastmod 2012-11-30
##

use warnings;
use strict;

# libauthen-sasl-perl libnet-xmpp-perl libxml-stream-perl
use Net::XMPP;

my $argc = @ARGV;
die "$0 <user> <pass> <addressee>\@<domain>, ex.\n$0 bob s3cret alice\@jabber.org\n" 
	if ($argc < 3);

my $user = shift;
my $pass = shift;
my $addr = shift;
my $body = shift || '';
my @rcpt = split(/@/, $addr);
my $addressee = $rcpt[0];
my $domain = $rcpt[1];

my $client = new Net::XMPP::Client();

$client->Connect(
    hostname => $domain, 
    port => '5222');
    
$client->AuthSend(
    hostname => $domain,
    username => $user,
    password => $pass);
    
my $msg = <<END;
<message to='$addressee\@$domain' type='headline'>
    <attention xmlns='urn:xmpp:attention:0'/>
    <body>
        $body
    </body>
</message>
END

$client->Send($msg);
$client->Disconnect();

