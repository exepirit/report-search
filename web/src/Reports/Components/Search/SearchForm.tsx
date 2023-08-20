import React from 'react';
import {Button, Form, InputGroup} from "react-bootstrap";

export const SearchForm = () => (
    <span>
        <InputGroup>
            <Form.Control type="text"></Form.Control>
            <Button variant="primary">Search</Button>
        </InputGroup>
    </span>
)