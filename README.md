#### Designing yojeeTW

1. What is yojeeTW?
    Implement a very basic version of Twitter. 
2. Requirements and Goals of the System
    - User can create a tweet on this page (anonymous)
    - Any tweet can be retweeted
    - The tweet will contain just 140 characters long string. 
    - Display top 10 tweets. (Ordering is based on the number of retweets the original tweet gets) 
    - The data can be maintained in memory

3. Capacity Estimation and Constraints
    How many total tweet-views will our system generate? 
    Let’s assume our system will generate 1B/day total tweet-views
    Let’s assume a use will retweet on time as 5 times view 1B/day / 5 -> 200M total re-tweet

    Storage Estimates 
    Let’s say each tweet has 140 characters and we need two bytes to store a character without compression
    Let’s assume we need 30 bytes to store metadata with each tweet (like ID, timestamp, user ID, etc.). Total storage we would need:

    1B * (30) bytes => 30GB/day
    We able to got 30GB/day which's kind of small due to assume to remove photo, videos. 

4. System APIs
    `tweet(tweetData, tweetID)`
    if `tweetID` is nil mean that we create new tweet
    
5. High Level System Design
    We need a system that can efficiently store all the new tweets, 1B/86400s => 11,709 tweets per second and read 28B/86400s => 325K tweets per second. 
    It is clear from the requirements that this will be a read-heavy system.

    Our yojeeTW going to be:
    Internet -> webFrontEnd -> server
    - Our client going to talk with webFrontEnd via websocket to keep there page refresh.
    - WebFrontEnd going to using gRPC to comunicated to our server for efficient. 

    Which that design we can scale webFrontEnd or server dependenly.

#### Prerequisities

    docker demon started

#### How to run:
For the quick setting up I commit the prebuild binary to the repo. Just run the follow command to start.

`./run.sh`

For create tweet: replace `tweet_data` with any text (even more than 140 char)

```
curl --location --request POST 'http://127.0.0.1:8080/tweet?tweet_data=Lorem%20ipsum%20dolor%20sit%20amet,%20consectetur%20adipiscing%20elit.%20Pellentesque%20interdum%20rutrum%20sodales.%20Nullam%20mattis%20fermentum%20libero,%20non%20volutpat.%20'

```

Get retweet:

```
curl --location --request GET 'http://127.0.0.1:8080/retweets'
```

Retweet (noted to replaced tweet_id with corrected id return form create tweet or from retweet API):

```
curl --location --request POST 'http://127.0.0.1:8080/tweet?tweet_id=bff3ce31-1b9f-4208-8983-baf72d8e2c2c'
```
