#!/usr/bin/python

import sys
import os
import argparse
import re
import requests
import tempfile
from datetime import datetime

description = """
    Find today's recurring event by filter in CalDAV collection.
    Input: .ics file or URL to .ics.
"""

BEGIN_VEVENT = "BEGIN:VEVENT"
END_VEVENT = "END:VEVENT"
SUMMARY = "SUMMARY:"
DTSTART = "DTSTART;"
today = datetime.today().strftime("%Y%m%d")
# 20YYMMdd
date_rex = re.compile("(20\d{6})")


def main():
    parser = argparse.ArgumentParser(description=description)
    parser.add_argument("-i", "--input", action='store', default=None, help="Input file/URL", required=True)
    parser.add_argument("-f", "--filter", action='store', default=None, help="Event filter (regex)", required=True)
    args = parser.parse_args()

    filter_rex = re.compile(args.filter)

    output = []
    in_event = False
    from_url = False
    input_file = args.input

    if input_file.lower().startswith("http"):
        from_url = True
        try:
            # default: do not verify SSL/TLS certificate
            r = requests.get(args.input, verify=False)
        except Exception as ex:
            print ex
            sys.exit(1)

        with tempfile.NamedTemporaryFile(delete=False) as t:
            t.write(r.text.encode('utf-8'))
            t.flush()
            input_file = t.name

    with open(input_file, 'r') as infile:

        for line in infile.readlines():
            line = line.strip()

            if line.startswith(BEGIN_VEVENT):
                in_event = True
                event = {
                    "date": None, # YYYYMMdd
                    "summary": None
                }
            elif line.startswith(SUMMARY):
                if len(filter_rex.findall(line)) > 0:
                    event["summary"] = line.split(SUMMARY)[-1]
            elif line.startswith(DTSTART):
                d = re.search(date_rex, line.split(DTSTART)[-1])
                try:
                    event["date"] = d.groups(1)[0]
                except AttributeError:
                    # no matching date string found
                    pass
            elif line.startswith(END_VEVENT):
                in_event = False

                if event["date"] == today and event["summary"] is not None:
                    print event["summary"]

    if from_url:
        # do not clutter /tmp
        os.unlink(input_file)


if __name__ == "__main__":
    try:
        main()
    except KeyboardInterrupt:
        pass
