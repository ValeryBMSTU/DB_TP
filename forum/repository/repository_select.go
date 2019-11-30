package repository

import (
	"bytes"
	"errors"
	"github.com/ValeryBMSTU/DB_TP/pkg/consts"
	"github.com/ValeryBMSTU/DB_TP/pkg/models"
	"github.com/jackc/pgx"
	xsort "sort"
	"strconv"
	"strings"
	"time"
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

func (rep *ReposStruct) SelectPostByID(ID int) (Post models.Post, Err error) {
	var posts []models.Post
	rows, err := rep.DataBase.Query(consts.SELECTPostsByID, ID)
	defer rows.Close()
	if err != nil {
		return models.Post{}, err
	}

	scanPost := models.Post{}
	for rows.Next() {
		var timetz time.Time
		err := rows.Scan(&scanPost.Author, &timetz, &scanPost.Forum,
			&scanPost.ID, &scanPost.IsEdited, &scanPost.Message, &scanPost.Parent,
			&scanPost.Thread)
		if err != nil {
			return models.Post{}, err
		}
		timetz.Format(time.RFC3339Nano)
		posts = append(posts, scanPost)
	}

	if len(posts) == 0 {
		return models.Post{}, errors.New("Can't find user by nickname")
	}
	return posts[0], nil
}

func (rep *ReposStruct) SelectPostByIDThreadID(ID int, threadID int) (Post models.Post, Err error) {
	var posts []models.Post
	rows, err := rep.DataBase.Query(consts.SELECTPostsByIDThreadID, ID, threadID)
	defer rows.Close()
	if err != nil {
		return models.Post{}, err
	}

	scanPost := models.Post{}
	for rows.Next() {
		var timetz time.Time
		err := rows.Scan(&scanPost.Author, &timetz, &scanPost.Forum,
			&scanPost.ID, &scanPost.IsEdited, &scanPost.Message, &scanPost.Parent,
			&scanPost.Thread)
		if err != nil {
			return models.Post{}, err
		}
		scanPost.Created = timetz.Format(time.RFC3339Nano)
		posts = append(posts, scanPost)
	}

	if len(posts) == 0 {
		return models.Post{}, errors.New("Can't find user by nickname")
	}
	return posts[0], nil
}

func (rep *ReposStruct) SelectPosts(threadID int, limit, since, sort, desc string) (Posts *models.Posts, Err error) {
	posts := models.Posts{}

	var rows *pgx.Rows
	var err error
	if sort == "flat" {
		if desc == "false" {
			rows, err = rep.DataBase.Query(consts.SELECTPostsFlat, threadID, limit, since)
		} else {
			rows, err = rep.DataBase.Query(consts.SELECTPostsFlatDesc, threadID, limit, since)
		}

	} else if sort == "tree" {
		if desc == "false" {
			if since != "0" && since != "999999999" {
				rows, err = rep.DataBase.Query(consts.SELECTPostsTree, threadID,  100000)
			} else {
				rows, err = rep.DataBase.Query(consts.SELECTPostsTree, threadID, limit)
			}
		} else {
			if since != "0" && since != "999999999" {
				rows, err = rep.DataBase.Query(consts.SELECTPostsTreeSinceDesc, threadID)
			} else {
				rows, err = rep.DataBase.Query(consts.SELECTPostsTreeDesc, threadID, limit, 1000000)
			}
		}
	} else if sort == "parent_tree" {
		if desc == "false" {
			rows, err = rep.DataBase.Query(consts.SELECTPostsParentTree, threadID)
		} else {
			rows, err = rep.DataBase.Query(consts.SELECTPostsParentTreeDesc, threadID)
		}
	}

	if sort != "parent_tree" {
		defer rows.Close()
		if err != nil {
			return &posts, err
		}

		for rows.Next() {
			scanPost := models.Post{}
			var timetz time.Time
			err := rows.Scan(&scanPost.Author, &timetz, &scanPost.Forum,
				&scanPost.ID, &scanPost.IsEdited, &scanPost.Message, &scanPost.Parent,
				&scanPost.Thread)
			if err != nil {
				return &posts, err
			}
			scanPost.Created = timetz.Format(time.RFC3339Nano)
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
			var timetz time.Time
			err := rows.Scan(&scanPost.Author, &timetz, &scanPost.Forum,
				&scanPost.ID, &scanPost.IsEdited, &scanPost.Message, &scanPost.Parent,
				&scanPost.Thread)
			if err != nil {
				return &posts, err
			}

			if scanPost.Parent == 0 {
				count = count + 1
			}
			if count > limitDigit && (since == "0" || since == "999999999") {
				break
			} else {
				scanPost.Created = timetz.Format(time.RFC3339Nano)
				posts = append(posts, &scanPost)
			}

		}
		rows.Close()
	}

	if since != "0" && since != "999999999" && sort == "tree"{
		limitDigit, _ := strconv.Atoi(limit)
		sinceDigit, _ := strconv.Atoi(since)
		sincePosts := models.Posts{}
		counter := 0
		//for ; posts[counter].ID <= sinceDigit && counter < len(posts); {
		//	counter++
		//}
		if desc == "false" {
			startIndex := 1000000000
			//postMinStartIndex
			minValue := 100000000000
			for i := 0; i < len(posts); i++ {
				if (posts[i].ID == sinceDigit) {
					startIndex = i + 1
					break
				}
				if (posts[i].ID > sinceDigit) && (posts[i].ID < minValue) {
					startIndex = i
					minValue = posts[i].ID
				}
			}
			sincePostsCount := 0
			counter = startIndex
			for ; sincePostsCount < limitDigit && counter < len(posts); {
				scanPost := models.Post{}
				scanPost = *posts[counter]
				sincePosts = append(sincePosts, &scanPost)
				if sort == "tree" {
					sincePostsCount++
				} else {
					if scanPost.Parent == 0 {
						sincePostsCount++
					}
				}
				counter++
			}
		} else {
			startIndex := -1000000000
			//postMinStartIndex
			maxValue := 0
			for i := len(posts) - 1; i >= 0; i-- {
				if (posts[i].ID == sinceDigit) {
					startIndex = i - 1
					break
				}
				if (posts[i].ID < sinceDigit) && (posts[i].ID > maxValue) {
					startIndex = i
					maxValue = posts[i].ID
				}
			}

			//xsort.Slice(posts[0:startIndex], func(i, j int) bool { return posts[i].ID < posts[j].ID})
			sincePostsCount := 0
			counter = startIndex
			for ; sincePostsCount < limitDigit && counter >= 0; {
				scanPost := models.Post{}
				scanPost = *posts[counter]
				sincePosts = append(sincePosts, &scanPost)
				if sort == "tree" {
					sincePostsCount++
				} else {
					if scanPost.Parent == 0 {
						sincePostsCount++
					}
				}
				counter--
			}
		}
		return  &sincePosts, nil
	}

	if since != "0" && since != "999999999" && sort == "parent_tree" {
		limitDigit, _ := strconv.Atoi(limit)
		sinceDigit, _ := strconv.Atoi(since)
		sincePosts := models.Posts{}
		counter := 0
		if desc == "false" {
			startIndex := 1000000000
			minValue := 100000000000
			for i := 0; i < len(posts); i++ {
				if (posts[i].ID == sinceDigit) {
					startIndex = i + 1
					break
				}
				if (posts[i].ID > sinceDigit) && (posts[i].ID < minValue) {
					startIndex = i
					minValue = posts[i].ID
				}
			}
			sincePostsCount := 0
			counter = startIndex
			for ; sincePostsCount < limitDigit && counter < len(posts); {
				scanPost := models.Post{}
				scanPost = *posts[counter]
				sincePosts = append(sincePosts, &scanPost)
				sincePostsCount++
				counter++
			}
		} else {
			startIndex := -1000000000
			//postMinStartIndex
			maxValue := 100000000000
			for i := len(posts) - 1; i >= 0; i-- {
				if (posts[i].ID == sinceDigit) {
					startIndex = i + 1
					break
				}
				if (posts[i].ID < sinceDigit) && (posts[i].ID < maxValue) {
					startIndex = i
					maxValue = posts[i].ID
				}
			}

			//xsort.Slice(posts[0:startIndex], func(i, j int) bool { return posts[i].ID < posts[j].ID})
			sincePostsCount := 0
			counter = startIndex
			for ; sincePostsCount < limitDigit && counter < len(posts); {
				scanPost := models.Post{}
				scanPost = *posts[counter]
				sincePosts = append(sincePosts, &scanPost)
				if sort == "tree" {
					sincePostsCount++
				} else {
					if scanPost.Parent == 0 {
						sincePostsCount++
					}
				}
				counter++
			}
		}
		return  &sincePosts, nil
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
		var timetz time.Time
		err := rows.Scan(&scanThread.Author, &timetz, &scanThread.Forum,
			&scanThread.ID, &scanThread.Message, &scanThread.Slug, &scanThread.Title,
			&scanThread.Votes)
		if err != nil {
			return &threads, err
		}
		scanThread.Created = timetz.Format(time.RFC3339Nano)
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
		var timetz time.Time
		err := rows.Scan(&scanThread.Author, &timetz, &scanThread.Forum,
			&scanThread.ID, &scanThread.Message, &scanThread.Slug, &scanThread.Title,
			&scanThread.Votes)
		if err != nil {
			return &threads, err
		}
		scanThread.Created = timetz.Format(time.RFC3339Nano)
		threads = append(threads, &scanThread)
	}
	return &threads, nil
}


func (rep *ReposStruct) SelectThreadsByForum(forum string, limit string, since string, desc string) (Threads *models.Threads, Err error) {
	threads := models.Threads{}
	var rows *pgx.Rows
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
		var timetz time.Time
		err := rows.Scan(&scanThread.Author, &timetz, &scanThread.Forum,
			&scanThread.ID, &scanThread.Message, &scanThread.Slug, &scanThread.Title,
			&scanThread.Votes)
		if err != nil {
			return &threads, err
		}
		scanThread.Created = timetz.Format(time.RFC3339Nano)
		threads = append(threads, &scanThread)
	}
	return &threads, nil
}

