package models

type Marks struct {
	DayOweek string `json:"dayOweek"`
	LessonTime string `json:"les_time"`
	Room int `json:"room"`
	Subject string `json:"subject"`
	Mark int `json:"mark"`
}
