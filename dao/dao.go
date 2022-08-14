package dao

import (
	"QQbot/global"
	"fmt"
)

//SelectQA 获取所有的问题
func SelectQA() error {
	var qas []global.QA
	err := global.DB.Find(&qas).Error
	if err != nil {
		return err
	}
	for i := 0; i < len(qas); i++ {
		global.QAs = append(global.QAs, &qas[i])
	}

	for _, v := range global.QAs {
		fmt.Println(v)
	}
	return nil
}

//AddQA 向数据库中增加问答
func AddQA(qa global.QA) error {
	return global.DB.Create(&qa).Error
}
