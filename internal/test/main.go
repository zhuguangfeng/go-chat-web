package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a uint = 1
	fmt.Println(aa(a))
}

func aa(val any) bool {
	// 获取 val 的反射值
	valValue := reflect.ValueOf(val)
	return val == nil || valValue.Kind() == reflect.Ptr && valValue.IsNil() || valValue.IsZero()
}
