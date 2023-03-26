package tms

import "time"

func Format(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
func FormatYMD(t time.Time) string {
	return t.Format("2006-01-02")
}
func FormatByMill(t int64) string {
	return time.UnixMilli(t).Format("2006-01-02 15:04:05")
}

// ParseTime 经典到让人恶心的go的时间解析
func ParseTime(str string) int64 {
	parse, _ := time.Parse("2006-01-02 15:04", str)
	return parse.UnixMilli()
}
