package storage

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

const TEST_DB_NAME = "test.db"

func getSession() Session {
	session := Session{}
	session.ChatID = 1
	session.OrderId = 1
	session.FareId = 1
	session.Address = "address"
	session.Phone = "123"
	session.State = "state"
	return session
}

func TestStoreAndGetSession(t *testing.T) {
	db := initDB()
	session1 := getSession()
	db.Create(&session1)
	session2 := Session{}
	db.First(&session2, "phone = ?", "123")
	assert.Equal(t, "123", session2.Phone)
	deleteDB()
}

func TestGetSessionByChatID_AndUpdate(t *testing.T) {
	db := initDB()
	session1 := getSession()
	db.Create(&session1)

	session2 := GetSessionByChatID(db, 1)
	session2.Phone = "456"
	db.Save(&session2)

	session3 := GetSessionByChatID(db, 1)
	assert.Equal(t, "456", session3.Phone)
	deleteDB()
}

func TestGetSessionByChatID_DoesNotExist(t *testing.T) {
	db := initDB()
	session := GetSessionByChatID(db, 1)
	assert.Equal(t, Session{}, session)
	deleteDB()
}

func TestGetSessionByChatID_Delete(t *testing.T) {
	db := initDB()

	session1 := getSession()
	db.Create(&session1)

	session2 := GetSessionByChatID(db, 1)
	assert.Equal(t, "123", session2.Phone)

	db.Delete(session2)

	session3 := GetSessionByChatID(db, 1)
	assert.Equal(t, int64(0), session3.ChatID)
	assert.Equal(t, Session{}, session3)

	deleteDB()
}

func TestGetSessionByChatID_DeleteNonExistentSession(t *testing.T) {
	db := initDB()

	session := GetSessionByChatID(db, 100)
	db.Delete(session)

	deleteDB()
}

