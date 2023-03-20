package main

// 钉钉oa类型消息通知
type DingOa struct {
	Msgtype string `json:"msgtype"`
	Oa      struct {
		MessageUrl string `json:"message_url"` //消息点击链接地址，当发送消息为小程序时支持小程序跳转链接。
		Head       struct {
			Bgcolor string `json:"bgcolor"` //消息头部的背景颜色。 长度限制为8个英文字符，其中前2为表示透明度，后6位表示颜色值。不要添加0x。
			Text    string `json:"text"`    //消息的头部标题。
		} `json:"head"`
		Body struct {
			Title string `json:"title"`
			Form  []*kv  `json:"form"`
			Rich  struct {
				Num  string `json:"num"`
				Unit string `json:"unit"`
			} `json:"rich"`
			Content   string `json:"content"`
			Image     string `json:"image"`
			FileCount string `json:"file_count"`
			Author    string `json:"author"`
		} `json:"body"`
	} `json:"oa"`
}

type kv struct {
	Key   string `json:"key"`   //消息体的关键字。
	Value string `json:"value"` //消息体的关键字对应的值。
}
