export default function joinPath(parts: string[]): string  {
    return "/" + parts.join("/");
}