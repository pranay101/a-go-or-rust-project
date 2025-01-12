# GoCraft - Minecraft in Go

A simple Minecraft-like game built with Go and Ebitengine.

## Prerequisites

### Go Installation
1. Install Go from [golang.org](https://golang.org/dl/) (version 1.19 or later recommended)
2. Verify installation with:
   ```bash
   go version
   ```

### System Dependencies

Depending on your operating system, you'll need to install some additional dependencies:

#### Ubuntu/Debian
```

#### Fedora/RHEL
```bash
sudo dnf install mesa-libGL-devel xorg-x11-server-devel
```

#### Manjaro/Arch
```bash
sudo pacman -S mesa libxrandr libxcursor libxinerama libxi libxxx
```

If you need additional build tools:
```bash
sudo pacman -S base-devel
```

#### Windows
No additional dependencies required. Just make sure Go is installed.

#### macOS
No additional dependencies required if you have Xcode installed.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/gocraft.git
   cd gocraft
   ```

2. Install Go dependencies:
   ```bash
   go mod tidy
   ```

## Running the Game

To run the game, use:
```bash
go run main.go
```

## Game Controls

Currently, the game displays a simple 10x10 grid of green blocks.

## Development

The game is built using:
- [Go](https://golang.org/)
- [Ebitengine](https://ebitengine.org/) - A dead simple 2D game engine for Go

## Project Structure

- `main.go` - Main game logic and rendering
- More files to be added as the project grows

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details