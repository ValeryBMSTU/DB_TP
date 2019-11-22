package consts

const (
	SELECTUserIDUsernameEmailByUsernameOrEmail = "SELECT u.id, u.username, u.email from sunrise.user as u where u.username = $1 OR u.email = $2"
	SELECTAllUsers                             = "SELECT * from sunrise.user"
	UPDATEUserByID                             = "UPDATE sunrise_db.sunrise.user SET username = $1, name = $2, 	surname = $3," +
		" hashpassword = $4,email = $5, age = $6, status = $7 where id = $8"
	UPDATEUserAvatarDirByID = "UPDATE sunrise_db.sunrise.user SET avatardir = $1 where id = $2"

	INSERTForum = `INSERT INTO forum.forum (slug,title,"user") values ($1,$2,$3) RETURNING id;`
	SELECTForumsBySlug   = `SELECT f.posts, f.slug, f.threads, f.title, f.user ` +
		`FROM forum.forum as f WHERE lower(f.slug) = lower($1);`

	INSERTPost = "INSERT INTO forum.post (author, message, parent, thread, forum) " +
		"VALUES ($1,$2,$3,$4,$5) RETURNING id, thread;"
	SELECTPostsByID = "SELECT p.author, p.created, p.forum, p.id, p.isEdited, p.message, p.parent, p.thread " +
		"FROM forum.post as p WHERE p.id = $1"
	SELECTPostsByIDThreadID = "SELECT p.author, p.created, p.forum, p.id, p.isEdited, p.message, p.parent, p.thread " +
		"FROM forum.post as p WHERE p.id = $1 AND p.thread = $2"


	SELECTPostsFlat = "SELECT p.author, p.created, p.forum, p.id, p.isEdited, p.message, p.parent, p.thread " +
		"FROM forum.post as p WHERE p.thread = $1 AND p.id > $3 ORDER BY p.id LIMIT $2"
	SELECTPostsFlatDesc = "SELECT p.author, p.created, p.forum, p.id, p.isEdited, p.message, p.parent, p.thread " +
		"FROM forum.post as p WHERE p.thread = $1 AND p.id < $3 ORDER BY p.id DESC LIMIT $2"


	SELECTPostsTree = "WITH RECURSIVE temp1 (author, created, forum, id, isEdited, message, parent, thread, PATH, LEVEL, root ) AS ( " +
"SELECT T1.author, T1.created, T1.forum, T1.id, T1.isEdited, T1.message, T1.parent, T1.thread, CAST (10000 + T1.id AS VARCHAR (50)) as PATH, 1, T1.id as root " +
"FROM forum.post as T1 WHERE T1.parent = 0 and T1.thread = $1 " +
"union " +
"select T2.author, T2.created, T2.forum, T2.id, T2.isEdited, T2.message, T2.parent, T2.thread, CAST ( temp1.PATH ||'->'|| 10000 + T2.id AS VARCHAR(50)), LEVEL + 1, root " +
"FROM forum.post T2 INNER JOIN temp1 ON( temp1.id = T2.parent) " +
") " +
"select author, created, forum, id, isEdited, message, parent, thread from temp1 ORDER BY root, PATH LIMIT $2;"

	SELECTPostsTreeSince = "WITH RECURSIVE temp1 (author, created, forum, id, isEdited, message, parent, thread, PATH, LEVEL ) AS ( " +
		"SELECT T1.author, T1.created, T1.forum, T1.id, T1.isEdited, T1.message, T1.parent, T1.thread, CAST (1000000 + T1.id AS VARCHAR (50)) as PATH, 1 " +
		"FROM forum.post as T1 WHERE T1.parent = 0 AND T1.thread = $1"+
		"union " +
		"select T2.author, T2.created, T2.forum, T2.id, T2.isEdited, T2.message, T2.parent, T2.thread, CAST ( temp1.PATH ||'->'|| T2.id AS VARCHAR(50)), LEVEL + 1 " +
		"FROM forum.post T2 INNER JOIN temp1 ON( temp1.id = T2.parent) " +
		") " +
		"select author, created, forum, id, isEdited, message, parent, thread from temp1 WHERE id > $3 ORDER BY PATH LIMIT $2;"

	SELECTPostsTreeDesc = "WITH RECURSIVE temp1 (author, created, forum, id, isEdited, message, parent, thread, PATH, LEVEL, root ) AS ( " +
"SELECT T1.author, T1.created, T1.forum, T1.id, T1.isEdited, T1.message, T1.parent, T1.thread, CAST (1000000 + T1.id AS VARCHAR (50)) as PATH, 1, T1.id as root " +
"FROM forum.post as T1 WHERE T1.parent = 0 AND T1.thread = $1" +
"union " +
"select  T2.author, T2.created, T2.forum, T2.id, T2.isEdited, T2.message, T2.parent, T2.thread, CAST (temp1.PATH ||'->'|| T2.id AS VARCHAR(50)), LEVEL + 1, root " +
"FROM forum.post as T2 INNER JOIN temp1 ON (temp1.id = T2.parent) " +
") " +
"select author, created, forum, id, isEdited, message, parent, thread from temp1 WHERE id < $3 ORDER BY PATH DESC LIMIT $2;"
	SELECTPostsTreeSinceDesc = "WITH RECURSIVE temp1 (author, created, forum, id, isEdited, message, parent, thread, PATH, LEVEL, root ) AS ( " +
		"SELECT T1.author, T1.created, T1.forum, T1.id, T1.isEdited, T1.message, T1.parent, T1.thread, CAST (1000000 + T1.id AS VARCHAR (50)) as PATH, 1, T1.id as root " +
		"FROM forum.post as T1 WHERE T1.parent = 0 AND T1.thread = $1" +
		"union " +
		"select  T2.author, T2.created, T2.forum, T2.id, T2.isEdited, T2.message, T2.parent, T2.thread, CAST (temp1.PATH ||'->'|| T2.id AS VARCHAR(50)), LEVEL + 1, root " +
		"FROM forum.post as T2 INNER JOIN temp1 ON (temp1.id = T2.parent) " +
		") " +
		"select author, created, forum, id, isEdited, message, parent, thread from temp1 ORDER BY PATH;"

	SELECTPostsParentTree = "WITH RECURSIVE temp1 (author, created, forum, id, isEdited, message, parent, thread, PATH, LEVEL, root ) AS ( " +
		"SELECT T1.author, T1.created, T1.forum, T1.id, T1.isEdited, T1.message, T1.parent, T1.thread, CAST (10000 + T1.id AS VARCHAR (50)) as PATH, 1, T1.id as root " +
		"FROM forum.post as T1 WHERE T1.parent = 0 AND T1.thread = $1"+
		"union " +
		"select T2.author, T2.created, T2.forum, T2.id, T2.isEdited, T2.message, T2.parent, T2.thread, CAST ( temp1.PATH ||'->'|| 10000 + T2.id AS VARCHAR(50)), LEVEL + 1, root " +
		"FROM forum.post T2 INNER JOIN temp1 ON( temp1.id = T2.parent) " +
		") " +
		"select author, created, forum, id, isEdited, message, parent, thread from temp1 ORDER BY root, PATH;"

	SELECTPostsParentTreeDesc = "WITH RECURSIVE temp1 (author, created, forum, id, isEdited, message, parent, thread, PATH, LEVEL, root ) AS ( " +
"SELECT T1.author, T1.created, T1.forum, T1.id, T1.isEdited, T1.message, T1.parent, T1.thread, CAST (10000 + T1.id AS VARCHAR (50)) as PATH, 1, T1.id as root " +
		"FROM forum.post as T1 WHERE T1.parent = 0 AND T1.thread = $1" +
		"union " +
		"select  T2.author, T2.created, T2.forum, T2.id, T2.isEdited, T2.message, T2.parent, T2.thread, CAST ( temp1.PATH ||'->'|| 10000 + T2.id AS VARCHAR(50)), LEVEL + 1, root " +
		"FROM forum.post as T2 INNER JOIN temp1 ON( temp1.id = T2.parent) " +
		") " +
		"select author, created, forum, id, isEdited, message, parent, thread  from temp1 ORDER BY root desc, PATH;"

	UPDATEPostByID = "UPDATE forum.post SET message = $1, isEdited = $2 WHERE id = $3;"


	INSERTThread = `INSERT INTO forum.thread (author, created, message, title, forum) values ($1,$2,$3,$4,$5) RETURNING id;`
	INSERTThreadWithoutCreated = `INSERT INTO forum.thread (author, message, title, forum) values ($1,$2,$3,$4) RETURNING id;`
	INSERTThreadWithSlugWithoutCreated = `INSERT INTO forum.thread (author, message, title, forum, slug) values ($1,$2,$3,$4,$5) RETURNING id;`
	INSERTThreadWithSlug = `INSERT INTO forum.thread (author, created, message, title, forum, slug) values ($1,$2,$3,$4,$5,$6) RETURNING id;`
	UPDATEThreadByID = "UPDATE forum.thread SET message = $1, title = $2 WHERE id = $3;"

	SELECTThreadsByForum = `SELECT t.author, t.created, t.forum, t.id, t.message, t.slug, t.title, t.votes ` +
		`FROM forum.thread as t WHERE lower(t.forum) = lower($1) ORDER BY created LIMIT $2;`
	SELECTThreadsByForumSince = `SELECT t.author, t.created, t.forum, t.id, t.message, t.slug, t.title, t.votes ` +
		`FROM forum.thread as t WHERE lower(t.forum) = lower($1) AND t.created >= $3 ORDER BY created LIMIT $2;`
	SELECTThreadsByForumDesc = `SELECT t.author, t.created, t.forum, t.id, t.message, t.slug, t.title, t.votes ` +
		`FROM forum.thread as t WHERE lower(t.forum) = lower($1) ORDER BY created DESC LIMIT $2;`
	SELECTThreadsByForumSinceDesc = `SELECT t.author, t.created, t.forum, t.id, t.message, t.slug, t.title, t.votes ` +
		`FROM forum.thread as t WHERE lower(t.forum) = lower($1) AND t.created <= $3 ORDER BY created DESC LIMIT $2;`
	SELECTThreadsBySlug = `SELECT t.author, t.created, t.forum, t.id, t.message, t.slug, t.title, t.votes ` +
		`FROM forum.thread as t WHERE lower(t.slug) = lower($1);`
	SELECTThreadsByID = `SELECT t.author, t.created, t.forum, t.id, t.message, t.slug, t.title, t.votes ` +
		`FROM forum.thread as t WHERE t.id = $1;`

	INSERTUser              = "INSERT INTO forum.user (about, email, fullname, nickname) values ($1,$2,$3,$4) RETURNING id;"


	SELECTUsersByNickname   = `SELECT u.about, u.email, u.fullname, u.nickname ` +
		`FROM forum.user as u WHERE lower(u.nickname) = lower($1);`
	SELECTUsersByEmail = `SELECT u.about, u.email, u.fullname, u.nickname ` +
		`FROM forum.user as u WHERE lower(u.email) = lower($1);`
	SELECTUsersByNicknameOrEmail   = `SELECT u.about, u.email, u.fullname, u.nickname ` +
		`FROM forum.user as u WHERE lower(u.email) = lower($1) OR lower(u.nickname) = lower($2);`
	SELECTUsersByForumSlug =   "SELECT u.about, u.email, u.fullname, u.nickname " +
		`FROM forum."user" as u ` +
		"WHERE u.nickname IN ( " +
		"SELECT t.author AS nickname " +
		"FROM forum.thread as t " +
		"WHERE lower(t.forum) = lower($1) " +
		"UNION " +
		"SELECT p.author AS nickname " +
		"FROM forum.post as p " +
		"WHERE lower(p.forum) = lower($1) ) " +
		"ORDER BY lower(u.nickname) " +
		"LIMIT 100;"
	SELECTUsersByForumSlugDesc =   "SELECT u.about, u.email, u.fullname, u.nickname " +
		`FROM forum."user" as u ` +
		"WHERE u.nickname IN ( " +
		"SELECT t.author AS nickname " +
		"FROM forum.thread as t " +
		"WHERE lower(t.forum) = lower($1) " +
		"UNION " +
		"SELECT p.author AS nickname " +
		"FROM forum.post as p " +
		"WHERE lower(p.forum) = lower($1) ) " +
		"ORDER BY u.nickname DESC " +
		"LIMIT 100;"


UPDATEUserByNickname = "UPDATE forum.user SET about = $1, email = $2, fullname = $3 WHERE nickname = $4"



	INSERTVote = "INSERT INTO forum.vote (nickname, voice, thread) " +
		"VALUES ($1,$2,$3) RETURNING id;"
	UPDATEVote = "UPDATE forum.vote SET voice = $1 WHERE nickname = $2 AND thread = $3;"


	SELECTStatus =  "SELECT " +
"(SELECT COALESCE(SUM(forum.posts), 0) FROM forum.forum WHERE posts > 0) AS post, " +
"(SELECT COALESCE(SUM(forum.threads), 0) FROM forum.forum WHERE threads > 0) AS thread, " +
"(SELECT COUNT(*) FROM forum.user) AS user, " +
"(SELECT COUNT(*) FROM forum.forum) AS forum;"

	CLEARE = "TRUNCATE forum.vote, forum.post, forum.thread, forum.forum, forum.user RESTART IDENTITY CASCADE;"

	INSERTSession           = "INSERT INTO sunrise.usersession (userid, cookiesvalue, cookiesexpiration)	values ($1,$2,$3) RETURNING id"

	SELECTUserByCookieValue = "SELECT U.id, U.username, U.name, U.surname, U.hashpassword, U.email, U.age, U.status," +
		" U.avatardir, U.isactive, U.salt, U.created_time from sunrise.User as U JOIN sunrise.usersession as s on U.id = s.userid " +
		"where s.cookiesvalue = $1"
	SELECTCookiesByCookieValue = "SELECT s.cookiesvalue, s.cookiesexpiration from sunrise.usersession" +
		" as s where s.cookiesvalue = $1"
	SELECTUsersByUsername = "SELECT U.id, U.username, U.name, U.surname, U.hashpassword, U.email, U.age, U.status," +
		" U.avatardir, U.isactive, U.salt, U.created_time from sunrise.User as U where U.username = $1"

	DELETESessionByKey = "DELETE FROM sunrise.usersession as s WHERE s.cookiesvalue = $1"

	SELECTCategoryByName = "SELECT c.name FROM sunrise.category as c WHERE c.name = $1"
	INSERTBoard          = "INSERT INTO sunrise.board (owner_id, title, description, category, createdTime) VALUES ($1,$2,$3,$4,$5) RETURNING id"
	SELECTBoardByID      = "SELECT b.id, b.owner_id, b.title, b.description, b.category, b.createdTime, b.isDeleted " +
		"FROM sunrise.board as b WHERE b.id = $1"
	SELECTBoardsByOwnerId = "SELECT b.id, b.owner_id, b.title, b.description, b.category, b.createdTime, b.isDeleted " +
		"FROM sunrise.board as b WHERE b.owner_id = $1"

	INSERTPin = "INSERT INTO sunrise.pin (owner_id, author_id, board_id, title, description, pindir, createdTime)" +
		" VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id"
	SELECTPinByID = "SELECT p.id, o.username, a.username, p.board_id, p.title, p.description, p.pindir, p.createdTime, p.isDeleted " +
		"FROM sunrise.pin as p join sunrise.user as a on p.author_id = a.id join sunrise.user as o on p.owner_id = o.id WHERE p.id = $1"
	SELECTPinsByBoardID = "SELECT p.id, p.owner_id, p.author_id, p.board_id, p.title," +
		" p.description, p.pindir, p.createdTime, p.isDeleted " +
		"FROM sunrise.pin as p WHERE p.board_id = $1"
	SELECTPinsDisplayByBoardId   = "SELECT p.id, p.title, p.pindir FROM sunrise.pin as p WHERE p.isDeleted = false AND p.board_id = $1"
	SELECTNewPinsDisplayByNumber = "SELECT p.id, p.pindir, p.title FROM (select id, pindir, title, isdeleted, ROW_NUMBER() OVER (ORDER BY createdtime desc) " +
		"from sunrise.pin WHERE isdeleted = false) as p WHERE p.ROW_NUMBER BETWEEN $1 AND $2;"
	//SELECTMyPinsByNumber = "SELECT p.id, p.pindir FROM (select id, pindir, isdeleted, ROW_NUMBER() OVER (ORDER BY createdtime) " +
	//	"from sunrise.pin WHERE owner_id = $2 AND isdeleted = false) as p WHERE p.ROW_NUMBER BETWEEN 0 AND $1;"
	SELECTMyPinsDisplayByNumber        = "SELECT p.id, p.pindir, p.title FROM sunrise.pin as p WHERE p.isdeleted = false and p.owner_id = $2 LIMIT $1;"
	SELECTSubscribePinsDisplayByNumber = "SELECT p.id, p.pindir, p.title FROM (select pin.id, pin.pindir, pin.title, pin.isdeleted, ROW_NUMBER() OVER (ORDER BY createdtime desc) " +
		"from sunrise.pin as pin join sunrise.subscribe as s on s.subscriber_id = $3 " +
		"AND s.followee_id = pin.owner_id " +
		"AND isdeleted = false) as p " +
		"WHERE p.ROW_NUMBER BETWEEN $1 AND $2;"
	SELECTCommentsByPinId = "SELECT c.text, u.username, u.avatardir, c.created_time FROM sunrise.comment as c join sunrise.pin as p on p.id = $1 " +
		"join sunrise.user as u on u.id = c.author_id where c.pin_id = $1 ORDER BY c.created_time"

	INSERTNotice          = "INSERT INTO sunrise.notice (user_id, receiver_id, message, createdTime) VALUES ($1,$2,$3,$4) RETURNING id"
	INSERTComment         = "INSERT INTO sunrise.comment (pin_id, text, author_id, created_time) VALUES ($1,$2,$3,$4) RETURNING id"
	INSERTSubscribeByName = "INSERT INTO sunrise.subscribe (subscriber_id, followee_id) " +
		"select $1, u.id from sunrise.user as u " +
		"where u.username = $2 " +
		"RETURNING id;"

	INSERTChatMessage = "INSERT INTO sunrise.chat_message (sender_id, receiver_id, text, send_time) " +
		"SELECT $1, u.id, $3, $4 from sunrise.user as u where u.username = $2 RETURNING id"

	DELETESubscribeByName = "DELETE FROM sunrise.subscribe as s WHERE s.subscriber_id = $1 and s.followee_id IN " +
		"(select u.id from sunrise.user as u " +
		"where u.username = $2);"

	SELECTPinsByTag = "SELECT DISTINCT p.id, p.pindir, p.title FROM sunrise.pin as p " +
		"JOIN sunrise.pinandtag as pt ON p.id = pt.pin_id " +
		"WHERE pt.tag_name = $1 AND p.isdeleted = false;"

	SELECTSessionByCookieValue = "SELECT s.id, s.userid FROM sunrise.usersession as s " +
		"WHERE s.cookiesvalue = $1;"
	SELECTMySubscribeByUsername = "SELECT s.id, s.subscriber_id, s.followee_id FROM sunrise.subscribe as s " +
		"join sunrise.user as u on u.id = s.followee_id WHERE s.subscriber_id = $1 AND u.username = $2;"

	SELECTNoticesByUserID = "SELECT n.id, n.user_id, n.receiver_id, n.message, n.createdtime, isread FROM sunrise.notice as n " +
		"WHERE n.receiver_id = $1 and n.isread = false;"

	SELECTTagAll = "SELECT t.name from sunrise.tag as t;"
	INSERTTag = "INSERT INTO sunrise.tag (name)" +
		" VALUES ($1) RETURNING name"
	INSERTPinAndTag = "INSERT INTO sunrise.pinandtag (pin_id, tag_name)" +
" VALUES ($1,$2) RETURNING id"

	DeletePosts = "DELETE FROM forum.post CASCADE"
	DeleteVotes = "DELETE FROM forum.vote CASCADE"
	DeleteThreads = "DELETE FROM forum.thread CASCADE"
	DeleteForums = "DELETE FROM forum.forum CASCADE"
	DeleteUsers = "DELETE FROM forum.user CASCADE"
)
