package middlewares

import "net/http"

// CORSMiddleware setup the propper headers for CORS
func CORSMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Access-Control-Allow-Origin", "https://danielrg.co")
		rw.Header().Set("Access-Control-Allow-Methods", "POST")
		rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")
		handler.ServeHTTP(rw, r)
	})
}
