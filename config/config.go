package config

import (
	"QQbot/global"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type OtherConfig struct {
	TimeLimit int64 `yaml:"time_limit"` // 配置协程最大空闲时间

	RasaURL string `yaml:"rasa_url"` // 后端链接rasa机器人传输问题的接口

	SendMsgURL string `yaml:"send_msg_url"` // 发送QQ信息的接口

	MyName string `yaml:"my_name"`  // qq机器人的自称(名字)
	MYQQID string `yaml:"my_qq_id"` // QQ机器人的qq号码

	MaxPoolNumber int `yaml:"max_pool_number"` // 连接池最大数量

	MysqlId       string `yaml:"mysql_id"`       // 数据库账号
	MysqlPassword string `yaml:"mysql_password"` // 数据库密码
	MysqlAddr     string `yaml:"mysql_addr"`     // 数据库url
	Database      string `yaml:"database"`       // 库的名称

	Father []string `yaml:"father"` // 控制权限所有者
}

func LoadOtherConfig() error {
	var config OtherConfig
	yamlFile, err := ioutil.ReadFile(global.URLTOOTHERConfig)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return err
	}

	// 赋初值
	global.TimeLimit = config.TimeLimit         // 协程最大空闲时间
	global.RasaURL = config.RasaURL             // rasa机器人webhook的地址
	global.SendMsgURL = config.SendMsgURL       // 向QQ发送消息的地址
	global.MyName = config.MyName               // QQ机器人的自称(名字)
	global.MYQQID = config.MYQQID               // QQ机器人的QQ号
	global.MaxPoolNumber = config.MaxPoolNumber // 协程池的最大数量
	global.Fathers = config.Father

	// 数据库登录信息
	global.Mysql = global.MysqlMsg{
		UId:      config.MysqlId,
		Password: config.MysqlPassword,
		Database: config.Database,
		Address:  config.MysqlAddr,
	}
	return nil
}
