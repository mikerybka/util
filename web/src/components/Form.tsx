import React, { useState } from "react";
import Button from "./Button";

export default function Form(props: {
    onSubmit: () => Promise<any>,
    onSuccess: (res: any) => void,
    children: React.ReactNode,
}) {
    const [err, setErr] = useState('');
    const onError = (e: string) => {
        setErr(e);
    }
    const onSuccess = async (res: any) => {
        setErr("");
        props.onSuccess(res);
    }
    return <div>
        {props.children}
        {err && <p className="text-red">{err}</p>}
        <Button text="Submit" onClick={props.onSubmit} onSuccess={onSuccess} onError={onError} />
    </div>
}
