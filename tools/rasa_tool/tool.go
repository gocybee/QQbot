package rasa_tool

import (
	"QQbot/global"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
)

//GetRasaAnswer 向rasa发送问题并收到回复
func GetRasaAnswer(session string, text string) (string, error) {
	q := global.RasaPost{Sender: session, Message: text}
	reader, err := json.Marshal(q)
	if err != nil {
		return "", err
	}
	resp, err := http.Post(global.RasaURL+"/webhooks/rest/webhook", "application/json", strings.NewReader(string(reader)))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var an []map[string]interface{}
	err = json.Unmarshal(data, &an)
	if err != nil {
		return "", err
	}

	var ans string //储存回答的语句
	for _, v := range an {
		if a, ok := v["text"]; ok {
			ans, err = strconv.Unquote(strings.Replace(strconv.Quote(a.(string)), `\\u`, `\u`, -1))
			if err != nil {
				return "", err
			}
		}
	}

	return ans, nil
}
