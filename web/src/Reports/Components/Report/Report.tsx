import React from 'react';
import {Report as ReportModel} from "../../Models/report";
import {Card} from "react-bootstrap";
import moment from 'moment';

export interface ReportProps {
  report: ReportModel;
}

export const Report = ({report}: ReportProps) => {
  const formatDate = (dateStr: string): string =>
    moment(dateStr).format('DD.MM.YYYY');

  return <Card>
    <Card.Header>
      <span dangerouslySetInnerHTML={{__html: report.subjectName}}/>
      <span>({formatDate(report.period.startDate)} - {formatDate(report.period.finishDate)})</span>
    </Card.Header>
    <Card.Body>
      {report.parts.map(part => (
        <p dangerouslySetInnerHTML={{__html: part.content}}/>
      ))}
    </Card.Body>
  </Card>
};