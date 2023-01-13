package test // import "github.com/kshdb/test"

import (
	"fmt"
	"time"
)

func Gtest1() string {
	fmt.Println("hello")
	return time.Now().Format("2006-01-02 15:04:05")
}
