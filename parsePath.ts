export default function parsePath(s: string): string[] {
	var path: string[] = [];
	s.split("/").forEach(part => {
		if (part !== "") {
			path = [...path, part]
		}
 	})
	return path
}
