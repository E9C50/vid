# vid - Docker 容器文件编辑工具

[English](./README_en.md) | 简体中文

<p align="center">
  <strong>厌倦了 Docker 中没有 vim？</strong><br>
  <strong>厌倦了手动 docker cp 来回复制配置文件？</strong><br>
  <strong>vid 让一切变得简单！</strong>
</p>

---

## 💡 项目初衷

在使用 Docker 容器时，我们经常遇到这些烦恼：

- 🔧 **容器里没有 vim**：许多精简的容器镜像为了减小体积，不包含文本编辑器
- 📝 **修改配置文件很麻烦**：需要手动执行 `docker cp` 将文件复制出来，编辑后再复制回去
- ⏱️ **重复劳动**：每次修改都要经历 复制→编辑→复制 的三步流程

**vid** 就是为了解决这些问题而生！它让你可以使用本地的 vim 编辑器直接编辑容器中的文件，**一条命令搞定一切**。

## ✨ 特性

- 🚀 **一键编辑**：一条命令直接编辑容器内文件，无需手动复制
- 💻 **使用本地编辑器**：利用你熟悉的本地 vim 配置和插件
- 🔄 **自动化流程**：自动处理文件复制、编辑、回写的完整流程
- 🧹 **自动清理**：编辑完成后自动清理临时文件
- 📊 **详细日志**：支持 `-v` 选项查看详细操作日志
- ⚡ **轻量快速**：单个二进制文件，无额外依赖

## 📦 安装

### 方式 1：下载预编译二进制（推荐）

从 [Releases](https://github.com/E9C50/vid/releases) 页面下载适用于您平台的预编译二进制文件。

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

从 [Releases](https://github.com/E9C50/vid/releases) 页面下载 `vid-windows-amd64.zip` 文件，解压后将 `vid.exe` 添加到系统 PATH 环境变量中。

### 方式 2：从源码构建

需要 Go 1.24 或更高版本：

```bash
# 克隆仓库
git clone https://github.com/E9C50/vid.git
cd vid

# 构建
go build -o vid main.go

# 安装到系统路径（可选）
sudo mv vid /usr/local/bin/
```

### 方式 3：使用 go install

```bash
go install github.com/E9C50/vid@latest
```

## 🚀 使用方法

### 基本用法

```bash
vid <容器名称或ID> <容器内文件路径>
```

### 选项

- `-v` - 启用详细日志输出，查看完整操作流程

### 使用示例

```bash
# 编辑 nginx 配置文件
vid my_nginx_container /etc/nginx/nginx.conf

# 编辑应用配置文件
vid web_app /app/config/app.conf

# 使用详细模式查看操作过程
vid -v my_container /etc/config.json

# 使用容器 ID
vid a1b2c3d4e5f6 /var/log/app.log
```

### 对比传统方式

**传统方式**（需要 3 个命令）：
```bash
docker cp my_container:/etc/nginx/nginx.conf ./nginx.conf
vim ./nginx.conf
docker cp ./nginx.conf my_container:/etc/nginx/nginx.conf
rm ./nginx.conf
```

**使用 vid**（只需 1 个命令）：
```bash
vid my_container /etc/nginx/nginx.conf
```

## 🔧 工作原理

```
┌─────────────────────────────────────────────────────────────┐
│  1. 从容器复制文件到本地临时文件 (docker cp)                 │
│                          ↓                                   │
│  2. 使用本地 vim 编辑器打开临时文件                          │
│                          ↓                                   │
│  3. 保存退出后自动将文件复制回容器 (docker cp)                │
│                          ↓                                   │
│  4. 自动清理本地临时文件                                      │
└─────────────────────────────────────────────────────────────┘
```

## 📋 前置要求

- Docker（已安装并可运行 `docker` 命令）
- vim 编辑器（已安装在本地系统）
- 对目标容器的访问权限

## ⚠️ 注意事项

- 确保 Docker 容器正在运行
- 确保目标文件路径在容器中存在
- 确保对容器和文件有相应的读写权限
- 使用 `-v` 选项可以查看详细的错误信息

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

## 📄 许可证

本项目采用 [MIT License](LICENSE) 开源协议。

## 🌟 Star History

如果这个项目对你有帮助，欢迎给个 Star ⭐

---

<p align="center">
  使用 ❤️ 和 Go 构建
</p>
