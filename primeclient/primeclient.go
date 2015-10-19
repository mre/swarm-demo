package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"math/rand"
)

// PrimeServerResponse contains the data received from the server
type PrimeServerResponse struct {
	number   int
	response *http.Response
	err      error
}

// NaturalNumber contains a number and a boolean that says if that number is prime
type NaturalNumber struct {
	number  int
	isPrime bool
}

func requestNumbers(ch chan *PrimeServerResponse, server string, first, last int){
	for i := first; i <= last; i++ {
		url := server + "/" + strconv.Itoa(i)
		resp, err := http.Get(url)
		ch <- &PrimeServerResponse{i, resp, err}
		time.Sleep(time.Duration(rand.Intn(50) + 50) * time.Millisecond)

	}
}

func extractNumber(r *PrimeServerResponse) (*NaturalNumber, error) {
	contents, err := ioutil.ReadAll(r.response.Body)
	defer r.response.Body.Close()
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

func receiveNumbers(ch chan *PrimeServerResponse, last int) {
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
			fmt.Printf("%d: %t\n", n.number, n.isPrime)
			if received == last {
				return
			}
		case <-time.After(300 * time.Millisecond):
			// Wait a bit for the server response
			fmt.Printf("/")
			time.Sleep(100*time.Millisecond)
			fmt.Printf("\r")
			fmt.Printf("\\")
			time.Sleep(100*time.Millisecond)
			fmt.Printf("\r")
		}
	}
}

func checkPrime(server string, first, last, batchSize int) {
	ch := make(chan *PrimeServerResponse, batchSize) // buffered

	go requestNumbers(ch, server, first, last)
	receiveNumbers(ch, last);
}

func main() {

	hostName := flag.String("host", "localhost:80", "The host name of the prime server")

	start := flag.Int("start", 1000000000, "The first number to check for being prime")
	stop := flag.Int("stop", 10000000000, "The last number to check for being prime")
	step := flag.Int("step", 10, "Check primes in batches of this size")

	flag.Parse()

	hostAddress := "http://" + *hostName
	fmt.Printf("Trying to connect to %s\n", hostAddress)

	checkPrime(hostAddress, *start, *stop, *step)
}
