package main

type DingCard struct {
	Msgtype    string `json:"msgtype"`
	ActionCard struct {
		Title          string    `json:"title"`
		Markdown       string    `json:"markdown"`
		BtnOrientation string    `json:"btn_orientation"`
		BtnJsonList    []*CardKv `json:"btn_json_list"`
	} `json:"action_card"`
}

type CardKv struct {
	Title     string `json:"title"`
	ActionUrl string `json:"action_url"`
}
