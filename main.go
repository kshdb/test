package test // import "git.qnsoft.cn/sxh/test"

import (
	"fmt"
	"time"
)

func Gtest1() string {
	fmt.Println("hello a11")
	return time.Now().Format("2006-01-02 15:04:05")
}
