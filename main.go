package main

import (
	"fmt"

	"github.com/wachayathorn/quiz-problems/design"
)

func main() {
	// twitter := design.Constructor()
	// twitter.PostTweet(1, 5)
	// userFeed := twitter.GetNewsFeed(1)
	// fmt.Println("user feed ids", userFeed)
	// twitter.Follow(1, 2)
	// twitter.PostTweet(2, 6)
	// userFeed = twitter.GetNewsFeed(1)
	// fmt.Println("user feed ids", userFeed)
	// twitter.Unfollow(1, 2)
	// userFeed = twitter.GetNewsFeed(1)
	// fmt.Println("user feed ids", userFeed)

	// Input : ["Twitter","postTweet","postTweet","postTweet","postTweet","postTweet","postTweet","postTweet","postTweet","postTweet","postTweet","postTweet","getNewsFeed"]
	// [[],[1,5],[1,3],[1,101],[1,13],[1,10],[1,2],[1,94],[1,505],[1,333],[1,22],[1,11],[1]]
	twitter := design.Constructor()
	twitter.PostTweet(1, 5)
	twitter.PostTweet(1, 3)
	twitter.PostTweet(1, 101)
	twitter.PostTweet(1, 13)
	twitter.PostTweet(1, 10)
	twitter.PostTweet(1, 2)
	twitter.PostTweet(1, 94)
	twitter.PostTweet(1, 505)
	twitter.PostTweet(1, 333)
	twitter.PostTweet(1, 22)
	twitter.PostTweet(1, 11)

	userFeed := twitter.GetNewsFeed(1)
	fmt.Println("feed ", userFeed)
}
