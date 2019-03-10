package main

import (
	"os"
	"fmt"
)

func main() { // main方法不能传入参数，但是os.Args可以获取命令行参数
	if (len(os.Args) > 1) {
		fmt.Println(os.Args[1]) 
	}
	fmt.Println(os.Args)
	fmt.Println("hello world")
	os.Exit(10) // 返回程序退出的状态
}

// 应用程序入口
// 必须是main包：package main
// 必须是main方法：func main(){}
// 文件名不一定是main.go