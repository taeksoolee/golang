package main

import (
	"log"
	"os"
)

var myLogger *log.Logger

func main() {
	// New(stream, prefix, flag)
	myLogger = log.New(os.Stdout, "INFO: ", log.LstdFlags)

	run()

	myLogger.Println("End of Program")
}

func run() {
	myLogger.Println("Test")
}