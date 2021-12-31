package utils

import (
	"github.com/realotz/mstore/internal/conf"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

// 计算日期相差多少月
func SubMonth(t1, t2 time.Time) (month int) {
	y1 := t1.Year()
	y2 := t2.Year()
	m1 := int(t1.Month())
	m2 := int(t2.Month())
	d1 := t1.Day()
	d2 := t2.Day()

	yearInterval := y1 - y2
	// 如果 d1的 月-日 小于 d2的 月-日 那么 yearInterval-- 这样就得到了相差的年数
	if m1 < m2 || m1 == m2 && d1 < d2 {
		yearInterval--
	}
	// 获取月数差值
	monthInterval := (m1 + 12) - m2
	if d1 < d2 {
		monthInterval--
	}
	monthInterval %= 12
	month = yearInterval*12 + monthInterval
	return
}

func Timestamppb(now *time.Time) *timestamppb.Timestamp {
	if now == nil {
		return nil
	}
	return timestamppb.New(*now)
}

// 获取月初时间
func GetFirstOfMonth(now time.Time) time.Time {
	return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)
}

// 获取月初时间
func GetFirstOfQuarter(now time.Time) time.Time {
	if now.Month() >= 1 && now.Month() <= 3 {
		return time.Date(now.Year(), 1, 1, 0, 0, 0, 0, time.Local)
	}
	if now.Month() >= 4 && now.Month() <= 6 {
		return time.Date(now.Year(), 4, 1, 0, 0, 0, 0, time.Local)
	}
	if now.Month() >= 7 && now.Month() <= 9 {
		return time.Date(now.Year(), 7, 1, 0, 0, 0, 0, time.Local)
	}
	if now.Month() >= 10 && now.Month() <= 12 {
		return time.Date(now.Year(), 10, 1, 0, 0, 0, 0, time.Local)
	}
	return time.Date(now.Year(), 1, 1, 0, 0, 0, 0, time.Local)
}

// 获取月末时间
func GetLastOfMonth(now time.Time) time.Time {
	return time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, time.Local).Add(-1 * time.Second)
}

func GetFirstOfMonthForStr(month string) time.Time {
	t, _ := time.Parse(conf.MonthFormat, month)
	return GetFirstOfMonth(t)
}

func GetLastOfMonthForStr(month string) time.Time {
	t, _ := time.Parse(conf.MonthFormat, month)
	return GetLastOfMonth(t)
}

// 获取某个周一时间
func GetFirstOfWeek(now time.Time) time.Time {
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
}

// 获取本周周一时间
func GetFirstDateOfWeek() time.Time {
	return GetFirstOfWeek(time.Now())
}

//获取上周的周一日期
func GetLastWeekFirstDate() time.Time {
	return GetFirstDateOfWeek().AddDate(0, 0, -7)
}
