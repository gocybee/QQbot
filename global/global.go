package global

import (
	"github.com/jinzhu/gorm"
)

const (
	CfgFileURL = "D:/GithubLibrary/QQbot/config/config.yml"
	MyName     = "我"
	MYQQID     = "3403191872"
	_          = iota
	FirMsg
	BeeMsg
	BlogMsg
	ThreeG
	HappyAn
	FearAn
)

var (
	DB  *gorm.DB
	Add = []string{"嘿嘿,", "emmm", "啊哈,", "qwq"} //供句子美化
	//其中表情的ID是1-221.
)
