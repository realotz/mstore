package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"
)

var defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var intLetters = []rune("0123456789")

func UnitSplit(s, sep string) []uint32 {
	var res []uint32
	resStr := strings.Split(s, sep)
	for _, v := range resStr {
		if v == "" {
			continue
		}
		i, _ := strconv.Atoi(v)
		res = append(res, uint32(i))
	}
	return res
}

func UnitJoin(elems []uint32, sep string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return fmt.Sprint(elems[0])
	}
	n := len(sep) * (len(elems) - 1)
	for i := 0; i < len(elems); i++ {
		n += len(fmt.Sprint(elems[i]))
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(fmt.Sprint(elems[0]))
	for _, s := range elems[1:] {
		b.WriteString(sep)
		b.WriteString(fmt.Sprint(s))
	}
	return b.String()
}

func RandomIntStr(n int) string {
	b := make([]rune, n)
	for i := range b {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(intLetters))))
		b[i] = intLetters[n.Int64()]
	}
	return string(b)
}

func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(defaultLetters))))
		b[i] = defaultLetters[n.Int64()]
	}
	return string(b)
}

func UnsafeStringToInt64(str string) int64 {
	i64, _ := strconv.ParseInt(str, 10, 64)
	return i64
}

func UnsafeStringToUInt32(str string) uint32 {
	return uint32(UnsafeStringToInt64(str))
}

func GetEndMonth(end time.Time) uint32 {
	nowTime := time.Now()
	var buyNum = end.Month() - nowTime.Month()
	if nowTime.AddDate(0, int(buyNum), 0).Unix() < end.Unix() {
		buyNum += 1
	}
	return uint32(buyNum)
}

func GetMonthDiff(start, end time.Time) uint32 {
	var buyNum = end.Month() - start.Month()
	return uint32(buyNum)
}

// 下划线字符转小写驼峰
func CamelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	data[0] += 32
	return string(data[:])
}
