package main

import "github.com/google/uuid"

const TOPTWEETS = 10

// Tweet is data struct for each tweet create
type Tweet struct {
	ID      string `json:"id,omitempty"`
	Data    string `json:"data,omitempty"`
	ReTweet int    `json:"re_tweet"`
}

// ByReTweet implements sort.Interface based on the ReTweet field.
type ByReTweet []Tweet

func (t ByReTweet) Len() int           { return len(t) }
func (t ByReTweet) Less(i, j int) bool { return t[i].ReTweet < t[j].ReTweet }
func (t ByReTweet) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }

func makeTweet(tweetData string) (Tweet, error) {
	UUID, err := uuid.NewRandom()

	if err != nil {
		return Tweet{}, err
	}

	tweetID := UUID.String()
	tweet := Tweet{
		ID:   tweetID,
		Data: tweetData,
	}

	tweets[tweetID] = tweet

	return tweet, nil
}
