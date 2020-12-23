package databaseutils

import (
	"config"
	"database/sql"
	"log"
	"models"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
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
	rows, err := db.Query("select c.class_id, c.name from class as c")
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
	_, err := db.Exec("insert into student(class, nickname, password, name) values ($1,$2,$3,$4)", student.Class, student.Username, student.Password, student.Name)
	if err != nil {
		log.Printf("Can't insert student to db %v\n",err)
	}
	return err
}
