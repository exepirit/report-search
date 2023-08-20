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
      <b dangerouslySetInnerHTML={{__html: report.subjectName}}/>
      <div>Written by&nbsp;
        <span dangerouslySetInnerHTML={{__html: report.author.shortName}} />&nbsp;
        in period {formatDate(report.period.startDate)} - {formatDate(report.period.finishDate)}
      </div>
    </Card.Header>
    <Card.Body>
      {report.parts.map(part => (
        <p dangerouslySetInnerHTML={{__html: part.content}}/>
      ))}
    </Card.Body>
  </Card>
};