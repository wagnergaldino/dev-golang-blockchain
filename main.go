package main

import (
	"os"

	// "github.com/wagnergaldino/dev-golang-blockchain/cli"
	"github.com/wagnergaldino/dev-golang-blockchain/wallet"
)

func main() {
	defer os.Exit(0)
	// cmd := cli.CommandLine{}
	// cmd.Run()

	w := wallet.MakeWallet()
	w.Address()

}
