package logs

import (
	"fmt"
	"gin_boot/config"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// dailyRotateWriter 按天轮转的写入器
type dailyRotateWriter struct {
	dir         string
	filename    string
	maxAge      int
	maxBackups  int
	compress    bool
	localTime   bool
	timeFormat  string
	currentDate string
	currentFile *os.File
}

// NewDailyRotateWriter 创建按天轮转的写入器
func NewDailyRotateWriter(cfg config.FileConfig) io.WriteCloser {
	// 确保日志目录存在
	if err := os.MkdirAll(cfg.Dir, 0755); err != nil {
		panic(fmt.Sprintf("创建日志目录失败: %v", err))
	}

	writer := &dailyRotateWriter{
		dir:        cfg.Dir,
		filename:   cfg.Filename,
		maxAge:     cfg.MaxAge,
		maxBackups: cfg.MaxBackups,
		compress:   cfg.Compress,
		localTime:  cfg.LocalTime,
		timeFormat: cfg.TimeFormat,
	}

	return writer
}

// Write 写入日志
func (w *dailyRotateWriter) Write(p []byte) (n int, err error) {
	// 获取当前日期
	var now time.Time
	if w.localTime {
		now = time.Now()
	} else {
		now = time.Now().UTC()
	}

	currentDate := now.Format(w.timeFormat)

	// 检查是否需要轮转
	if w.currentDate != currentDate {
		w.rotate(currentDate)
	}

	// 写入当前文件
	if w.currentFile == nil {
		if err := w.openFile(currentDate); err != nil {
			return 0, err
		}
	}

	return w.currentFile.Write(p)
}

// Close 关闭写入器
func (w *dailyRotateWriter) Close() error {
	if w.currentFile != nil {
		return w.currentFile.Close()
	}
	return nil
}

// rotate 执行日志轮转
func (w *dailyRotateWriter) rotate(newDate string) {
	// 关闭当前文件
	if w.currentFile != nil {
		w.currentFile.Close()
		w.currentFile = nil
	}

	// 清理过期文件
	w.cleanup()
	w.currentDate = newDate
}

// openFile 打开日志文件
func (w *dailyRotateWriter) openFile(date string) error {
	filename := fmt.Sprintf("%s_%s.log", w.filename, date)
	filepath := filepath.Join(w.dir, filename)

	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	w.currentFile = file
	w.currentDate = date
	return nil
}

// cleanup 清理过期和超量的日志文件
func (w *dailyRotateWriter) cleanup() {
	files, err := w.getLogFiles()
	if err != nil {
		return
	}

	// 按时间排序（最新的在前面）
	sort.Sort(sort.Reverse(logFileSlice(files)))

	// 删除过期文件
	cutoff := time.Now().AddDate(0, 0, -w.maxAge)
	for _, file := range files {
		if file.ModTime.Before(cutoff) {
			os.Remove(filepath.Join(w.dir, file.Name))
		}
	}

	// 删除超量文件
	if w.maxBackups > 0 && len(files) > w.maxBackups {
		for _, file := range files[w.maxBackups:] {
			os.Remove(filepath.Join(w.dir, file.Name))
		}
	}

	// 压缩文件（如果启用）
	if w.compress {
		w.compressOldFiles()
	}
}

// logFile 日志文件信息
type logFile struct {
	Name    string
	ModTime time.Time
}

type logFileSlice []logFile

func (s logFileSlice) Len() int           { return len(s) }
func (s logFileSlice) Less(i, j int) bool { return s[i].ModTime.Before(s[j].ModTime) }
func (s logFileSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// getLogFiles 获取日志文件列表
func (w *dailyRotateWriter) getLogFiles() ([]logFile, error) {
	entries, err := os.ReadDir(w.dir)
	if err != nil {
		return nil, err
	}

	var files []logFile
	prefix := w.filename + "_"
	suffix := ".log"

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()
		if strings.HasPrefix(name, prefix) && strings.HasSuffix(name, suffix) {
			info, err := entry.Info()
			if err != nil {
				continue
			}

			files = append(files, logFile{
				Name:    name,
				ModTime: info.ModTime(),
			})
		}
	}

	return files, nil
}

// compressOldFiles 压缩旧文件
func (w *dailyRotateWriter) compressOldFiles() {
	// 这里可以实现压缩逻辑
	// 为了简单起见，这里不实现具体的压缩功能
	// 可以使用 gzip 包来实现
}

// LumberjackWriter 使用lumberjack实现按天轮转（简化版本）
func NewLumberjackWriter(cfg config.FileConfig) io.WriteCloser {
	// 生成当天的文件名
	date := time.Now().Format(cfg.TimeFormat)
	filename := fmt.Sprintf("%s_%s.log", cfg.Filename, date)
	filepath := filepath.Join(cfg.Dir, filename)
	return &lumberjack.Logger{
		Filename:   filepath,
		MaxSize:    cfg.MaxSize,
		MaxAge:     cfg.MaxAge,
		MaxBackups: cfg.MaxBackups,
		Compress:   cfg.Compress,
		LocalTime:  cfg.LocalTime,
	}
}
