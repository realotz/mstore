package utils

import jsoniter "github.com/json-iterator/go"

func JsonBytes(v interface{}) []byte {
	bytes, _ := jsoniter.Marshal(v)
	return bytes
}

func JsonString(v interface{}) string {
	str, _ := jsoniter.MarshalToString(v)
	return str
}
