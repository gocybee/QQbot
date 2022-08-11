//数据示例，最后得删

package global

import "time"

//GroupQA 描述群聊的信息格式
type GroupQA struct {
	Anonymous   string `json:"anonymous,omitempty" yaml:"anonymous"`
	Font        int32  `json:"font,omitempty" yaml:"font"`
	GroupId     string `json:"group_id,omitempty" yaml:"group_id"`
	Message     string `json:"message,omitempty" yaml:"message"`
	MessageId   string `json:"message_id,omitempty" yaml:"message_id"`
	MessageSeq  int64  `json:"message_seq,omitempty" yaml:"message_seq"`
	MessageType string `json:"message_type,omitempty" yaml:"message_type"` //代表群聊信息或者私聊信息
	PostType    string `json:"post_type,omitempty" yaml:"post_type"`
	RawMessage  string `json:"raw_message,omitempty" yaml:"raw_message"` //@我的人所发表的信息
	SelfId      int64  `json:"self_id,omitempty" yaml:"self_id"`
	Sender      struct {
		Age      int32  `json:"age,omitempty" yaml:"age"`
		Area     string `json:"area,omitempty" yaml:"area"`
		Card     string `json:"card,omitempty" yaml:"card"`
		Level    int32  `json:"level,omitempty" yaml:"level"`
		Nickname string `json:"nickname,omitempty" yaml:"nickname"`
		Role     string `json:"role,omitempty" yaml:"role"`
		Sex      string `json:"sex,omitempty" yaml:"sex"`
		Title    string `json:"title,omitempty" yaml:"title"`
		UserId   string `json:"user_id,omitempty" yaml:"user_id"`
	} `json:"sender" yaml:"sender" yaml:"g_sender"`
	SubType string    `json:"sub_type,omitempty" yaml:"sub_type"`
	Time    time.Time `json:"time" yaml:"time"`
	UserId  int64     `json:"user_id,omitempty" yaml:"user_id"`
}
