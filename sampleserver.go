package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func getMyIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "unknown"
}

func sayhello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from %s!", getMyIP())
}

func main() {
	http.HandleFunc("/", sayhello)           // set router
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
