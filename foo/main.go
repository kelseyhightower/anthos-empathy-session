package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	bar string
)

func main() {
	flag.StringVar(&bar, "bar", "http://127.0.0.1:8081", "The bar service HTTP endpoint")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		res, err := http.Get(bar)
		if err != nil {
			log.Fatalf("error calling the bar service: %v", err)
		}

		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprint(w, string(data))
	})

	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatal(err)
	}
}
