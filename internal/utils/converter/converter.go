package converter

// 类型转换
import "strconv"

// ======================
// String => 数值类型
// ======================

// StringToInt 字符串转 int
func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

// StringToInt64 字符串转 int64
func StringToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

// StringToUint 字符串转 uint
func StringToUint(s string) (uint, error) {
	i, err := strconv.ParseUint(s, 10, 0)
	return uint(i), err
}

// StringToUint64 字符串转 uint64
func StringToUint64(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 64)
}

// ======================
// 数值类型 => String
// ======================

// IntToString int 转字符串
func IntToString(i int) string {
	return strconv.Itoa(i)
}

// Int64ToString int64 转字符串
func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

// UintToString uint 转字符串
func UintToString(u uint) string {
	return strconv.FormatUint(uint64(u), 10)
}

// Uint64ToString uint64 转字符串
func Uint64ToString(u uint64) string {
	return strconv.FormatUint(u, 10)
}

// ======================
// Must 系列（不处理错误，出错时 panic，慎用！）
// ======================

// MustStringToInt 字符串转 int，出错 panic
func MustStringToInt(s string) int {
	i, err := StringToInt(s)
	if err != nil {
		panic(err)
	}
	return i
}

// MustStringToInt64 字符串转 int64，出错 panic
func MustStringToInt64(s string) int64 {
	i, err := StringToInt64(s)
	if err != nil {
		panic(err)
	}
	return i
}

// MustStringToUint 字符串转 uint，出错 panic
func MustStringToUint(s string) uint {
	u, err := StringToUint(s)
	if err != nil {
		panic(err)
	}
	return u
}

// MustStringToUint64 字符串转 uint64，出错 panic
func MustStringToUint64(s string) uint64 {
	u, err := StringToUint64(s)
	if err != nil {
		panic(err)
	}
	return u
}
