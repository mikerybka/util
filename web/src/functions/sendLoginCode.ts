import API_ENDPOINT from "../constants/API_ENDPOINT"

export default async function sendLoginCode(phone: string) {
    const endpoint = `${API_ENDPOINT}/auth/send-login-code`;
    const res = await fetch(endpoint, {
        method: "POST",
        body: JSON.stringify({
            phone,
        }),
    });
    if (!res.ok) {
        const body = await res.text();
        const err = `${res.status}: ${body}`;
        throw new Error(err);
    }
    return res.json();
}
