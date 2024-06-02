import API_ENDPOINT from "../constants/API_ENDPOINT"

export default async function logout(phone: string, token: string) {
    const endpoint = `${API_ENDPOINT}/auth/logout`;
    const res = await fetch(endpoint, {
        method: "POST",
        body: JSON.stringify({
            phone,
            token,
        }),
    });
    if (!res.ok) {
        const body = await res.text();
        const err = `${res.status}: ${body}`;
        throw new Error(err);
    }
    return res.json();
}
