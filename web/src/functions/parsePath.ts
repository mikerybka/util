export default function parsePath(path: string):string[] {
    return path.split('/').filter(Boolean);
}
