import React, { useEffect, useState } from "react";
import Type from "../types/Type";
import Session from "../types/Session";
import StringView from "./StringView";
import BoolView from "./BoolView";
import IntView from "./IntView";

export default function Entity(props: {
    jsonURL: string;
    type: Type;
}) {
    const [loading, setLoading] = useState(false);
    const [err, setErr] = useState("");
    const [data, setData] = useState<any>(null);

    const session: Session = JSON.parse(localStorage.getItem('session'));

    useEffect(() => {
        setLoading(true);
        fetch(props.jsonURL, {
            headers: {
                Authorization: "Token "+session.Token,
            },
        }).then((res) => {
            if (res.ok) {
                res.json().then(d => {
                    setData(d);
                });
            } else {
                res.text().then(t => {
                    setErr(t)
                });
            }
        }).catch(e => {
            setErr(e.message);
        }).finally(() => {
            setLoading(false);
        })
    }, [props.jsonURL]);

    if (err !== "") {
        return <div>{err}</div>
    }

    if (props.type.IsScalar) {
        switch (props.type.Kind) {
            case "string":
                return  <StringView value={data} onChange={setData} />
            case "bool":
                return <BoolView value={data} onChange={setData} />
            case "int":
                return <IntView value={data} onChange={setData} />
            default:
                throw new Error("unknown type '"+props.type.Kind+"'")
        }
    }

    if (props.type.IsPointer) {
        throw new Error("pointers not yet implemented")
    }

    if (props.type.IsArray) {
        throw new Error("arrays not implemented")
    }
    
    if (props.type.IsMap) {
        throw new Error("arrays not implemented")
    }

    if (props.type.IsStruct) {
        <div>
            {
                props.type.Fields.map(f => {
                    const url = window.location + "/" + f.ID;
                    return <div><a href={url}>{f.Name}</a></div>
                })
            }
        </div>
    }

    return <div></div>
}