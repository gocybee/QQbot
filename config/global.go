package config

type Config struct {
	Mysql `yaml:"mysql"` //数据库配置
	Res   []QA           `yaml:"res"` //问答初始化
}

type Mysql struct {
	DBName   string `yaml:"dbname"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int32  `yaml:"port"`
}

//QA 用于创建数据库并储存相关的信息
type QA struct {
	ID     int32  `gorm:"AUTO_INCREMENT" gorm:"id"`
	Q1     string `gorm:"type:char(25)" yaml:"q1" gorm:"q1"`
	Q2     string `gorm:"type:char(25)" yaml:"q2" gorm:"q2"`
	Q3     string `gorm:"type:char(25)" yaml:"q3" gorm:"q3"`
	Answer string `gorm:"type:char(255)" yaml:"answer" gorm:"answer"`
}
