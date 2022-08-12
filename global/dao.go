package global

//QA 用于创建数据库并储存相关的信息
type QA struct {
	ID     int32  `gorm:""`
	Q1     string `gorm:""`
	Q2     string `gorm:""`
	Q3     string `gorm:""`
	Answer string `gorm:""`
}
