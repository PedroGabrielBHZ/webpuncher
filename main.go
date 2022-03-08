package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://xyz.com/",
	}

	// Website crasher.
	// for i := 0; i <= 100000; i++ {
	// 	links = append(links, links[0])
	// }

	fmt.Println(links)

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down...")
		c <- link
		return
	}

	c <- link
	fmt.Println(link, "is up.")
}
