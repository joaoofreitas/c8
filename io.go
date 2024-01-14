package main

import (
    "bufio"
    "os"
    "fmt"
)


// ReadRom reads a rom file and returns a slice of uint16 representing the instruction
func ReadRom(rom string) ([]uint16, error) {
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

    data := make([]uint16, len(bytes) / 2)
    for i := 0; i < len(bytes); i += 2 {
	data[i / 2] = uint16(bytes[i]) << 8 | uint16(bytes[i + 1])
    }

    return data, err
}
