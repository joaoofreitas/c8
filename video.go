package main

import (
    "github.com/veandco/go-sdl2/sdl"
)

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
    w.white_map = ACCENT_COLOR
    w.black_map = BG_COLOR 
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
			    vm.peripherals[0x1] = 1
			    break
			case sdl.K_2:
			    vm.peripherals[0x2] = 1
			    break
		        case sdl.K_3:
			    vm.peripherals[0x3] = 1
			    break
		        case sdl.K_q:
			    vm.peripherals[0x4] = 1
			    break   
		        case sdl.K_w:
			    vm.peripherals[0x5] = 1
			    break
			case sdl.K_e:
			    vm.peripherals[0x6] = 1
			    break
			case sdl.K_a:
			    vm.peripherals[0x7] = 1
			    break
			case sdl.K_s:
			    vm.peripherals[0x8] = 1
			    break
			case sdl.K_d:
			    vm.peripherals[0x9] = 1
			    break
			case sdl.K_z:
			    vm.peripherals[0xA] = 1
			    break
			case sdl.K_x:
			    vm.peripherals[0x0] = 1
			    break
			case sdl.K_c:
			    vm.peripherals[0xB] = 1
			    break
			case sdl.K_4:
			    vm.peripherals[0xC] = 1
			    break
			case sdl.K_r:
			    vm.peripherals[0xD] = 1
			    break
			case sdl.K_f:
			    vm.peripherals[0xE] = 1
			    break
			case sdl.K_v:
			    vm.peripherals[0xF] = 1
			    break
		    }
		} else if key_event.Type == sdl.KEYUP {
		    switch key_event.Keysym.Sym {
			case sdl.K_1:
			    vm.peripherals[0x1] = 0
			    break
			case sdl.K_2:
			    vm.peripherals[0x2] = 0
			    break
			case sdl.K_3:
			    vm.peripherals[0x3] = 0
			    break
			case sdl.K_q:
			    vm.peripherals[0x4] = 0
			    break
			case sdl.K_w:
			    vm.peripherals[0x5] = 0
			    break
			case sdl.K_e:
			    vm.peripherals[0x6] = 0
			    break
			case sdl.K_a:
			    vm.peripherals[0x7] = 0
			    break
			case sdl.K_s:
			    vm.peripherals[0x8] = 0
			    break
			case sdl.K_d:
			    vm.peripherals[0x9] = 0
			    break
			case sdl.K_z:
			    vm.peripherals[0xA] = 0
			    break
			case sdl.K_x:
			    vm.peripherals[0x0] = 0
			    break
			case sdl.K_c:
			    vm.peripherals[0xB] = 0
			    break
			case sdl.K_4:
			    vm.peripherals[0xC] = 0
			    break
			case sdl.K_r:
			    vm.peripherals[0xD] = 0
			    break
			case sdl.K_f:
			    vm.peripherals[0xE] = 0
			    break
			case sdl.K_v:
			    vm.peripherals[0xF] = 0
			    break
			}
		}
		break
	}
	break;
    }
}
