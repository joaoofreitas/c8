package main

// Guide taken from  
// https://tobiasvl.github.io/blog/write-a-chip-8-emulator/#fetchdecodeexecute-loop

import (
    "github.com/veandco/go-sdl2/sdl"
    "time"
)

func main() {
    if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
	panic(err)
    }
    defer sdl.Quit()
    
    vm := NewEmulator()
    vm.load_rom("roms/ibm.ch8")
    vm.print_memory()


    window, err := sdl.CreateWindow("C8 Emulator", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		DISPLAY_WIDTH * VIDEO_SCALE, DISPLAY_HEIGHT * VIDEO_SCALE, sdl.WINDOW_SHOWN)
    if err != nil {
    	panic(err)
    }
    defer window.Destroy()

    surface, err := window.GetSurface()
    if err != nil {
    	panic(err)
    }
    surface.FillRect(nil, 0)

   running := true
    for running {
	vm.execute()
	time.Sleep(time.Second / time.Duration(vm.frequency))
	// Draw white pixel for display 1 and black for 0
    	for i := 0; i < DISPLAY_WIDTH * DISPLAY_HEIGHT; i++ {
    	    pixel := sdl.Rect{}
    	    white := sdl.Color{R: 255, G: 255, B: 255, A: 255}
    	    white_map := sdl.MapRGBA(surface.Format, white.R, white.G, white.B, white.A)	
	    black := sdl.Color{R: 0, G: 0, B: 0, A: 255}
	    black_map := sdl.MapRGBA(surface.Format, black.R, black.G, black.B, black.A)
    	    pixel.X = int32(i % DISPLAY_WIDTH) * VIDEO_SCALE
    	    pixel.Y = int32(i / DISPLAY_WIDTH) * VIDEO_SCALE
    	    pixel.W = VIDEO_SCALE
    	    pixel.H = VIDEO_SCALE

    	    if vm.display[i] > 0 {
    	        surface.FillRect(&pixel, white_map)
    	    } else {
		surface.FillRect(&pixel, black_map)
	    }
    	}

    	window.UpdateSurface()
 

    	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
	    switch event.(type) {
		case sdl.QuitEvent:
		    running = false
		    break
		//case sdl.KeyboardEvent:
		//    key_event := event.(sdl.KeyboardEvent)
		//    if key_event.Type == sdl.KEYDOWN {
		//	switch key_event.Keysym.Sym {
		//	    case sdl.K_1:
		//		fmt.Println("1")
		//	    case sdl.K_q:

	   }
	}
    }
    sdl.Delay(15)
}
