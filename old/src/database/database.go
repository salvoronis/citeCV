package databaseutils

import (
	"config"
	"context"
	"database/sql"
	"log"
	"models"
	"utils"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
	ctx = context.Background()
)

func init() {
	dbtmp, err := sql.Open("postgres", config.GetDbConnStr())
	if err != nil {
		log.Printf("Can't connect to database %v\n", err)
	}
	db = dbtmp
	log.Println("Connected sucsessfully to database (but don't hope too much, it still can fall :) )")
}

func GetDB() *sql.DB {
	return db
}

func GetClasses() []models.Class {
	result := []models.Class{}
	rows, err := db.QueryContext(ctx, "select c.class_id, c.name from class as c")
	if err != nil {
		log.Printf("Can't get classes %v\n",err)
	}
	defer rows.Close()

	for rows.Next() {
		tmp := models.Class{}
		err := rows.Scan(&tmp.Id, &tmp.Name)
		if err != nil {
			log.Printf("Can't scan row")
			continue
		}
		result = append(result, tmp)
	}
	return result
}

func SaveStudent(student models.Student) error {
	_, err := db.ExecContext(ctx, "insert into student(class, nickname, password, name) values ($1,$2,$3,$4)", student.Class, student.Username, student.Password, student.Name)
	if err != nil {
		log.Printf("Can't insert student to db %v\n",err)
	}
	return err
}

func CheckStudent(nickname, password string) bool {
	var tmp string
	db.QueryRowContext(ctx, "select password from student where nickname = $1", nickname).Scan(&tmp)
	password = utils.GetSHA256(password)
	return tmp == password
}

func GetStudentByNickname(nickname string) models.Student {
	var student models.Student
	db.QueryRowContext(ctx, "select s.student_id, s.nickname, s.name, c.name as class from student as s join class as c on s.class = c.class_id where s.name = $1", nickname).Scan(&student.ID, &student.Username, &student.Name, &student.Class)
	return student
}
