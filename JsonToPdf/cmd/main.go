package main

import (
	"JsonToPdf/cmd/commands"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.NewSource(time.Now().UnixNano())
	JsonToPdfCmd := commands.NewJsonToPdfCmd()
	if err := JsonToPdfCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
