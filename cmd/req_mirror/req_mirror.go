package main

import (
	"os"

	"github.com/moznion/req_mirror"
)

func main() {
	reqMirror.Run(os.Args[1:])
}
