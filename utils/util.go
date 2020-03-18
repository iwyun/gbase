package utils

import (
	"unsafe"
)

func StrToBytes(str string) []byte {

	return *(*[]byte)(unsafe.Pointer(&str))
}

func BytesToStr(bytes []byte) string {

	return *(*string)(unsafe.Pointer(&bytes))
}

//remove duplicate
func ArrayDuplicatedFilter(list []string) (data []string) {
	filter := make(map[string]bool)

	for _, v := range list {
		_, ok := filter[v]
		if ok {
			continue
		} else {
			filter[v] = true
			data = append(data, v)
		}
	}
	return
}
