package msgpush

import (
	"github.com/imroc/req"
)

const wecomBase = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="

type weComSendTextContent struct {
	Msgtype string `json:"msgtype"`
	Text    struct {
		Content             string   `json:"content"`
		MentionedList       []string `json:"mentioned_list"`
		MentionedMobileList []string `json:"mentioned_mobile_list"`
	} `json:"text"`
}

type weComSendMDContent struct {
	Msgtype  string `json:"msgtype"`
	Markdown struct {
		Content string `json:"content"`
	} `json:"markdown"`
}

type WeCom struct {
	Token  string
	ReqUrl string
}

// NewWeCom
func NewWeCom(token string) *WeCom {
	return &WeCom{Token: token, ReqUrl: wecomBase + token}
}

func (w *WeCom) Send(content string) error {
	return w.SendMd(content)
}

func (w *WeCom) SendMd(content string) error {
	_, err := req.Post(w.ReqUrl, req.BodyJSON(&weComSendMDContent{
		Msgtype: "markdown",
		Markdown: struct {
			Content string `json:"content"`
		}{
			Content: content,
		},
	}))
	return err
}

func (w *WeCom) SendText(content string, ms []string) error {
	_, err := req.Post(w.ReqUrl, req.BodyJSON(&weComSendTextContent{
		Msgtype: "text",
		Text: struct {
			Content             string   `json:"content"`
			MentionedList       []string `json:"mentioned_list"`
			MentionedMobileList []string `json:"mentioned_mobile_list"`
		}{
			Content:       content,
			MentionedList: ms,
		},
	}))
	return err
}

func (w *WeCom) String() string {
	return "wecom"
}
