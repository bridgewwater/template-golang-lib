//go:build !test

package main

import (
	_ "embed"
	"flag"
	"log"
	"os"
)

const cliVersion = "0.1.2"

var serverPort = flag.String("serverPort", "49002", "http service address")

func main() {
	log.Printf("-> env:CI_DEBUG %s", os.Getenv("CI_DEBUG"))
	flag.Parse()
	log.Printf("-> run serverPort %v", *serverPort)
	log.Printf("=> now version %v", cliVersion)
}
