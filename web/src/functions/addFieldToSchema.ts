import API_ENDPOINT from "../constants/API_ENDPOINT"

export default async function addFieldToSchema(schemaID: string, fieldID: string, fieldName: string, fieldType: string) {
    const endpoint = `${API_ENDPOINT}/schemas/${schemaID}/add_field`;
    const res = await fetch(endpoint, {
        method: "POST",
        body: JSON.stringify({
            id: fieldID,
            name: fieldName,
            type: fieldType,
        }),
    });
    if (!res.ok) {
        const body = await res.text();
        const err = `${res.status}: ${body}`;
        throw new Error(err);
    }
    return res.json();
}
