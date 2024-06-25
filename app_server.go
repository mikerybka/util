package util

type AppServer struct {
	Apps map[string]App
}

type App struct {
	Meta AppMeta
}

type AppMeta struct {
	Title  string
	Desc   string
	Author string
}
