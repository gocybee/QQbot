package config

import (
	"QQbot/global"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

type Config struct {
	TimeLimit             time.Duration `yaml:"time_limit"`                //配置协程最大空闲时间
	PostQuestionToRasaURL string        `yaml:"post_question_to_rasa_url"` //后端链接rasa机器人传输问题的接口
	QuestionAnalysisURL   string        `yaml:"question_analysis_url"`     //后端语义分析接口
	GetRasaAnswerURL      string        `yaml:"get_rasa_answer_url"`       //后端从rasa获取答案的接口
	SendMsgURL            string        `yaml:"send_msg_url"`              //发送QQ信息的接口
	RefuseFileName        string        `yaml:"refuse_file_name"`          //复读打断消息所需图片的文件名
	RefuseURL             string        `yaml:"refuse_url"`                //复读打断消息所需图片的地址
	MyName                string        `yaml:"my_name"`                   //qq机器人的自称(名字)
	MYQQID                string        `yaml:"my_qq_id"`                  //QQ机器人的qq号码
	MaxPoolNumber         int           `yaml:"max_pool_number"`           //连接池最大数量
}

func init() {
	var config Config
	yamlFile, err := ioutil.ReadFile(global.URLTOConfig)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}
	//赋初值
	global.TimeLimit = config.TimeLimit
	global.PostQuestionToRasaURL = config.PostQuestionToRasaURL
	global.QuestionAnalysisURl = config.QuestionAnalysisURL
	global.GetRasaAnswerURL = config.GetRasaAnswerURL
	global.SendMsgURL = config.SendMsgURL
	global.RefuseFileName = config.RefuseFileName
	global.RefuseURL = config.RefuseURL
	global.MyName = config.MyName
	global.MYQQID = config.MYQQID
	global.MaxPoolNumber = config.MaxPoolNumber
}
