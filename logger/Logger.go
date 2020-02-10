package logger

import (
	"log"
	"net/http"
	"os"
	"time"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		pid := os.Getpid()

		hostname, _ := os.Hostname()

		log.Printf(
			"%s\t%s\t%s\t%s\t%v\t%s\t%v\t%s\t%v\t%v",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start).String(),
			pid,
			hostname,
			r.Header.Get("User-Agent"),
			r.Host,
			r.RemoteAddr,
			r.UserAgent(),
			// w.Header,
			// r.Response.Body,
		)
	})
}
