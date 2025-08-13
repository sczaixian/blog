package response

import (
	"blog/server/models"
)

type UserResponse struct {
	User models.User `json:"user"`
}

type LoginResponse struct {
	User      models.User `json:"user"`
	Token     string      `json:"token"`
	ExpiresAt int64       `json:"expires_at"`
}
