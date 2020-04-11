package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	tem := &templateHandler{fileName: "tweet.html"}
	host := port()
	fmt.Println("Serving at port ", host)

	http.Handle("/", tem)
	http.HandleFunc("/tweet", protectTweet(serveTweet))
	http.HandleFunc("/retweets", protectTweets(serveTweets))

	log.Fatal(http.ListenAndServe(host, nil))
}

func port() string {
	return "localhost:8080"
}
