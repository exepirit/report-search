import {Report} from "../Models/report";
import axios from 'axios';

export const searchReports = (text: string): Promise<Report[]> =>
  axios.get<Report[]>(`/api/v1/search?text=${text}`)
    .then(result => result.data);