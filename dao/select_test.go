package dao

import (
	"QQbot/tools/dao_tool"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	msg := "觉得这门语言好难学哦"
	fmt.Println(dao_tool.CalculateAnswer(msg)) // 打印结果
}
