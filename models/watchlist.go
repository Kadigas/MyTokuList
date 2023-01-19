package models

type Watchlist struct {
	ID       int64 `json:"id"`
	Users_id int64 `json:"users_id"`
	Movie_id int64 `json:"movie_id"`
}
