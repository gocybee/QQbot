package global

import "fmt"

// CQAt 表示@的信息
type CQAt struct {
	Type string // 表明CQ类型at
	Data struct {
		QQ   int64  // @群内的人的qq
		Name string // 此人不在群内则输入名字
	}
}

// CQFace 表示QQ表情
type CQFace struct {
	Type string // 表明CQ类型 face
	Data struct {
		Id int64 // 表情的序号
	}
}

// CQUrl 描述QQ的分享链接
type CQUrl struct {
	Type string // 表明CQ类型share
	Data struct {
		Url    string
		Tittle string
	}
}

// CodeCQAt 包装at信息,
func CodeCQAt(qq *int64) string {
	return fmt.Sprintf("[CQ:%s,qq=%v]", AT, *qq)
}

// CodeCQFace 包装表情信息
func CodeCQFace(id int64) string {
	return fmt.Sprintf("[CQ:%s,id=%v]", FACE, id)
}

// CodeCQUrl 包装链接信息
func CodeCQUrl(url string, title string) string {
	return fmt.Sprintf("[CQ:%s,url=%s,title=%s]", SHARE, url, title)
}
