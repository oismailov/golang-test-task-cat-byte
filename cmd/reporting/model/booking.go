package model

type Message struct {
	UUID     string `json:"uuid" gorm:"unique"`
	Sender   string `valid:"required,alpha"  json:"sender"`
	Receiver string `valid:"required,alpha"  json:"receiver"`
	Message  string `valid:"required,alpha"  json:"message"`
}
