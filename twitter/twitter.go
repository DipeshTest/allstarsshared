package twitter

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/ChimeraCoder/anaconda"
)

func PostTweet(consumerKey, consumerSecret, accessToken, accessTokenSecret, tweet string) (int, string) {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
	twt, err := api.PostTweet(tweet, nil)
	if err != nil {
		abc := err.(*anaconda.ApiError)
		return abc.StatusCode, abc.Error()
	}
	return 200, twt.IdStr
}
func ReTweet(consumerKey, consumerSecret, accessToken, accessTokenSecret string, tweetId int64) (int, string) {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
	twt, err := api.Retweet(tweetId, false)
	if err != nil {
		abc := err.(*anaconda.ApiError)
		return abc.StatusCode, abc.Error()
	}
	return 200, twt.IdStr
}
func DirectMessage(consumerKey, consumerSecret, accessToken, accessTokenSecret, directmsg, user string) (int, string) {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
	twt, err := api.PostDMToScreenName(directmsg, user)
	if err != nil {
		abc := err.(*anaconda.ApiError)
		return abc.StatusCode, abc.Error()
	}
	return 200, twt.IdStr
}
func UnFollow(consumerKey, consumerSecret, accessToken, accessTokenSecret, twitterHandle string) (int, string) {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
	resp, err := api.UnfollowUser(twitterHandle)
	if err != nil {
		abc := err.(*anaconda.ApiError)
		return abc.StatusCode, abc.Error()
	}
	return 200, resp.IdStr
}
func BlockUser(consumerKey, consumerSecret, accessToken, accessTokenSecret, twitterHandle string) (int, string) {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
	resp, err := api.BlockUser(twitterHandle, nil)
	if err != nil {
		abc := err.(*anaconda.ApiError)
		return abc.StatusCode, abc.Error()
	}
	return 200, resp.IdStr
}
func UnBlockUser(consumerKey, consumerSecret, accessToken, accessTokenSecret, twitterHandle string) (int, string) {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
	resp, err := api.UnblockUser(twitterHandle, nil)
	if err != nil {
		abc := err.(*anaconda.ApiError)
		return abc.StatusCode, abc.Error()
	}
	return 200, resp.IdStr
}
func GetUserTimeline(consumerKey, consumerSecret, accessToken, accessTokenSecret string, pageCount, sinceId int) (int, string) {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
	uriValues := url.Values{}
	if pageCount == 0 {
		uriValues.Add("count", strconv.Itoa(pageCount))
	} else {
		cnt := strconv.Itoa(pageCount)
		uriValues.Add("count", cnt)
	}
	uriValues.Add("trim_user", "true")
	uriValues.Add("include_rts", "true")
	if sinceId > 0 {
		uriValues.Add("since_id", strconv.Itoa(sinceId))
	}
	twt, err := api.GetUserTimeline(uriValues)
	if err != nil {
		abc := err.(*anaconda.ApiError)
		return abc.StatusCode, abc.Error()
	}
	resp, _ := json.Marshal(twt)
	return 200, string(resp)
}

func GetTrendsByPlace(consumerKey, consumerSecret, accessToken, accessTokenSecret string, placeId int64) (int, string) {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
	if placeId == 0 {
		placeId = 1
	}
	twt, err := api.GetTrendsByPlace(placeId, nil)
	if err != nil {
		abc := err.(*anaconda.ApiError)
		return abc.StatusCode, abc.Error()
	}
	resp, _ := json.Marshal(twt)
	return 200, string(resp)
}

func Follow(consumerKey, consumerSecret, accessToken, accessTokenSecret, twitterHandle string) (int, string) {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
	resp, err := api.FollowUser(twitterHandle)
	if err != nil {
		abc := err.(*anaconda.ApiError)
		return abc.StatusCode, abc.Error()
	}
	return 200, resp.IdStr
}

//This functinality is currenlty not impremented as an activity for hackathon however we can leverage the below code to create an activity
//func GetFollowers(consumerKey, consumerSecret, accessToken, accessTokenSecret, nextPgcursor string, pageCount int) (int, string, int, string) {
// 	anaconda.SetConsumerKey(consumerKey)
// 	anaconda.SetConsumerSecret(consumerSecret)
// 	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
// 	uriValues := url.Values{}
// 	if pageCount == 0 {
// 		uriValues.Add("count", "200")
// 	} else {
// 		cnt := strconv.Itoa(pageCount)
// 		uriValues.Add("count", cnt)
// 	}
// 	uriValues.Add("skip_status", "true")
//
// 	twt, err := api.GetFollowersList(uriValues)
// 	if err != nil {
// 		abc := err.(*anaconda.ApiError)
// 		return abc.StatusCode, abc.Error(), 0, ""
// 	}
// 	resp, _ := json.Marshal(twt)
// 	return 200, string(resp), len(twt.Users), twt.Next_cursor_str
// }

