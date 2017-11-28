package models

import "time"

type GetInfoCast struct {
	// - ブログは、キャストIDから配列取得し、Length制限
	// - スケジュールも、キャストIDから取得
	ID      int      `json:"id"`
	NameG   string   `json:"nameG"` // 源氏名
	NameE   string   `json:"nameE"` // ヘボンor英字
	Age     int      `json:"age"`
	Comment string   `json:"comment"` // 公開用メモ
	Hidden  struct { // 非公開情報
		Name  string `json:"name"`
		Tel   string `json:"tel"`
		Email string `json:"email"`
		Memo  string `json:"memo"` // 内部用メモ
	} `json:"hidden"`
	Body struct { // 身体情報
		Height int    `json:"h"`   // 身長
		B      int    `json:"b"`   // バストサイズ
		Cup    string `json:"cup"` // バストカップ
		W      int    `json:"w"`   // ウエスト
		H      int    `json:"h"`   // ヒップ
	} `json:"body"`
	Meta struct { // メタ情報
		Birth         string   `json:"birth"`         // 誕生日
		Blood         string   `json:"blood"`         // 血液型
		Constellation string   `json:"constellation"` // 星座
		Job           string   `json:"job"`           // 前職
		Hobby         string   `json:"hobby"`         // 趣味
		Char          string   `json:"char"`          // 性格
		Charm         string   `json:"charm"`         // チャームポイント
		Smoke         string   `json:"smoke"`         // 喫煙有無
		Erog          string   `json:"erog"`          // 性感帯
		Play          string   `json:"play"`          // 得意プレイ
		Sm            string   `json:"sm"`            // S or M
		Options       []string `json:"options"`       // 可能オプション
	} `json:"meta"`

	Filenames []string `json:"filenames"` // 画像
	Moviename string   `json:"moviename"` // 動画

	Public    bool      `json:"public"` // 公開設定
	CreatedAt time.Time `json:"createdAt"`
	ChangedAt time.Time `json:"changedAt"`
}
