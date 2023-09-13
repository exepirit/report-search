package index

import (
	"encoding/json"
	"fmt"
	"github.com/exepirit/report-search/internal/data"
)

type Report struct {
	ID          string            `json:"id"`
	SubjectID   string            `json:"subjectId"`
	SubjectName string            `json:"subjectName"`
	Period      ReportPeriod      `json:"period"`
	Author      data.User         `json:"author"`
	Parts       []data.ReportPart `json:"-"`
	RawParts    []json.RawMessage `json:"parts"`
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
		var basePart data.BaseReportPart
		err = json.Unmarshal(rawPart, &basePart)
		if err != nil {
			return err
		}

		var part data.ReportPart
		switch basePart.Type {
		case data.ReportPartTypeText:
			part = &data.TextReportPart{}
		case data.ReportPartTypeImage:
			part = &data.ImageReportPart{}
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
	ID         string `json:"id"`
	StartDate  int64  `json:"startDate"`
	FinishDate int64  `json:"finishDate"`
	Deadline   int64  `json:"deadline"`
}
