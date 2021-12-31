package utils

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func PBTimeString(timestamp *timestamppb.Timestamp) string {
	if timestamp == nil {
		return ""
	}
	return timestamp.AsTime().Local().Format("2006-01-02 15:04:05")
}

func ParseLocalTime(layout string, value string) (time.Time, error) {
	return time.ParseInLocation(layout, value, time.Local)
}
