import {ReportPeriod} from './reportPeriod';

export interface Report {
    id: string;
    subjectId: string;
    subjectName: string;
    period: ReportPeriod;
    parts: ReportPart[];
}

interface ReportPart {
    id: string;
    content: string;
}