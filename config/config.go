package config

import (
	"QQbot/global"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"time"
)

type OtherConfig struct {
	TimeLimit      int    `yaml:"time_limit"`       // 配置协程最大空闲时间
	RasaURL        string `yaml:"rasa_url"`         // 后端链接rasa机器人传输问题的接口
	SendMsgURL     string `yaml:"send_msg_url"`     // 发送QQ信息的接口
	RefuseFileName string `yaml:"refuse_file_name"` // 复读打断消息所需图片的文件名
	RefuseURL      string `yaml:"refuse_url"`       // 复读打断消息所需图片的地址
	MyName         string `yaml:"my_name"`          // qq机器人的自称(名字)
	MYQQID         string `yaml:"my_qq_id"`         // QQ机器人的qq号码
	MaxPoolNumber  int    `yaml:"max_pool_number"`  // 连接池最大数量
}

type AnaConfig struct {
	StudioKey       []string `yaml:"studio_key"`         // 其他工作室相关
	QffKey          []string `yaml:"qff_key"`            // 勤奋蜂相关
	FreshmenKey     []string `yaml:"freshmen_key"`       // 勤奋蜂-零基础相关
	QffStayKey      []string `yaml:"qff_stay_key"`       // 勤奋蜂-刷人相关
	QffRecruitKey   []string `yaml:"qff_recruit_key"`    // 勤奋蜂-招新相关
	QffSeniorStuKey []string `yaml:"qff_senior_stu_key"` // 勤奋蜂-学长学姐相关
	QffExam         []string `yaml:"qff_exam_key"`       // 勤奋蜂-考核相关
	QffClass        []string `yaml:"qff_class_key"`      // 勤奋蜂-上课相关
	SchoolKey       []string `yaml:"school_key"`         // 学校相关
	LikeKey         []string `yaml:"like_key"`           // “喜欢”情感倾向相关
	// 3G的关键词默认
}

// AnswerCfg 读取答案
type AnswerCfg struct {
	Studio   []string `yaml:"studio"`
	Qff      []string `yaml:"qff"`
	Freshmen []string `yaml:"freshmen"`
	Stay     []string `yaml:"stay"`
	Recruit  []string `yaml:"recruit"`
	Senior   []string `yaml:"senior"`
	School   []string `yaml:"school"`
	Like     []string `yaml:"like"`
	ThreeG   []string `yaml:"3G"`
	Exam     []string `yaml:"exam"`
	Class    []string `yaml:"class"`
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

	// 赋初值
	global.TimeLimit = time.Duration(config.TimeLimit) // 协程最大空闲时间
	global.RasaURL = config.RasaURL                    // rasa机器人webhook的地址
	global.SendMsgURL = config.SendMsgURL              // 向QQ发送消息的地址
	global.RefuseFileName = config.RefuseFileName      // 拒绝复读的图片文件名
	global.RefuseURL = config.RefuseURL                // 拒绝复读的图片url
	global.MyName = config.MyName                      // QQ机器人的自称(名字)
	global.MYQQID = config.MYQQID                      // QQ机器人的QQ号
	global.MaxPoolNumber = config.MaxPoolNumber        // 协程池的最大数量
}

func loadAnalysisConfig() {
	var config AnaConfig
	yamlFile, err := ioutil.ReadFile(global.URLTOANALYSISConfig)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	// 赋初值
	global.IntentionKey = global.IntentionKeys{
		StudioKey:       config.StudioKey,
		QffKey:          config.QffKey,
		QffFreshmenKey:  config.FreshmenKey,
		QffStayKey:      config.QffStayKey,
		QffRecruitKey:   config.QffRecruitKey,
		QffSeniorStuKey: config.QffSeniorStuKey,
		QffExam:         config.QffExam,
		QffClass:        config.QffClass,
		SchoolKey:       config.SchoolKey,
		LikeKey:         config.LikeKey,
	}

}

func loadAnswerConfig() {
	var config AnswerCfg
	yamlFile, err := ioutil.ReadFile(global.URLTOANSWERConfig)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	// 赋初值
	global.AnswerMap[global.STUDIO] = config.Studio
	global.AnswerMap[global.QFF] = config.Qff
	global.AnswerMap[global.QffFreshmen] = config.Freshmen
	global.AnswerMap[global.QffStay] = config.Stay
	global.AnswerMap[global.QffRecruit] = config.Recruit
	global.AnswerMap[global.QffSenior] = config.Senior
	global.AnswerMap[global.QffExam] = config.Exam
	global.AnswerMap[global.QffClass] = config.Class
	global.AnswerMap[global.SCHOOL] = config.School
	global.AnswerMap[global.LIKE] = config.Like
	global.AnswerMap[global.THREE] = config.ThreeG
}

func init() {
	loadOtherConfig()
	loadAnalysisConfig()
	loadAnswerConfig()
}
