package main

import (
    "testing"
)

func TestNewEmulator(t *testing.T) {
    e := NewEmulator()
    if e == nil {
	t.Error("Expected emulator to be created")
    }

    if len(e.memory) != MEMORY_SIZE {
	t.Error("Expected memory to be created")
    }

    if len(e.registers) != 16 {
	t.Error("Expected registers to be created")
    }
}

func TestFontset(t *testing.T) {
    e := NewEmulator()
    if e.memory[FONTSET_ADDR] != 0xF0 {
	t.Error("Expected fontset to be loaded into memory")
    }
    
    if e.memory[FONTSET_ADDR + 79] != 0x80 {
	t.Error("Expected fontset to be loaded into memory")
    }
}

func TestPC (t *testing.T) {
    e := NewEmulator()
    if e.pc != START_ADDRESS {
        t.Error("Expected pc to be set to START_ADDRESS")
    }
}

func TestMemoryPrint(t *testing.T) {
	e := NewEmulator()
	e.load_rom("roms/test.ch8")
	e.print_memory()
}

func TestLoadROM(t *testing.T) {
    e := NewEmulator()
    e.load_rom("roms/test.ch8")

    if e.memory[START_ADDRESS] != 0 {
	t.Error("Expected ROM to be loaded into memory")
    }
}
