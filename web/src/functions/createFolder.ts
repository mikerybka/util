import API_ENDPOINT from "../constants/API_ENDPOINT"

export default async function createFolder(path: string) {
    const endpoint = `${API_ENDPOINT}/schemas/${path}`;
    const res = await fetch(endpoint, {
        method: "PUT",
        body: JSON.stringify({
            isFolder: true,
        }),
    });
    if (!res.ok) {
        const body = await res.text();
        const err = `${res.status}: ${body}`;
        throw new Error(err);
    }
    return res.json();
}
