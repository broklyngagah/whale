package utils

import (
	"time"
)

//秒数
const (
	MinuteSecsHalf = 30
	MinuteSecs     = MinuteSecsHalf * 2
	HourSecs       = 60 * MinuteSecs
	DaySecs        = 24 * HourSecs
)

// 将mysql默认的时间字符串转变成unix的timestamp(秒)
func TimeStrToUnix(s string) int64 {
	t, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		return 0
	}
	return t.Unix()
}

//返回距离午夜的秒数
func GetMidnightSeconds() int64 {
	now := time.Now()
	midnight := (23-now.Hour())*60*60 + (59-now.Minute())*60 + 59 - now.Second() + 1
	return int64(midnight)
}

func GetZeroTime() int64 {
	now := time.Now()
	zero := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return zero.Unix()
}

func GetHourTime() int64 {
	now := time.Now()
	zero := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, now.Location())
	return zero.Unix()
}

func IsThisHour(t int64) bool {
	hour := time.Unix(t, 0)
	now := time.Now()
	timeHour := time.Date(hour.Year(), hour.Month(), hour.Day(), hour.Hour(), 0, 0, 0, now.Location())
	nowHour := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, now.Location())
	return nowHour.Equal(timeHour)
}

func GetDailyMidnightRefreshTime() time.Time {
	now := time.Now()
	//func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

}

func GetDailyDawnRefreshTime() time.Time {
	now := time.Now()
	//func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
	return time.Date(now.Year(), now.Month(), now.Day(), 5, 0, 0, 0, now.Location())

}

func GetWeekMidnightRefreshTime() time.Time {
	refreshTime := time.Now()

	weekDay := refreshTime.Weekday()
	if weekDay == 0 {
		weekDay = 7
	}

	refreshTime = refreshTime.AddDate(0, 0, -int(weekDay-1))

	return time.Date(refreshTime.Year(), refreshTime.Month(), refreshTime.Day(), 0, 0, 0, 0, refreshTime.Location())
}

func GetWeekDawnRefreshTime() time.Time {
	refreshTime := time.Now()

	weekDay := refreshTime.Weekday()
	if weekDay == 0 {
		weekDay = 7
	}

	refreshTime = refreshTime.AddDate(0, 0, -int(weekDay-1))

	return time.Date(refreshTime.Year(), refreshTime.Month(), refreshTime.Day(), 5, 0, 0, 0, refreshTime.Location())
}

func GetMonthMidnightRefreshTime() time.Time {
	now := time.Now()

	return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
}

func GetMonthDawnRefreshTime() time.Time {
	now := time.Now()

	return time.Date(now.Year(), now.Month(), 1, 5, 0, 0, 0, now.Location())
}

func IsPassZeroTime() bool {
	now := time.Now()
	refreshTime := time.Unix(GetZeroTime(), 0)
	if now.UnixNano() <= refreshTime.UnixNano() {
		return false
	}
	return true
}

func IsPassRefreshTime() bool {
	now := time.Now()
	refreshTime := GetDailyDawnRefreshTime()
	if now.UnixNano() <= refreshTime.UnixNano() {
		return false
	}
	return true
}

func canRefresh(lastRefreshTime, targetRefreshTime time.Time) bool {

	now := time.Now()

	if now.UnixNano() <= targetRefreshTime.UnixNano() {
		return false
	}
	return lastRefreshTime.UnixNano() < targetRefreshTime.UnixNano()
}

func CanRefreshMidnightDaily(lastRefreshTime time.Time) bool {
	return canRefresh(lastRefreshTime, GetDailyMidnightRefreshTime())
}

func CanRefreshDawnDaily(lastRefreshTime time.Time) bool {
	return canRefresh(lastRefreshTime, GetDailyDawnRefreshTime())
}

func CanRefreshMidnightWeekly(lastRefreshTime time.Time) bool {
	return canRefresh(lastRefreshTime, GetWeekMidnightRefreshTime())
}

func CanRefreshDawnWeekly(lastRefreshTime time.Time) bool {
	return canRefresh(lastRefreshTime, GetWeekDawnRefreshTime())
}

func CanRefreshMidnightMonthly(lastRefreshTime time.Time) bool {
	return canRefresh(lastRefreshTime, GetMonthMidnightRefreshTime())
}

func CanRefreshDawnMonthly(lastRefreshTime time.Time) bool {
	return canRefresh(lastRefreshTime, GetMonthDawnRefreshTime())
}

func IsYeasterDay(t int64) bool {
	date := time.Unix(t, 0)
	now := time.Now()
	//func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
	timeDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, now.Location())
	yesterDay := now.AddDate(0, 0, -1)
	yearDayTime := time.Date(yesterDay.Year(), yesterDay.Month(), yesterDay.Day(), 0, 0, 0, 0, now.Location())
	return yearDayTime.Equal(timeDay)
}

func IsToday(t int64) bool {
	date := time.Unix(t, 0)
	now := time.Now()
	//func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
	timeDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, now.Location())
	todayDateTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return todayDateTime.Equal(timeDay)
}

func GetDaysPass(t int64) int {
	date := time.Unix(t, 0)
	now := time.Now()
	//func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
	timeDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, now.Location())
	passDay := time.Since(timeDay)
	return int(passDay.Hours() / 24)
}

