package gfile

import (
	"io"
	"os"
	"path/filepath"
	"time"
)

// ======================
// 文件是否存在
// ======================

// FileExists 判断文件是否存在
func FileExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

// DirExists 判断目录是否存在
func DirExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// ======================
// 创建文件夹（目录）
// ======================

// MkdirAll 创建目录（包括多级目录），类似 mkdir -p
func MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

// ======================
// 读写文件
// ======================

// ReadFile 读取文件内容为 []byte
func ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

// WriteFile 写入 []byte 到文件（覆盖）
func WriteFile(path string, data []byte, perm os.FileMode) error {
	return os.WriteFile(path, data, perm)
}

// AppendToFile 追加内容到文件
func AppendToFile(path string, data []byte) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	return err
}

// ======================
// 删除文件 / 文件夹
// ======================

// RemoveFile 删除文件
func RemoveFile(path string) error {
	return os.Remove(path)
}

// RemoveDir 删除空目录
func RemoveDir(path string) error {
	return os.Remove(path) // 只能删空目录
}

// RemoveAll 删除目录及所有内容（慎用！）
func RemoveAll(path string) error {
	return os.RemoveAll(path)
}

// ======================
// 文件信息
// ======================

// GetFileSize 获取文件大小（字节）
func GetFileSize(path string) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}

// GetFileModTime 获取文件修改时间
func GetFileModTime(path string) (time.Time, error) {
	info, err := os.Stat(path)
	if err != nil {
		return time.Time{}, err
	}
	return info.ModTime(), nil
}

// ======================
// 文件操作（高级）
// ======================

// CopyFile 复制文件（不是文件夹）
func CopyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}

// WalkDir 遍历目录（类似 filepath.Walk，但使用 fs.FS 更现代，这里用 os）
func WalkDir(root string, fn func(path string, isDir bool) error) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		return fn(path, info.IsDir())
	})
}
