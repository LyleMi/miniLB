package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	configFile := flag.String("config", "config.yaml", "Path to the config file")
	flag.Parse()

	config, err := ReadConfig(*configFile)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	for index, _ := range config.Entries {
		// If you directly obtain a value and pass it by reference, it will pass the same pointer, which can lead to issues.
		go HandleEntry(&config.Entries[index])
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	log.Printf("Ctrl+C pressed")
	<-sig
}
