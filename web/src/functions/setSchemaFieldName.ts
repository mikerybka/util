import API_ENDPOINT from "../constants/API_ENDPOINT"

export default async function setSchemaFieldName(schemaID: string, fieldID: string, newName: string) {
    const endpoint = `${API_ENDPOINT}/schemas/${schemaID}/fields/${fieldID}/name`;
    const res = await fetch(endpoint, {
        method: "PUT",
        body: JSON.stringify(newName),
    });
    if (!res.ok) {
        const body = await res.text();
        const err = `${res.status}: ${body}`;
        throw new Error(err);
    }
    return res.json();
}
