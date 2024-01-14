package main

import (
    "testing"
)

func TestReadRomFile(t *testing.T) {
    // Make sure this does not exist
    rom := "roms/test.ch7"
    _, err := ReadRom(rom)
    if err == nil {
        t.Errorf("Expected file, got nil")
    }
}

func TestReadRom(t *testing.T) {
    rom := "roms/test.ch8"
    data, err := ReadRom(rom)
    if err != nil {
        t.Errorf("Expected nil, got %s", err)
    }
    if len(data) != 478 {
        t.Errorf("Expected 478 bytes, got %d", len(data))
    }
}
