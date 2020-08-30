package vkapi

type (
	Album struct {
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

	Wall struct {
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

	Photo struct {
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
)
