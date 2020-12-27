package main

import (
	"log"
	"os"
	"io"
)

var myLogger *log.Logger

func main() {
	// setOutputTest()
	multiWriterLogTest()
}


func loggerTest() {
	fpLog, err := os.OpenFile("log/log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	defer fpLog.Close()

	myLogger = log.New(fpLog, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	run1()
	myLogger.Println("End of Program")
}

func run1() {	
	myLogger.Print("Test2")
}

func setOutputTest() {
	fpLog, err := os.OpenFile("log/log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	log.SetOutput(fpLog)
	run2()
	log.Println("End of Program")
}

func multiWriterLogTest() {
	fpLog, err := os.OpenFile("log/log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	multiWriter := io.MultiWriter(fpLog, os.Stdout)
	log.SetOutput(multiWriter)

	run2()
	log.Println("End of Program")
}

func run2() {
	log.Println("Test2")
}

