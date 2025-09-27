# vid - Docker容器文件编辑工具

vid是一个简单的命令行工具，允许您使用本地vim编辑器直接编辑Docker容器中的文件，而无需手动执行`docker cp`命令。

## 功能特点

- 直接在Docker容器中编辑文件
- 使用熟悉的vim编辑器
- 自动处理文件复制过程
- 静默模式运行，可选详细日志输出
- 错误处理和调试信息提示

## 安装

```bash
# 克隆项目
git clone <repository-url>

# 进入项目目录
cd vid

# 编译
go build -o vid main.go

# 将vid添加到PATH中（可选）
sudo mv vid /usr/local/bin/
```

## 使用方法

基本用法：
```bash
vid <container> <file_path>
```

例如：
```bash
vid my_container /etc/nginx/nginx.conf
```

启用详细日志输出：
```bash
vid -v my_container /etc/nginx/nginx.conf
```

## 工作原理

1. 使用`docker cp`从指定容器复制目标文件到本地临时文件
2. 使用本地vim编辑器打开临时文件
3. 保存更改后，使用`docker cp`将修改后的文件复制回容器
4. 清理临时文件

## 依赖

- Docker
- vim编辑器
- Go 1.24+（仅编译时需要）

## 许可证

MIT