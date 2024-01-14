package main

// Guide taken from  
// https://tobiasvl.github.io/blog/write-a-chip-8-emulator/#fetchdecodeexecute-loop

import (
    "time"
    "os"
    "math/rand"
)

var draw chan bool

func random_uint8() uint8 {
	return uint8(rand.Intn(255))
}

func main() {
    // Get arguments from command line
    args := os.Args[1:]
    rom_location := args[0]

    w := SDL_WINDOW{}
    w.InitSDLWindow()
    defer w.window.Destroy()

    w.InitColors()
   
    vm := NewEmulator()
    vm.load_rom(rom_location)
    draw = make(chan bool) 

    // Run the machine at 1MHz 
    go func() {
	for vm.running {
	    vm.execute()	
    	    time.Sleep(time.Second / time.Duration(vm.frequency))
	}
    }()
   
    // Making user input async
    go func() {
	for vm.running {
	    w.HandleEvents(vm);
	}
	close(draw)	// Closing channel when VM is stopped so there's nothing more to draw
    }()
    
    // Drawing only when having a draw signal
    for vm.running {
	_ = <- draw
    	w.Draw(vm)
    }
}
