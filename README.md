# http-get

A very minimal command line application to perform simple HTTP GET requests, written in Go

## Usage

```
http-get <url>

flags

-o output - write to a file (otherwise stdout if not specified)
-t maxDownloadTime - timeout before abandoning download (default 10s)
-T maxConnectTime - timeout when trying to connect to server (default 5s)
-u userAgent - User-Agent to use when making requests (default is Golang default http client)

```

## Features

* Control maximum timeout before abandoning connections
* Set User-Agent

## Non-features

Other HTTP verbs - this is intended to be minimal - use another client if you need more options.

## Contact

Contact me at luke [at] marmaladefoo [dot] com


