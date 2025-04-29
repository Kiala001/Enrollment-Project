package db

import (
	"Enrollment/internal/config"
	"Enrollment/internal/domain/entities"
	valueobject "Enrollment/internal/domain/value_object"
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

type SQLiteStudentRepository struct {
	*config.Database
}

func NewSQLiteStudentRepository(db *config.Database) (*SQLiteStudentRepository) {
	repo := &SQLiteStudentRepository{Database: db}
	return repo
}

func (r *SQLiteStudentRepository) Save(student entities.Student) {
	_, err := r.Db.Exec(`INSERT INTO student (id, name, email) VALUES (?, ?, ?)`, student.Id, student.Name.String(), student.Email.String())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Estudante inscrito com sucesso!")
}

func (r *SQLiteStudentRepository) Count() int {
	var count int
	err := r.Db.QueryRow(`SELECT COUNT(*) FROM student`).Scan(&count)

	if err != nil {
		log.Fatal("Erro ao pegar a quantidade ",err)
	}

	return count
}

func (r *SQLiteClassroomRepository) GetStudentsNotNotifieds() []entities.Student {
	rows, err := r.Db.Query(`SELECT id, name, email FROM student WHERE notified = 0`)
	if err != nil {
		log.Fatal(err)
	}

	var students []entities.Student

	for rows.Next() {
		var id string
		var nameStr string
		var emailStr string

		err := rows.Scan(&id, &nameStr, &emailStr)
		if err != nil {
			log.Fatal(err)
		}

		name := valueobject.Name{nameStr}
		email := valueobject.Email{emailStr}

		student := entities.Student{
			Id:    id,
			Name:  name,
			Email: email,
		}

		students = append(students, student)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return students

}

func (r *SQLiteStudentRepository) CountStudentWithoutClass() int {
	var count int
	err := r.Db.QueryRow(`SELECT COUNT(*) FROM student WHERE classroom_id IS NULL`).Scan(&count)
	if err != nil  {
		log.Fatal("Erro ao pegar alunos sem turma ",err)
	}

	return count
}

func (r *SQLiteStudentRepository) StudentExists(id string) bool {
	var exists bool

	query := `SELECT COUNT(1) > 0 FROM student WHERE id = ?`
	err := r.Db.QueryRow(query, id).Scan(&exists)
	if err != nil {
		log.Fatal(err)
	}

	return exists
}

func (r *SQLiteStudentRepository) UpdateStudentWithoutClass(classroomID string) {
	_, err := r.Db.Exec(`UPDATE student SET classroom_id = ?, notified = ? WHERE classroom_id IS NULL`, classroomID, 1)
	if err != nil {
		log.Fatal("Erro ao actualizar ",err)
	}
}
	
func (r *SQLiteStudentRepository) GetAllStudents() []entities.Student {
	rows, err := r.Db.Query(`SELECT id, name, email, classroom_id FROM student`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var students []entities.Student

	for rows.Next() {
		var id string
		var nameStr string
		var emailStr string
		var classroomID sql.NullString

		err := rows.Scan(&id, &nameStr, &emailStr, &classroomID)
		if err != nil {
			log.Fatal(err)
		}

		var classroom string
		if classroomID.Valid {
			classroom = classroomID.String
		}

		name := valueobject.Name{nameStr}
		email := valueobject.Email{classroom + "---" + emailStr}

		student := entities.Student{
			Id:    id,
			Name:  name,
			Email: email,
		}

		students = append(students, student)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return students
}

