package main

// Guide taken from  
// https://tobiasvl.github.io/blog/write-a-chip-8-emulator/#fetchdecodeexecute-loop

import (
    "time"
)

var draw chan bool


func main() {
    w := SDL_WINDOW{}
    w.InitSDLWindow()
    defer w.window.Destroy()

    w.InitColors()
   
    vm := NewEmulator()
    vm.load_rom("roms/ibm.ch8")
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
