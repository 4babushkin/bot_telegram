package storage

import "github.com/jinzhu/gorm"

type Session struct {
	gorm.Model
	ChatID int64
	Phone string
	Address string
	FareId int
	State string
	OrderId int
}

const (
	STATE_NEED_PHONE    = "need phone"
	STATE_NEED_ADDRESS  = "need address"
	STATE_NEED_FARE     = "need fare"
	STATE_ORDER_CREATED = "order created"

	STATE_NEED_COMMAND = "need command"
	SEARCH_BY_NUMBER = "search number"
	SEARCH_TP = "search tp"
	SEARCH_BY_NAME = "search name"
	SEARCH_BY_ABB = "search abb"
	ADD_PHONE = "add phone"
)

func GetSessionByChatID(db *gorm.DB, chatID int64) Session {
	session := Session{}
	db.First(&session, "chat_id = ?", chatID)
	return session
}
