package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	cnf :=GetConfig()
	//fmt.Println(cnf.Config.Tunnel)
	//fmt.Println(cnf.Config.Destinations)
	//fmt.Println(cnf.Config.Password)

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
		// Start the server in the background. You will need to wait a
		// small amount of time for it to bind to the localhost port
		// before you can start sending connections.

		//println(_dst_1_src[0])
		//println(_dst_1_src[1])

		go tunnel.Start()

		fmt.Printf("Successfully  :) %s \n", dst1Src)
	}

	select {}
}
