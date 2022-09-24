package global

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/panjf2000/ants/v2"
	"go.uber.org/atomic"
)

const (
	URLTOOTHERConfig = "./config/config.yml" // 到配置文件的路径

	AT    = "at"
	FACE  = "face"
	SHARE = "share"
	IMAGE = "image"
	POKE  = "poke"

	RepeatLimit = 3
)

var (
	Debug bool //是否为debug状态-标志聊天白名单是否开启

	// RasaURL rasa机器人url
	RasaURL string

	// SendMsgURL QQ机器人配置信息
	SendMsgURL string // 发送QQ信息的接口
	MyName     string // qq机器人的自称(名字)
	MYQQID     string // QQ机器人的qq号码

	// Mysql 数据库信息配置
	Mysql MysqlMsg // 储存数据库登录信息
	DB    *gorm.DB // 数据库链接

	// Pool 协程池配置
	Pool          *ants.PoolWithFunc // 协程池
	PoolNumber    atomic.Int32       // 记录已使用的连接池数量
	TimeLimit     int64              // 配置对话最长保持时间(s 单位)
	MaxPoolNumber int                // 连接池最大数量

	// Routing 每个用户聊天的协程描述
	Routing = make(map[string]*RoutingMsg)

	// Repeated 全局复读记录表-key是发送方的OppositeId
	Repeated = make(map[string]*Repeat, 1)

	// 其中表情的ID是1-221.

	Controller []string // 对机器人有设置权限的人的qq号

	ChatList []string //聊天白名单
)

func PrintVars() {
	fmt.Println()
	fmt.Println("=====配置选项如下=====")
	fmt.Println("Debug", Debug)
	fmt.Println("RasaURL:", RasaURL)
	fmt.Println("SendMsgURL:", SendMsgURL)
	fmt.Println("BotName:", MyName)
	fmt.Println("BotQQNumber:", MYQQID)
	fmt.Println("MaxPoolNumber:", MaxPoolNumber)
	fmt.Println("TimeLimit:", TimeLimit)
	fmt.Println()
}
