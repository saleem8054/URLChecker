package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

func checkAndSaveBody(url string, wg *sync.WaitGroup) {
	if resp, err := http.Get(url); err != nil {
		fmt.Println(err)
		fmt.Printf("%s is down\n", url)
	} else {
		defer resp.Body.Close()
		fmt.Printf("%s -> status code: %d\n", url, resp.StatusCode)
		if resp.StatusCode == 200 {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			file := strings.Split(url, "//")[1]
			file += ".txt"
			fmt.Printf("Writing response body to %s\n", url)

			err = ioutil.WriteFile(file, bodyBytes, 0664)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	wg.Done()

}

func main() {
	urlNames := []string{"https://www.google.com", "https://www.tesla.com", "https://www.apple.com"}

	var wg sync.WaitGroup
	wg.Add(len(urlNames))
	for _, url := range urlNames {
		go checkAndSaveBody(url, &wg)
	}
	wg.Wait()
}
