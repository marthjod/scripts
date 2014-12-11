watchmakr
=========

Watch directory and trigger external command on file modification events


## Example use case

- Edit text-based source files (Markdown...) which get converted to PDF ([pandoc](http://johnmacfarlane.net/pandoc/)...) as laid out in a Makefile.
- Every time you save the source file, you want _make_ to run automatically, building stuff anew and causing an already opened PDF viewer to refresh, showing the newly-generated output.
- Only included file patterns should trigger re-builds. (These are provided as regular expressions, not glob expressions.) Without any inclusion pattern, every file change acts as a trigger.

```bash
% ./watchmakr -h
Usage of ./watchmakr:
  -include="": Filename pattern to include (regex)
  -watch="none": Directory to watch for modification events

 % ./watchmakr -include=".*\.(md|bib)$" -watch=.
```

## Dependencies

- [github.com/howeyc/fsnotify](https://github.com/howeyc/fsnotify)
