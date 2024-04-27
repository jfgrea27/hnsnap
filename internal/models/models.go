package models

type Story struct {
	By          string  `json:"by"`
	Descendants int64   `json:"descendants"`
	Id          int64   `json:"id"`
	Kids        []int64 `json:"kids"`
	Score       int64   `json:"score"`
	Time        int64   `json:"time"`
	Title       string  `json:"title"`
	Url         string  `json:"url"`
}

type Comment struct {
	By     string  `json:"by"`
	Id     int64   `json:"id"`
	Kids   []int64 `json:"kids"`
	Parent int64   `json:"parent"`
	Text   int64   `json:"text"`
	Time   int64   `json:"time"`
}

type Ask struct {
	By          string  `json:"by"`
	Descendants int64   `json:"descendants"`
	Id          int64   `json:"id"`
	Kids        []int64 `json:"kids"`
	Score       int64   `json:"score"`
	Text        string  `json:"text"`
	Time        int64   `json:"time"`
	Title       string  `json:"title"`
}

type Job struct {
	By    string `json:"by"`
	Id    int64  `json:"id"`
	Score int64  `json:"score"`
	Text  string `json:"text"`
	Time  int64  `json:"time"`
	Title string `json:"title"`
	Url   string `json:"url"`
}

type Poll struct {
	By          string  `json:"by"`
	Descendants int64   `json:"descendants"`
	Id          int64   `json:"id"`
	Kids        []int64 `json:"kids"`
	Parts       []int64 `json:"parts"`
	Score       int64   `json:"score"`
	Text        string  `json:"text"`
	Time        int64   `json:"time"`
	Title       string  `json:"title"`
}

type PollOpt struct {
	By    string  `json:"by"`
	Id    int64   `json:"id"`
	Poll  []int64 `json:"poll"`
	Score int64   `json:"score"`
	Text  string  `json:"text"`
	Time  int64   `json:"time"`
}
