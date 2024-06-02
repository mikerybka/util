import React from "react";

export default function BoolView(props: {
    value: boolean;
    onChange: (v: boolean) => void;
}) {
    return <input
        type="checkbox"
        value={JSON.stringify(props.value)}
        onChange={e => props.onChange(e.target.value ? true : false)}
    />
}
