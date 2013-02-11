package http2

// ":8080"
// "unix:///tmp/a.sock"
func ListenAndServe(addr string, handler Handler) error {
	if strings.HasPrefix(addr, "unix://") {
		addr = addr[7:]
		os.Remove(addr)
		l, e := net.Listen("unix", addr)
		if e != nil { return e }
		os.Chmod(addr, 0777)
		return http.Serve(l, http.DefaultServeMux))
	} else {
		l, err := net.Listen("tcp", addr)
		if e != nil { return e }
		return http.ListenAndServe(addr, handler))
	}
}