func (rep *ReposStruct) SelectUsersByForum(slug, limit, since, desc string) (Users *models.Users, Err error) {
	users := models.Users{}
	var rows *pgx.Rows
	var err error
	//if since == "" {
		if desc == "false" {
			rows, err = rep.DataBase.Query(consts.SELECTUsersByForumSlug, slug)
		} else {
			rows, err = rep.DataBase.Query(consts.SELECTUsersByForumSlugDesc, slug)
		}
	//} else {
	//	if desc == "false" {
	//		rows, err = rep.DataBase.Query(consts.SELECTUsersByForumSlugSince, slug)
	//	} else {
	//		rows, err = rep.DataBase.Query(consts.SELECTUsersByForumSlugSinceDesc, slug)
	//	}
	//}
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

	//ab := []byte(strings.ToLower(users[0].Nickname))
	//	println(ab)



	resUsers := models.Users{}

	limitDigit, _ := strconv.Atoi(limit)

	if desc == "false" {

		xsort.Slice(users, func(i, j int) bool { return bytes.Compare([]byte(strings.ToLower(users[i].Nickname)),[]byte(strings.ToLower(users[j].Nickname))) < 0})


		if since == "" {
			for i := 0; i < limitDigit && i < len(users); i++ {
				resUsers = append(resUsers, users[i])
			}
		} else {
			j := 0
			for i := 0; j < limitDigit && i < len(users); {
				if bytes.Compare([]byte(strings.ToLower(users[i].Nickname)), []byte(strings.ToLower(since))) > 0 {
					resUsers = append(resUsers, users[i])
					j++
				}
				i++
			}
		}
	} else {

		xsort.Slice(users, func(i, j int) bool { return bytes.Compare([]byte(strings.ToLower(users[i].Nickname)),[]byte(strings.ToLower(users[j].Nickname))) > 0})


		if since == "" {
			for i := 0; i < limitDigit && i < len(users); i++ {
				resUsers = append(resUsers, users[i])
			}
		} else {
			j := 0
			for i := 0; j < limitDigit && i < len(users); {
				if bytes.Compare([]byte(strings.ToLower(users[i].Nickname)), []byte(strings.ToLower(since))) < 0 {
					resUsers = append(resUsers, users[i])
					j++
				}
				i++
			}
		}
	}



	return &resUsers, nil
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

func (rep *ReposStruct) SelectStatus() (Status models.Status, Err error) {
	err := rep.DataBase.QueryRow(consts.SELECTStatus).Scan(&Status.Post, &Status.Thread, &Status.User, &Status.Forum)
	if err != nil {
		return Status, err
	}
	return
}

