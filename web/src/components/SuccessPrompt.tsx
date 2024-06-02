import React from 'react';
import Action from '../types/Action';
import Card from './Card';
import SuccessIcon from './SuccessIcon';
import Button from './Button';

export default function SuccessPrompt(props: {
    message: string;
    actions: Action[];
}) {
    return <Card>
        <SuccessIcon />
        <p>{props.message}</p>
        {props.actions.map(a => {
            return <Button
                primary={a.primary}
                text={a.text}
                onClick={a.onClick}
                onSuccess={() => { }}
                onError={() => { }}
            />
        })}
    </Card>
}
