package notify

import (
	"strconv"

	"github.com/wxpusher/wxpusher-sdk-go"
	wxpusher_model "github.com/wxpusher/wxpusher-sdk-go/model"
)

var _ Sender = (*WxPusher)(nil)

type WxPusher struct {
	appToken string
}

func NewWxPusher(appToken string) *WxPusher {
	return &WxPusher{
		appToken: appToken,
	}
}

func (w *WxPusher) Channel() string {
	return "wxpusher"
}

func (w *WxPusher) Send(topic, wechatUserID, title, content string) error {
	msg := wxpusher_model.NewMessage(w.appToken)
	msg = msg.SetSummary(title).SetContent(content)
	if topic != "" {
		topicID, err := strconv.Atoi(topic)
		if err != nil {
			return err
		}
		msg = msg.AddTopicId(topicID)
	} else if wechatUserID != "" {
		msg = msg.AddUId(wechatUserID)
	}
	_, err := wxpusher.SendMessage(msg)
	return err
}
