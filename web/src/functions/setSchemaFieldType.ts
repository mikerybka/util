import API_ENDPOINT from "../constants/API_ENDPOINT"

export default async function setSchemaFieldType(schemaID: string, fieldID: string, fieldType: string) {
    const endpoint = `${API_ENDPOINT}/schemas/${schemaID}/fields/${fieldID}/type`;
    const res = await fetch(endpoint, {
        method: "PUT",
        body: JSON.stringify(fieldType),
    });
    if (!res.ok) {
        const body = await res.text();
        const err = `${res.status}: ${body}`;
        throw new Error(err);
    }
    return res.json();
}
