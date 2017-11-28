package models

// 店舗情報
type ShopInfo struct {
	Name    string `json:"name"`
	NameE   string `json:"nameE"`
	Manager string `json:"manager"` // 管理者や代表者名
	Email   string `json:"email"`
	Tel     string `json:"tel"`  // 営業用
	TelR    string `json:"telR"` // 求人用
	// 営業時間
	Start   string `json:"start"`
	End     string `json:"end"`
	Holiday string `json:"holiday"`
}
