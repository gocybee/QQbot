package runtime

//NormalAnswer 一般语句和QQ表情的组合回答
type NormalAnswer struct {
	Head Head `yaml:"head"` //消息的头部，一般为语气词或者@其他人 config初始化部分
	Body Body `yaml:"body"` //主要的回答信息，config文件中初始化的部分
	Tail Tail `yaml:"tail"` //消息的尾部，通常包含表情包
}

type Head struct {
	ThreeG          []string `yaml:"three_g"`          //3G芯片的故事
	HappyAnswer     []string `yaml:"happy_answer"`     //开心
	FearAnswer      []string `yaml:"fear_answer"`      //害怕
	AngerAnswer     []string `yaml:"anger_answer"`     //生气
	AnonymousAnswer []string `yaml:"anonymous_answer"` //面对匿名的回答
}

type Body struct {
	FriendsMessage  string   `yaml:"friends_message"`  //勤奋蜂的成员介绍
	BeesMessage     string   `yaml:"bees_message"`     //勤奋蜂组织介绍
	ThreeGURL       string   `yaml:"threeG_url"`       //3G芯片的故事
	HappyAnswer     []string `yaml:"happy_answer"`     //开心
	FearAnswer      []string `yaml:"fear_answer"`      //害怕
	AngerAnswer     []string `yaml:"anger_answer"`     //生气
	AnonymousAnswer []string `yaml:"anonymous_answer"` //面对匿名的回答
}

type Tail struct {
	last []string `yaml:"last"` //尾部信息
}

//AnonAnswer 回答匿名消息
func (n *NormalAnswer) AnonAnswer() string {
	return ""
}

//HappyAnswer 开心倾向的回答
func (n *NormalAnswer) HappyAnswer() string {
	return ""
}

//FearAnswer 害怕倾向的回答
func (n *NormalAnswer) FearAnswer() string {
	return ""
}

//AngerAnswer 生气倾向的回答
func (n *NormalAnswer) AngerAnswer() string {
	return ""
}

//ThreeGStory 3G芯片的故事
func (n *NormalAnswer) ThreeGStory() string {
	return ""
}
