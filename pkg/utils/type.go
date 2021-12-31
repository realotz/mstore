package utils

type StringTypeMap []string

func (s StringTypeMap) CheckKey(v string) bool {
	for _, sv := range s {
		if sv == v {
			return true
		}
	}
	return false
}
