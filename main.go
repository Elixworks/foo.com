package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	h "format-foo.com/controllers"
	"format-foo.com/lib"
)

const (
	PORT = ":8000"
)

// - Golang内ViewTemplate設定
type Template struct {
	templates *template.Template
}

// Render レシーバ
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// Template Functions
	f := funcMap.FunctionForTemplate()

	// filenameの前置詞で
	// Template判断
	var bytes1 []byte
	var err error
	if strings.Index(name, "primary") == 0 {
		bytes1, err = ioutil.ReadFile("./_views/_layouts/primary.gohtml")
		if err != nil {
			log.Fatal(err)
		}
	} else if strings.Index(name, "secondary") == 0 {
		bytes1, err = ioutil.ReadFile("./_views/_layouts/secondary.gohtml")
		if err != nil {
			log.Fatal(err)
		}
	} else if strings.Index(name, "admin") == 0 {
		bytes1, err = ioutil.ReadFile("./_views/_layouts/admin.gohtml")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		bytes1, err = ioutil.ReadFile("./_views/_layouts/default.gohtml")
		if err != nil {
			log.Fatal(err)
		}
	}

	// c.Renderが持ってくるOverlayを読む
	path := strings.Join([]string{"./_views/pages/", name}, "")
	bytes2, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	// ガッチャンコ
	tpl, err := template.New("template").Funcs(f).Parse(string(bytes1))
	if err != nil {
		log.Fatal(err)
	}
	overlay, err := template.Must(tpl.Clone()).Parse(string(bytes2))
	if err != nil {
		log.Fatal(err)
	}

	return overlay.Execute(w, data)
}

func main() {
	// Echoの初期化
	e := echo.New()

	// Middleware使用宣言
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 2,
	}))
	// 末尾のスラッシュ[/]有りのURIでもアクセス可能に
	e.Pre(middleware.RemoveTrailingSlash())
	// CORS クロスアクセス
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// Access設定
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// テンプレートレンダー
	t := &Template{
		templates: nil,
	}
	e.Renderer = t

	// 静的ディレクトリのルーティング
	e.Static("/", "_statics")

	// ページルーティング
	e.GET("/", h.Index)

	// APIルーティング
	// e.GET("/", h.GetInfo)

	// エターナルサーバー
	// TSLでhttps実装したい
	e.Logger.Fatal(e.Start(PORT))
}
