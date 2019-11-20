package repository

import (
	"database/sql"
	"errors"
	"github.com/ValeryBMSTU/DB_TP/pkg/consts"
	"github.com/ValeryBMSTU/DB_TP/pkg/models"
	"sort"
	"strconv"
	"strings"
)

func (rep *ReposStruct) SelectForumsBySlug(slug string) (Forum []models.Forum, Err error) {
	var forums []models.Forum
	rows, err := rep.DataBase.Query(consts.SELECTForumsBySlug, slug)
	defer rows.Close()
	if err != nil {
		return forums, err
	}

	scanForum := models.Forum{}
	for rows.Next() {
		err := rows.Scan(&scanForum.Posts, &scanForum.Slug, &scanForum.Thread,
			&scanForum.Title, &scanForum.User)
		if err != nil {
			return forums, err
		}
		forums = append(forums, scanForum)
	}

	return forums, nil
}

func (rep *ReposStruct) SelectPosts(threadID int, limit, since, sort, desc string) (Posts *models.Posts, Err error) {
	posts := models.Posts{}

	var rows *sql.Rows
	var err error
	if sort == "flat" {
		if desc == "false" {
			rows, err = rep.DataBase.Query(consts.SELECTPostsFlat, threadID, limit, since)
		} else {
			rows, err = rep.DataBase.Query(consts.SELECTPostsFlatDesc, threadID, limit, since)
		}

	} else if sort == "tree" {
		if desc == "false" {
			rows, err = rep.DataBase.Query(consts.SELECTPostsTree, threadID, limit, since)
		} else {
			rows, err = rep.DataBase.Query(consts.SELECTPostsTreeDesc, threadID, limit, since)
		}
	} else if sort == "parent_tree" {
		if desc == "false" {
			rows, err = rep.DataBase.Query(consts.SELECTPostsParentTree, threadID, since)
		} else {
			rows, err = rep.DataBase.Query(consts.SELECTPostsParentTreeDesc, threadID, since)
		}
	}

	if sort != "parent_tree" {
		defer rows.Close()
		if err != nil {
			return &posts, err
		}

		for rows.Next() {
			scanPost := models.Post{}
			err := rows.Scan(&scanPost.Author, &scanPost.Created, &scanPost.Forum,
				&scanPost.ID, &scanPost.IsEdited, &scanPost.Message, &scanPost.Parent,
				&scanPost.Thread)
			if err != nil {
				return &posts, err
			}
			posts = append(posts, &scanPost)
		}
	} else {
		if err != nil {
			rows.Close()
			return &posts, err
		}

		count := 0
		limitDigit, _ := strconv.Atoi(limit)

		for rows.Next() {
			scanPost := models.Post{}
			err := rows.Scan(&scanPost.Author, &scanPost.Created, &scanPost.Forum,
				&scanPost.ID, &scanPost.IsEdited, &scanPost.Message, &scanPost.Parent,
				&scanPost.Thread)
			if err != nil {
				return &posts, err
			}

			posts = append(posts, &scanPost)

			if scanPost.Parent == 0 {
				count = count + 1
			}
			if count >= limitDigit {
				rows.Close()
				return &posts, nil
			}
		}
		rows.Close()
	}

	return &posts, nil
}

func (rep *ReposStruct) SelectThreadsBySlug(slug string) (Threads *models.Threads, Err error) {
	threads := models.Threads{}


	rows, err := rep.DataBase.Query(consts.SELECTThreadsBySlug, slug)

	defer rows.Close()
	if err != nil {
		return &threads, err
	}

	for rows.Next() {
		scanThread := models.Thread{}
		err := rows.Scan(&scanThread.Author, &scanThread.Created, &scanThread.Forum,
			&scanThread.ID, &scanThread.Message, &scanThread.Slug, &scanThread.Title,
			&scanThread.Votes)
		if err != nil {
			return &threads, err
		}
		threads = append(threads, &scanThread)
	}
	return &threads, nil
}

func (rep *ReposStruct) SelectThreadsByID(id int) (Threads *models.Threads, Err error) {
	threads := models.Threads{}


	rows, err := rep.DataBase.Query(consts.SELECTThreadsByID, id)

	defer rows.Close()
	if err != nil {
		return &threads, err
	}

	for rows.Next() {
		scanThread := models.Thread{}
		err := rows.Scan(&scanThread.Author, &scanThread.Created, &scanThread.Forum,
			&scanThread.ID, &scanThread.Message, &scanThread.Slug, &scanThread.Title,
			&scanThread.Votes)
		if err != nil {
			return &threads, err
		}
		threads = append(threads, &scanThread)
	}
	return &threads, nil
}


