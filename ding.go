package main

type DingSend struct {
	Msgtype string `json:"msgtype"`
	Text    struct {
		Content Result `json:"content"`
	} `json:"text"`
}
