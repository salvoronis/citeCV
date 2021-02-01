package models

type Schedule struct {
	ClassName string `json:"classname"`
	DayOweek string `json:"dayOweek"`
	LessonTime string `json:"lessonTime"`
	Room int `json:"room"`
	Subject string `json:"subject"`
	TeacherLogin string `json:"t_login"`
	TeacherFName string `json:"t_fname"`
	TeacherSName string `json:"t_sname"`
}
