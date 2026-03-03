package design

import (
	"sort"
	"time"
)

type Twitter struct {
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

func Constructor() Twitter {
	return Twitter{
		users: make(map[int]user),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.initUserInfo(userId)
	// Add new tweet
	authorInfo := this.users[userId]
	authorInfo.TweetIDs = append(authorInfo.TweetIDs, tweetInfo{
		ID:   tweetId,
		Time: time.Now(),
	})
	this.users[userId] = authorInfo
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	this.initUserInfo(userId)
	authorInfo := this.users[userId]
	authorTweets := authorInfo.TweetIDs
	tweetsInfo := []tweetInfo{}
	tweetsInfo = append(tweetsInfo, authorTweets...)

	for followingID := range authorInfo.FollowingIDs {
		followingInfo := this.users[followingID]
		tweetsInfo = append(tweetsInfo, followingInfo.TweetIDs...)
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

// User 1 follows user 2.
func (this *Twitter) Follow(followerId int, followeeId int) {
	this.initUserInfo(followerId)
	this.initUserInfo(followeeId)

	// init follower to target user
	followeeInfo := this.users[followeeId]
	followeeInfo.FollowerIDs[followerId] = time.Now()
	this.users[followeeId] = followeeInfo

	// init following to author user
	followerInfo := this.users[followerId]
	followerInfo.FollowingIDs[followeeId] = time.Now()
	this.users[followerId] = followerInfo
}

// User 1 unfollows user 2.
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	this.initUserInfo(followerId)
	this.initUserInfo(followeeId)

	// remove follower to target user
	followeeInfo := this.users[followeeId]
	delete(followeeInfo.FollowerIDs, followerId)
	this.users[followeeId] = followeeInfo

	// remove following from author user
	followerInfo := this.users[followerId]
	delete(followerInfo.FollowingIDs, followeeId)
	this.users[followerId] = followerInfo
}

func (this *Twitter) initUserInfo(userId int) {
	if _, isExisting := this.users[userId]; !isExisting {
		this.users[userId] = user{
			UserID:       userId,
			FollowerIDs:  make(map[int]time.Time),
			FollowingIDs: make(map[int]time.Time),
		}
	}
}
