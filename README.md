# vid - Docker容器文件编辑工具

vid 是一个命令行工具，允许您使用本地 vim 编辑器直接编辑 Docker 容器中的文件，而无需手动执行 `docker cp` 命令。

## 功能特点

使用本地 vim 编辑器直接编辑 Docker 容器中的文件

## 安装

### 发布版本安装（推荐）

从 [Releases](https://github.com/E9C50/vid/releases) 页面下载适用于您平台的预编译二进制文件：

#### Linux

```bash
wget https://github.com/E9C50/vid/releases/download/vX.X.X/vid-linux-amd64.tar.gz
tar -xzf vid-linux-amd64.tar.gz
sudo mv vid /usr/local/bin/
```

#### macOS

```bash
# 从 Releases 页面下载
wget https://github.com/E9C50/vid/releases/download/vX.X.X/vid-darwin-amd64.tar.gz
tar -xzf vid-darwin-amd64.tar.gz
sudo mv vid /usr/local/bin/
```

#### Windows

从 [Releases](https://github.com/<username>/vid/releases) 页面下载 `vid-windows-amd64.zip` 文件，解压后将 `vid.exe` 添加到系统 PATH 环境变量中。

### 从源码构建

需要安装 Go 1.24 或更高版本：

```bash
# 克隆仓库
git clone https://github.com/E9C50/vid

cd vid

# 构建二进制文件
go build -o vid main.go

# 可选：移动到 PATH 中的目录
# Linux/macOS:
sudo mv vid /usr/local/bin/

# Windows (以管理员权限运行):
move vid.exe C:\Windows\System32\
```



## 使用方法

```bash
vid [选项] <容器> <文件路径>
```

### 选项

- `-v`: 启用详细日志

### 示例

```bash
# 编辑容器中的文件
vid my_container /etc/nginx/nginx.conf

# 带详细日志的编辑
vid -v my_container /etc/nginx/nginx.conf
```

## 工作原理

1. 使用 `docker cp` 将目标文件从容器复制到本地临时文件
2. 在本地 vim 编辑器中打开临时文件
3. 保存并退出后将修改后的文件复制回容器
4. 清理临时文件

## 依赖

- Docker
- vim 编辑器
- Go 1.24+ (仅用于从源码构建)

## 许可证

[MIT LICENSE](https://github.com/E9C50/vid/LICENSE)