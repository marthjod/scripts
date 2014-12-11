#!/bin/bash

##
# claimtime.sh
# Check regularly which application [window] has (had) 
# focus for how long.
# Useful for keeping track of how much overall time you spend with
# what applications.
#
# marthjod@gmail.com lastmod 2012-11-10
# Ex. claimtime.sh > $(date +%s).csv
#
# Use piegraph.py for totals of all distinct entries;
# From the totals you can then generate a pie graph with LibreOffice, f.ex.
#
# If only interested in a specific application, use the shorthand:
# cat <csv file> | \
# perl -ne 'if( /<programe name>:(\d+)/ ) { $total += $1; print "$total\n"; }'
##

# how often should we check if another window has focus (seconds)?
CHECK_INTERVAL=5
# csv separator
SEPARATOR=":"

# initialized once
last_active_window=""
last_switch_time=""

while :; do
     
    # now
    interval_start=$(date +%s)
    
    # get active window id from window manager
    active_window_id=$(xprop -root | \
    grep "_NET_ACTIVE_WINDOW(WINDOW)" | \
    perl -pe 's/.*window id # (0x\d+)/$1/')

    # get active window string via id
    active_window_string=$(xprop -id $active_window_id | \
    grep "WM_CLASS(STRING)")

    # parse out 'application name'
    active_window_string=$(echo $active_window_string | \
    perl -pe 's/WM_CLASS\(STRING\) = "(.+)",.*/$1/')    
    
    # if application name differs from last check
    if [ "$last_active_window" != "$active_window_string" ]; then

        # calculate time interval
        time_spent=$((interval_start-last_switch_time))
        
        # show line in csv format <program name><separator><seconds spent>
        
        # not the first time, though
        # i.e. only if non-zero
        # (use quotes)
        if [ -n "$last_switch_time" ]; then
            echo $time_spent
        fi
        
        echo -n "$active_window_string$SEPARATOR"               
        
        # save current time 
        last_switch_time=$interval_start
        
    fi
    
    # keep record of last active window for comparison next time
    last_active_window=$active_window_string
    
    sleep $CHECK_INTERVAL
done 
