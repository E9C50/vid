package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var verbose = flag.Bool("v", false, "Enable verbose logging")

func main() {
	flag.Parse()

	if len(flag.Args()) < 2 {
		fmt.Println("Usage: vid [-v] <container> <file_path>")
		fmt.Println("Example: vid my_container /etc/nginx/nginx.conf")
		fmt.Println("  -v  Enable verbose logging")
		os.Exit(1)
	}

	container := flag.Args()[0]
	filePath := flag.Args()[1]

	// 验证容器是否存在和运行
	log("Checking container status: %s\n", container)
	checkContainerCmd := exec.Command("docker", "inspect", container)
	if err := checkContainerCmd.Run(); err != nil {
		fmt.Printf("Error: Container '%s' does not exist or is not accessible\n", container)
		os.Exit(1)
	}

	// 创建临时文件
	tempFile, err := os.CreateTemp("", "vid-*")
	if err != nil {
		fmt.Printf("Error creating temporary file: %v\n", err)
		os.Exit(1)
	}
	tempFileName := tempFile.Name()
	_ = tempFile.Close() // 立即关闭文件句柄，以便 'docker cp' 能够写入该路径

	// 确保程序运行结束后删除临时文件
	defer os.Remove(tempFileName)

	log("Copying file from container: %s:%s -> %s\n", container, filePath, tempFileName)

	// 从容器中复制文件到临时位置
	copyFromCmd := exec.Command("docker", "cp", container+":"+filePath, tempFileName)
	copyFromCmd.Stdout = os.Stdout
	copyFromCmd.Stderr = os.Stderr
	if err := copyFromCmd.Run(); err != nil {
		fmt.Printf("Error copying file from container: %v\n", err)
		fmt.Println("Please check:")
		fmt.Println("1. Is the container running? Try 'docker ps' to verify.")
		fmt.Println("2. Does the file exist in the container?")
		fmt.Println("3. Do you have permission to access the container and file?")
		fmt.Println("4. Try manually running: docker cp " + container + ":" + filePath + " " + tempFileName)
		os.Exit(1)
	}

	// 使用vim编辑临时文件
	log("Opening file in vim: %s\n", tempFileName)
	vimCmd := exec.Command("vim", tempFileName)
	vimCmd.Stdin = os.Stdin
	vimCmd.Stdout = os.Stdout
	vimCmd.Stderr = os.Stderr

	if err := vimCmd.Run(); err != nil {
		fmt.Printf("Error running vim: %v\n", err)
		os.Exit(1)
	}

	// 将修改后的文件复制回容器
	log("Copying file back to container: %s -> %s:%s\n", tempFileName, container, filePath)
	copyToCmd := exec.Command("docker", "cp", tempFileName, container+":"+filePath)
	copyToCmd.Stdout = os.Stdout
	copyToCmd.Stderr = os.Stderr
	if err := copyToCmd.Run(); err != nil {
		fmt.Printf("Error copying file to container: %v\n", err)
		fmt.Println("Please check:")
		fmt.Println("1. Do you have permission to write to the container?")
		fmt.Println("2. Is the target directory writable?")
		fmt.Println("3. Try manually running: docker cp " + tempFileName + " " + container + ":" + filePath)
		os.Exit(1)
	}

	fmt.Printf("Successfully edited %s in container %s\n", filePath, container)
}

// log prints formatted message only when verbose mode is enabled
func log(format string, args ...interface{}) {
	if *verbose {
		fmt.Printf(format, args...)
	}
}
