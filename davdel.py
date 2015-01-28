import httplib
import fileinput

for line in fileinput.input():
	h=httplib.HTTPConnection('myhostname')
	print '/dav/'+line.strip()
	h.request('DELETE', '/dav/' + line.strip())


