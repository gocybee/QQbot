package tools

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	x := "刚步入大学的你是想虚度四年，还是想和优秀的人一起变得更优秀？你想要更好的学习氛围吗？那就来了解一下极客&勤奋峰科技团队吧"
	status := Send(881902822, &x, "group")

	fmt.Println(status)
}
