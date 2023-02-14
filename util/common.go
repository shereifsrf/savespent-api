package util

import (
	"fmt"
	"strconv"
)

func GetInt64(value string) int64 {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		fmt.Printf("[GetInt64] strconv.Atoi(%s), err: %v", value, err)
	}

	return int64(intValue)
}