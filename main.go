package main

import (
	"loglol/foo"
)

func main() {
	foo.HTTPGet("www.google.com")
	foo.HTTPGet("http://www.google.com")
}
