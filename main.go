package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"email_sender/config"
	"email_sender/db"
	"email_sender/email"

	"github.com/joho/godotenv"
)

func main() {
	// Загрузка переменных окружения из .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("❌ Не удалось загрузить .env файл (он должен быть рядом с main.go)")
	}
	config.Init() // 👈 ВАЖНО: подгружаем переменные

	// Проверка SMTP-переменных
	if config.FromEmail == "" || config.SMTPUser == "" || config.SMTPPass == "" {
		log.Fatal("❌ SMTP-настройки не заданы. Проверь .env или config/config.go")
	}

	// Настройка логирования в файл
	logFile, err := os.OpenFile("logs/email_sender.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("❌ Не удалось открыть лог-файл: %v", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	log.Println("🚀 Запуск рассылки...")

	rand.Seed(time.Now().UnixNano())

	database := db.InitDB()
	defer database.Close()

	entrepreneurs := db.GetPendingEntrepreneurs(10)
	if len(entrepreneurs) == 0 {
		log.Println("Нет подходящих записей для отправки писем.")
		return
	}

	tmpl, err := email.LoadTemplate("templates/base.html")
	if err != nil {
		log.Fatalf("Ошибка загрузки шаблона: %v", err)
	}

	for _, e := range entrepreneurs {
		fullName := fmt.Sprintf("%s %s %s", e.LastName, e.FirstName, e.MiddleName)
		log.Printf("⏳ Подготовка письма для: %s <%s>", fullName, e.Email)

		body, err := email.RenderTemplate(tmpl, fullName)
		if err != nil {
			log.Printf("❌ Ошибка рендера шаблона для ID %d: %v", e.ID, err)
			db.MarkAsError(e.ID, err.Error())
			continue
		}

		err = email.SendEmail(config.FromEmail, e.Email, "ИИ для бизнеса — SmartStack.AI", body)
		if err != nil {
			log.Printf("❌ Ошибка отправки письма ID %d (%s): %v", e.ID, e.Email, err)
			db.MarkAsError(e.ID, err.Error())
			continue
		}

		log.Printf("✅ Письмо успешно отправлено: %s", e.Email)
		db.MarkAsSent(e.ID)

		// Рандомная задержка между 45 и 90 сек
		delay := time.Duration(45+rand.Intn(46)) * time.Second
		log.Printf("🕒 Задержка перед следующей отправкой: %v", delay)
		time.Sleep(delay)
	}

	log.Println("🎉 Рассылка завершена.")
}
