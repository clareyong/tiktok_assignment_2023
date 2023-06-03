package main

type Message struct {
	Sender   string `gorm:"sender"`
	Receiver string `gorm:"receiver"`
	Text     string `gorm:"text"`
}

func (Message) TableName() string {
	return "message"
}
