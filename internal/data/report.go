package data

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Report struct {
	ID          uuid.UUID         `json:"id"`
	SubjectID   uuid.UUID         `json:"subjectId"`
	SubjectName string            `json:"subjectName"`
	Period      ReportPeriod      `json:"period"`
	Author      User              `json:"author"`
	Parts       []ReportPart      `json:"-"`
	RawParts    []json.RawMessage `json:"parts"`
}

func (r Report) GetID() string {
	return r.ID.String()
}

func (r *Report) MarshalJSON() ([]byte, error) {
	type report Report // json.Marshall goes infinite loop without it

	for _, part := range r.Parts {
		raw, err := json.Marshal(part)
		if err != nil {
			return nil, err
		}
		r.RawParts = append(r.RawParts, raw)
	}

	return json.Marshal((*report)(r))
}

func (r *Report) UnmarshalJSON(b []byte) error {
	type report Report // json.Unmarshall goes infinite loop without it

	err := json.Unmarshal(b, (*report)(r))
	if err != nil {
		return err
	}

	for _, rawPart := range r.RawParts {
		var basePart BaseReportPart
		err = json.Unmarshal(rawPart, &basePart)
		if err != nil {
			return err
		}

		var part ReportPart
		switch basePart.Type {
		case ReportPartTypeText:
			part = &TextReportPart{}
		case ReportPartTypeImage:
			part = &ImageReportPart{}
		default:
			return fmt.Errorf("unknown type %q", basePart.Type)
		}

		err = json.Unmarshal(rawPart, part)
		if err != nil {
			return err
		}
	}

	return nil
}

type ReportPeriod struct {
	ID         uuid.UUID `json:"id"`
	StartDate  time.Time `json:"startDate"`
	FinishDate time.Time `json:"finishDate"`
	Deadline   time.Time `json:"deadline"`
}

type ReportPartType string

const (
	ReportPartTypeUnknown ReportPartType = ""
	ReportPartTypeText    ReportPartType = "text"
	ReportPartTypeImage   ReportPartType = "image"
)

type ReportPart interface {
	GetType() ReportPartType
}

type BaseReportPart struct {
	ID   uuid.UUID      `json:"id"`
	Type ReportPartType `json:"type"`
}

func (rp BaseReportPart) GetType() ReportPartType {
	return rp.Type
}

type TextReportPart struct {
	BaseReportPart
	Content string `json:"content"`
}

type ImageReportPart struct {
	BaseReportPart
	URL   string `json:"url"`
	Label string `json:"link"`
}
