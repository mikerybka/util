import API_ENDPOINT from "../constants/API_ENDPOINT"

export default async function login(phone: string, code: string) {
    const endpoint = `${API_ENDPOINT}/auth/login`;
    const res = await fetch(endpoint, {
        method: "POST",
        body: JSON.stringify({
            phone,
            code,
        }),
    });
    if (!res.ok) {
        const body = await res.text();
        const err = `${res.status}: ${body}`;
        throw new Error(err);
    }
    return res.json();
}
