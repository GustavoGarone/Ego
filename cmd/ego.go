package main

import (
	"log"
	"os"

	"github.com/GustavoGarone/ego/internal/emulator"
	"github.com/GustavoGarone/ego/internal/nes"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: ego <path>")
	}

	path := os.Args[1]
	log.Printf("Loading cartridge with path %s\n", path)

	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Failed to load file data: %v", err)
	}

	emulator, err := loadEmulator(bytes)
	if err != nil {
		log.Fatalf("Failed to start emulator: %v", err)
	}

	emulator.Run()
}

func loadEmulator(data []byte) (emulator.Emulator, error) {
	// forcefully loading a NES emulator because we
	// dont support others at the moment.
	emulator, err := nes.New(data)
	if err != nil {
		return nil, err
	}

	return emulator, nil
}
