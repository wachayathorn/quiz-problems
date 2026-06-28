package twitter

import (
	"sort"
	"time"
)

type Feed struct {
	users map[int]user
}

type tweetInfo struct {
	ID   int
	Time time.Time
}

type user struct {
	UserID       int
	TweetIDs     []tweetInfo
	FollowerIDs  map[int]time.Time
	FollowingIDs map[int]time.Time
}

func NewFeed() *Feed {
	return &Feed{
		users: make(map[int]user),
	}
}

func (f *Feed) PostTweet(userID, tweetID int) {
	f.initUser(userID)

	author := f.users[userID]
	author.TweetIDs = append(author.TweetIDs, tweetInfo{
		ID:   tweetID,
		Time: time.Now(),
	})
	f.users[userID] = author
}

func (f *Feed) GetNewsFeed(userID int) []int {
	f.initUser(userID)

	author := f.users[userID]
	tweetsInfo := append([]tweetInfo{}, author.TweetIDs...)

	for followingID := range author.FollowingIDs {
		following := f.users[followingID]
		tweetsInfo = append(tweetsInfo, following.TweetIDs...)
	}

	sort.Slice(tweetsInfo, func(i, j int) bool {
		return tweetsInfo[i].Time.After(tweetsInfo[j].Time)
	})

	tweets := make([]int, 0, 10)
	for _, t := range tweetsInfo {
		if len(tweets) == 10 {
			break
		}
		tweets = append(tweets, t.ID)
	}

	return tweets
}

func (f *Feed) Follow(followerID, followeeID int) {
	f.initUser(followerID)
	f.initUser(followeeID)

	followee := f.users[followeeID]
	followee.FollowerIDs[followerID] = time.Now()
	f.users[followeeID] = followee

	follower := f.users[followerID]
	follower.FollowingIDs[followeeID] = time.Now()
	f.users[followerID] = follower
}

func (f *Feed) Unfollow(followerID, followeeID int) {
	f.initUser(followerID)
	f.initUser(followeeID)

	followee := f.users[followeeID]
	delete(followee.FollowerIDs, followerID)
	f.users[followeeID] = followee

	follower := f.users[followerID]
	delete(follower.FollowingIDs, followeeID)
	f.users[followerID] = follower
}

func (f *Feed) initUser(userID int) {
	if _, exists := f.users[userID]; !exists {
		f.users[userID] = user{
			UserID:       userID,
			FollowerIDs:  make(map[int]time.Time),
			FollowingIDs: make(map[int]time.Time),
		}
	}
}
