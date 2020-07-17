package storage

import "github.com/jinzhu/gorm"

type Phone struct {
	gorm.Model
	ChatID int64
	Number string
	Ph string
}

func GetLastPhonesByChatID(db *gorm.DB, chatID int64) []Phone {
	phones := []Phone{}
	db.Order("id desc").Limit(3).Where("chat_id = ?", chatID).Find(&phones)
	return phones
}

func GetPhonesByNumber(db *gorm.DB, chatID int64) []Phone {
	phones := []Phone{}
	db.Order("id desc").Limit(3).Where("chat_id = ?", chatID).Find(&phones)
	return phones
}

func ChekPhoneByNumber(db *gorm.DB, ph string) bool {
	phones := []Phone{}

	db.First(&phones, "number = ?", ph)

	//db.Order("id desc").Limit(3).Where("number = ?", ph).Find(&phones)
	if len(phones) == 0 {
		return false
	}

	return true

}
