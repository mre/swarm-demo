package main

import (
	"fmt"
	"github.com/armon/go-metrics"
	"log"
	"net/http"
	"strconv"
	"time"
)

// isPrime checks if a number is prime
func isPrime(x int) bool {
	if x < 2 {
		return false
	}
	for i := 2; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

// checkPrime handles the prime check request
func checkPrime(w http.ResponseWriter, r *http.Request) {
	defer metrics.MeasureSince([]string{"runtime"}, time.Now())
	metrics.IncrCounter([]string{"requests"}, 1)
	i, err := strconv.Atoi(r.URL.Path[1:])
	if err != nil {
		fmt.Fprintf(w, "That's not a number. You make server cry :,(")
		return
	}
	fmt.Fprintf(w, "%t", isPrime(i))

}

// checkPrime handles the prime check request
func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func main() {
	// Setup metrics endpoint
	sink, err := metrics.NewStatsdSink("0.0.0.0:8125")
	if err != nil {
		log.Fatal(err)
	}
	metrics.NewGlobal(metrics.DefaultConfig("primeserver"), sink)

	http.HandleFunc("/", checkPrime)        // set router
	http.HandleFunc("/health", health)        // set router
	err = http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
