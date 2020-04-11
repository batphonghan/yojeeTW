Designing yojeeTW

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
    tweet(tweetData)
    reTweet(tweetID, tweetData)
    
5. High Level System Design
    We need a system that can efficiently store all the new tweets, 1B/86400s => 11,709 tweets per second and read 28B/86400s => 325K tweets per second. 
    It is clear from the requirements that this will be a read-heavy system.

    Our yojeeTW going to be:
    Internet -> webFrontEnd -> server
    - Our client going to talk with webFrontEnd via websocket to keep there page refresh.
    - WebFrontEnd going to using gRPC to comunicated to our server for efficient. 

    Which that design we can scale webFrontEnd or server dependenly.