## `gb` a HTTP benchmarking tool written in Go

This repo is an attempt to learn more about concurrency and Go in general.
I've used ChatGPT-4 as an instructor for this project. It gave me suggestion on
code and explained concepts that were unfamiliar to me. 


### Install

Clone the repo and run `go run .` with following options:

```bash
Usage:
  gb [flags]

Flags:
  -c, --concurrent int         number of concurrent request (default 10)
  -d, --data string            request body
  -H, --header strings         Header to pass to request. This flag can be used multiple times
  -h, --help                   help for gb
  -l, --limit int              limit of concurrent requests per second (default 10)
  -m, --method string          request method [GET, POST, PUT, PATCH, DELETE] (default "GET")
  -n, --number int             number of request (default 10)
  -O, --output-file string     output filename format [json, html, csv] (default "gb")
  -o, --output-format string   output format [json, html, csv] (default "std")
  -T, --timeout int            Timeout for each requests (in seconds) (default 30)
  -u, --url string             URL to benchmark

```

example:

```bash
$ go run . --url https://httpstat.us/Random -n 20 -c 10 -H "Foo: Bar"

Start benchmarking:  https://httpstat.us/Random
Average Time: 815.652016ms
Total Time: 16.313040321s
Fastest Time: 141.653101ms
Slowest Time: 1.493196755s
Error Count: 0
Success Count: 20
200: 7
3xx: 1
4xx: 3
5xx: 9
Bytes Received: 50348b


```


## contributing

Feel free to send a PR or open an issue.
