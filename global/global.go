package global

import (
	"github.com/jinzhu/gorm"
)

// QA 用于创建数据库并储存相关的信息
type QA struct {
	ID     int32  `gorm:"AUTO_INCREMENT" gorm:"id"`
	Q1     string `gorm:"type:nvarchar(25)" yaml:"q1" gorm:"q1"`
	Q2     string `gorm:"type:nvarchar(25)" yaml:"q2" gorm:"q2"`
	Q3     string `gorm:"type:nvarchar(25)" yaml:"q3" gorm:"q3"`
	Answer string `gorm:"type:nvarchar(255)" yaml:"answer" gorm:"answer"`
}

type AI struct {
	Result  int    `json:"result,omitempty"`
	Content string `json:"content,omitempty"`
}

type Repeat struct {
	Flag    string // group or private
	Content string
	Id      int64 // 指向群号
	Times   int   // 重复次数
}

const (
	SendMsgURL  = "http://127.0.0.1:5700"
	ResourceURL = "./resource/"
	CfgFileURL  = "./config/config.yml"

	RefuseFileName = "e185eceb199babeb7fe1061df5a88236.image"
	RefuseURL      = "https://gchat.qpic.cn/gchatpic_new/2505772098/881902822-3083931220-E185ECEB199BABEB7FE1061DF5A88236/0?term=3"
	MyName         = "我"
	MYQQID         = "3403191872"
	DistanceLimit  = 10

	AT    = "at"
	FACE  = "face"
	SHARE = "share"
	IMAGE = "image"
	POKE  = "poke"

	PrivateFlag = "private"
	GroupFlag   = "group"

	RepeatLimit = 2
)

var (
	DB  *gorm.DB
	QAs []*QA
	Re  = make(map[string]*Repeat, 1) // 储存可能是复读的句子 索引为群号或QQ号
	// 其中表情的ID是1-221.
)
