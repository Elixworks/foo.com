package models

import (
	"testing"
)

func TestGetBlogFromDB(t *testing.T) {
	blogs := GetBlogFromDB()

	if len(blogs) == 0 {
		t.Errorf("DBが空ですやん")
	}
}

func TestSaveBlogToDB(t *testing.T) {
	a := Blog{
		ID:    1,
		Title: "タイトル",
	}

	if err := SaveBlogToDB(a); err != nil {
		t.Errorf("セットできず \n %v", err)
	}
}

func TestGetBlogs(t *testing.T) {
	blogs := GetBlogs()

	if len(blogs) == 0 {
		t.Errorf("DBが空ですやん")
	}
}

func TestGetBlogByID(t *testing.T) {
	ar := []int{0, 1, 5, 7, 9, 11, 13, 17, 19}
	for _, v := range ar {
		blog := GetBlogByID(v)
		if v != blog.ID {
			t.Errorf("求めるID: %v, But データのID: %v", v, blog.ID)
			if blog.Title == "" {
				t.Errorf("データが空だからね")
			}
			break
		}
	}
}

func TestGetBlogsByWhois(t *testing.T) {
	ar := []int{0, 1, 5, 7, 9, 11, 13, 17, 19}

	for _, v := range ar {
		blogs := GetBlogsByWhois(v)
		if len(blogs) != 0 {
			t.Logf("削除できず\nwhois id: %dは存在しない", v)
		}

		for _, value := range blogs {
			if v != value.WhoIs {
				t.Errorf("求めるID: %v, Cause データのID: %v", v, value.ID)
				if value.Title == "" {
					t.Errorf("だって、データが空だからね")
				}
				break
			}
		}
	}
}

func TestPostBlog(t *testing.T) {
	title := "タイトルを投稿"
	a := Blog{
		WhoIs: 2,
		Title: title,
	}

	if err := PostBlog(a); err != nil {
		t.Errorf("投稿できず: %s", err)
	}

	// 検知用
	var this []Blog
	blogs := GetBlogsByWhois(2)
	for _, v := range blogs {
		if title == v.Title {
			this = append(this, v)
		}
	}

	// 合致したものだけで検証
	for _, v := range this {
		if title != v.Title {
			t.Logf("合致せず: %s", v.Title)
		}
	}
}

func TestPutBlog(t *testing.T) {
	id := 1
	title := "タイトルに変更"
	a := Blog{
		ID:    id,
		Title: title,
	}

	if err := PutBlog(a); err != nil {
		t.Errorf("PutBlog: %s", err)
	}

	blog := GetBlogByID(id)
	if blog.Title != title {
		t.Errorf("変更できず")
	}
}

func TestDelBlog(t *testing.T) {
	var ar = []int{0, 1, 5, 7, 9, 11, 13, 17, 19}

	for _, v := range ar {
		if err := DelBlog(v); err != nil {
			t.Errorf("why")
			break
		}
	}
}
