package common

import (
	"fmt"
	"strings"
	"time"
)

type timeTool struct{}

var TimeTool = timeTool{}

// 将中国时区文本转成时间戳
func (cls *timeTool) Str2Time(timeStr string) (int64, error) {
	layout := "2006-01-02 15:04:05"
	//获取中国时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return 0, err
	}
	t, err := time.ParseInLocation(layout, timeStr, loc)

	if err != nil {
		return 0, err
	}

	// 获取时间戳
	timestamp := t.Unix()
	return timestamp, nil
}

// 将时间戳转成中国时区文本
//
// layout 2006-01-02 15:04:05
func (cls *timeTool) Time2Str(timestamp int64, layout string) (string, error) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return "", err
	}
	t := time.Unix(timestamp, 0).In(loc)
	return t.Format(layout), nil
}

// 将时间戳分割成 Y m d H i s
//
// []string{"2006","01","02","15","04","05"}
func (cls *timeTool) YmdHis(timestamp int64) ([]string, error) {
	dateStr, err := cls.Time2Str(timestamp, "2006-01-02-15-04-05")
	if err != nil {
		return nil, err
	}
	result := strings.Split(dateStr, "-")
	return result, nil
}

// 获取时间戳
func (cls *timeTool) Time() (int64, error) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return 0, err
	}
	// 获取中国时区的当前时间
	now := time.Now().In(loc)
	// 获取秒级别时间戳
	// now.Unix()
	// 获取毫秒级别时间戳
	// now.UnixMilli()
	// 获取纳秒级别时间戳
	// now.UnixNano()
	return now.Unix(), nil
}

// 当前时间转文本
//
// layout 2006-01-02 15:04:05
func (cls *timeTool) Date(layout string) (string, error) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return "", err
	}
	now := time.Now().In(loc)
	formatted := now.Format(layout)
	return formatted, nil
}

type TimeLog struct {
	start   time.Time
	refresh time.Time
	logMap  map[string]time.Time
}

func NewTimeLog() *TimeLog {
	start := time.Now()
	logMap := map[string]time.Time{}
	return &TimeLog{
		start:   start,
		refresh: start,
		logMap:  logMap,
	}
}

func (cls *TimeLog) Start(key string) {
	cls.logMap[key] = time.Now()
}

func (cls *TimeLog) End(key string) {
	start, ok := cls.logMap[key]
	if !ok {
		start = cls.start
	}
	cls.log(key, start)
}

func (cls *TimeLog) Log(tag string) {
	cls.log(tag, cls.start)
}

func (cls *TimeLog) Refresh() {
	cls.refresh = time.Now()
}

func (cls *TimeLog) LogRefresh(tag string) {
	cls.log(tag, cls.refresh)
	cls.Refresh()
}

func (cls *TimeLog) log(tag string, start time.Time) {
	elapsed := time.Since(start)
	fmt.Println(tag, elapsed)
}
