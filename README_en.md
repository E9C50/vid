# vid - Docker Container File Editor

English | [简体中文](./README.md)

<p align="center">
  <strong>Tired of containers without vim?</strong><br>
  <strong>Tired of manually copying config files back and forth with docker cp?</strong><br>
  <strong>vid makes it simple!</strong>
</p>

---

## 💡 Motivation

When working with Docker containers, we often face these frustrations:

- 🔧 **No vim in containers**: Many minimal container images don't include text editors to reduce size
- 📝 **Editing configs is tedious**: Have to manually `docker cp` files out, edit them, then copy them back
- ⏱️ **Repetitive work**: Every edit requires the copy → edit → copy back workflow

**vid** was created to solve these problems! It lets you edit files inside containers using your local vim editor, **all with a single command**.

## ✨ Features

- 🚀 **One-command editing**: Edit container files directly with a single command, no manual copying
- 💻 **Use your local editor**: Leverage your familiar vim configuration and plugins
- 🔄 **Automated workflow**: Automatically handles copying, editing, and writing back files
- 🧹 **Auto cleanup**: Automatically removes temporary files after editing
- 📊 **Verbose logging**: Use `-v` flag to see detailed operation logs
- ⚡ **Lightweight & fast**: Single binary, no additional dependencies

## 📦 Installation

### Option 1: Download Precompiled Binary (Recommended)

Download precompiled binaries for your platform from the [Releases](https://github.com/E9C50/vid/releases) page.

#### Linux

```bash
wget https://github.com/E9C50/vid/releases/download/vX.X.X/vid-linux-amd64.tar.gz
tar -xzf vid-linux-amd64.tar.gz
sudo mv vid /usr/local/bin/
sudo chmod +x /usr/local/bin/vid
```

#### macOS

```bash
wget https://github.com/E9C50/vid/releases/download/vX.X.X/vid-darwin-amd64.tar.gz
tar -xzf vid-darwin-amd64.tar.gz
sudo mv vid /usr/local/bin/
sudo chmod +x /usr/local/bin/vid
```

#### Windows

Download `vid-windows-amd64.zip` from the [Releases](https://github.com/E9C50/vid/releases) page, extract it, and add `vid.exe` to your system PATH.

### Option 2: Build from Source

Requires Go 1.24 or later:

```bash
# Clone the repository
git clone https://github.com/E9C50/vid.git
cd vid

# Build
go build -o vid main.go

# Install to system path (optional)
sudo mv vid /usr/local/bin/
```

### Option 3: Using go install

```bash
go install github.com/E9C50/vid@latest
```

## 🚀 Usage

### Basic Usage

```bash
vid <container_name_or_id> <file_path_in_container>
```

### Options

- `-v` - Enable verbose logging to see the full operation flow

### Examples

```bash
# Edit nginx configuration file
vid my_nginx_container /etc/nginx/nginx.conf

# Edit application config file
vid web_app /app/config/app.conf

# Use verbose mode to see operation details
vid -v my_container /etc/config.json

# Use container ID
vid a1b2c3d4e5f6 /var/log/app.log
```

### Comparison with Traditional Approach

**Traditional approach** (requires 3 commands):
```bash
docker cp my_container:/etc/nginx/nginx.conf ./nginx.conf
vim ./nginx.conf
docker cp ./nginx.conf my_container:/etc/nginx/nginx.conf
rm ./nginx.conf
```

**Using vid** (only 1 command):
```bash
vid my_container /etc/nginx/nginx.conf
```

## 🔧 How It Works

```
┌─────────────────────────────────────────────────────────────┐
│  1. Copy file from container to local temp file (docker cp) │
│                          ↓                                   │
│  2. Open temp file in local vim editor                      │
│                          ↓                                   │
│  3. Auto copy file back to container after save (docker cp) │
│                          ↓                                   │
│  4. Auto cleanup local temp file                            │
└─────────────────────────────────────────────────────────────┘
```

## 📋 Requirements

- Docker (installed and `docker` command available)
- vim editor (installed on local system)
- Access permissions to target container

## ⚠️ Notes

- Ensure the Docker container is running
- Ensure the target file path exists in the container
- Ensure you have appropriate read/write permissions for the container and file
- Use the `-v` flag to see detailed error information

## 🤝 Contributing

Issues and Pull Requests are welcome!

## 📄 License

This project is licensed under the [MIT License](LICENSE).

## 🌟 Star History

If this project helps you, please give it a star ⭐

---

<p align="center">
  Built with ❤️ and Go
</p>
