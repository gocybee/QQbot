package global

import "github.com/jinzhu/gorm"

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

var DB *gorm.DB
