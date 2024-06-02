import API_ENDPOINT from "../constants/API_ENDPOINT"

export default async function removeFieldFromSchema(schemaID: string, fieldID: string) {
    const endpoint = `${API_ENDPOINT}/schemas/${schemaID}/fields/${fieldID}`;
    const res = await fetch(endpoint, {
        method: "DELETE",
    });
    if (!res.ok) {
        const body = await res.text();
        const err = `${res.status}: ${body}`;
        throw new Error(err);
    }
    return res.json();
}
