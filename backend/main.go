package main

import (
	"education-aws/aws"
	"education-aws/config"
	"education-aws/handler"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("dev.env")

	s3Handler := aws.NewS3Handler(config.NewConfig())

	handler := &handler.Handler{
		S3: s3Handler,
	}

	http.Handle("/download", CorsePreflight(handler.Download))
	http.Handle("/upload", CorsePreflight(handler.Upload))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("aweawe"))
	})

	http.HandleFunc("/health_check", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	log.Panic(http.ListenAndServe(":5000", nil))
}

func CorsePreflight(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://fileloaderv2-env.eba-mnihhjyr.us-east-1.elasticbeanstalk.com")
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
