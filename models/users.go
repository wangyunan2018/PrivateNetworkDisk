package models

// user表结构体
type Users struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// 把用户信息存储到切片中
var Slice []Users

// 用于临时存储用户登录信息的Map
var State = make(map[string]interface{})
