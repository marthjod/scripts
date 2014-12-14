#!/usr/bin/python

import argparse

description = """
    Interactively clean up CalDAV collection file
"""

BEGIN_VEVENT = "BEGIN:VEVENT"
END_VEVENT = "END:VEVENT"
SUMMARY = "SUMMARY:"


def main():
    parser = argparse.ArgumentParser(description=description)
    parser.add_argument("-i", "--input", action='store', default=None, help="Input file", required=True)
    parser.add_argument("-o", "--output", action='store', default=None, help="Output file", required=True)
    args = parser.parse_args()

    output = []
    in_event = False
    drop_event = False
    event = []

    with open(args.input, 'r') as infile:

        for line in infile.readlines():
            line = line.strip()

            if line.startswith(BEGIN_VEVENT):
                in_event = True
                event = [line]
            elif line.startswith(SUMMARY):
                drop = raw_input('Drop "%s"? [y/n] ' % line.split(SUMMARY)[-1])
                if drop in ["y", "Y"]:
                    drop_event = True
                else:
                    event.append(line)
            elif line.startswith(END_VEVENT):
                if not drop_event:
                    event.append(line)
                    output.extend(event)

                in_event = False
                drop_event = False
            else:
                if in_event:
                    event.append(line)
                else:
                    output.append(line)

    with open(args.output, 'w') as outfile:
        outfile.write("\n".join(output))


if __name__ == "__main__":
    try:
        main()
    except KeyboardInterrupt:
        pass