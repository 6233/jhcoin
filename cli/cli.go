package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/6233/jhcoin/explorer"
	"github.com/6233/jhcoin/rest"
)

func usage() {
	fmt.Printf("Welcome to jhcoin\n\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port:		Set the PORT of the server\n")
	fmt.Printf("-mode:		Choose between 'html' and 'rest'\n\n")
	os.Exit()
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set prot of the Server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)		
	case "html":
		explorer.Start(*port)	
	default:
		usage()
	}
}