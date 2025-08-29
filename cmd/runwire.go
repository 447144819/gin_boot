package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	// 获取项目根目录
	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Println("无法获取工作目录:", err)
		os.Exit(1)
	}

	// 1. 运行 go run .\cmd\wire_set.go
	wireSetCmd := exec.Command("go", "run", filepath.Join("cmd", "wire_set.go"))
	wireSetCmd.Stdout = os.Stdout
	wireSetCmd.Stderr = os.Stderr
	wireSetCmd.Dir = rootDir // 设置工作目录为项目根目录
	fmt.Println("生成wire: go run ./cmd/wire_set.go")
	if err := wireSetCmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "运行失败 wire_set.go: %v\n", err)
		os.Exit(1)
	}

	// 2. 运行 wire（在 internal/controller 目录下）
	controllerDir := filepath.Join(rootDir, "cmd", "wire")
	wireCmd := exec.Command("wire")
	//wireCmd := exec.Command("go", "run", filepath.Join("cmd", "wire_set.go"))
	wireCmd.Stdout = os.Stdout
	wireCmd.Stderr = os.Stderr
	wireCmd.Dir = controllerDir // 设置工作目录为 internal/controller
	fmt.Println("运行: wire")
	if err := wireCmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to run wire: %v\n", err)
		os.Exit(1)
	}

	// 3. 移动 wire_gen.go 到 ../internal/ioc/
	//srcFile := filepath.Join(controllerDir, "wire_gen.go")
	//dstDir := filepath.Join(rootDir, "internal", "ioc")
	//dstFile := filepath.Join(dstDir, "wire_gen.go")

	//// 确保目标目录存在
	//if err := os.MkdirAll(dstDir, 0755); err != nil {
	//	fmt.Fprintf(os.Stderr, "Failed to create directory %s: %v\n", dstDir, err)
	//	os.Exit(1)
	//}

	// 移动文件（Windows 和 Linux/macOS 兼容）
	//fmt.Printf("Moving %s to %s\n", srcFile, dstFile)
	//if err := os.Rename(srcFile, dstFile); err != nil {
	//	fmt.Fprintf(os.Stderr, "Failed to move wire_gen.go: %v\n", err)
	//	os.Exit(1)
	//}

	fmt.Println("生成成功")
}