//This functinality is currenlty not impremented as an activity for hackathon however we can leverage the below code to create an activity
// func GetHomeTimeline(consumerKey, consumerSecret, accessToken, accessTokenSecret string, count, sinceId int) (int, string) {
// 	anaconda.SetConsumerKey(consumerKey)
// 	anaconda.SetConsumerSecret(consumerSecret)
// 	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
// 	uriValues := url.Values{}
// 	if count == 0 {
// 		uriValues.Add("count", "200")
// 	} else {
// 		cnt := strconv.Itoa(count)
// 		uriValues.Add("count", cnt)
// 	}
// 	uriValues.Add("trim_user", "true")
// 	uriValues.Add("include_rts", "true")
// 	if sinceId > 0 {
// 		uriValues.Add("since_id", strconv.Itoa(sinceId))
// 	}
// 	twt, err := api.GetHomeTimeline(uriValues)
// 	if err != nil {
// 		abc := err.(*anaconda.ApiError)
// 		return abc.StatusCode, abc.Error()
// 	}
// 	resp, _ := json.Marshal(twt)
// 	return 200, string(resp)
// }

// This functinality is currenlty not impremented as an activity for hackathon however we can leverage the below code to create an activity
// func Search(consumerKey, consumerSecret, accessToken, accessTokenSecret, searchStr string, pageCount, sinceId int) (int, string) {
// 	anaconda.SetConsumerKey(consumerKey)
// 	anaconda.SetConsumerSecret(consumerSecret)
// 	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
// 	uriValues := url.Values{}
// 	if pageCount == 0 {
// 		uriValues.Add("count", strconv.Itoa(pageCount))
// 	} else {
// 		cnt := strconv.Itoa(pageCount)
// 		uriValues.Add("count", cnt)
// 	}
// 	if sinceId > 0 {
// 		uriValues.Add("since_id", strconv.Itoa(sinceId))
// 	}
// 	twt, err := api.GetSearch(searchStr, uriValues)
// 	if err != nil {
// 		abc := err.(*anaconda.ApiError)
// 		return abc.StatusCode, abc.Error()
// 	}
// 	resp, _ := json.Marshal(twt)
// 	return 200, string(resp)
// }

//This functinality is currenlty not impremented as an activity for hackathon however we can leverage the below code to create an activity
// func UserSearch(consumerKey, consumerSecret, accessToken, accessTokenSecret, searchStr string, count, page int) (int, string) {
// 	anaconda.SetConsumerKey(consumerKey)
// 	anaconda.SetConsumerSecret(consumerSecret)
// 	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
// 	uriValues := url.Values{}
// 	if count == 0 {
// 		uriValues.Add("count", strconv.Itoa(count))
// 	} else {
// 		cnt := strconv.Itoa(count)
// 		uriValues.Add("count", cnt)
// 	}
// 	if page > 0 {
// 		uriValues.Add("page", strconv.Itoa(page))
// 	}
// 	//uriValues.Add("q", searchStr)
// 	twt, err := api.GetUserSearch(searchStr, uriValues)
// 	if err != nil {
// 		abc := err.(*anaconda.ApiError)
// 		return abc.StatusCode, abc.Error()
// 	}
// 	resp, _ := json.Marshal(twt)
// 	return 200, string(resp)
// }

// func Follow(consumerKey, consumerSecret, accessToken, accessTokenSecret, twitterHandle string) (int, string) {
// 	anaconda.SetConsumerKey(consumerKey)
// 	anaconda.SetConsumerSecret(consumerSecret)
// 	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
// 	resp, err := api.FollowUser(twitterHandle)
// 	if err != nil {
// 		abc := err.(*anaconda.ApiError)
// 		return abc.StatusCode, abc.Error()
// 	}
// 	return 200, resp.IdStr
// }

// This functinality is currenlty not impremented as an activity for hackathon however we can leverage the below code to create an activity
//func GetStats(consumerKey, consumerSecret, accessToken, accessTokenSecret, hashTag, timetoTrack string) (int, int, string) {
// 	anaconda.SetConsumerKey(consumerKey)
// 	anaconda.SetConsumerSecret(consumerSecret)
// 	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
// 	stream, err := api.GetSearch(hashTag, url.Values{
// 		"count": []string{"1000"},
// 	})
// 	count := len(stream.Statuses)
// 	tobreak := true
// 	for tobreak {
// 		stream2, err2 := stream.GetNext(api)
// 		if err2 != nil {
// 			fmt.Println(err2)
// 		} else {
// 			fmt.Println("Size of next:", len(stream2.Statuses))
// 			nxt := len(stream2.Statuses)
// 			count = count + nxt
// 			if nxt == 0 {
// 				tobreak = false
// 			}
// 		}
//
// 	}
// 	if count > 0 && err == nil {
// 		return 200, count, "Success"
// 	}
// 	return 201, 0, "No Tweets found"
// }
