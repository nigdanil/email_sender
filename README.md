## 📧 Email Sender — Автоматическая SMTP-рассылка HTML-писем

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

### 💡 Полезные советы

* Рассылку можно запускать по расписанию (`cron`, `systemd`, Task Scheduler)
* Можно подключить генератор токенов для персональной ссылки отписки
* Используйте отдельный SMTP-пароль приложения (а не основной)
