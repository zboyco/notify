package notify

// 消息发送器
var senders = make(map[string]Sender)

// 消息发送器接口
type Sender interface {
	Send(topic, wechatUserID, title, content string) error
	Channel() string
}

// 注册消息发送器
func RegisterSender(sender Sender) {
	senders[sender.Channel()] = sender
}

// 获取消息发送器
func GetSender(channel string) Sender {
	return senders[channel]
}
