package config

import (
	"database/sql"
	"log"
)

type Database struct {
	Db *sql.DB
}

func NewDatabase () *Database {
	db, err := sql.Open("sqlite", "enrollment.db")

    if err != nil {
        log.Fatal("Erro com a conexão com a base de dados: ",err)
    }
	
	database := &Database{Db: db}
	database.CreateTableClassroom()
	database.CreateTableStudent()
	database.CreateTablePaymentNote()

	return database
}

func (r *Database) CreateTableClassroom() error {
	_, err := r.Db.Exec(`
		CREATE TABLE IF NOT EXISTS classroom (
			id TEXT PRIMARY KEY
		)
	`)
	return err
}

func (r *Database) CreateTableStudent() error {
	_, err := r.Db.Exec(`
        CREATE TABLE IF NOT EXISTS student (
            id TEXT PRIMARY KEY,
            name TEXT,
            email TEXT,
			notified BOOLEAN DEFAULT 0,
			classroom_id TEXT,
    		FOREIGN KEY (classroom_id) REFERENCES classroom(id)
        )
    `)
	return err
}

func (r *Database) CreateTablePaymentNote() {
	_, err := r.Db.Exec(`
		CREATE TABLE IF NOT EXISTS payment_note (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			student_id TEXT,
			price INTEGER,
			date TEXT,   		
			FOREIGN KEY (student_id) REFERENCES student(id)
		)
	`)
	if err != nil {
		log.Fatal("Erro ao criar a tabela: ",err)
	}
}

