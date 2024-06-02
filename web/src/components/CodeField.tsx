import React from "react";

export default function CodeField(props: {
    digits: number,
    id: string,
    name: string,
    value: string,
    onChange: (v: string) => void;
}) {
    return <div>
        <label htmlFor={props.id}>{props.name}:</label>
        <input type="text" id={props.id} value={props.value} onChange={e => props.onChange(e.target.value)} />
    </div>
}
