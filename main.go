package main

import (
	"github.com/6233/jhcoin/blockchain"
	"github.com/6233/jhcoin/cli"
)

func main() {
	blockchain.Blockchain()
	cli.Start()
}