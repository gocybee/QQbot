package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"testing"
)

func Test(t *testing.T) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root", "030831", "localhost", "qqbot",
	)

	fmt.Println("dsn::", dsn)

	_, err := gorm.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}
}
