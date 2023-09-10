package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	go forever()

	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel
	//time for cleanup before exit
	fmt.Println("Adios!") // TODO: add logging package
}

func forever() {
	for {
		fmt.Printf("%v+\n", time.Now())
		time.Sleep(time.Second) // TODO: add sleep duration
	}
}
