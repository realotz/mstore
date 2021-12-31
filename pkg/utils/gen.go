package utils

import (
	"crypto/md5"
	"fmt"
	"math"
	"sync/atomic"
	"time"
)

var objectIdCounter uint32 = 0

func GenIdStringAsDateFormat() string {
	i := atomic.AddUint32(&objectIdCounter, 1)
	atomic.CompareAndSwapUint32(&objectIdCounter, math.MaxInt16, 1)
	date := time.Now().Format("20060102150405")
	id := fmt.Sprintf("%v%v", date, i)
	return id
}

func MD5String(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}
