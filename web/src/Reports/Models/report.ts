import {ReportPeriod} from './reportPeriod';

export interface Report {
    id: string;
    subjectId: string;
    subjectName: string;
    author: Author;
    period: ReportPeriod;
    parts: ReportPart[];
}

interface Author {
    id: number;
    shortName: string;
}

interface ReportPart {
    id: string;
    content: string;
}