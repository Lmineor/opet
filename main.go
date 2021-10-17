package main

import (
	"fmt"
	"math/rand"
	"opet/cmd"
	"opet/pkg/flag"
	"opet/pkg/logs"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	flags := flag.NewFlags()

	command := cmd.NewOpEtCmd(flags)
	logs.InitLogs()
	defer logs.FlushLogs()
	if err := cmd.Execute(command); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to start opet %v\n", err)
		os.Exit(1)
	}
}
