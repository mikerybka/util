package util

func LatestGoVersion() string {
	v, err := GetLatestGoVersion()
	if err != nil {
		panic(err)
	}
	return v
}
