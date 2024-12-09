package main

import (
	"errors"
	"fmt"
	"github.com/zhuguangfeng/go-chat/internal/common"
	"github.com/zhuguangfeng/go-chat/pkg/errorx"
)

type A struct {
	Name string
	Age  int
}

type B struct {
	Name string
	Age  int
}

func main() {
	err1 := errorx.NewBizError(common.SystemInternalError).WithError(errors.New("err1"))
	err2 := errorx.NewBizError(common.SystemInternalError)

	fmt.Println(err1)
	fmt.Println(err2)
	fmt.Println(err1.Error())
	fmt.Println(err2.Error())
	fmt.Println(err1 == nil)
	fmt.Println(err2 == nil)

}

func demo(b []*B) {

}
