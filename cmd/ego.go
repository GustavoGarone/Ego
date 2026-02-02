package main

import (
	"log"
	"os"

	"github.com/GustavoGarone/ego/internal/bus"
	"github.com/GustavoGarone/ego/internal/cartridge"
	"github.com/GustavoGarone/ego/internal/cpu"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: ego <path>")
	}

	path := os.Args[1]
	log.Printf("Loading cartridge with path %s\n", path)

	cart, err := cartridge.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read cartridge: %v", err)
	}

	bus := bus.New(cart)
	cpu := cpu.New(bus)

	cpu.Run()
}
