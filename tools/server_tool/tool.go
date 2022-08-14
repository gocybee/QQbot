package server_tool

import (
	"QQbot/global"
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
)

// GetIdAndMsg 从初始结构体中获取信息和发送的id
func GetIdAndMsg(form *map[string]interface{}, flag string) (*int64, *string, error) {
	var id int64
	msg := (*form)["raw_message"].(string)
	if flag == "group" {

		id = int64((*form)["group_id"].(float64)) //获取群号

	} else if flag == "private" {

		id = int64((*form)["user_id"].(float64)) //获取QQ号

	} else {
		return nil, nil, errors.New("flag error")
	}
	return &id, &msg, nil
}

//AIHelp 获取AI帮助
func AIHelp(msg *string) (string, error) {
	url := fmt.Sprintf("http://api.qingyunke.com/api.php?key=free&appid=0&msg=%s", *msg)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var AIType global.AI
	err = json.Unmarshal(data, &AIType)
	if err != nil {
		return "", err
	}
	return AIType.Content, nil
}

//ExportSqlMsg 导出所有的学习的信息
func ExportSqlMsg() error {
	data, err := yaml.Marshal(global.QAs)
	if err != nil {
		return err
	}
	return ioutil.WriteFile("oldMsg.yml", data, 0777)
}

//GetIdFromMap 从接受到的表单中提取出用户或者群聊Id
func GetIdFromMap(id interface{}) int64 {
	return int64(id.(float64))
}

//DoOrNot 生成随机数换算为概率--输入小数,现两位，默认0.5
func DoOrNot(p float32) bool {
	i := rand.Int() % 100
	if i < int(p*100) {
		return true
	}
	return false
}

//GetUsefulMsg 删去@自己部分（CQcode部分），获取消息的可被分析部分
func GetUsefulMsg(msg *string) *string {
	var x [2]int
	str := *msg
	res := []rune(str)

	for i := 0; i < len(res); i++ {
		if res[i] == '[' {
			x[0] = i
		}
		if res[i] == ']' {
			x[1] = i
		}
		if x[0] != 0 || x[1] != 0 {
			res = []rune(string(res[:x[0]]) + string(res[x[1]+1:]))
			x[0], x[1] = 0, 0
		}
	}
	an := strings.TrimSpace(string(res))
	return &an
}

//Beautify 为句子的头和尾美化
func Beautify(ctx *string) {
	//60%的概率做尾部美化
	if DoOrNot(0.6) {
		i := rand.Int()%221 + 1
		//避开奇怪的表情
		if (i > 40 && i < 92) || (i > 111 && i < 172) || i > 183 {
			i = 179
		}
		*ctx += global.CodeCQFace(int64(i))
	}
}
