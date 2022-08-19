package global

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
