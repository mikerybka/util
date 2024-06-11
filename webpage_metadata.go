package util

type WebpageMetadata struct {
	Title    string
	Desc     string
	Author   *Person
	Keywords []string
	Favicon  []byte
}
