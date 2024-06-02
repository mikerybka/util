import API_ENDPOINT from "../constants/API_ENDPOINT"

export default async function move(fromPath: string, toPath: string) {
    const endpoint = `${API_ENDPOINT}/schemas/move`;
    const res = await fetch(endpoint, {
        method: "POST",
        body: JSON.stringify({
            from: fromPath,
            to: toPath,
        }),
    });
    if (!res.ok) {
        const body = await res.text();
        const err = `${res.status}: ${body}`;
        throw new Error(err);
    }
    return res.json();
}
