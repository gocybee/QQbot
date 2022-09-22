package global

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/panjf2000/ants/v2"
	"go.uber.org/atomic"
)

const (
	URLTOOTHERConfig = "./config/other_conf.yml" // 到配置文件的路径

	AT    = "at"
	FACE  = "face"
	SHARE = "share"
	IMAGE = "image"
	POKE  = "poke"

	RepeatLimit = 3
)

var (
	RasaURL       string // 后端链接rasa机器人传输问题的接口
	SendMsgURL    string // 发送QQ信息的接口
	MyName        string // qq机器人的自称(名字)
	MYQQID        string // QQ机器人的qq号码
	MaxPoolNumber int    // 连接池最大数量
	TimeLimit     int64  // 配置对话最长保持时间(s 单位)

	Mysql MysqlMsg //储存数据库登录信息

	DB *gorm.DB //数据库链接

	Fathers []string //对机器人有设置权限的人的qq号

	PoolNumber atomic.Int32 // 记录已使用的连接池数量

	Pool *ants.PoolWithFunc // 协程池

	Routing = make(map[string]*RoutingMsg) // 每一个用户的对话储存为一个Logic协程

	Repeated = make(map[string]*Repeat, 1) // 储存可能是复读的句子 索引为群号或QQ号

	// 其中表情的ID是1-221.
)

func PrintVars() {
	fmt.Println()
	fmt.Println("=====配置选项如下=====")
	fmt.Println("RasaURL:", RasaURL)
	fmt.Println("SendMsgURL:", SendMsgURL)
	fmt.Println("BotName:", MyName)
	fmt.Println("BotQQNumber:", MYQQID)
	fmt.Println("MaxPoolNumber:", MaxPoolNumber)
	fmt.Println("TimeLimit:", TimeLimit)
	fmt.Println()
}
