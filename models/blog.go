package models

import (
	"fmt"
	"time"
)

var (
	// DB代用
	DB_Blog []Blog
)

type Blog struct {
	ID int `json:"id"`
	// [0] Cast
	// [1] Staff
	// [2] Event
	WhoIs   int `json:"whois"`   // Cast or Staffなど判別
	Name_id int `json:"name_id"` // 上記カテゴリのid

	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Filenames []string `json:"filenames"`

	Public    bool      `json:"public"` // 公開設定
	CreatedAt time.Time `json:"createdAt"`
	ChangedAt time.Time `json:"changedAt"`
}

type dbChange interface {
	GetBlogFromDB() []Blog
	SaveBlogToDB() error
}

// DBがnilなのでサンプルのIDだけのものを作る
func init() {
	var ids [20]int
	for i := range ids {
		var a Blog
		a.ID = i
		a.WhoIs = i
		DB_Blog = append(DB_Blog, a)
	}
}

// Request Blogs from DB
func GetBlogFromDB() []Blog {

	//
	//

	return DB_Blog
}

// Blog set to DB
// データベースに登録する関数を一箇所で管理するため
func SaveBlogsToDB(a []Blog) error {
	DB_Blog = a

	return nil
}

// Blog set to DB
// データベースに登録する関数を一箇所で管理するため
func SaveBlogToDB(a Blog) error {
	DB_Blog = append(DB_Blog, a)

	return nil
}

// 全ブログを取得する
func GetBlogs() []Blog {
	blogs := GetBlogFromDB()

	return blogs
}

// 投球されたIDからブログを特定する
func GetBlogByID(id int) Blog {
	blogs := GetBlogFromDB()
	for i, v := range blogs {
		if id == v.ID {
			// idとBlog_idが合致したらBreak
			id = i
			break
		}
	}

	return blogs[id]
}

// 投稿者IDから投稿ブログを取得する
func GetBlogsByWhois(whois int) []Blog {
	blogs := GetBlogFromDB()
	var posts []Blog
	for _, v := range blogs {
		if whois == v.WhoIs {
			// idとBlog_idが合致したらBreak
			posts = append(posts, v)
		}
	}

	// 合致がなかったら、空っぽ
	return posts
}

func PostBlog(a Blog) error {
	if err := SaveBlogToDB(a); err != nil {
		return err
	}

	return nil
}

// ブログ情報を変更するため
func PutBlog(a Blog) error {
	blogs := GetBlogs()
	for i, v := range blogs {
		if a.ID == v.ID {
			blogs[i] = a

			if err := SaveBlogsToDB(blogs); err != nil {
				return err
			}
			break
		}
	}

	return nil
}

// ブログを削除する
func DelBlog(id int) error {
	// 削除有無判断
	var b bool
	blogs := GetBlogs()
	for i, v := range blogs {
		if id == v.ID {
			blogs = append(blogs[:i], blogs[i+1:]...)

			if err := SaveBlogsToDB(blogs); err != nil {
				return err
			}
			b = true
			break
		}
	}

	if b != false {
		return nil
	} else {
		return fmt.Errorf("削除要望のIDが存在しない")
	}
}
