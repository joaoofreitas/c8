package main

import (
    "fmt"
)

type Emulator struct {
    memory[MEMORY_SIZE] uint16
    registers[16] uint8
    pc uint16
    display[DISPLAY_WIDTH * DISPLAY_HEIGHT] bool
}

// Constructor for Emulator, fontset loading, initialization of PC, etc...
func NewEmulator() *Emulator {
    e := new(Emulator)

    // Load fontset into memory	
    fontset := [80]uint8 {
	0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
	0x20, 0x60, 0x20, 0x20, 0x70, // 1
	0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
	0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
	0x90, 0x90, 0xF0, 0x10, 0x10, // 4
	0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
	0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
	0xF0, 0x10, 0x20, 0x40, 0x40, // 7
	0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
	0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
	0xF0, 0x90, 0xF0, 0x90, 0x90, // A
	0xE0, 0x90, 0xE0, 0x90, 0xE0, // B
	0xF0, 0x80, 0x80, 0x80, 0xF0, // C
	0xE0, 0x90, 0x90, 0x90, 0xE0, // D
	0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
	0xF0, 0x80, 0xF0, 0x80, 0x80, // F
    }
    
    // Copy fontset into memory address 0x50
    for i := 0; i < 80; i++ {
	e.memory[FONTSET_ADDR + i] = uint16(fontset[i])
    }

    e.pc = START_ADDRESS
    return e
}

// Load ROM into memory
func (e *Emulator) load_rom(rom string) {
    data, err := ReadRom(rom)
    if err != nil {
	panic(err)
    }
    
    for i := 0; i < len(data); i++ {
	e.memory[START_ADDRESS + i] = data[i]
    }
}

// Prints memory in a pretty table of addresses last address digit will be on collumns and rest as lines
func (e *Emulator) print_memory() {
    fmt.Println("Memory Dump:")
    fmt.Printf("Addr |  00  |  01  |  02  |  03  |  04  |  05  |  06  |  07  |  08  |  09  |  0A  |  0B  |  0C  |  0D  |  0E  |  0F  |\n")
    fmt.Printf("======================================================================================================================\n")

    for i := 0; i < MEMORY_SIZE; i += 16 {
        empty := true 	// Check if the entire row is 0000 and skip it if true
        for j := 0; j < 16; j++ {
    	if e.memory[i+j] != 0 {
    	    empty = false
                break
            }
        }
        if empty {
            continue
        }
    
        // Print the memory address
        fmt.Printf("%04X | ", i)
        // Print each memory cell
        for j := 0; j < 16; j++ {
            fmt.Printf("%04X | ", e.memory[i+j])
        }
    	fmt.Println()
    }
}

// Run the vm 
func (e *Emulator) run() {
    fmt.Println("Running...")
}
