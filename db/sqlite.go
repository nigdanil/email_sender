package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type Entrepreneur struct {
	ID         int
	LastName   string
	FirstName  string
	MiddleName string
	Email      string
}

func InitDB() *sql.DB {
	var err error
	db, err = sql.Open("sqlite3", "./data/potential_customers.sqlite")
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("База данных недоступна: %v", err)
	}

	return db
}

func GetPendingEntrepreneurs(limit int) []Entrepreneur {
	rows, err := db.Query(`
		SELECT id, last_name, first_name, middle_name, email
		FROM individual_entrepreneur
		WHERE email != '' AND send_email = 0
		LIMIT ?`, limit)
	if err != nil {
		log.Printf("Ошибка запроса к базе: %v", err)
		return nil
	}
	defer rows.Close()

	var result []Entrepreneur
	for rows.Next() {
		var e Entrepreneur
		if err := rows.Scan(&e.ID, &e.LastName, &e.FirstName, &e.MiddleName, &e.Email); err != nil {
			log.Printf("Ошибка чтения строки: %v", err)
			continue
		}
		result = append(result, e)
	}
	return result
}

func MarkAsSent(id int) {
	_, err := db.Exec(`
		UPDATE individual_entrepreneur
		SET send_email = 1,
		    last_sent = datetime('now'),
		    error = ''
		WHERE id = ?`, id)
	if err != nil {
		log.Printf("Ошибка обновления записи ID %d: %v", id, err)
	}
}

func MarkAsError(id int, errMsg string) {
	_, err := db.Exec(`
		UPDATE individual_entrepreneur
		SET error = ?
		WHERE id = ?`, errMsg, id)
	if err != nil {
		log.Printf("Ошибка записи ошибки для ID %d: %v", id, err)
	}
}
