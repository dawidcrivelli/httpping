package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func doPing(address string, done chan<- float64, index <-chan int) {
	ind := <-index
	startTime := time.Now()
	resp, err := http.Get(address)
	if err != nil {
		fmt.Println(err)
		done <- -1
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	timeTaken := time.Since(startTime).Seconds() * 1000

	fmt.Printf("(#%4d) Reply [%s]: %d bytes, %.0f ms\n", ind, resp.Status, len(body), timeTaken)
	done <- timeTaken
}

func main() {
	const defaultAddress = "https://api.kontakt.io/healthcheck"
	address := flag.String("address", defaultAddress, "URL resource to ping")
	count := flag.Int("n", 10, "Repetitions")
	workers := flag.Int("workers", 10, "Concurrent workers")
	flag.Parse()

	done := make(chan float64, *count)
	index := make(chan int, *workers)

	// warm up
	fmt.Println("Going to ping", *address)
	index <- 0
	doPing(*address, make(chan float64, 1), index)

	startTime := time.Now()
	defer func() {
		fmt.Printf("Wall-clock time taken: %.3f s\n", time.Since(startTime).Seconds())
	}()

	workUnit := 0
	for ; workUnit < *workers; workUnit++ {
		index <- workUnit
	}
	for i := 0; i < *count; i++ {
		go doPing(*address, done, index)
	}

	tot, n, fails := 0.0, 0.0, 0.0
	for i := 0; i < *count; i++ {
		val := <-done
		index <- workUnit
		workUnit++
		if val >= 0 {
			n++
			tot += val
		} else {
			fails++
		}
	}
	fmt.Printf("Average time taken: %.3fms, %.0f queries failed\n", tot/n, fails)
}
