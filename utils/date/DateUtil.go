package date

import (
	"encoding/binary"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	cst *time.Location
)

// CSTLayout China Standard Time Layout
const CSTLayout = "2006-01-02 15:04:05"

func ParseDuration(d string) (time.Duration, error) {
	d = strings.TrimSpace(d)
	dr, err := time.ParseDuration(d)
	if err == nil {
		return dr, nil
	}
	if strings.Contains(d, "d") {
		index := strings.Index(d, "d")

		hour, _ := strconv.Atoi(d[:index])
		dr = time.Hour * 24 * time.Duration(hour)
		ndr, err := time.ParseDuration(d[index+1:])
		if err != nil {
			return dr, nil
		}
		return dr + ndr, nil
	}

	dv, err := strconv.ParseInt(d, 10, 64)
	return time.Duration(dv), err
}

func init() {
	var err error
	if cst, err = time.LoadLocation("Asia/Shanghai"); err != nil {
		panic(err)
	}

	// 默认设置为中国时区
	time.Local = cst
}

// RFC3339ToCSTLayout convert rfc3339 value to china standard time layout
// 2020-11-08T08:18:46+08:00 => 2020-11-08 08:18:46
func RFC3339ToCSTLayout(value string) (string, error) {
	ts, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return "", err
	}

	return ts.In(cst).Format(CSTLayout), nil
}

// CSTLayoutString 格式化时间
// 返回 "2006-01-02 15:04:05" 格式的时间
func CSTLayoutString() string {
	ts := time.Now()
	return ts.In(cst).Format(CSTLayout)
}

// ParseCSTInLocation 格式化时间
func ParseCSTInLocation(date string) (time.Time, error) {
	return time.ParseInLocation(CSTLayout, date, cst)
}

// CSTLayoutStringToUnix 返回 unix 时间戳
// 2020-01-24 21:11:11 => 1579871471
func CSTLayoutStringToUnix(cstLayoutString string) (int64, error) {
	stamp, err := time.ParseInLocation(CSTLayout, cstLayoutString, cst)
	if err != nil {
		return 0, err
	}
	return stamp.Unix(), nil

}

func CSTLayoutStringToUnixByte(cstLayoutString string) ([]byte, error) {
	stamp, err := time.ParseInLocation(CSTLayout, cstLayoutString, cst)
	if err != nil {
		return Int64ToBytes(0), err
	}
	return Int64ToBytes(stamp.Unix()), nil

}

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

// GMTLayoutString 格式化时间
// 返回 "Mon, 02 Jan 2006 15:04:05 GMT" 格式的时间
func GMTLayoutString() string {
	return time.Now().In(cst).Format(http.TimeFormat)
}

// ParseGMTInLocation 格式化时间
func ParseGMTInLocation(date string) (time.Time, error) {
	return time.ParseInLocation(http.TimeFormat, date, cst)
}

// SubInLocation 计算时间差
func SubInLocation(ts time.Time) float64 {
	return math.Abs(time.Now().In(cst).Sub(ts).Seconds())
}
