import React from "react";

export default function StringView(props : {
    value: string;
    onChange: (s: string) => void;
}) {
    return <input type="text" value={props.value} onChange={e => props.onChange(e.target.value)} />
}
