// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

// input要素
type NewTodo struct {
	// タスクの説明
	Text string `json:"text"`
	// ユーザid
	UserID string `json:"userId"`
}

// ユーザ情報
type User struct {
	// ユーザid
	ID string `json:"id"`
	// 名前
	Name string `json:"name"`
}