package index

type Report struct {
	ID          string       `json:"id"`
	SubjectID   string       `json:"subjectId"`
	SubjectName string       `json:"subjectName"`
	Period      ReportPeriod `json:"period"`
	Author      User         `json:"author"`
	Parts       []ReportPart `json:"parts"`
}

type ReportPeriod struct {
	ID         string `json:"id"`
	StartDate  int64  `json:"startDate"`
	FinishDate int64  `json:"finishDate"`
	Deadline   int64  `json:"deadline"`
}

type ReportPart struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}
