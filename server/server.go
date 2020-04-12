package main

import (
	"context"
	"log"
	"yojeeTW/pb"
)

type server struct {
}

func (*server) Tweet(ctx context.Context, request *pb.TweetRequest) (*pb.TweetResponse, error) {
	log.Printf("Serving Tweet: ID[%s] Data:[%s] \n", request.TweetID, request.TweetData)
	if tweetID := request.TweetID; len(tweetID) > 0 {
		err := reTweet(tweetID)
		if err != nil {
			log.Printf("StatusNotFound ID: %s \n", tweetID)
			return &pb.TweetResponse{}, nil
		}
		return &pb.TweetResponse{
			Tweet: &pb.Tweet{
				TweetID: tweetID,
			},
		}, nil
	}

	if tweetData := request.TweetData; len(tweetData) > 0 {
		tweet, err := makeTweet(tweetData)
		if err != nil {
			log.Printf("Error tweet with err: %v \n", err)
			return &pb.TweetResponse{}, nil
		}
		return &pb.TweetResponse{
			Tweet: &pb.Tweet{
				TweetData: tweet.Data,
				TweetID:   tweet.ID,
				Retweets:  tweet.ReTweet,
			},
		}, nil
	}

	return &pb.TweetResponse{}, nil
}

func (*server) TopRetweets(ctx context.Context, request *pb.TopRetweetsRequest) (*pb.TopRetweetsResponse, error) {
	b, err := makeTopTweetsResponse()
	if err != nil {
		log.Println("Error when make toptweet response")
		return nil, nil
	}
	return &pb.TopRetweetsResponse{Data: b}, nil
}
