#!/usr/bin/python

##
# piegraph.py
# Generates total sums from multiple lines of the form
# <string1>:<number1>
# <string2>:<number2>
# <string1>:<number3> etc.
# Useful together with claimtime.sh to accumulate amounts
# of time spent for distinct applications so they can be 
# fed to a pie graph.
#
# marthjod@gmail.com lastmod 2012-12-01
##

import sys

if len(sys.argv) < 2:
    print "filename missing"
    sys.exit(1)

pie = {}

with open(sys.argv[1], 'r') as f:

    for line in f:
        program_name = line.split(':')[0]
        program_time = line.split(':')[1]
        if (program_time == ''):
            program_time = '0'
        program_time = int(program_time)
        
        if (program_name in pie.keys()):
            pie[program_name] = int(pie[program_name]) + program_time
        else:
            pie[program_name] = program_time

# print totals
for program in pie:
    print "%s:%s" % (program, pie[program]) 
