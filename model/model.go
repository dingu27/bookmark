package model

//Readmark ...
type Readmark struct {
	Email    string `json:"email"`
	URL      string `json:"imgUrl"`
	Bookmark string `json:"bookmark"`
}

//Bookmark ...
type Bookmark struct {
	Email string `json:"email"`
	URL   string `json:"imgUrl"`
}

//Bookmarks ...
type Bookmarks []Bookmark

//GetBookmarks ...
type GetBookmarks struct {
	Email string   `json:"email"`
	URLs  []string `json:"imgUrls"`
}
