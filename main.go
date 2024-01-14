package main

// Guide taken from  
// https://tobiasvl.github.io/blog/write-a-chip-8-emulator/#fetchdecodeexecute-loop

import (
    "github.com/veandco/go-sdl2/sdl"
    "time"
    "fmt"
)

var draw chan bool

type SDL_WINDOW struct {
    // Window
    window *sdl.Window
    surface *sdl.Surface
   
    // Colors
    white_map uint32
    black_map uint32 

    // Current pixel being drawn
    pixel sdl.Rect
}

func (w *SDL_WINDOW) InitSDLWindow()  {
	var err error
	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
	    sdl.Quit()
    	    panic(err)
    	}

	w.window, err = sdl.CreateWindow("C8 Emulator", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		DISPLAY_WIDTH * VIDEO_SCALE, DISPLAY_HEIGHT * VIDEO_SCALE, sdl.WINDOW_SHOWN)
	if err != nil {
		sdl.Quit()
		panic(err)
	}
}

func (w *SDL_WINDOW) InitColors() {
    var err error
    w.surface, err = w.window.GetSurface()
    if err != nil {
    	panic(err)
    }
    w.surface.FillRect(nil, 0)

    w.pixel = sdl.Rect{}
    w.white_map = 0xFFFFFFFF
    w.black_map = 0x00000000
    w.pixel.W = VIDEO_SCALE
    w.pixel.H = VIDEO_SCALE
}

func (w *SDL_WINDOW) Draw(vm *Emulator) {
    // Draw white pixel for display 1 and black for 0
    for i := 0; i < DISPLAY_WIDTH * DISPLAY_HEIGHT; i++ {
       w.pixel.X = int32(i % DISPLAY_WIDTH) * VIDEO_SCALE
       w.pixel.Y = int32(i / DISPLAY_WIDTH) * VIDEO_SCALE
        if vm.display[i] > 0 {
            w.surface.FillRect(&w.pixel, w.white_map)
        } else {
    	w.surface.FillRect(&w.pixel, w.black_map)
        }
    }
    w.window.UpdateSurface()
}

func (w *SDL_WINDOW) HandleEvents(vm *Emulator) {
    for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
	switch event.(type) {
	    case sdl.QuitEvent:
		vm.running = false
		break
	    case sdl.KeyboardEvent:
		key_event := event.(sdl.KeyboardEvent)
		if key_event.Type == sdl.KEYDOWN {
		    switch key_event.Keysym.Sym {
			case sdl.K_1:
			    fmt.Println("1")
			    break
			case sdl.K_2:
			    fmt.Println("2")
			    break
		        case sdl.K_3:
			    fmt.Println("3")
			    break
		        case sdl.K_q:
			    fmt.Println("4")
			    break   
		        case sdl.K_w:
			    fmt.Println("5")
			    break
			case sdl.K_e:
			    fmt.Println("6")
			    break
			case sdl.K_a:
			    fmt.Println("7")
			    break
			case sdl.K_s:
			    fmt.Println("8")
			    break
			case sdl.K_d:
			    fmt.Println("9")
			    break
			case sdl.K_z:
			    fmt.Println("A")
			    break
			case sdl.K_x:
			    fmt.Println("0")
			    break
			case sdl.K_c:
			    fmt.Println("B")
			    break
			case sdl.K_4:
			   fmt.Println("C")
			    break
			case sdl.K_r:
			   fmt.Println("D")
			    break
			case sdl.K_f:
			    fmt.Println("E")
			    break
			case sdl.K_v:
			    fmt.Println("F")
			    break
		    }
		}
		break
	}
	break;
    }
}



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
