package utils

import (
	"time"
)

var layout string = "2006-01-02 15:04:05"

func GetTimeStampMillSecond() int64 {
	return time.Now().UnixMilli()
}

//	func GetTimeStampMills() int64 {
//		return time.Now().UnixMilli()
//	}
//
//	func FormatNowTime() string {
//		return time.Now().Format("20060102150405")
//	}
//
//	func FormatNowTime() string {
//		return time.Now().Format(format)
//	}
//
//	func FormatTime(format string, timestamp int64) string {
//		return time.Unix(timestamp, 0).Format(format)
//	}
func FormatTimeString(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006-01-02 15:04")
}

func FormatToTimeStamp(timeStr string) (int64, error) {
	// 按照指定格式解析时间字符串
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		return 0, err
	}
	return t.UnixMilli(), nil
}
