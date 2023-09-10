package main

import (
	"U6143-ssd1306-golang/ipaddr"
	"fmt"
	"log"
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
	ipv4Addr := ipaddr.Address{}.GetDisplayValueForInterface("eth0")
	log.Printf("Found ipv4 address for eth0 as: %s", ipv4Addr)
	for {
		time.Sleep(time.Second) // TODO: add sleep duration
	}
}
