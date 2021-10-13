package main

import (
	"fmt"
	"math/rand"
	cmd "opet/pkg/cmd"
	"opet/pkg/flag"
	"opet/pkg/logs"
	"os"
	"time"
)

func main(){
	rand.Seed(time.Now().UnixNano())
	command := cmd.NewOpEtCmd()
	flag.InitFlags()

	logs.InitLogs()
	defer logs.FlushLogs()
	if err:= command.Execute(); err!= nil{
		fmt.Fprintf(os.Stderr, "Failed to start opet %v\n", err)
		os.Exit(1)
	}
}


