# http-get

A very minimal command line application to perform simple HTTP GET requests, written in Go

```
usage

http-get <url>

flags

-o output - write to a file (otherwise stdout)
-t maxDownloadTime - timeout before abandoning download
-T maxConnectTime - timeout when trying to connect to server
-u userAgent - User-Agent to use when making requests

```

## Features

* Control maximum timeout before abandoning connections
* Set User-Agent

## Non-features

Other HTTP verbs - this is intended to be minimal - use another client if you need more options.

## Contact

Contact me at luke [at] marmaladefoo [dot] com


