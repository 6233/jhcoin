package main

import (
	"github.com/6233/jhcoin/explorer"
	"github.com/6233/jhcoin/rest"
)

func main() {
	go explorer.Start(3000)

	rest.Start(4000)
}