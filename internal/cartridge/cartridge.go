package cartridge

import (
	"errors"
	"os"
)

// Cartridge contains all the data needed to render or process the game.
type Cartridge struct {
	ProgramData  []byte
	GraphicsData []byte
	MapperID     byte

	// 1 for true, 0 for false
	Mirroring byte
}

func New(programData, graphicsData []byte) *Cartridge {
	return &Cartridge{
		ProgramData:  programData,
		GraphicsData: graphicsData,
	}
}

// NoGraphics returns a cartridge with an empty graphics memory. Useful for testing
func NoGraphics(data []byte) *Cartridge {
	return New(data, []byte{})
}

// Reads the content in the address. Useful for when the cartridge contains its own mappers.
func (c *Cartridge) Read(address uint16) byte {
	if address < 0x8000 {
		return 0
	}

	// mirrors for 16kb roms, handles 32kb roms normally
	normalizedAddress := (address - 0x8000) % uint16(len(c.ProgramData))
	if int(normalizedAddress) < len(c.ProgramData) {
		return c.ProgramData[normalizedAddress]
	}

	return 0
}

// Writes the content to the address. Note that this should not be called normally -- only if
// the cartridge contains mappers.
func (c *Cartridge) Write(address uint16, data byte) {
	// TODO: do nothing for now
}

// ReadFile returns a cartridge from file data. Useful for CLI or GUI uses.
func ReadFile(path string) (*Cartridge, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if len(bytes) < 16 || string(bytes[0:3]) != "NES" || bytes[3] != 0x1a {
		return nil, errors.New("Not a valid iNES file")
	}

	programSize := int(bytes[4]) * 16384 // units of 16kb
	graphicsSize := int(bytes[5]) * 8192 // units of 8kb

	hasTrainer := bytes[6]&0b0000_0100 != 0
	mirroring := bytes[6] & 0b0000_0001

	mapperLow := bytes[6] >> 4
	mapperHigh := bytes[7] & 0xF0
	mapperID := mapperHigh | mapperLow

	programStart := 16
	if hasTrainer {
		programStart += 512
	}

	programEnd := programStart + programSize
	graphicsEnd := programEnd + graphicsSize

	if len(bytes) < graphicsEnd {
		return nil, errors.New("File is smaller than the header claims")
	}

	return &Cartridge{
		ProgramData:  bytes[programStart:programEnd],
		GraphicsData: bytes[programEnd:graphicsEnd],
		MapperID:     mapperID,
		Mirroring:    mirroring,
	}, nil
}
