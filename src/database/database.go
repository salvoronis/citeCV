package databaseutils

import (
	"config"
	"context"
	"database/sql"
	"log"
	"models"
	"time"
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
	log.Println("Connected sucsessfully to database (but don't hope too much, it still can fall )")
}

//func GetDB() *sql.DB {
//	return db
//}

func CheckUser(login, password string) (bool, uint) {
	var tmp string
	var id uint
	db.QueryRowContext(ctx, "select password, userId from member where login = $1", login).
		Scan(&tmp,&id)
	password = utils.GetSHA256(password)
	return tmp == password, id
}

func SaveUser(user models.User) error {
	_, err := db.
	ExecContext(
		ctx,
		"insert into member(login,password,firstname,secondname,email) values ($1,$2,$3,$4,$5)",
		user.Login,
		utils.GetSHA256(user.Password),
		user.FirstName,
		user.SecondName,
		user.Email,
	)

	usr, err := GetUserByLogin(user.Login)
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = db.
	ExecContext(
		ctx,
		"insert into class_user(student_id, class_id) values ($1, $2)",
		usr.Id,
		user.Class,
	)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func GetUserByLogin(login string) (models.User,error) {
	var user models.User
	db.QueryRowContext(ctx, "select * from member where login = $1", login).
		Scan(
			&user.Id,
			&user.Login,
			&user.Password,
			&user.FirstName,
			&user.SecondName,
			&user.Email,
		)
	return user, nil
}

func GetClasses() []models.Class {
	result := []models.Class{}
	rows, err := db.QueryContext(ctx, "select c.classId, c.name from class as c")
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

func GetMarks(studentId int, from time.Time, to time.Time) []models.Marks {
	result := []models.Marks{}
	rows, err := db.QueryContext(ctx, "select * from get_schedule_marks($1, $2, $3)",
		studentId,
		from,
		to,
	)
	if err != nil {
		log.Printf("Can't get marks %v\n", err)
	}
	defer rows.Close()

	for rows.Next() {
		tmp := models.Marks{}
		err := rows.Scan(
			&tmp.DayOweek,
			&tmp.LessonTime,
			&tmp.Room,
			&tmp.Subject,
			&tmp.Mark,
		)
		if err != nil {
			log.Printf("Cant scan rows")
			continue
		}
		result = append(result, tmp)
	}
	return result
}

func GetSchedule(classId int) []models.Schedule {
	result := []models.Schedule{}
	rows, err := db.QueryContext(ctx, "select * from get_schedule_by_class( $1 )",
		classId,
	)
	if err != nil {
		log.Printf("Can't get schedule %v\n", err)
	}
	defer rows.Close()

	for rows.Next() {
		tmp := models.Schedule{}
		err := rows.Scan(
			&tmp.ClassName,
			&tmp.DayOweek,
			&tmp.LessonTime,
			&tmp.Room,
			&tmp.Subject,
			&tmp.TeacherLogin,
			&tmp.TeacherFName,
			&tmp.TeacherSName,
		)
		if err != nil {
			log.Printf("Cant scan rows schedule")
			continue
		}
		result = append(result, tmp)
	}
	return result
}
