package main

import (
		"gopkg.in/telegram-bot-api.v4"

    "log"
    "os"
    "./chat"
    "./storage"
    "./askuemssql"
    "strings"

)

var (
	db  = storage.GetGormDB("./bot.sqlite3")
)


func main() {

	mainRun()
}


func mainRun() {
    //os.Setenv("HTTP_PROXY", "ip:port")
    storage.MigrateAll(db)

    TelegramBotToken := os.Getenv("BOT_TOKEN")
	

    bot, err := tgbotapi.NewBotAPI(TelegramBotToken)
    log.Printf("Authorized on account %s", TelegramBotToken)


    if err != nil {
    log.Panic(err)
    }

    bot.Debug = false

    log.Printf("Authorized on account %s", bot.Self.UserName)

    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60

    updates, err := bot.GetUpdatesChan(u)

    if err != nil {
        log.Panic(err)
    }
    // В канал updates будут приходить все новые сообщения.
    for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		chatStateMachine(update, bot)
	}
}

func chatStateMachine(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
  //basicKeyboard := chat.GetBasicKeyboard()
  //orderKeyboard := chat.GetOrderKeyboard()
  session := storage.GetSessionByChatID(db, update.Message.Chat.ID)

  //phone := update.Message.Text
    if update.Message.Contact != nil {
      phone := "+" + update.Message.Contact.PhoneNumber
      log.Printf("[%s] %s", update.Message.Contact, phone)


  //storage.ChekPhoneByNumber(db, phone) ==

      if  storage.ChekPhoneByNumber(db,phone) {
    //  if  storage.ChekPhoneByNumber(db,phone){

        session := &storage.Session{}
  		  session.ChatID = update.Message.Chat.ID
  		  session.State = storage.STATE_NEED_COMMAND
        session.Phone = phone

        msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Слушаю мой повелитель")
        //msg.ReplyToMessageID = update.Message.MessageID

        msg.ReplyMarkup = chat.GetAskueKeyboard()
        bot.Send(msg)
        db.Create(&session)
        return
      }  else   {
        msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Доступ запрещен")
        //msg.ReplyToMessageID = update.Message.MessageID
        msg.ReplyMarkup = chat.GetPhoneKeyboard()
        bot.Send(msg)
        return
       }
    }

  log.Printf("[%s] %s - %s", "Сессия", session.Phone, session.State)
  log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Chat.ID)

  if session.ChatID != int64(0) {

    switch session.State {
        case storage.STATE_NEED_COMMAND:

          if update.Message.Text == "Поиск по №" {
          		//msg := tgbotapi.NewMessage(update.Message.Chat.ID, "_КаналЫ_")
          		//msg.ReplyMarkup = chat.GetPhoneKeyboard()
          		//bot.Send(msg)
              session.State = storage.SEARCH_BY_NUMBER
              log.Printf("[[%s]]", session.State)
              db.Save(&session)
              msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите номер счетчика")
              msg.ReplyMarkup = chat.BackKeyboard()
              bot.Send(msg)
          		return
          	}
            if update.Message.Text == "Поиск по названию" {
              session.State = storage.SEARCH_BY_NAME
              log.Printf("[[%s]]", session.State)
              db.Save(&session)
              msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите название (выводит только 5 ПУ)")
              msg.ReplyMarkup = chat.BackKeyboard()
              bot.Send(msg)
            		return
          	}
            if update.Message.Text == "Поиск по Абоненту" {
              session.State = storage.SEARCH_BY_ABB
              log.Printf("[[%s]]", session.State)
              db.Save(&session)
              msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите Абонента")
              msg.ReplyMarkup = chat.BackKeyboard()
              bot.Send(msg)
                return
            }

            if update.Message.Text == "Подстанции" {
              session.State = storage.SEARCH_TP
              db.Save(&session)
              msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите место")
              msg.ReplyMarkup = chat.GetTPKeyboard()
              bot.Send(msg)
              return
            }
            if update.Message.Text == "/exit" {
              db.Delete(&session)
              //db.Save(&session)
              msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вышли")
              msg.ReplyMarkup = chat.GetPhoneKeyboard()
              bot.Send(msg)
              return
            }

            if update.Message.Text == "Add1692" {
              session.State = storage.ADD_PHONE
              log.Printf("[[%s]]", session.State)
              db.Save(&session)
              msg := tgbotapi.NewMessage(update.Message.Chat.ID, "введите номер телефона")
              msg.ReplyMarkup = chat.BackKeyboard()
              bot.Send(msg)
            		return
            }


        return

      case storage.ADD_PHONE:

        if update.Message.Text == "назад" {
          session.State = storage.STATE_NEED_COMMAND
          db.Save(&session)
          msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Что делаем?")
          msg.ReplyMarkup = chat.GetAskueKeyboard()
          bot.Send(msg)

          return
        }


          phone4 :=  &storage.Phone{}
          phone4.ChatID = 2
          phone4.Number = update.Message.Text
          db.Create(&phone4)

          session.State = storage.STATE_NEED_COMMAND
          db.Save(&session)
          msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Добавил")
          msg.ReplyMarkup = chat.GetAskueKeyboard()
          bot.Send(msg)


      return



        case storage.SEARCH_BY_NUMBER:

					if update.Message.Text == "/exit" {
						db.Delete(&session)
						//db.Save(&session)
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вышли")
						msg.ReplyMarkup = chat.GetPhoneKeyboard()
						bot.Send(msg)
						return
					}

				  if update.Message.Text == "назад" {
            session.State = storage.STATE_NEED_COMMAND
            db.Save(&session)
            msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Что делаем?")
            msg.ReplyMarkup = chat.GetAskueKeyboard()
            bot.Send(msg)

            return
          }

          askuemssql.ConnectDb()
          text := askuemssql.FindAll(update.Message.Text, storage.SEARCH_BY_NUMBER)
          if text == "" {
            text ="нет такого счетчика"
          }
          msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
          msg.ParseMode = "HTML"
          bot.Send(msg)
          askuemssql.CloseDb()

        return

      case storage.SEARCH_BY_NAME:

				if update.Message.Text == "/exit" {
					db.Delete(&session)
					//db.Save(&session)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вышли")
					msg.ReplyMarkup = chat.GetPhoneKeyboard()
					bot.Send(msg)
					return
				}

        if update.Message.Text == "назад" {
          session.State = storage.STATE_NEED_COMMAND
          db.Save(&session)
          msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Что делаем?")
          msg.ReplyMarkup = chat.GetAskueKeyboard()
          bot.Send(msg)

          return
        }

        askuemssql.ConnectDb()
        text := askuemssql.FindAll(update.Message.Text, storage.SEARCH_BY_NAME)
        if text == "" {
          text ="нет такого счетчика"
        }
        msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
        msg.ParseMode = "HTML"
        bot.Send(msg)
        askuemssql.CloseDb()

      return

    case storage.SEARCH_BY_ABB:

			if update.Message.Text == "/exit" {
				db.Delete(&session)
				//db.Save(&session)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вышли")
				msg.ReplyMarkup = chat.GetPhoneKeyboard()
				bot.Send(msg)
				return
			}

      if update.Message.Text == "назад" {
        session.State = storage.STATE_NEED_COMMAND
        db.Save(&session)
        msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Что делаем?")
        msg.ReplyMarkup = chat.GetAskueKeyboard()
        bot.Send(msg)

        return
      }

      askuemssql.ConnectDb()
      text := askuemssql.FindAll(update.Message.Text, storage.SEARCH_BY_ABB)
      if text == "" {
        text ="нет такого счетчика"
      }
      msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
      msg.ParseMode = "HTML"
      bot.Send(msg)
      askuemssql.CloseDb()

    return




      case storage.SEARCH_TP:

				if update.Message.Text == "/exit" {
					db.Delete(&session)
					//db.Save(&session)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вышли")
					msg.ReplyMarkup = chat.GetPhoneKeyboard()
					bot.Send(msg)
					return
				}

        if update.Message.Text == "назад" {
          session.State = storage.STATE_NEED_COMMAND
          db.Save(&session)
          msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Что делаем?")
          msg.ReplyMarkup = chat.GetAskueKeyboard()
          bot.Send(msg)
          return

        }

        askuemssql.ConnectDb()
        msg := tgbotapi.NewMessage(update.Message.Chat.ID, askuemssql.ShowTP(update.Message.Text))
        msg.ParseMode = "HTML"
        msg.ReplyMarkup = chat.GetTPKeyboard()


        bot.Send(msg)
        askuemssql.CloseDb()
      return


        case storage.STATE_NEED_PHONE:
    			phone := update.Message.Text
    			if update.Message.Contact != nil {
    				phone = "+" + update.Message.Contact.PhoneNumber
    			}

    			if !strings.HasPrefix(phone, "+7913") {
    				msg := tgbotapi.NewMessage(update.Message.Chat.ID, chat.BOT_PHONE_START_996)
    				phones := storage.GetLastPhonesByChatID(db, session.ChatID)
    				if len(phones) > 0 {
    					msg.ReplyMarkup = chat.GetPhonesKeyboard(phones)
    				} else {
    					msg.ReplyMarkup = chat.GetPhoneKeyboard()
    				}
    				bot.Send(msg)
    				return
    			}
    			session.Phone = phone
    			session.State = storage.STATE_NEED_FARE
    			db.Save(&session)
    			msg := tgbotapi.NewMessage(update.Message.Chat.ID, chat.BOT_ASK_FARE)
    			msg.ReplyMarkup = chat.GetFaresKeyboard()
    			bot.Send(msg)
    			return

    		case storage.STATE_ORDER_CREATED:
    			if update.Message.Text == "Отменить мой заказ" {
    			////	chat.HandleOrderCancel(nambaTaxiAPI, &session, db, update, bot)
    				return
    			}

    			return

    		default:
    			//db.Delete(&session)
          session.State = storage.STATE_NEED_COMMAND
          db.Save(&session)

    			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "дефолт")
    			//msg.ReplyToMessageID = update.Message.MessageID
    			//msg.ReplyMarkup = chat.GetPhoneKeyboard()
    			bot.Send(msg)
    			return
    		}



  }

  if update.Message.Text == "/start" {
  		msg := tgbotapi.NewMessage(update.Message.Chat.ID, chat.BOT_WELCOME_MESSAGE)
  		msg.ReplyMarkup = chat.GetPhoneKeyboard()
  		bot.Send(msg)
  		return
  	}
  if update.Message.Text == "/ep" {
    msg := tgbotapi.NewMessage(update.Message.Chat.ID, chat.BOT_WELCOME_MESSAGE)
    msg.ReplyMarkup = chat.GetPhoneKeyboard()
    bot.Send(msg)
    return
  }


  msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ты кто?")
  //msg.ReplyToMessageID = update.Message.MessageID
  msg.ReplyMarkup = chat.GetPhoneKeyboard()
  bot.Send(msg)
 }
