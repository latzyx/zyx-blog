package moudels

type Post struct {
	Content     string `json:"content"`
	Slug        string `json:"slug"`
	Title       string `json:"title"`
	Pid         string `json:"pid"`
	Markdown    string `json:"markdown"`
	CategoryId  int    `json:"category_id"`
	UserId      int    `json:"user_id"`
	ViewerCount int    `json:"viewer_count"`
	Type        int    `json:"type"`
	CreateAt    string `json:"create_at"`
	UpdateAt    string `json:"update_at"`
}

type PostMore struct {
	Content     string `json:"content"`
	Slug        string `json:"slug"`
	Title       string `json:"title"`
	Pid         string `json:"pid"`
	Markdown    string `json:"markdown"`
	CategoryId  int    `json:"category_id"`
	UserId      int    `json:"user_id"`
	UserName    int    `json:"user_name"`
	ViewerCount int    `json:"viewer_count"`
	Type        int    `json:"type"`
	CreateAt    string `json:"create_at"`
	UpdateAt    string `json:"update_at"`
}
