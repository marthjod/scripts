# time-gaps

Prints durations between (log file) timestamps.
Unsorted, but including absolute timestamps for reference.

## Building

```bash
go build -o time-gaps timegaps.go
```

## Usage

```bash
-f="none": File to parse
-min-duration="1s": Minimum duration to consider in output
-t="": Timestamp pattern in file (regex)

./time-gaps -f /var/log/auth.log -t "\d{2}:\d{2}:\d{2}" -min-duration="20s"
```

- Example output:

		1m1s		(19:48:59-19:50:00)
		1m58s		(20:15:02-20:17:00)
		44s			(20:26:17-20:27:01)
		31s			(20:28:30-20:29:01)
		59s			(23:20:02-23:21:01)
		39s			(19:50:22-19:51:01)
		1m58s		(19:51:02-19:53:00)
		1m57s		(19:57:03-19:59:00)
		
## Known bugs

- Beyond a certain time threshold, this overflows, of course:

		2562047h47m16.854775807s