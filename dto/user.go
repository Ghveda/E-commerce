package dto

type AuthBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
