package chat

import (
	"gopkg.in/telegram-bot-api.v4"
	"fmt"
	"log"
	"errors"
	"../api"
	"../storage"
)

var NambaTaxiApi api.NambaTaxiApi

func GetBasicKeyboard() tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Поиск"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Показания"),
		),
	)
	keyboard.OneTimeKeyboard = true
	return keyboard
}
func GetAskueKeyboard() tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Поиск по №"),
			tgbotapi.NewKeyboardButton("Поиск по названию"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Поиск по Абоненту"),
			tgbotapi.NewKeyboardButton("Подстанции"),
		),
	)
	keyboard.OneTimeKeyboard = true
	return keyboard
}

func BackKeyboard() tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("назад"),
		),
	)
	keyboard.OneTimeKeyboard = true
	return keyboard
}

func GetTPKeyboard() tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(

		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Дивногорск"),
			tgbotapi.NewKeyboardButton("ДивПромзона"),
				),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Овсянка"),
			tgbotapi.NewKeyboardButton("УстьМана"),
			tgbotapi.NewKeyboardButton("Манский"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Слизнево"),
			tgbotapi.NewKeyboardButton("Бирюса"),
			tgbotapi.NewKeyboardButton("назад"),
		),
	)
	keyboard.OneTimeKeyboard = true
	return keyboard
}



func GetPhoneKeyboard() tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButtonContact("Отправить ваш номер телефона"),
		),
	)
	keyboard.OneTimeKeyboard = true
	return keyboard
}

func GetOrderKeyboard() tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Узнать статус моего заказа"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Отменить мой заказ"),
		),
	)
	keyboard.OneTimeKeyboard = true
	return keyboard
}

func GetFaresKeyboard() tgbotapi.ReplyKeyboardMarkup {
	fares, err := NambaTaxiApi.GetFares()
	if err != nil {
		log.Printf("error getting fares: %v", err)
		return tgbotapi.NewReplyKeyboard()
	}

	var rows [][]tgbotapi.KeyboardButton
	for _, fare := range fares.Fare {
		rows = append(rows, tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton(fare.Name)))
	}

	keyboard := tgbotapi.NewReplyKeyboard(rows...)
	keyboard.OneTimeKeyboard = true
	return keyboard
}

func GetAddressKeyboard(addresses []storage.Address) tgbotapi.ReplyKeyboardMarkup {
	var rows [][]tgbotapi.KeyboardButton
	for _, address := range addresses {
		rows = append(rows, tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton(address.Text)))
	}
	keyboard := tgbotapi.NewReplyKeyboard(rows...)
	keyboard.OneTimeKeyboard = true
	return keyboard
}

func GetPhonesKeyboard(phones []storage.Phone) tgbotapi.ReplyKeyboardMarkup {
	var rows [][]tgbotapi.KeyboardButton
	for _, phone := range phones {
		rows = append(rows, tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton(phone.Number)))
	}
	keyboard := tgbotapi.NewReplyKeyboard(rows...)
	keyboard.OneTimeKeyboard = true
	return keyboard
}

func GetFareIdByName(fareName string) (int, error) {
	fares, err := NambaTaxiApi.GetFares()
	if err != nil {
		log.Printf("error getting fares: %v", err)
		return 0, err
	}
	for _, fare := range fares.Fare {
		if fare.Name == fareName {
			return fare.Id, nil
		}
	}
	return 0, errors.New(fmt.Sprintf("Cannot find fare with name %v", fareName))
}
