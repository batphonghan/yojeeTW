package main

import (
	"encoding/json"
	"log"
	"net/http"
	"yojeeTW/pb"
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

func serveTopTweets(rw http.ResponseWriter, rq *http.Request) {
	log.Printf("ServeTopTweets Tweet with url: %v \n", discoveryURL)
	res, err := client.TopRetweets(rq.Context(), &pb.TopRetweetsRequest{})
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	b := res.Data

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(b)
}

func serveTweet(rw http.ResponseWriter, rq *http.Request) {
	log.Printf("Serving Tweet with url: %v \n", discoveryURL)
	tweetID := rq.FormValue("tweet_id")
	tweetData := rq.FormValue("tweet_data")
	request := pb.TweetRequest{
		TweetID:   tweetID,
		TweetData: tweetData,
	}
	resp, err := client.Tweet(rq.Context(), &request)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error tweet with err: %v \n", err)
		return
	}
	if tweet := resp.Tweet; tweet != nil {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusCreated)
		b, err := json.Marshal(tweet)
		if err != nil {
			log.Println(err)
		}
		rw.Write(b)
		return
	}

	log.Printf("Error bad request")
	rw.WriteHeader(http.StatusBadRequest)
}
