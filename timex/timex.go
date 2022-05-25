package timex

import (
	"encoding/binary"
	"fmt"
	"time"
)

type T struct {
	t int64
}

func (t *T) Int64() int64 {
	return t.t
}

func (t *T) Int() int {
	return int(t.t)
}

func (t *T) String() string {
	return fmt.Sprint(t.t)
}

func (t *T) Bytes() []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(t.t))
	return buf
}

func (t *T) Float32() float32 {
	return float32(t.t)
}

func (t *T) Float64() float64 {
	return float64(t.t)
}

// CNs Current NanoSecond
func CNs() *T {
	return &T{time.Now().UnixNano()}
}

// CMs Current MicroSecond
func CMs() *T {
	return &T{time.Now().UnixNano() / 1e6}
}

// CUs Current MilliSecond
func CUs() *T {
	return &T{time.Now().UnixNano() / 1e3}
}

// Cs Current Second
func Cs() *T {
	return &T{time.Now().Unix()}
}

// After Return the Time From Now After Duration
func After(duration time.Duration) *T {
	return &T{time.Now().Add(duration).Unix()}
}

// CFormat Current time format
func CFormat(layout string) string {
	return time.Now().Format(layout)
}

// Format Any time format
func Format(timestamp int64, layout string) string {
	return time.Unix(timestamp, 0).Format(layout)
}

// Time Any timestamp to time
func Time(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

// AtToday Check timestamp is in today time range
func AtToday(timestamp int64) bool {
	t := time.Now()
	today := []int64{
		time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix(),
		time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location()).Unix(),
	}
	ts := time.Unix(timestamp, 0).Unix()
	return ts >= today[0] && ts <= today[1]
}

// Parse DataTime parse with layout to unix
// set compatible true can make func try some preset layouts
var presetLayouts = []string{
	"2006-01-02 15:04:05",
	"2006-01-2 15:04:05",
	"2006-1-2 15:04:05",
	"2006-1-2 15:04",
	"2006/1/2 15:04",
	"2006/1/2 15:04:05",
	"2006/1/02 15:04:05",
	"2006/01/02 15:04:05",
}

func Parse(dt, layout string, compatible ...bool) *T {
	if dt == "" {
		return &T{0}
	}
	t, err := time.ParseInLocation(layout, dt, time.Local)
	if err != nil {
		if len(compatible) != 0 && compatible[0] {
			for _, presetLayout := range presetLayouts {
				if t, e := time.ParseInLocation(presetLayout, dt, time.Local); e == nil {
					return &T{t.Unix()}
				}
			}
		}
		fmt.Println(err)
		return &T{0}
	}
	return &T{t.Unix()}
}

var WeekZhou = map[string]string{
	"Monday":    "周一",
	"Tuesday":   "周二",
	"Wednesday": "周三",
	"Thursday":  "周四",
	"Friday":    "周五",
	"Saturday":  "周六",
	"Sunday":    "周日",
}
var WeekXingQi = map[string]string{
	"Monday":    "星期一",
	"Tuesday":   "星期二",
	"Wednesday": "星期三",
	"Thursday":  "星期四",
	"Friday":    "星期五",
	"Saturday":  "星期六",
	"Sunday":    "星期日",
}

//Week Any timestamp to chinese week
func Week(timestamp int64, weekMap map[string]string) string {
	return weekMap[time.Unix(timestamp, 0).Weekday().String()]
}

//Zero Any(Current) timestamp's zero time at that(to-) day
func Zero(timestamp ...int64) int64 {
	var t time.Time
	if len(timestamp) == 0 {
		t = time.Now()
	} else {
		t = time.Unix(timestamp[0], 0)
	}
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix()
}

//IsFormattedDate Determine the date is input format
func IsFormattedDate(date, format string) bool {
	_, err := time.Parse(date, format)
	return err == nil
}
