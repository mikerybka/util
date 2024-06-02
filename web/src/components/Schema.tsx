import React, { useEffect, useState } from "react";
import Schema from "../types/Schema";
import Loading from "./Loading";

export default function Schema(props: {
    id: string;
}) {
    const [data, setData] = useState<Schema|null>(null);
    useEffect(() => {}, [props.id]);
    if (!data) {
        return <Loading />
    }
    
}