package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"sync"
)

var tweetsMT sync.RWMutex
var tweets = make(map[string]Tweet)
var topTweets = make(map[string]Tweet)

func makeTopTweetsResponse() ([]byte, error) {
	topTweetMT.RLock()
	defer topTweetMT.RUnlock()
	reTweets := make([]Tweet, 0, len(topTweets))

	for _, v := range topTweets {
		reTweets = append(reTweets, v)
	}

	sort.Sort(ByReTweet(reTweets))

	return json.Marshal(reTweets)
}

func reTweet(tweetID string) error {
	tweetsMT.Lock()
	defer tweetsMT.Unlock()
	tweet, ok := tweets[tweetID]

	if !ok {
		return fmt.Errorf("TweetID %s not exist", tweetID)
	}

	tweet.ReTweet++
	tweets[tweetID] = tweet

	go updateTopTweets(tweet)

	return nil
}

var topTweetMT sync.RWMutex

func updateTopTweets(t Tweet) {
	topTweetMT.Lock()
	defer topTweetMT.Unlock()

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

func adTweet(t Tweet) {
	tweetsMT.Lock()
	defer tweetsMT.Unlock()

	tweets[t.ID] = t
}
