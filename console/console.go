// Package console 命令行辅助方法
package console

import (
	"fmt"
	"os"
)

// Success 打印一条成功消息，绿色输出
func Success(format string, v ...interface{}) {
	fmt.Println(Green(fmt.Sprintf(format, v...)))
}

// Danger 打印一条报错消息，红色输出
func Danger(format string, v ...interface{}) {
	fmt.Println(Red(fmt.Sprintf(format, v...)))
}

// Warning 打印一条提示消息，黄色输出
func Warning(format string, v ...interface{}) {
	fmt.Println(Yellow(fmt.Sprintf(format, v...)))
}

// Info 打印一条普通消息，蓝色输出
func Info(format string, v ...interface{}) {
	fmt.Println(Blue(fmt.Sprintf(format, v...)))
}

// Exit 打印一条报错消息，并退出 os.Exit(1)
func Exit(format string, v ...interface{}) {
	Danger(format, v...)
	os.Exit(1)
}

// ExitIf 语法糖，自带 err != nil 判断
func ExitIf(err error) {
	if err != nil {
		Exit(err.Error())
	}
}
