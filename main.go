package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {
	flag.Parse()
	url := strings.Join(flag.Args(), " ")

	for true {
		resp, _ := http.Get(url)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(byteArray))

		time.Sleep(time.Second * 1)
	}
}
