import Field from "./Field";

export default interface Schema {
    ID: string;
    Name: string;
    Fields: Field[];
}