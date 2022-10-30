package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)
func main() {
	cnf :=GetConfig()
	for _,dstsrc :=range cnf.Config.Destinations {
		dst1Src :=strings.Split(dstsrc,"->")
		// Setup the tunnel, but do not yet start it yet.
		tunnel := NewSSHTunnel(
			// User and host of tunnel server, it will default to port 22
			// if not specified.
			cnf.Config.Tunnel,
			// Pick ONE of the following authentication methods:
			//PrivateKeyFile("path/to/private/key.pem"), // 1. private key
			cnf.Config.Password,                  // 2. password
			// The destination host and port of the actual server.
			dst1Src[0],
			dst1Src[1],

		)
		// You can provide a logger for debugging, or remove this line to
		// make it silent.
		tunnel.Log = log.New(os.Stdout, "", log.Ldate | log.Lmicroseconds)

		go tunnel.Start()

		fmt.Printf("Successfully  :) %s \n", dst1Src)
	}

	select {}
}
