package helpers

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
