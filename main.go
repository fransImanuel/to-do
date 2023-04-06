package main

import (
	"flag"
	"fmt"
	"os"

	"to-do/config"
	"to-do/db"
	"to-do/server"
)

func main() {
	config.InitEnv()
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	db.Init()
	server.Init()
}
