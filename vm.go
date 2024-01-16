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

    delay_timer uint8
    sound_timer uint8

    display[DISPLAY_WIDTH * DISPLAY_HEIGHT] uint8

    peripheral_release uint8 // Internal flag
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

    e.peripheral_release = 255
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

func (e *Emulator) print_registers() {
	fmt.Println("Registers")
	for i := 0; i < 16; i++ {
	fmt.Printf("V%X: %02X\n", i, e.registers[i])
	}
}

func (e *Emulator) print_timer() {
	fmt.Println("Timers")
	fmt.Printf("Delay: %02X\n", e.delay_timer)
	fmt.Printf("Sound: %02X\n", e.sound_timer)
}

func (e *Emulator) print_peripherals() {
	fmt.Println("Peripherals")
	for i := 0; i < 16; i++ {
	    fmt.Printf("P%X: %02X\n", i, e.peripherals[i])
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
    e.print_timer()
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
		case 0x00EE:
		    // Return from subroutine
		    e.pop()
		    e.pc = e.stack[e.sp]
		    break
	    }
	    break
	case 0x1000:
	    // Jump to address NNN
	    e.pc = instruction & 0x0FFF
	    break
	case 0x2000:
	    // Call subroutine at NNN
	    e.push(e.pc)
	    e.pc = instruction & 0x0FFF
	    break
	case 0x3000:
	    // Skip next instruction if VX == NN
	    vx := (instruction & 0x0F00) >> 8
	    nn := instruction & 0x00FF
	    if e.registers[vx] == uint8(nn) {
		e.pc += 2
	    }
	    break
	case 0x4000:
	    // Skip next instruction if VX != NN
	    vx := (instruction & 0x0F00) >> 8
	    nn := instruction & 0x00FF
	    if e.registers[vx] != uint8(nn) {
		e.pc += 2
	    }
	    break
	case 0x5000:
	    // Skip next instruction if VX == VY
	    vx := (instruction & 0x0F00) >> 8
	    vy := (instruction & 0x00F0) >> 4
	    if e.registers[vx] == e.registers[vy] {
		e.pc += 2
	    }
	    break
	case 0x6000:
	    // Set VX to NN
	    vx := (instruction & 0x0F00) >> 8
	    nn := instruction & 0x00FF
	    e.registers[vx] = uint8(nn)
	    break
	case 0x7000:
	    // Add NN to VX
	    vx := (instruction & 0x0F00) >> 8
	    nn := instruction & 0x00FF
	    e.registers[vx] += uint8(nn)
	    break
	case 0x8000:
	    vx := (instruction & 0x0F00) >> 8
	    vy := (instruction & 0x00F0) >> 4
	    switch instruction & 0x000F {
		case 0x0: 	// Set
		    e.registers[vx] = e.registers[vy]
		    break
		case 0x1: 	// OR
		    e.registers[vx] |= e.registers[vy]
		    break
		case 0x2:	// AND
		    e.registers[vx] &= e.registers[vy]
		    break
		case 0x3: 	// XOR
		    e.registers[vx] ^= e.registers[vy]
		case 0x4: 	// ADD
		    sum := uint16(e.registers[vx]) + uint16(e.registers[vy])
		    e.registers[vx] += e.registers[vy] 
		    e.registers[0xF] = 0	// Reset carry
		    if sum > 0xFF {
			e.registers[0xF] = 1	// Carry for register overflow
		    } 
		    break
		case 0x5: 	// SUB VX = VX - VY
		    sub := e.registers[vx] - e.registers[vy]
		    c := 0
		    if e.registers[vx] < e.registers[vy] { 
			c = 0
		    } else {
			c = 1
		    }
		    e.registers[vx] = sub
		    e.registers[0xF] = uint8(c)
		    break;
		case 0x7: 	// SUBN VX = VY - VX
		    e.registers[vx] = e.registers[vy] - e.registers[vx]
		    if e.registers[vy] > e.registers[vx] {
			e.registers[0xF] = 1
		    }  else {
			e.registers[0xF] = 0
		    }
		    break
		case 0x6: 	// SHIFTR VX = VX >> 1 (Has quircks)
		    bit := e.registers[vx] & 0x1
		    e.registers[vx] = e.registers[vx] >> 1
		    e.registers[0xF] = bit
		    break
		case 0xE: 	// SHIFTL VX = VX << 1 (Has quircks)
		    bit := (e.registers[vx] >> 7) & 0x1
		    e.registers[vx] <<= 1
		    e.registers[0xF] = bit
		    break
		}
	    break
	case 0x9000:
	    // Skip next instruction if VX != VY
	    vx := (instruction & 0x0F00) >> 8
	    vy := (instruction & 0x00F0) >> 4
	    if e.registers[vx] != e.registers[vy] {
		e.pc += 2
	    }
	    break
	case 0xA000:
	    // Set I to NNN
	    e.i = instruction & 0x0FFF 
	    break
	case 0xB000:
	    // Jump to NNN + V0 // Fix if CHIP-48 &| SUPERCHIP (Same with ALU)
	    e.pc = (instruction & 0x0FFF) + uint16(e.registers[0])
	    break
	case 0xC000:
	    // Set VX to random number & NN
	    vx := (instruction & 0x0F00) >> 8
	    nn := instruction & 0x00FF
	    e.registers[vx] = random_uint8() & uint8(nn)
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
	case 0xE000:
	    // Skip if key pressed
	    vx := (instruction & 0x0F00) >> 8
	    switch instruction & 0x00FF {
		case 0x9E:
		    if e.peripherals[e.registers[vx]] == 1 {
			e.pc += 2
		    }
		    break
		case 0xA1:
		    if e.peripherals[e.registers[vx]] == 0 {
			e.pc += 2
		    }
		    break
	    }
	    break
	case 0xF000:
	    vx := (instruction & 0x0F00) >> 8
	    switch instruction & 0x00FF {
		// Timers
		case 0x07:
		    e.registers[vx] = e.delay_timer
		    break
		case 0x15:
		    e.delay_timer = e.registers[vx]
		    break
		case 0x18:
		    e.sound_timer = e.registers[vx]
		    break
		case 0x1E:	// Add to I
		    e.i += uint16(e.registers[vx])
		    // Set VF to 1 if overflow
		    if e.i > 0xFFF {
			e.registers[0xF] = 1
		    } else {
			e.registers[0xF] = 0
		    }
		    break
		case 0x0A:	// Get Key 
		    var key_released bool = false
		    // Wait for X key press to be released
		    if (e.peripheral_release < 16) {
			if(e.peripherals[e.peripheral_release] == 0) {
			    key_released = true
			}
		    } else {
			for i, state := range e.peripherals {
		    	    if state == 1 {
		    	        e.peripheral_release = uint8(i)
		    	        break
		    	    }
		    	}
		    }

		    if !key_released  {
			e.pc -= 2
		    } else {
			e.registers[vx] = e.peripheral_release
			e.peripheral_release = 255
		    }

		    break
		case 0x29:	// Set I to location of sprite for digit VX
		    e.i = FONTSET_ADDR + uint16(e.registers[vx]) * 5
		    break
		case 0x33:
		    vx := (instruction & 0x0F00) >> 8
		    e.memory[e.i] = e.registers[vx] / 100
		    e.memory[e.i + 1] = (e.registers[vx] / 10) % 10
		    e.memory[e.i + 2] = (e.registers[vx] % 100) % 10
		    break
		case 0x55:
		    // Store registers V0 through VX in memory starting at location I
		    for i := uint16(0); i <= vx; i++ {
			e.memory[e.i + i] = e.registers[i]
		    }
		    break
		case 0x65:
		    // Read registers V0 through VX from memory starting at location 
		    for i := uint16(0); i <= vx; i++ {
			e.registers[i] = e.memory[e.i + i]
		    }
		    break
	    }
	    break
    }
}
