package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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

	// 创建一个临时文件来存储容器中的文件内容
	tempDir := os.TempDir()
	tempFile := filepath.Join(tempDir, filepath.Base(filePath))

	// 确保临时文件路径有效
	if filepath.Base(filePath) == "." || filepath.Base(filePath) == "/" {
		tempFile = filepath.Join(tempDir, "vid_temp_file")
	}

	log("Copying file from container: %s:%s -> %s\n", container, filePath, tempFile)

	// 从容器中复制文件到临时位置
	copyFromCmd := exec.Command("docker", "cp", container+":"+filePath, tempFile)
	copyFromCmd.Stdout = os.Stdout
	copyFromCmd.Stderr = os.Stderr
	if err := copyFromCmd.Run(); err != nil {
		fmt.Printf("Error copying file from container: %v\n", err)
		fmt.Println("Please check:")
		fmt.Println("1. Is the container running? Try 'docker ps' to verify.")
		fmt.Println("2. Does the file exist in the container?")
		fmt.Println("3. Do you have permission to access the container and file?")
		fmt.Println("4. Try manually running: docker cp " + container + ":" + filePath + " " + tempFile)
		os.Exit(1)
	}

	// 使用vim编辑临时文件
	log("Opening file in vim: %s\n", tempFile)
	vimCmd := exec.Command("vim", tempFile)
	vimCmd.Stdin = os.Stdin
	vimCmd.Stdout = os.Stdout
	vimCmd.Stderr = os.Stderr

	if err := vimCmd.Run(); err != nil {
		fmt.Printf("Error running vim: %v\n", err)
		// 清理临时文件
		os.Remove(tempFile)
		os.Exit(1)
	}

	// 将修改后的文件复制回容器
	log("Copying file back to container: %s -> %s:%s\n", tempFile, container, filePath)
	copyToCmd := exec.Command("docker", "cp", tempFile, container+":"+filePath)
	copyToCmd.Stdout = os.Stdout
	copyToCmd.Stderr = os.Stderr
	if err := copyToCmd.Run(); err != nil {
		fmt.Printf("Error copying file to container: %v\n", err)
		fmt.Println("Please check:")
		fmt.Println("1. Do you have permission to write to the container?")
		fmt.Println("2. Is the target directory writable?")
		fmt.Println("3. Try manually running: docker cp " + tempFile + " " + container + ":" + filePath)
		os.Remove(tempFile)
		os.Exit(1)
	}

	// 清理临时文件
	os.Remove(tempFile)

	fmt.Printf("Successfully edited %s in container %s\n", filePath, container)
}

// log prints formatted message only when verbose mode is enabled
func log(format string, args ...interface{}) {
	if *verbose {
		fmt.Printf(format, args...)
	}
}
