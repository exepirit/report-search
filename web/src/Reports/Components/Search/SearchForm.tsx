import React from 'react';
import {Button, Form, InputGroup} from "react-bootstrap";

export interface SearchFormProps {
  buttonDisabled: boolean
  onInput: (val: string) => void
  onClick: () => void
}

export const SearchForm = (props: SearchFormProps) => (
    <span>
        <InputGroup>
            <Form.Control type="text"
                          onInput={(event: React.FormEvent<HTMLInputElement>) =>
                            props.onInput(event.currentTarget.value)} />
            <Button variant="primary"
                    disabled={props.buttonDisabled}
                    onClick={() => props.onClick()}>Search</Button>
        </InputGroup>
    </span>
)