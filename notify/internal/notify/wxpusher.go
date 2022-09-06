package notify

import (
	"github.com/wxpusher/wxpusher-sdk-go"
	wxpusher_model "github.com/wxpusher/wxpusher-sdk-go/model"
)

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

func (w *WxPusher) Send(topicID int, wechatUserID, title, content string) error {
	msg := wxpusher_model.NewMessage(w.appToken)
	msg = msg.SetSummary(title).SetContent(content)
	if topicID != 0 {
		msg = msg.AddTopicId(topicID)
	} else if wechatUserID != "" {
		msg = msg.AddUId(wechatUserID)
	}
	_, err := wxpusher.SendMessage(msg)
	return err
}
