package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	// s3Handler := aws.NewS3Handler(config.NewConfig())

	// handler := &handler.Handler{
	// 	S3: s3Handler,
	// }

	// http.Handle("/download", CorsePreflight(handler.Download))
	// http.Handle("/upload", CorsePreflight(handler.Upload))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("aweawe"))
	})

	http.HandleFunc("/health_check", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	fmt.Println("start lintening on port 5000")

	log.Panic(http.ListenAndServe(":5000", nil))
}

func CorsePreflight(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, *")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
