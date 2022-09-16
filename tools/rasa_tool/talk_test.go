package rasa_tool

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	str, _ := GetRasaAnswer("2505772068", "你是谁")
	fmt.Println("回答", str)
}
