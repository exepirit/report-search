import React, {useEffect, useState} from 'react';
import {Container, Row} from "react-bootstrap";
import {Search} from "./Search";
import {ReportsList} from "./Report";
import {Report} from "../Models/report";
import {searchReports} from "../Api";

export const ReportsPage = () => {
  const [reports, setReports] = useState<Report[]>([]);
  const [loaded, setLoaded] = useState(false);

  useEffect(() => {
    searchReports('dolores')
      .then(reports => {
        setReports(reports);
        setLoaded(true);
      });
  }, []);

  return <Container>
    <Row>
      <Search/>
    </Row>
    <Row>
      <ReportsList reports={reports} />
    </Row>
  </Container>
}