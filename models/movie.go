package models

type Movie struct {
	ID             int64  `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Image_url      string `json:"image_url"`
	Air_date       string `json:"air_date"`
	Total_episode  string `json:"total_episode"`
	Current_status string `json:"current_status"`
	Category_id    int64  `json:"category_id"`
	Type_id        int64  `json:"type_id"`
	Created_at     string `json:"created_at"`
	Updated_at     string `json:"updated_at"`
}
