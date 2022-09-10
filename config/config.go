package config

import (
	"QQbot/global"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"time"
)

type OtherConfig struct {
	TimeLimit      time.Duration `yaml:"time_limit"`       //配置协程最大空闲时间
	RasaURL        string        `yaml:"rasa_url"`         //后端链接rasa机器人传输问题的接口
	SendMsgURL     string        `yaml:"send_msg_url"`     //发送QQ信息的接口
	RefuseFileName string        `yaml:"refuse_file_name"` //复读打断消息所需图片的文件名
	RefuseURL      string        `yaml:"refuse_url"`       //复读打断消息所需图片的地址
	MyName         string        `yaml:"my_name"`          //qq机器人的自称(名字)
	MYQQID         string        `yaml:"my_qq_id"`         //QQ机器人的qq号码
	MaxPoolNumber  int           `yaml:"max_pool_number"`  //连接池最大数量
}

type AnaConfig struct {
	I global.IntentionKeys
}

func init() {
	loadOtherConfig()
	loadAnalysisConfig()
}

func loadOtherConfig() {
	var config OtherConfig
	yamlFile, err := ioutil.ReadFile(global.URLTOOTHERConfig)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	//赋初值
	global.TimeLimit = config.TimeLimit           //协程最大空闲时间
	global.RasaURL = config.RasaURL               //rasa机器人webhook的地址
	global.SendMsgURL = config.SendMsgURL         //向QQ发送消息的地址
	global.RefuseFileName = config.RefuseFileName //拒绝复读的图片文件名
	global.RefuseURL = config.RefuseURL           //拒绝复读的图片url
	global.MyName = config.MyName                 //QQ机器人的自称(名字)
	global.MYQQID = config.MYQQID                 //QQ机器人的QQ号
	global.MaxPoolNumber = config.MaxPoolNumber   //协程池的最大数量
}

func loadAnalysisConfig() {
	var config AnaConfig
	yamlFile, err := ioutil.ReadFile(global.URLTOANALYSISConfig)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &(config.I))
	if err != nil {
		panic(err)
	}

	//赋初值
	global.IntentionKey = config.I
}
