package helper

import "time"

func Timestamp2S(timeUnix int64) string {
	layout := "2006-01-02 15:04:05"
	timeStr := time.Unix(timeUnix, 0).Format(layout)
	return timeStr
}

func GetTodayStartTimeStamp() int64 {
	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	now := time.Now()
	timeStr := now.Format("2006-01-02")
	res, _ := time.ParseInLocation("2006-01-02", timeStr, cstSh)
	return res.Unix()
}

func GetTodayEndTimeStamp() int64 {
	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	now := time.Now()
	timeStr := now.Format("2006-01-02")
	res, _ := time.ParseInLocation("2006-01-02", timeStr, cstSh)
	res = res.AddDate(0, 0, 1)
	res = res.Add(-1)
	return res.Unix()
}

func GetNDaysBefore(n int) int64 {
	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	now := time.Now()
	timeStr := now.Format("2006-01-02")
	res, _ := time.ParseInLocation("2006-01-02", timeStr, cstSh)
	res = res.AddDate(0, 0, -n)
	return res.Unix()
}
