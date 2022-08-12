package config

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	err := loadMysqlCfg()
	if err != nil {
		fmt.Println("loadMysqlCfg():", err)
	}
	fmt.Printf("%#v\n", Cfg)

	err = initDB()
	if err != nil {
		fmt.Println("loadMysqlCfg():", err)
	}
	fmt.Println("链接成功")
}
