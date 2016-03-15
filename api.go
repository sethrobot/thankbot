package main

import "github.com/ChimeraCoder/anaconda"

const maxAllowedFollowers = 70000

type userData struct {
	ScreenName     string `json:"screen_name"`
	FollowersCount int    `json:"followers_count"`
}

type earliestFollowers struct {
	Self      userData   `json:"self"`
	Followers []userData `json:"followers"`
}

type api struct {
	*anaconda.TwitterApi
}

func minimalUser(input anaconda.User) userData {
	return userData{
		ScreenName:     input.ScreenName,
		FollowersCount: input.FollowersCount,
	}
}

func minimalUsers(input []anaconda.User) []userData {
	ret := make([]userData, len(input), len(input))
	for i, el := range input {
		ret[i] = minimalUser(el)
	}
	return ret
}

func (a *api) earliestFollowers(count int) (*earliestFollowers, error) {
	ret := new(earliestFollowers)
	self, err := a.GetSelf(nil)
	if err != nil {
		return nil, errNoTwitterAccess
	}
	ret.Self = minimalUser(self)
	if self.FollowersCount > maxAllowedFollowers {
		return ret, nil
	}
	followers, err := a.firstFollowers(count)
	if err != nil {
		return nil, err
	}
	ret.Followers = minimalUsers(followers)
	return ret, nil
}

func (a *api) firstFollowers(count int) ([]anaconda.User, error) {
	return a.GetUsersLookupByIds(a.firstFollowerIDs(count), nil)
}

func (a *api) firstFollowerIDs(count int) []int64 {
	idBuf := newCircularBuffer(count)
	for followersPage := range a.GetFollowersIdsAll(nil) {
		for _, followerID := range followersPage.Ids {
			idBuf.add(followerID)
		}
	}
	return idBuf.safeBuffer()
}
