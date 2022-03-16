package main

import (
	"github.com/6233/jhcoin/cli"
	"github.com/6233/jhcoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}