package models

// デザイン関係
type Design struct {
	Logo            string `json:"log"`
	Icon            string `json:"icon"`
	HeaderPrimary   string `json:"headerPrimary"`
	HeaderSecondary string `json:"headerSecondary"`

	// テンプレート名（cssなどを分岐させるために使用）
	Template string `json:"template"`
}
