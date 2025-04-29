package db

import (
	"Enrollment/internal/config"
	"Enrollment/internal/domain/entities"
	"log"
)

type SQLiteClassroomRepository struct {
	*config.Database
}

func NewSQLiteClassroomRepository (Db *config.Database) (*SQLiteClassroomRepository) {
	repo := &SQLiteClassroomRepository{Database: Db}
	return repo
}

func (r *SQLiteClassroomRepository)  Save(classroom entities.ClassRoom) {
	_, err := r.Db.Exec(`INSERT INTO classroom (id) VALUES (?)`, classroom.ClassId)
	if err != nil {
		log.Fatal(err)
	}
}

func (r *SQLiteClassroomRepository) Count() (int) {
	var count int
	err := r.Db.QueryRow(`SELECT COUNT(*) FROM classroom`).Scan(&count)

	if err != nil {
		log.Fatal("Erro ao pegar a quantidade ",err)
	}

	return count
}
