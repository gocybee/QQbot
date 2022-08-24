package rasa_tool

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	str := GetRasaAnswer("2505772098", "你是谁")
	fmt.Println("回答", str)
}

//func Test2(t *testing.T) {
//	str := GetAnalysisId("Hello")
//	fmt.Println("回答", str)
//}
