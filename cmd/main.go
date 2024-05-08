package main

import (
	"github.com/sharework-cn/idp/pkg/cmd"
	"github.com/sharework-cn/idp/pkg/cmd/idp"
)

// 主程序入口
func main() {
	// 执行RootCommand，并获取执行过程中的错误
	err := idp.Execute()
	cmd.CheckError(err) // 检查并处理错误
}
