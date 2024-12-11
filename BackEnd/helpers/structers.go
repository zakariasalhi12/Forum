package helpers

type Posts struct {
	User_ID    int
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	Categories []string `json:"categories"`
}

type Comment struct {
	UserID  int
	PostId  int    `json:"postid"`
	Content string `json:"content"`
}

type Likes struct {
	Counter int
	IsLiked bool
}

type Dislikes struct {
	Counter   int
	IsDislike bool
}

type AllPosts struct {
	Id         int
	User_id    int
	UserName   string
	Title      string
	Content    string
	Categories []string
	Comments   []Comments
	Likes      Likes
	Dislikes   Dislikes
	CreatedAt  string
}

type Comments struct {
	Id        int
	UserID    int
	UserName  string
	Content   string
	Likes     Likes
	Dislikes  Dislikes
	CreatedAt string
}
