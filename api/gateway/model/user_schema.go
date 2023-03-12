package model

type (
	// usersテーブルのレコード
	User struct {
		ManualIDModel
		Password  string
		Admin     bool
		Documents []Document
	}

	// documentsテーブルのレコード
	Document struct {
		AutoUUIDModel
		UserID string // 所有者（ユーザー）のID
		Title  string // タイトル
		Body   string // 本文
	}
)
