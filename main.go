package main

// Guide taken from  
// https://tobiasvl.github.io/blog/write-a-chip-8-emulator/#fetchdecodeexecute-loop

import (
    "github.com/veandco/go-sdl2/sdl"
)

func main() {
    if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
	panic(err)
    }
    defer sdl.Quit()

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

    //// Draw white pixel for display 1 and black for 0
    //for i := 0; i < DISPLAY_WIDTH * DISPLAY_HEIGHT; i++ {
    //    pixel := sdl.Rect{}
    //    color := sdl.Color{R: 255, G: 255, B: 255, A: 255}
    //    color_map := sdl.MapRGBA(surface.Format, color.R, color.G, color.B, color.A)	
    //    if display[i] {
    //        pixel.X = int32(i % DISPLAY_WIDTH) * VIDEO_SCALE
    //        pixel.Y = int32(i / DISPLAY_WIDTH) * VIDEO_SCALE
    //        pixel.W = VIDEO_SCALE
    //        pixel.H = VIDEO_SCALE
    //        surface.FillRect(&pixel, color_map)
    //    }
    //}

    window.UpdateSurface()
    running := true
    for running {
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
