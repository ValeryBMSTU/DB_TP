package models



type NewForum struct {
	Slug string `json:"slug"`
	Title string `json:"title"`
	User string	`json:"user"`
}

type Forum struct {
	Posts int `json:"posts"`
	Slug string `json:"slug"`
	Thread int `json:"threads"`
	Title string `json:"title"`
	User string	`json:"user"`
}

// ==============================

type NewPost struct {
	Author string `json:"author"`
	Message string `json:"message"`
	Parent int `json:"parent"`
}

type ChangePost struct {
	Message string `json:"message"`
}

type Post struct {
	Author string `json:"author"`
	Created string `json:"created"`
	Forum string `json:"forum"`
	ID int `json:"id"`
	IsEdited bool `json:"isEdited"`
	Message string `json:"message"`
	Parent int `json:"parent"`
	Thread int `json:"thread"`
}

type EditedPost struct {
	Author string `json:"author"`
	Created string `json:"created"`
	Forum string `json:"forum"`
	ID int `json:"id"`
	Message string `json:"message"`
	Parent int `json:"parent"`
	Thread int `json:"thread"`
}

type PostDetails struct {
	Forum interface{}  `json:"forum,omitempty"`
	Thread interface{} `json:"thread,omitempty"`
 	User interface{} `json:"author,omitempty"`
	Post interface{} `json:"post"`
}

type NewPosts []*NewPost

type Posts []*Post

// ==============================

type NewThread struct {
	Author string `json:"author"`
	Created string `json:"created"`
	Message string `json:"message"`
	Slug string `json:"slug"`
	Title string `json:"title"`
}

type ChangeThread struct {
	Message string `json:"message"`
	Title string `json:"title"`
}

type Thread struct {
	Author string `json:"author"`
	Created string `json:"created"`
	Forum string `json:"forum"`
	ID int `json:"id"`
	Message string `json:"message"`
	Slug string `json:"slug"`
	Title string `json:"title"`
	Votes int `json:"votes"`
}

type Threads []*Thread

// ==============================

type NewUser struct {
	About string `json:"about"`
	Email string `json:"email"`
	Fullname string `json:"fullname"`
}

type User struct {
	About string `json:"about"`
	Email string `json:"email"`
	Fullname string `json:"fullname"`
	Nickname string `json:"nickname"`
}

type Users []*User

// ==============================


type Status struct {
	Post int `json:"post"`
	Thread int `json:"thread"`
	User	int `json:"user"`
	Forum int `json:"forum"`
}

// ==============================

type NewVote struct {
	Nickname string `json:"nickname"`
	Voice int `json:"voice"`
}

// ==============================

type Error struct {
	Message string `json:"message"`
}

// ==============================

type Body struct {
	Body interface{}
}

