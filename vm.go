package main

import (
    "fmt"
)

type Emulator struct {
    memory[MEMORY_SIZE] uint8
    registers[16] uint8
    pc uint16
    i uint16
    sp uint8

    display[DISPLAY_WIDTH * DISPLAY_HEIGHT] uint8
    peripherals[16] uint8
    stack[16] uint16


    frequency uint64
    running bool
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
	e.memory[FONTSET_ADDR + i] = fontset[i]
    }

    e.pc = START_ADDRESS
    e.frequency = 1000000 // 1MHz
    e.running = true
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
    fmt.Printf("Addr | 00 | 01 | 02 | 03 | 04 | 05 | 06 | 07 | 08 | 09 | 0A | 0B | 0C | 0D | 0E | 0F |\n")
    fmt.Printf("======================================================================================\n")

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
            fmt.Printf("%02X | ", e.memory[i+j])
        }
    	fmt.Println()
    }
}

func (e *Emulator) print_display() {
    fmt.Println("Display")
    for y := 0; y < DISPLAY_HEIGHT; y++ {
	fmt.Println() 
	for x := 0; x < DISPLAY_WIDTH; x++ {
	    if e.display[x + (y * DISPLAY_WIDTH)] == 1 {
		fmt.Printf("X")
	    } else {
		fmt.Printf(" ")
	    }
	}
    }
}

func (e *Emulator) push(addr uint16) {
    e.stack[e.sp] = addr
    e.sp++
}

func (e *Emulator) pop() {
    e.stack[e.sp] = 0
    e.sp--
}


//    X: The second nibble. Used to look up one of the 16 registers (VX) from V0 through VF.
//    Y: The third nibble. Also used to look up one of the 16 registers (VY) from V0 through VF.
//    N: The fourth nibble. A 4-bit number.
//    NN: The second byte (third and fourth nibbles). An 8-bit immediate number.
//    NNN: The second, third and fourth nibbles. A 12-bit immediate memory address.

func (e *Emulator) execute() {
    // Fetch
    var i_byte_1 , i_byte_2 uint16
    i_byte_1 = uint16(e.memory[e.pc])
    i_byte_2 = uint16(e.memory[e.pc+1])

    e.pc += 2

    var instruction uint16 = (i_byte_1 << 8) | (i_byte_2)

    // Decode
    var opcode uint16 = instruction & 0xF000

    // execute

    switch opcode {
	case 0x0000:
	    switch instruction {
		case 0x00E0:
		    // Clear the screen
		    for i := 0; i < DISPLAY_WIDTH * DISPLAY_HEIGHT; i++ {
			e.display[i] = 0 
		    }
		    break
	    }
	    break
	case 0x1000:
	    // Jump to address NNN
	    e.pc = instruction & 0x0FFF
	    break
	case 0x6000:
	    // Set VX to NN
	    x := (instruction & 0x0F00) >> 8
	    nn := instruction & 0x00FF
	    e.registers[x] = uint8(nn)
	    break
	case 0x7000:
	    // Add NN to VX
	    x := (instruction & 0x0F00) >> 8
	    nn := instruction & 0x00FF
	    e.registers[x] += uint8(nn)
	    break
	case 0xA000:
	    // Set I to NNN
	    e.i = instruction & 0x0FFF 
	    break
	case 0xD000:
	    // Draw sprite at VX, VY with height N 
	    vx := uint16(e.registers[(instruction & 0x0F00)>>8])
	    vy := uint16(e.registers[(instruction & 0x00F0)>>4])

	    n := instruction & 0x000F
	    e.registers[0xF] = 0 
	    
	    for yline := uint16(0); yline < n; yline++ {
    	        spriteRow := e.memory[e.i + yline]
    	        for xline := uint16(0); xline < 8; xline++ {
    	            if (spriteRow & (0x80 >> xline)) != 0 {
    	                if vx+xline >= DISPLAY_WIDTH || vy+yline >= DISPLAY_HEIGHT {
    	                    continue 
    	                }
    	                addr := (vx + xline) + ((vy + yline) * DISPLAY_WIDTH)
    	                if e.display[addr] == 1 {
    	                    e.registers[0xF] = 1
    	                }
    	                e.display[addr] ^= 1
    	            }
    	        }
    	    }
	    draw <- true
    	    break
    }
}

// Run the vm 
func (e *Emulator) run() {
    fmt.Println("Running...")
}
