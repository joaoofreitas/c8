# C8: A Chip-8 Emulator

## Introduction

C8 is a Chip-8 emulator written in Go. Chip-8 is a simple, interpreted programming language used in the 1970s for programming on computers like the COSMAC VIP and the Telmac 1800. This project aims to emulate the Chip-8 system, allowing users to run Chip-8 programs on modern hardware.

## Features

- Chip-8 Emulation: Accurately emulates the Chip-8 system.
- Written in Go: Utilizes the Go programming language for efficient and readable code.
- SDL Integration: Uses SDL for graphics and input handling.

## Installation

To build and run the emulator, you'll need to have Go installed on your system. Follow these steps:

1. Clone the repository:
   git clone https://github.com/joaoofreitas/c8.git
   cd c8

2. Build the project:
   go build

## Usage

To run a Chip-8 ROM, use the built executable followed by the path to the ROM file:

./c8 <path_to_rom_file>

For example:
./c8 roms/PONG.ch8

## Contributing

Contributions are welcome! If you'd like to contribute, please fork the repository and use a feature branch. Pull requests are warmly welcome.

1. Fork the repository.
2. Create your feature branch (git checkout -b feature/AmazingFeature).
3. Commit your changes (git commit -m 'Add some AmazingFeature').
4. Push to the branch (git push origin feature/AmazingFeature).
5. Open a Pull Request.
