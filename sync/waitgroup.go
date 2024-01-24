package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var urls = []string{
	"http://google.com",
	"http://twitter.com",
	"http://yahoo.com",
}

func fetchStatus(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Gorouitne waitgroup running....!")
	var wg sync.WaitGroup

	for _, url := range urls {

		wg.Add(1)
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Fprintf(w, "%+v\n", err)
			}
			fmt.Fprintf(w, "%+v\n", resp.Status)
			wg.Done()
		}(url)
	}
	wg.Wait()
}

func main() {

	fmt.Println("Goroutine Waitgroup URl")

	http.HandleFunc("/", fetchStatus)

	log.Fatal(http.ListenAndServe(":8050", nil))

}
