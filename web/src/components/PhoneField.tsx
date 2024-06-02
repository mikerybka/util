import React from "react";
import CodeField from "./CodeField";

export default function PhoneField(props: {
    id: string,
    name: string,
    value: string,
    onChange: (v: string) => void;
}) {
    return <CodeField digits={10} {...props} />
}
