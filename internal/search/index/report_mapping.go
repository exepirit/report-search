package index

import (
	"fmt"
	"github.com/exepirit/report-search/internal/data"
	"github.com/google/uuid"
	"time"
)

func MapReportToIndex(report data.Report) Report {
	return Report{
		ID:          report.ID.String(),
		SubjectID:   report.SubjectID.String(),
		SubjectName: report.SubjectName,
		Period:      MapReportPeriodToIndex(report.Period),
		Author:      report.Author,
		Parts:       report.Parts,
	}
}

func MapReportPeriodToIndex(reportPeriod data.ReportPeriod) ReportPeriod {
	return ReportPeriod{
		ID:         reportPeriod.ID.String(),
		StartDate:  reportPeriod.StartDate.Unix(),
		FinishDate: reportPeriod.FinishDate.Unix(),
		Deadline:   reportPeriod.Deadline.Unix(),
	}
}

func MapReportFromIndex(report Report) (data.Report, error) {
	reportId, err := uuid.Parse(report.ID)
	if err != nil {
		return data.Report{}, fmt.Errorf("invalid report ID: %w", err)
	}

	subjectId, err := uuid.Parse(report.SubjectID)
	if err != nil {
		return data.Report{}, fmt.Errorf("invalid subject ID: %w", err)
	}

	period, err := MapReportPeriodFromIndex(report.Period)
	if err != nil {
		return data.Report{}, fmt.Errorf("invalid report period: %w", err)
	}

	return data.Report{
		ID:          reportId,
		SubjectID:   subjectId,
		SubjectName: report.SubjectName,
		Period:      period,
		Author:      report.Author,
		Parts:       report.Parts,
	}, nil
}

func MapReportPeriodFromIndex(reportPeriod ReportPeriod) (data.ReportPeriod, error) {
	periodId, err := uuid.Parse(reportPeriod.ID)
	if err != nil {
		return data.ReportPeriod{}, fmt.Errorf("invalid report period ID: %w", err)
	}

	return data.ReportPeriod{
		ID:         periodId,
		StartDate:  time.Unix(reportPeriod.StartDate, 0),
		FinishDate: time.Unix(reportPeriod.FinishDate, 0),
		Deadline:   time.Unix(reportPeriod.Deadline, 0),
	}, nil
}
