package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
)

func protectTweet(f http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, rq *http.Request) {
		if rq.Method != http.MethodPost {
			log.Printf("Error  StatusMethodNotAllowed")
			rw.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		f(rw, rq)
	}
}

func protectTweets(f http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, rq *http.Request) {
		if rq.Method != http.MethodGet {
			log.Printf("Error  StatusMethodNotAllowed")
			rw.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		f(rw, rq)
	}
}

func makeTopTweetsResponse() ([]byte, error) {
	reTweets := make([]Tweet, 0, len(topTweets))

	for _, v := range topTweets {
		reTweets = append(reTweets, v)
	}

	sort.Sort(ByReTweet(reTweets))

	return json.Marshal(reTweets)
}

func serveTweets(rw http.ResponseWriter, rq *http.Request) {
	rs, err := makeTopTweetsResponse()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(rs)
}

func serveTweet(rw http.ResponseWriter, rq *http.Request) {
	if tweetID := rq.FormValue("tweet_id"); len(tweetID) > 0 {
		if err := reTweet(tweetID); err != nil {
			log.Printf("StatusNotFound ID: %s \n", tweetID)
			rw.WriteHeader(http.StatusNotFound)
			return
		}

		log.Printf("Success retweetID: %s \n", tweetID)
		rw.WriteHeader(http.StatusOK)
		return
	}

	if tweetData := rq.FormValue("tweet_data"); len(tweetData) > 0 {
		tweet, err := makeTweet(tweetData)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)

			log.Printf("Error tweet with err: %v \n", err)
			return
		}

		go updateTopTweets(tweet)

		log.Printf("Success tweet %s", tweet.ID)
		rw.WriteHeader(http.StatusCreated)
		return
	}

	log.Printf("Error bad request")
	rw.WriteHeader(http.StatusBadRequest)
}
