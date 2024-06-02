import API_ENDPOINT from "../constants/API_ENDPOINT"

export default async function moveSchemaFieldPos(schemaID: string, fieldIndex: number, newFieldIndex: number) {
    const endpoint = `${API_ENDPOINT}/schemas/${schemaID}/move_field_pos`;
    const res = await fetch(endpoint, {
        method: "POST",
        body: JSON.stringify({
            from: fieldIndex,
            to: newFieldIndex,
        }),
    });
    if (!res.ok) {
        const body = await res.text();
        const err = `${res.status}: ${body}`;
        throw new Error(err);
    }
    return res.json();
}
