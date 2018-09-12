package ints

import (
	"strconv"
	"strings"
)

func Join(ints []interface{}, sep string) string {
	var strs []string
	for _, i := range ints {
		switch a := i.(type) {
		case int:
			strs = append(strs, strconv.Itoa(a))
		case int8:
			strs = append(strs, strconv.Itoa(int(a)))
		case int16:
			strs = append(strs, strconv.Itoa(int(a)))
		case int32:
			strs = append(strs, strconv.Itoa(int(a)))
		case int64:
			strs = append(strs, strconv.Itoa(int(a)))
		case uint:
			strs = append(strs, strconv.Itoa(int(a)))
		case uint8:
			strs = append(strs, strconv.Itoa(int(a)))
		case uint16:
			strs = append(strs, strconv.Itoa(int(a)))
		case uint32:
			strs = append(strs, strconv.Itoa(int(a)))
		case uint64:
			strs = append(strs, strconv.Itoa(int(a)))
		}
	}
	return strings.Join(strs, sep)
}
