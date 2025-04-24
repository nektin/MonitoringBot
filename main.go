package main

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"time"
)

var siteStatus = make(map[string]int)

func main() {
	file, err := os.OpenFile("monitor.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Неудалось открыть файл с логами %v", err)
	}
	log.SetOutput(file)

	bot, err := tgbotapi.NewBotAPI(getBotToken())
	if err != nil {
		log.Fatalf("Ошибка при создании бота: %v", err)
	}
	bot.Debug = true
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	updates := bot.GetUpdatesChan(updateConfig)

	timer := time.NewTimer(time.Duration(getCheckInterval()) * time.Second)
	defer timer.Stop()
	for {
		select {
		case <-timer.C:
			newTicker := time.NewTimer(time.Duration(getCheckInterval()) * time.Second)
			if timer != newTicker {
				timer.Reset(time.Duration(getCheckInterval()) * time.Second)
			}
			newStatus := RequestToSite()
			ProcessStatusChanges(bot, newStatus)

		case update := <-updates:
			if update.Message != nil {
				handleMessage(bot, update.Message)
			}
		}

	}

}
func handleMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	log.Printf("Получено сообщение: %+v", message)
	switch message.Text {
	case "/start":
		log.Println("Обработка команды /start")

		msg := tgbotapi.NewMessage(message.Chat.ID, "Выберите действие:")
		msg.ReplyMarkup = createPersistentMenu()
		if _, err := bot.Send(msg); err != nil {
			log.Printf("Ошибка при отправке сообщения: %v", err)
		} else {
			log.Println("Приветственное сообщение с постоянным меню успешно отправлено")
		}
	default:
		if message.Text == "Показать статус сайтов" {
			log.Println("Обработка запроса на показ статуса сайтов")
			statusText := "Статус сайтов:\n"
			for url, status := range RequestToSite() {
				statusText += fmt.Sprintf("- %s: %d\n", url, status)
			}

			msg := tgbotapi.NewMessage(message.Chat.ID, statusText)
			if _, err := bot.Send(msg); err != nil {
				log.Printf("Ошибка при отправке сообщения: %v", err)
			} else {
				log.Println("Сообщение со статусом сайтов успешно отправлено")
			}
		} else {
			// Ответ на неизвестные команды
			log.Printf("Получена неизвестная команда: %s", message.Text)
			msg := tgbotapi.NewMessage(message.Chat.ID, "Неизвестная команда. Используйте /start.")
			if _, err := bot.Send(msg); err != nil {
				log.Printf("Ошибка при отправке сообщения: %v", err)
			} else {
				log.Println("Сообщение о неизвестной команде успешно отправлено")
			}
		}
	}
}
func createPersistentMenu() tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Показать статус сайтов"),
		),
	)
	keyboard.ResizeKeyboard = true
	keyboard.OneTimeKeyboard = false
	return keyboard
}
func ProcessStatusChanges(bot *tgbotapi.BotAPI, newStatus map[string]int) {
	for key, value := range newStatus {
		var message tgbotapi.MessageConfig

		if value != 200 && siteStatus[key] == 200 {
			log.Printf("Сайт %s стал недоступен (Код: %d)", key, value)
			message = tgbotapi.NewMessage(int64(getChatID()), fmt.Sprintf(
				"⚠️ %s Сайт %s недоступен (Код: %d)",
				time.Now().Format("2006-01-02 15:04:05"), key, value,
			))
		}

		if value == 200 && siteStatus[key] != 200 {
			log.Printf("Сайт %s снова доступен", key)
			message = tgbotapi.NewMessage(int64(getChatID()), fmt.Sprintf(
				"✅ %s Сайт %s снова доступен",
				time.Now().Format("2006-01-02 15:04:05"), key,
			))
		}

		if _, exists := siteStatus[key]; !exists {
			if value == 200 {
				log.Printf("Новый сайт %s добавлен и доступен", key)
				message = tgbotapi.NewMessage(int64(getChatID()), fmt.Sprintf(
					"🆕 %s Новый сайт %s добавлен и доступен",
					time.Now().Format("2006-01-02 15:04:05"), key,
				))
			} else {
				log.Printf("Новый сайт %s добавлен, но недоступен (Код: %d)", key, value)
				message = tgbotapi.NewMessage(int64(getChatID()), fmt.Sprintf(
					"⚠️ %s Новый сайт %s добавлен, но недоступен (Код: %d)",
					time.Now().Format("2006-01-02 15:04:05"), key, value,
				))
			}
		}

		if message.Text != "" {
			if _, err := bot.Send(message); err != nil {
				log.Printf("Ошибка при отправке сообщения для сайта %s: %v", key, err)
			} else {
				log.Printf("Сообщение успешно отправлено для сайта %s", key)
			}
		}
	}

	for key := range siteStatus {
		if _, exists := newStatus[key]; !exists {
			log.Printf("Сайт %s удален из мониторинга", key)
			message := tgbotapi.NewMessage(int64(getChatID()), fmt.Sprintf(
				"❌ %s Сайт %s удален из мониторинга",
				time.Now().Format("2006-01-02 15:04:05"), key,
			))

			if _, err := bot.Send(message); err != nil {
				log.Printf("Ошибка при отправке сообщения для удаленного сайта %s: %v", key, err)
			} else {
				log.Printf("Сообщение успешно отправлено для удаленного сайта %s", key)
			}
		}
	}
	siteStatus = newStatus
}
