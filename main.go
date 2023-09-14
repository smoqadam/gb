package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Start...")
	start := time.Now()
	res, err := http.Get("https://en.wikipedia.org/wiki/Base64")
	elapsed := time.Since(start)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Println("Res: ", elapsed)
}
