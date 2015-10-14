package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// PrimeServerResponse contains the data received from the server
type PrimeServerResponse struct {
	url      string
	number   int
	response *http.Response
	err      error
}

// NaturalNumber contains a number and a boolean that says if that number is prime
type NaturalNumber struct {
	number  int
	isPrime bool
}

func requestNumbers(server string, first, last int) <-chan *PrimeServerResponse {
	count := last - first
	ch := make(chan *PrimeServerResponse, count) // buffered

	for i := first; i < last; i++ {
		url := server + "/" + strconv.Itoa(i)
		go func(url string, i int) {
			resp, err := http.Get(url)
			ch <- &PrimeServerResponse{url, i, resp, err}
		}(url, i)
	}

	return ch
}

func extractNumber(r *PrimeServerResponse) (*NaturalNumber, error) {
	contents, err := ioutil.ReadAll(r.response.Body)
	if err != nil {
		return nil, err
	}
	// Convert []byte to string
	isPrime, err := strconv.ParseBool(string(contents[:]))
	if err != nil {
		return nil, err
	}
	return &NaturalNumber{r.number, isPrime}, nil
}

func receiveNumbers(ch <-chan *PrimeServerResponse, first, last int) []NaturalNumber {
	count := last - first
	numbers := []NaturalNumber{}
	received := 0
	for {
		select {
		case r := <-ch:
			received++
			// Do something with result
			if r.err != nil {
				fmt.Printf("Oops! %s\n", r.err)
				continue
			}
			n, err := extractNumber(r)
			if err != nil {
				fmt.Printf("Oops! %s\n", err)
				continue
			}
			numbers = append(numbers, *n)
			if received == count {
				return numbers
			}
		case <-time.After(50 * time.Millisecond):
			// Wait a bit for the server response
			fmt.Printf(".")
		}
	}
}

func checkPrime(server string, first, last, batchSize int) {
	for i := first; i <= last; i += batchSize {
		ch := requestNumbers(server, i, i+batchSize)
		for _, number := range receiveNumbers(ch, i, i+batchSize) {
			fmt.Printf("%d: %t\n", number.number, number.isPrime)
		}
		fmt.Printf("From %d to %d\n", i, i+batchSize)
	}
}

func main() {
	checkPrime("http://localhost:9090", 1000000000, 10000000000, 10)
}
