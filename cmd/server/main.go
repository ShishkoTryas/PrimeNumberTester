package main

import "PrimeNumberTester/pkg/delivery/api/http"

func main() {
	r := http.CreateRouter()
	r.Run()
}
