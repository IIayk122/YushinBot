package vkapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	API_METHOD_URL      = "https://api.vk.com/method/"
	AUTH_HOST           = "https://oauth.vk.com/authorize"
	AUTH_HOST_GET_TOKEN = "https://oauth.vk.com/access_token"
)

type Api struct {
	AccessToken string
	UserId      int
	ExpiresIn   int
	debug       bool
}

func (vk Api) Request(method_name string, params map[string]string) string {
	u, err := url.Parse(API_METHOD_URL + method_name)
	if err != nil {
		fmt.Print(err)
	}
	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	q.Set("access_token", vk.AccessToken)
	u.RawQuery = q.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err)
	}

	return string(content)
}

func GetResponse(m string, parametr string) interface{} {
	data := []byte(m)
	var parsed interface{}
	err := json.Unmarshal(data, &parsed)
	if err != nil {
		fmt.Print(err)
	}
	par, _ := parsed.(map[string]interface{})
	var va interface{}
	for _, v := range par {
		va = v
	}
	return (append([]interface{}{}, va))[0].([]interface{})[1].(map[string]interface{})[parametr]

}

type Album struct {
	Response struct {
		Count int `json:"count"`
		Items []struct {
			ID          int    `json:"id"`
			ThumbID     int    `json:"thumb_id"`
			OwnerID     int    `json:"owner_id"`
			Title       string `json:"title"`
			Description string `json:"description"`
			Created     int    `json:"created"`
			Updated     int    `json:"updated"`
			Size        int    `json:"size"`
			ThumbIsLast int    `json:"thumb_is_last,omitempty"`
			CanUpload   int    `json:"can_upload"`
		} `json:"items"`
	} `json:"response"`
}

type Wall struct {
	Response struct {
		Count int `json:"count"`
		Items []struct {
			ID          int    `json:"id"`
			FromID      int    `json:"from_id"`
			OwnerID     int    `json:"owner_id"`
			Date        int    `json:"date"`
			MarkedAsAds int    `json:"marked_as_ads"`
			PostType    string `json:"post_type"`
			Text        string `json:"text"`
			IsPinned    int    `json:"is_pinned"`
			Attachments []struct {
				Type  string `json:"type"`
				Photo struct {
					ID        int    `json:"id"`
					AlbumID   int    `json:"album_id"`
					OwnerID   int    `json:"owner_id"`
					UserID    int    `json:"user_id"`
					Photo75   string `json:"photo_75"`
					Photo130  string `json:"photo_130"`
					Photo604  string `json:"photo_604"`
					Photo807  string `json:"photo_807"`
					Photo1280 string `json:"photo_1280"`
					Photo2560 string `json:"photo_2560"`
					Width     int    `json:"width"`
					Height    int    `json:"height"`
					Text      string `json:"text"`
					Date      int    `json:"date"`
					PostID    int    `json:"post_id"`
					AccessKey string `json:"access_key"`
				} `json:"photo"`
			} `json:"attachments"`
			PostSource struct {
				Type string `json:"type"`
			} `json:"post_source"`
			Comments struct {
				Count   int `json:"count"`
				CanPost int `json:"can_post"`
			} `json:"comments"`
			Likes struct {
				Count      int `json:"count"`
				UserLikes  int `json:"user_likes"`
				CanLike    int `json:"can_like"`
				CanPublish int `json:"can_publish"`
			} `json:"likes"`
			Reposts struct {
				Count        int `json:"count"`
				UserReposted int `json:"user_reposted"`
			} `json:"reposts"`
		} `json:"items"`
	} `json:"response"`
}

type Photo struct {
	Response struct {
		Count int `json:"count"`
		Items []struct {
			ID      int `json:"id"`
			AlbumID int `json:"album_id"`
			OwnerID int `json:"owner_id"`
			UserID  int `json:"user_id"`
			Sizes   []struct {
				Type   string `json:"type"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"sizes"`
			Text string `json:"text"`
			Date int    `json:"date"`
		} `json:"items"`
	} `json:"response"`
}
