## Simple HttpPing reimplementation

A test program, mostly to learn Go, but also to write a simple concurrent pinger, that returns the HTTP status and the average time taken by the requests
```
Usage of ./httpping:
  -address string
    	URL resource to ping (default "https://api.kontakt.io/healthcheck")
  -n int
    	Repetitions (default 10)
  -workers int
    	Concurrent workers (default 10)
```