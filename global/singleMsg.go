//数据示例，最后得删

package global

import "time"

//SingleQA 描述私聊的信息格式
type SingleQA struct {
	Font        int32  `json:"font,omitempty" yaml:"font"`
	Message     string `json:"message,omitempty" yaml:"message"`
	MessageId   string `json:"message_id,omitempty" yaml:"message_id"`
	MessageType string `json:"message_type,omitempty" yaml:"message_type"` //区分群聊和私聊信息
	PostType    string `json:"post_type,omitempty" yaml:"post_type"`
	RawMessage  string `json:"raw_message,omitempty" yaml:"raw_message"` //信息原文
	SelfId      int64  `json:"self_id,omitempty" yaml:"self_id"`
	Sender      struct {
		Age      int32  `json:"age,omitempty" yaml:"age"`
		Nickname string `json:"nickname,omitempty" yaml:"nickname"` //发送方的QQ名称
		Sex      string `json:"sex,omitempty" yaml:"sex"`
		UserId   int64  `json:"user_id,omitempty" yaml:"user_id"`
	}
	SubType  string    `json:"sub_type,omitempty" yaml:"sub_type"`
	TargetId int64     `json:"target_id" yaml:"target_id"`
	Time     time.Time `json:"time" yaml:"time"`
	UserId   int64     `json:"user_id,omitempty" yaml:"user_id"`
}
