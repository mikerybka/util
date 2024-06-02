import React from "react";

export default function IntView(props : {
    value: number;
    onChange: (n: number) => void;
}) {
    return <input type="number" value={props.value} onChange={e => props.onChange(JSON.parse(e.target.value))} />
}
