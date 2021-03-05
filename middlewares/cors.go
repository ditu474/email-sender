package middlewares

import "net/http"

// CORS setup the propper headers for CORS
func CORS(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Access-Control-Allow-Origin", "https://www.danielrg.co")
		rw.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")
		handler.ServeHTTP(rw, r)
	})
}
