package notify

// 消息发送器
var Senders = []Sender{}

// 消息发送器接口
type Sender interface {
	Send(topicID int, wechatUserID, title, content string) error
	Channel() string
}

// 注册消息发送器
func RegisterSender(sender Sender) {
	Senders = append(Senders, sender)
}
