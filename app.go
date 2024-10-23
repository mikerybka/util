package util

import (
	"os"
	"path/filepath"
)

type App struct {
	Name        Name
	Description string
	Logo        []byte
	StaticFiles map[string][]byte
	Types       map[string]Type
	Functions   map[string]Function
	Constants   map[string]Value
	Variables   map[string]Value
	CoreType    string
}

func (a *App) GenSrc(dir string) error {
	// Write Dockerfile
	dockerfile := filepath.Join(dir, "Dockerfile")
	err := os.WriteFile(dockerfile, []byte(`FROM alpine:latest

		RUN apk update
		RUN apk add go nodejs npm
		
		COPY . /src
		WORKDIR /src
		RUN npx create-next-app@latest --typescript --eslint --tailwind --src-dir --app --import-alias "@/*" web
		
		# TODO: Write favicon.ico
		# TODO: Write static files
		# TODO: Write src/types
		# TODO: Write src/constants
		# TODO: Write src/functions
		# TODO: Write src/hooks
		# TODO: Write src/components
		# TODO: Write src/app
		
		WORKDIR /src/web
		RUN npm run build
		
		WORKDIR /src
		RUN go build -o /bin/server main.go
		
		ENTRYPOINT ["/bin/server"]
	`), os.ModePerm)
	if err != nil {
		return err
	}

	// Write main.go
	mainfile := filepath.Join(dir, "main.go")
	err = os.WriteFile(mainfile, []byte(`package main

	import (
		"fmt"
		"net/http"
		"net/http/httputil"
		"os"
		"os/exec"
	)
	
	func main() {
		// Start frontend
		fmt.Println("Starting frontend")
		cmd := exec.Command("npm", "run", "start")
		cmd.Dir = "/src/web"
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Start()
		if err != nil {
			panic(err)
		}
		defer func() {
			cmd.Process.Kill()
		}()
	
		http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
			// TODO
		})

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			// Proxy other requests to frontend
			proxy := &httputil.ReverseProxy{
				Rewrite: func(pr *httputil.ProxyRequest) {
					pr.Out.URL.Scheme = "http"
					pr.Out.URL.Host = "localhost:3000"
				},
			}
			proxy.ServeHTTP(w, r)
		})
	
		err = http.ListenAndServe(":8000", nil)
		if err != nil {
			panic(err)
		}
	}
	`), os.ModePerm)
	if err != nil {
		return err
	}

	// Write go.mod
	gomod := filepath.Join(dir, "Dockerfile")
	err = os.WriteFile(gomod, []byte(`module server

		go 1.22.1
	`), os.ModePerm)
	if err != nil {
		return err
	}

	// Generate Next.js frontend
	err = a.NextApp().Write(filepath.Join(dir, "web"))
	if err != nil {
		return err
	}

	return nil
}

// NextApp returns a structure that can be used to generate souce code for a Next.js application frontend.
// The app's name is used as the title on webpages and app labels.
// The description is used in the website metadata for seach engines and link sumarizers.
// The logo gets converted to favicon.ico file and placed in `src/app/favicon.ico`.
// Static files can be made available at any path.
// Custom react pages are copied from pages.
// All pages keys should start with a '/' and are the same path as in Next.js without the final '/page.js'.
// If no home page is defined one will be generated with a 'Coming soon...' message.
// TypeScript types are generated for each type in the applicaiton.
// TypeScript functions are generated for each function and method in the application.
// TypeScript constants are generated for each constant in the app.
// React hooks are generated for each variable in the app.
// React components are generated for each type in the app.
// Next.js pages are generated for auth, navigation, user profile and settings management and org management.
// CoreType determines the main type being worked with with the given application.
// The point of the entire app is simply to maintain lists of this "core type".
// Each org/user has a list of this type that they work on through the app interface.
// A single main layout is generated that includes the nav bar with navigation, search, settings and auth features.
func (a *App) NextApp() *NextApp {
	na := &NextApp{
		StaticFiles: a.StaticFiles,
		Types:       a.Types,
		Constants:   a.Constants,
	}

	// Convert logo to favicon

	// TODO: Add functions

	// TODO: Add hooks

	// TODO: Add components

	// TODO: Add home page

	// TODO: Add layouts

	// TODO: Add auth pages

	// TODO: Add org page

	// TODO: Add app pages

	return na
}
