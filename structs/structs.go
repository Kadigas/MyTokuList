package structs

type Category struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type Type struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type Movie struct {
	ID             int64  `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Image_url      string `json:"image_url"`
	Air_date       string `json:"aired"`
	Total_episode  string `json:"total_episode"`
	Current_status string `json:"current_status"`
	Category_id    int64  `json:"category_id"`
	Type_id        int64  `json:"type_id"`
	Created_at     string `json:"created_at"`
	Updated_at     string `json:"updated_at"`
}

type Account struct {
	ID         int64  `json:"id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type Users struct {
	ID         int64  `json:"id"`
	Username   string `json:"username"`
	Image_url  string `json:"image_url"`
	Account_id int64  `json:"account_id"`
	Updated_at string `json:"updated_at"`
}

type Watchlist struct {
	ID       int64 `json:"id"`
	Users_id int64 `json:"users_id"`
	Movie_id int64 `json:"movie_id"`
}