func GetRealDaysPass(t int64) int {
	date := time.Unix(t, 0)
	now := time.Now()
	//func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
	timeDay := time.Date(date.Year(), date.Month(), date.Day(), date.Hour(), date.Minute(), date.Second(), date.Nanosecond(), now.Location())
	passDay := time.Since(timeDay)
	return int(passDay.Hours() / 24)
}

func GetHoursPass(t int64) int {
	date := time.Unix(t, 0)
	timeDay := time.Date(date.Year(), date.Month(), date.Day(), date.Hour(), date.Minute(), date.Second(), date.Nanosecond(), time.Local)
	passDay := time.Since(timeDay)
	return int(passDay.Hours())
}

func IsThisWeek(t int64) bool {
	date := time.Unix(t, 0)
	now := time.Now()

	dateWeekDay := date.Weekday()
	if dateWeekDay == 0 {
		dateWeekDay = 7
	}
	dateWeekDay = -(dateWeekDay - 1)

	nowWeekDay := now.Weekday()
	if nowWeekDay == 0 {
		nowWeekDay = 7
	}
	nowWeekDay = -(nowWeekDay - 1)

	timeDay := date.AddDate(0, 0, int(dateWeekDay))
	todayDateTime := now.AddDate(0, 0, int(nowWeekDay))

	return todayDateTime.YearDay() == timeDay.YearDay()
}

func IsThisMonth(t int64) bool {
	date := time.Unix(t, 0)
	now := time.Now()
	return date.Month() == now.Month()
}

// 除去秒以下的时间内
func TimeNow() time.Time {
	now := time.Now()
	return time.Unix(now.Unix(), 0)
}

func MakeHourBegin(begin time.Time) int64 {
	return begin.Unix() - int64(begin.Minute()*60) - int64(begin.Second())
}

func Time(n int64) time.Time {
	t := time.Unix(n, 0)
	return t
}

func TimePtr(n int64) *time.Time {
	t := Time(n)
	return &t
}

func SecondsOfOneDay() int64 {
	return 24 * 60 * 60
}

func MakeDayBegin(begin time.Time) int64 {
	return begin.Unix() - int64(begin.Hour()*60*60) - int64(begin.Minute()*60) - int64(begin.Second())
}

func ISOWeek(t time.Time) (week int) {
	yday := t.YearDay() - 1
	year, month, day := t.Date()
	weekDay := t.Weekday()
	wday := int(weekDay+6) % 7 // weekday but Monday = 0.
	const (
		Mon int = iota
		Tue
		Wed
		Thu
		Fri
		Sat
		Sun
	)

	// Calculate week as number of Mondays in year up to
	// and including today, plus 1 because the first week is week 0.
	// Putting the + 1 inside the numerator as a + 7 keeps the
	// numerator from being negative, which would cause it to
	// round incorrectly.
	week = (yday - wday + 7) / 7

	// The week number is now correct under the assumption
	// that the first Monday of the year is in week 1.
	// If Jan 1 is a Tuesday, Wednesday, or Thursday, the first Monday
	// is actually in week 2.
	jan1wday := (wday - yday + 7*53) % 7
	if Tue <= jan1wday && jan1wday <= Thu {
		week++
	}

	// If the week number is still 0, we're in early January but in
	// the last week of last year.
	if week == 0 {
		year--
		week = 52
		// A year has 53 weeks when Jan 1 or Dec 31 is a Thursday,
		// meaning Jan 1 of the next year is a Friday
		// or it was a leap year and Jan 1 of the next year is a Saturday.
		if jan1wday == Fri || (jan1wday == Sat && isLeap(year)) {
			week++
		}
	}

	// December 29 to 31 are in week 1 of next year if
	// they are after the last Thursday of the year and
	// December 31 is a Monday, Tuesday, or Wednesday.
	if month == time.December && day >= 29 && wday < Thu {
		if dec31wday := (wday + 31 - day) % 7; Mon <= dec31wday && dec31wday <= Wed {
			year++
			week = 1
		}
	}

	// 如果是周一早上5点之前，算上周
	if weekDay == time.Monday && t.Hour() >= 0 && t.Hour() < 5 {
		_, week = t.AddDate(0, 0, -1).ISOWeek()
	}

	return
}

// 以凌晨5点为基准，返回几天和明天的date
func DayBase5() (today int, tomorrow int) {
	t := time.Now()
	// 算昨天
	if t.Hour() >= 0 && t.Hour() < 5 {
		_, _, today = t.AddDate(0, 0, -1).Date()
		_, _, tomorrow = t.Date()
		return
	}

	_, _, today = t.Date()
	_, _, tomorrow = t.AddDate(0, 0, 1).Date()
	return
}

func isLeap(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func FormatDateForChat(ts int64) string {
	t := GetZeroTime()
	if ts >= t {
		return "今天"
	} else if ts > t-86400 {
		return "昨天"
	} else {
		ti := time.Unix(ts, 0)
		return ti.Format("01-02")
	}
}
func FormatTimeForChat(ts int64) string {
	return time.Unix(ts, 0).Format("15:04")
}

func FormatDateTime(t time.Time) string {
	return t.Format("2016-01-02 15:04")
}
