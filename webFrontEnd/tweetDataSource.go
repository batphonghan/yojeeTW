package main

import (
	"fmt"
	"sync"
)

var tweets = make(map[string]Tweet)
var topTweets = make(map[string]Tweet)

func reTweet(tweetID string) error {
	tweet, ok := tweets[tweetID]

	if !ok {
		return fmt.Errorf("TweetID %s not exist", tweetID)
	}

	tweet.ReTweet++
	tweets[tweetID] = tweet

	go updateTopTweets(tweet)

	return nil
}

var mtUpdate sync.RWMutex

func updateTopTweets(t Tweet) {
	mtUpdate.Lock()
	defer mtUpdate.Unlock()

	_, ok := topTweets[t.ID]
	if ok {
		topTweets[t.ID] = t
		return
	}

	if len(topTweets) < TOPTWEETS {
		topTweets[t.ID] = t
		return
	}

	var lessRetweet Tweet
	for _, v := range topTweets {
		if v.ReTweet < lessRetweet.ReTweet {
			lessRetweet = v
		}
	}

	if lessRetweet.ReTweet < t.ReTweet {
		topTweets[lessRetweet.ID] = t
	}
}
