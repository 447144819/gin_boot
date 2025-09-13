package times

import "time"

const (
	// DefaultLayout 是默认的日期时间格式，如 "2024-06-01 14:30:00"
	DefaultLayout = "2006-01-02 15:04:05"

	// LayoutDate 只有日期部分，如 "2024-06-01"
	LayoutDate = "2006-01-02"

	// LayoutTime 只有时间部分，如 "14:30:05"
	LayoutTime = "15:04:05"
)

// GetNowTimestamp 获取当前时间的 Unix 时间戳（秒级）
func GetNowTimestamp() int64 {
	return time.Now().Unix()
}

// GetNowTimestampMilli 获取当前时间的 Unix 时间戳（毫秒级）
func GetNowTimestampMilli() int64 {
	return time.Now().UnixMilli()
}

// GetNowTimestampMicro 获取当前时间的 Unix 时间戳（微秒级）
func GetNowTimestampMicro() int64 {
	return time.Now().UnixMicro()
}

// GetNowTimestampNano 获取当前时间的 Unix 时间戳（纳秒级）
func GetNowTimestampNano() int64 {
	return time.Now().UnixNano()
}

// TimestampToDateTime 将 Unix 时间戳（秒级）转为格式化的日期时间字符串
func TimestampToDateTime(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(DefaultLayout)
}

// TimestampMilliToDateTime 将毫秒级时间戳转为格式化的日期时间字符串
func TimestampMilliToDateTime(timestamp int64) string {
	return time.UnixMilli(timestamp).Format(DefaultLayout)
}

// DateTimeToTimestamp 将日期时间字符串转为 Unix 时间戳（秒级）
// 输入格式应为： "2006-01-02 15:04:05"
func DateTimeToTimestamp(dateTimeStr string) (int64, error) {
	t, err := time.Parse(DefaultLayout, dateTimeStr)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

// DateTimeToTimestampMilli 将日期时间字符串转为毫秒级时间戳
func DateTimeToTimestampMilli(dateTimeStr string) (int64, error) {
	t, err := time.Parse(DefaultLayout, dateTimeStr)
	if err != nil {
		return 0, err
	}
	return t.UnixMilli(), nil
}

// FormatTime 格式化 time.Time 为指定布局的字符串
func FormatTime(t time.Time, layout string) string {
	return t.Format(layout)
}

// ParseTime 将字符串按指定布局解析为 time.Time
func ParseTime(timeStr, layout string) (time.Time, error) {
	return time.Parse(layout, timeStr)
}
