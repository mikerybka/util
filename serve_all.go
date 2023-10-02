package util

func ServeAll(handlers map[string]http.Handler) error {
	errs := make(chan error)
	for addr, h := range handlers {
		go serve(errs, addr, h)
	}
	return <-errs
}

func serve(errs chan error, addr string, h http.Handler) {
	errs <- http.ListenAndServe(addr, h)
}
