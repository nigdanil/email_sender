## 📧 Email Sender — Автоматическая SMTP-рассылка HTML-писем
👉 [Посмотреть проект на GitHub](https://github.com/nigdanil/email_sender)

Проект на Go для отправки персонализированных HTML-писем предпринимателям из SQLite-базы через SMTP (Яндекс или любой другой).

---

### 🧾 Возможности

* Загрузка HTML-шаблона письма
* Подстановка имени и ссылки отписки
* Отправка через SMTP (`gomail.v2`)
* Выборка только нужных записей (`send_email = 0`)
* Логирование в файл (`logs/email_sender.log`)
* Рандомная задержка между письмами
* Отметка отправленных писем в базе
* Поддержка `.env` для конфиденциальных настроек

---

### 📂 Структура проекта

```
email_sender/
├── main.go                 # Точка входа
├── go.mod / go.sum         # Зависимости Go
├── .env                    # Конфиденциальные переменные
├── config/                 # Загрузка SMTP-конфига
│   └── config.go
├── db/                     # Работа с SQLite
│   └── sqlite.go
├── email/                  # SMTP и шаблоны
│   ├── smtp.go
│   └── template.go
├── templates/              # HTML шаблоны
│   ├── base.html
│   └── base_old.html
├── data/                   # База SQLite
│   └── potential_customers.sqlite
├── logs/                   # Файлы логов
│   └── email_sender.log
```

---

### ⚙️ Установка и запуск

#### 1. Установить Go ≥ 1.18

#### 2. Клонировать репозиторий

```bash
git clone https://github.com/yourname/email_sender.git
cd email_sender
```

#### 3. Установить зависимости

```bash
go mod tidy
```

#### 4. Создать `.env` в корне:

```env
SMTP_FROM=info@smart-stack.ru
SMTP_USER=info@smart-stack.ru
SMTP_PASS=your_smtp_app_password
```

#### 5. Создать лог-папку (если нет)

```bash
mkdir -p logs
```

#### 6. Запустить

```bash
go run .
```

---

### 📋 Структура таблицы в SQLite

```sql
CREATE TABLE individual_entrepreneur (
  id INTEGER PRIMARY KEY,
  last_name TEXT,
  first_name TEXT,
  middle_name TEXT,
  email TEXT,
  send_email INTEGER DEFAULT 0,
  last_sent TEXT,
  error TEXT
);
```

---

### 🛠 Зависимости

* [`gopkg.in/gomail.v2`](https://pkg.go.dev/gopkg.in/gomail.v2) — SMTP-отправка
* [`github.com/joho/godotenv`](https://github.com/joho/godotenv) — загрузка `.env`
* [`github.com/mattn/go-sqlite3`](https://github.com/mattn/go-sqlite3) — SQLite

---

## 🚀 Установка и автоматический запуск

### ✅ 1. Убедитесь, что Go установлен

```bash
which go
```

Если не установлен:

```bash
sudo apt update
sudo apt install golang-go
```

---

### ✅ 2. Соберите исполняемый файл

```bash
cd /email_sender
go build -o email_sender
ls -la email_sender
```

---

### ✅ 3. Протестируйте вручную

```bash
./email_sender
```

Проверьте, что в `logs/email_sender.log` появились записи, и рассылка работает.

---

### ✅ 4. Добавьте в cron запуск по будням в 08:00 по МСК

Откройте crontab:

```bash
crontab -e
```

Добавьте строку:

```
0 8 * * 1-5 /email_sender/run.sh
```

#### 📌 Расшифровка:

| Поле                   | Значение              |
| ---------------------- | --------------------- |
| `0`                    | Минута (00)           |
| `8`                    | Час (08:00)           |
| `*`                    | Каждый день           |
| `*`                    | Каждый месяц          |
| `1-5`                  | Понедельник – Пятница |
| `/email_sender/run.sh` | Скрипт запуска        |

Проверь:

```bash
crontab -l
```

---

### ✅ 5. Скрипт запуска `run.sh`

```bash
#!/bin/bash
cd "$(dirname "$0")"
echo "🚀 Запуск рассылки: $(date)" >> logs/email_sender.log
./email_sender >> logs/email_sender.log 2>&1
```

---

### ✅ 6. Скрипт остановки `stop.sh`

```bash
#!/bin/bash
echo "⛔ Остановка email_sender..."
pkill -f "./email_sender"
```

> Завершит любой запущенный процесс email\_sender

---

### ✅ 7. Сделайте скрипты исполняемыми

```bash
chmod +x run.sh stop.sh
```

---

### ✅ Структура проекта на сервере

```
email_sender/
├── run.sh
├── stop.sh
├── email_sender          # бинарник
├── main.go
├── config/
├── db/
├── email/
├── templates/
├── data/
│   └── potential_customers.sqlite
├── logs/
│   └── email_sender.log
└── .env
```

---

### 📜 Пример логов (logs/email_sender.log)

🚀 Запуск рассылки: 2025-05-15 08:00:00
⏳ Подготовка письма для: Иванов Иван Иванович test@example.com
✅ Письмо успешно отправлено: test@example.com
🕒 Задержка перед следующей отправкой: 1h22m0s
🎉 Рассылка завершена.

### 🛡 Рекомендации

* Убедитесь, что `.env` и лог-файлы доступны из того же каталога
* Если `cron` не видит `.env`, указывайте абсолютные пути
* Не забывайте про логирование и авто-рестарт в случае падения (через `systemd`, если потребуется)
* Для продакшена можно заменить cron на systemd-сервис (например, для мониторинга, автоматического рестарта и логов)
