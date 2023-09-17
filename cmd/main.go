package main

import (
	"U6143-ssd1306-golang/system"
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
	ipv4Addr := system.Address{}.GetDisplayValueForInterface("eth0")
	log.Printf("Found ipv4 address for eth0 as: %s", ipv4Addr)

	temperature := system.SystemTemperature{}.GetDisplayValueForSystemTemperature("thermal_zone0")
	log.Printf("found SystemTemperature for thermal_zone0 as %s C", temperature)

	for {
		time.Sleep(time.Second) // TODO: add sleep duration
	}
}
