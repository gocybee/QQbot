package tools

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	res, err := AIHelp(" 天气 ")
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
