package main

import (
    "bufio"
    "os"
    "fmt"
)


// ReadRom reads a rom file and returns a slice of uint16 representing the instruction
func ReadRom(rom string) ([]uint8, error) {
    var err error = nil;

    f, err := os.Open(rom)
    if err != nil {
	err = fmt.Errorf("Error opening file: %s", err)
	return nil, err
    }
    defer f.Close()
   
    stats, stats_err := f.Stat()
    if stats_err != nil {
	err = fmt.Errorf("Error getting file stats: %s", stats_err)
	return nil, err
    }


    bytes := make([]byte, stats.Size())
    bufr := bufio.NewReader(f)
    _, err = bufr.Read(bytes)

    if err != nil {
	err = fmt.Errorf("Error reading buffer file: %s", err)
	return nil, err
    }

    return bytes, err
}
