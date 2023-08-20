import React from 'react';
import {Report as ReportModel} from "../../Models/report";
import {Row} from "react-bootstrap";
import {Report} from "./Report";

export interface ReportsListProps {
    reports: ReportModel[];
}

export const ReportsList = (props: ReportsListProps) => {
    return <div>
      {props.reports.map(report => (
        <Row>
          <Report report={report} />
        </Row>
      ))}
    </div>
}