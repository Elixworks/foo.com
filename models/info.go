package models

type SiteInfo struct {
	// HTMLメタ情報
	Sitename  string `json:"sitename"`
	Domain    string `json:"domain"`
	MetaSeo   string `json:"metaSeo"`
	Analytics string `json:"analytics"`

	// ページ名
	Pagename string `json:"pagename"`
	// 管理者ログイン
	Admin bool `json:"admin"`

	// 店舗情報
	ShopInfo ShopInfo `json:"shopInfo"`

	// デザイン関係
	Design Design `json:"design"`

	// SNS関連
	Sns Sns `json:"sns"`
}

func (s SiteInfo) New() SiteInfo {
	s.Sitename = "サイト名"
	s.Domain = "http://abc.co.jp"
	s.Analytics = "UX-xxx"

	return s
}
