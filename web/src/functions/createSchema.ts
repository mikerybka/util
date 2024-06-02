import API_ENDPOINT from "../constants/API_ENDPOINT"

export default async function createSchema(path: string, name: string, type: string) {
    const endpoint = `${API_ENDPOINT}/schemas/${path}`;
    const res = await fetch(endpoint, {
        method: "PUT",
        body: JSON.stringify({
            isSchema: true,
            schema: {
                name,
                type,
            }
        }),
    });
    if (!res.ok) {
        const body = await res.text();
        const err = `${res.status}: ${body}`;
        throw new Error(err);
    }
    return res.json();
}
