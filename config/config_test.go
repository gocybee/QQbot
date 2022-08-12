package config

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	err := loadCfg()
	if err != nil {
		fmt.Println("loadCfg():", err)
	}
	fmt.Printf("%#v\n", cfg)

	err = initDB()
	if err != nil {
		fmt.Println("InitDB():", err)
	}
	fmt.Println("链接成功")

	err = loadQA()
	if err != nil {
		fmt.Println("loadQA():", err)
	}
	fmt.Println("加载成功")
}
