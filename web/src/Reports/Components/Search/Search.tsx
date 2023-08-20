import React, {useEffect, useState} from 'react';
import { SearchForm } from './SearchForm';
import {Report} from "../../Models/report";
import {searchReports} from "../../Api";

export interface SearchProps {
    onResultsChange: (reports: Report[]) => void
}

export const Search = ({onResultsChange}: SearchProps) => {
    const [searchText, setSearchText] = useState('');

    const doSearch = () => {
        if (searchText === '') {
          onResultsChange([]);
          return
        }

        searchReports(searchText)
          .then(reports => onResultsChange(reports));
    };

    return <SearchForm
        onClick={() => doSearch()}
        onInput={setSearchText}
        buttonDisabled={searchText === ''}/>
}