func (rep *ReposStruct) SelectThreadsByForum(forum string, limit string, since string, desc string) (Threads *models.Threads, Err error) {
	threads := models.Threads{}
	var rows *sql.Rows
	var err error
	if since == "" && desc == "false" {
		rows, err = rep.DataBase.Query(consts.SELECTThreadsByForum, forum, limit)
	} else if since != "" && desc == "false" {
		rows, err = rep.DataBase.Query(consts.SELECTThreadsByForumSince, forum, limit, since)
	} else if since == "" && desc == "true" {
		rows, err = rep.DataBase.Query(consts.SELECTThreadsByForumDesc, forum, limit)
	} else {
		rows, err = rep.DataBase.Query(consts.SELECTThreadsByForumSinceDesc, forum, limit, since)
	}
	defer rows.Close()
	if err != nil {
		return &threads, err
	}

	for rows.Next() {
		scanThread := models.Thread{}
		err := rows.Scan(&scanThread.Author, &scanThread.Created, &scanThread.Forum,
			&scanThread.ID, &scanThread.Message, &scanThread.Slug, &scanThread.Title,
			&scanThread.Votes)
		if err != nil {
			return &threads, err
		}
		threads = append(threads, &scanThread)
	}
	return &threads, nil
}

func (rep *ReposStruct) SelectUsersByForum(slug, limit, desc string) (Users *models.Users, Err error) {
	var users models.Users
	var rows *sql.Rows
	var err error
	if desc == "false" {
		rows, err = rep.DataBase.Query(consts.SELECTUsersByForumSlug, slug, limit)
	} else {
		rows, err = rep.DataBase.Query(consts.SELECTUsersByForumSlugDesc, slug, limit)
	}
	defer rows.Close()
	if err != nil {
		return &users, err
	}


	for rows.Next() {
		scanUser := models.User{}
		err := rows.Scan(&scanUser.About, &scanUser.Email, &scanUser.Fullname,
			&scanUser.Nickname)
		if err != nil {
			return &users, err
		}
		users = append(users, &scanUser)
	}


	sort.Slice(users, func(i, j int) bool { return strings.ToLower(users[i].Nickname) < strings.ToLower(users[j].Nickname) })

	return &users, nil
}

func (rep *ReposStruct) SelectUsersByNicknameOrEmail(email string, nickname string) (Users []models.User, Err error) {
	var users []models.User
	rows, err := rep.DataBase.Query(consts.SELECTUsersByNicknameOrEmail, email, nickname)
	defer rows.Close()
	if err != nil {
		return users, err
	}

	scanUser := models.User{}
	for rows.Next() {
		err := rows.Scan(&scanUser.About, &scanUser.Email, &scanUser.Fullname,
			&scanUser.Nickname)
		if err != nil {
			return users, err
		}
		users = append(users, scanUser)
	}
	return users, nil
}

func (rep *ReposStruct) SelectUserByNickname(nickname string) (user models.User, Err error) {
	var users []models.User
	rows, err := rep.DataBase.Query(consts.SELECTUsersByNickname, nickname)
	defer rows.Close()
	if err != nil {
		return models.User{}, err
	}

	scanUser := models.User{}
	for rows.Next() {
		err := rows.Scan(&scanUser.About, &scanUser.Email, &scanUser.Fullname,
			&scanUser.Nickname)
		if err != nil {
			return models.User{}, err
		}
		users = append(users, scanUser)
	}

	if len(users) == 0 {
		return models.User{}, errors.New("Can't find user by nickname")
	}
	return users[0], nil
}

func (rep *ReposStruct) SelectUsersByEmail(email string) (Users []models.User, Err error) {
	var users []models.User
	rows, err := rep.DataBase.Query(consts.SELECTUsersByEmail, email)
	defer rows.Close()
	if err != nil {
		return users, err
	}

	scanUser := models.User{}
	for rows.Next() {
		err := rows.Scan(&scanUser.About, &scanUser.Email, &scanUser.Fullname,
			&scanUser.Nickname)
		if err != nil {
			return users, err
		}
		users = append(users, scanUser)
	}
	return users, nil
}

