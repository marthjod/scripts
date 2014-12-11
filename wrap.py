import textwrap
import sys
import os

def usage():
    print sys.argv[0] + " <in-file> <out-file> <width>"

def main():
    if(len(sys.argv) < 3):
        usage()
        sys.exit(1)
    else:
        infile = sys.argv[1]
        outfile = sys.argv[2]
        width = int(sys.argv[3]) if len(sys.argv) > 3 else 80

        if not os.path.exists(infile):
            print "Input file %s does not exist" % infile
            sys.exit(1)
        print "Reading %s" % infile
        with open(infile, 'r') as i:
            line_string = i.read()
        with open(outfile, 'w') as o:
            w = textwrap.TextWrapper(replace_whitespace=False, \
                    break_on_hyphens=False, \
                    break_long_words=False, \
                    width=width)
            print "Writing %s wrapped at column %i" % (outfile, width)
            o.write( w.fill(line_string) )
        print "Done. Have a nice day."

if __name__ == "__main__":
    main()
