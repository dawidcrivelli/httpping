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

Running with the defaults, `go run httpping` results in:

```
Going to ping https://api.kontakt.io/healthcheck
(#  -1) Reply [200 OK]: 2 bytes, 398 ms
(#   0) Reply [200 OK]: 2 bytes, 68 ms
(#   2) Reply [200 OK]: 2 bytes, 70 ms
(#   1) Reply [200 OK]: 2 bytes, 71 ms
(#   4) Reply [200 OK]: 2 bytes, 71 ms
(#   6) Reply [200 OK]: 2 bytes, 72 ms
(#   9) Reply [200 OK]: 2 bytes, 134 ms
(#   8) Reply [200 OK]: 2 bytes, 134 ms
(#   5) Reply [200 OK]: 2 bytes, 134 ms
(#   3) Reply [200 OK]: 2 bytes, 134 ms
(#   7) Reply [200 OK]: 2 bytes, 134 ms
Average time taken: 102.260ms, 0 queries failed
Wall-clock time taken: 0.135 s
```
