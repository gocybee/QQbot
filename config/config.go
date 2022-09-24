package config

import (
	"QQbot/global"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	Pool       Pool     `yaml:"pool"`
	Rasa       Rasa     `yaml:"rasa"`
	QQ         QQ       `yaml:"qq"`
	Mysql      Mysql    `yaml:"mysql"`
	Debug      bool     `yaml:"debug"`      //控制chat_list的开关
	Controller []string `yaml:"controller"` //bot管理
	ChatList   []string `yaml:"chat_list"`  //回答的白名单
}

// Pool 协程池连接配置
type Pool struct {
	TimeLimit     int64 `yaml:"time_limit"`      // 配置协程最大空闲时间
	MaxPoolNumber int   `yaml:"max_pool_number"` // 连接池最大数量
}

// Rasa rasa机器人链接配置
type Rasa struct {
	RasaURL string `yaml:"rasa_url"` // 后端链接rasa机器人传输问题的接口
}

// QQ QQ机器人的相关配置
type QQ struct {
	SendMsgURL string `yaml:"send_msg_url"` // 发送QQ信息的接口
	Name       string `yaml:"name"`         // qq机器人的自称(名字)
	QUID       string `yaml:"qq_id"`        // QQ机器人的qq号码
}

// Mysql 数据库配置
type Mysql struct {
	User     string `yaml:"user"`     // 数据库账号
	Password string `yaml:"password"` // 数据库密码
	Addr     string `yaml:"addr"`     // 数据库url
	DBName   string `yaml:"dbname"`   // 库的名称
}

func Init() error {
	var config Config
	yamlFile, err := ioutil.ReadFile(global.URLTOOTHERConfig)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return err
	}

	//debug
	global.Debug = config.Debug

	//Pool
	global.TimeLimit = config.Pool.TimeLimit         // 协程最大空闲时间
	global.MaxPoolNumber = config.Pool.MaxPoolNumber // 协程池的最大数量

	//Rasa
	global.RasaURL = config.Rasa.RasaURL // rasa机器人http_hook的地址

	//QQ
	global.SendMsgURL = config.QQ.SendMsgURL // 向QQ发送消息的地址
	global.MyName = config.QQ.Name           // QQ机器人的自称(名字)
	global.MYQQID = config.QQ.QUID           // QQ机器人的QQ号

	//Mysql
	global.Mysql = global.MysqlMsg{
		User:     config.Mysql.User,
		Password: config.Mysql.Password,
		DbName:   config.Mysql.DBName,
		Address:  config.Mysql.Addr,
	}

	//Controller
	global.Controller = config.Controller

	//ChatList
	global.ChatList = config.ChatList

	return nil
}
