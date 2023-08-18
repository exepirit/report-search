package data

import (
	"github.com/google/uuid"
	"time"
)

type Report struct {
	ID          uuid.UUID    `json:"id"`
	SubjectID   uuid.UUID    `json:"subjectId"`
	SubjectName string       `json:"subjectName"`
	Period      ReportPeriod `json:"period"`
	Author      User         `json:"author"`
	Parts       []ReportPart `json:"parts"`
}

type ReportPeriod struct {
	ID         uuid.UUID `json:"id"`
	StartDate  time.Time `json:"startDate"`
	FinishDate time.Time `json:"finishDate"`
	Deadline   time.Time `json:"deadline"`
}

type ReportPart struct {
	ID      uuid.UUID `json:"id"`
	Content string    `json:"content"`
}
