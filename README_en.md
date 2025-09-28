# vid - Docker Container File Editor

`vid` is a command-line tool that allows you to edit files inside Docker containers using your local vim editor, without manually running `docker cp` commands.

## Features

Edit files directly inside Docker containers using your local vim editor

## Installation

### Precompiled Binaries (Recommended)

Download precompiled binaries for your platform from the [Releases](https://github.com/E9C50/vid/releases) page:

#### Linux

```bash
wget https://github.com/E9C50/vid/releases/download/vX.X.X/vid-linux-amd64.tar.gz
tar -xzf vid-linux-amd64.tar.gz
sudo mv vid /usr/local/bin/
```

#### macOS

```bash
# Download from Releases page
wget https://github.com/E9C50/vid/releases/download/vX.X.X/vid-darwin-amd64.tar.gz
tar -xzf vid-darwin-amd64.tar.gz
sudo mv vid /usr/local/bin/
```

#### Windows

Download `vid-windows-amd64.zip` from the [Releases](https://github.com/<username>/vid/releases) page, extract it, and add `vid.exe` to your system PATH.

### Building from Source

Requires Go 1.24 or later:

```bash
# Clone the repository
git clone https://github.com/E9C50/vid

cd vid

# Build the binary
go build -o vid main.go

# Optional: Move to a directory in your PATH
# Linux/macOS:
sudo mv vid /usr/local/bin/

# Windows (run as Administrator):
move vid.exe C:\Windows\System32\
```



## Usage

```bash
vid [OPTIONS] <container> <file_path>
```

### Options

- `-v`: Enable verbose logging

### Examples

```bash
# Edit a file in a container
vid my_container /etc/nginx/nginx.conf

# Edit with verbose output
vid -v my_container /etc/nginx/nginx.conf
```

## How It Works

1. Copies the target file from the container to a temporary local file using `docker cp`
2. Opens the temporary file in your local vim editor
3. Copies the modified file back to the container after you save and exit
4. Cleans up the temporary file

## Requirements

- Docker
- vim editor
- Go 1.24+ (for building from source only)

## License

[MIT LICENSE](https://github.com/E9C50/vid/LICENSE)