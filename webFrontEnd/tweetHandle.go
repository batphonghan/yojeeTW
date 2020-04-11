package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
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
func serveTweets(rw http.ResponseWriter, rq *http.Request) {
	for k, v := range tweets {
		fmt.Printf("ID %s Data %s \n", k, v.data)
	}
}
func serveTweet(rw http.ResponseWriter, rq *http.Request) {
	if tweetID := rq.FormValue("tweet_id"); len(tweetID) > 0 {
		if err := reTweet(tweetID); err != nil {
			fmt.Fprintln(rw, "Error when retweet")
			log.Printf("Error retweetID: %s \n", tweetID)
			return
		}

		log.Printf("Success retweetID: %s \n", tweetID)
		rw.WriteHeader(http.StatusOK)
		return
	}

	if tweetData := rq.FormValue("tweet_data"); len(tweetData) > 0 {
		if _, err := tweet(tweetData); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)

			log.Printf("Error tweet with err: %v \n", err)
			return
		}

		log.Printf("Success tweet")
		rw.WriteHeader(http.StatusCreated)
		return
	}

	log.Printf("Error bad request")
	rw.WriteHeader(http.StatusBadRequest)
}

var tweets = make(map[string]Tweet)
var topTweets []Tweet = make([]Tweet, 10)

// Tweet is data struct for each tweet create
type Tweet struct {
	ID      string
	data    string
	reTweet uint64
}

func reTweet(tweetID string) error {
	tweet, ok := tweets[tweetID]

	if !ok {
		return fmt.Errorf("TweetID %s not exist", tweetID)
	}

	tweet.reTweet++
	tweets[tweetID] = tweet

	return nil
}

func tweet(tweetData string) (string, error) {
	UUID, err := uuid.NewUUID()

	if err != nil {
		return "", err
	}

	tweetID := UUID.String()
	tweet := Tweet{
		ID:   tweetID,
		data: tweetData,
	}

	tweets[tweetID] = tweet

	topTweets = append(topTweets, tweet)

	return tweetID, nil
}
