package main

import (
	"io"
	"log"
	"os"
)

var (
	Info    *log.Logger
	Error   *log.Logger
	Discard *log.Logger
)

func init() {
	Info = log.New(os.Stdout, "INFO:", log.Ldate|log.Llongfile)

	file, err := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("Failed to open error.log:" + err.Error())
	}
	Error = log.New(io.MultiWriter(os.Stdout, file), "ERROR:", log.Ldate|log.Llongfile)

	Discard = log.New(io.Discard, "DISCARD:", 0)
}

func main() {
	Info.Println("something happened")
	Error.Println("error happened")
	Discard.Println("discard log message")
}
