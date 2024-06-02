import API_ENDPOINT from "../constants/API_ENDPOINT";

export default async function get(path: string) {
    const endpoint = `${API_ENDPOINT}/schemas/${path}`;
    const res = await fetch(endpoint, {
        method: "GET",
        headers: JSON.parse(localStorage.getItem('session')),
    });
    if (!res.ok) {
        const body = await res.text();
        const err = `${res.status}: ${body}`;
        throw new Error(err);
    }
    return res.json();
}
