package dao

import (
	"QQbot/config"
	"QQbot/global"
)

//SelectQA 获取所有的问题
func SelectQA() (*[]config.QA, error) {
	var qa []config.QA
	err := global.DB.Find(&qa).Error
	if err != nil {
		return nil, err
	}
	return &qa, nil
}
