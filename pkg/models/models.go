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

type NewThread struct {
	Author string `json:"author"`
	Created string `json:"created"`
	Message string `json:"message"`
	Slug string `json:"slug"`
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

// ==============================

type Error struct {
	Message string `json:"message"`
}

