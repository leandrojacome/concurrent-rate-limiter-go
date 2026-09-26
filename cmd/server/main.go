package main

import (
	"log"
	"net/http"

	"github.com/leandrojacome/concurrent-rate-limiter-go/adapters"
	"github.com/leandrojacome/concurrent-rate-limiter-go/application"
)

func main() {
	limiter := application.NewTokenBucket(10, 2, adapters.SystemClock{})
	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", adapters.NewRateLimitHandler(limiter)))
}
