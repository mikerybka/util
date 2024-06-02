import React,{ useState } from "react";

export default function Button(props: {
    text: string,
    onClick: () => Promise<any>,
    onSuccess: (res: any) => void,
    onError: (err: string) => void,
    primary?: boolean,
}) {
    const [loading, setLoading] = useState(false);
    const click = () => {
        setLoading(true);
        props.onClick().then((res) => {
            props.onSuccess(res);
        }).catch(e => {
            props.onError(e.message);
        }).finally(() => {
            setLoading(false);
        })
    }
    return <button
        disabled={loading}
        onClick={click}
    >
        {loading ? "Loading..." : props.text}
    </button>
}
