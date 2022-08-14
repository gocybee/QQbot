package dao

import (
	"QQbot/tools"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	msg := "觉得这门语言好难学哦"
	fmt.Println(tools.CalculateAnswer(msg)) // 打印结果
}
