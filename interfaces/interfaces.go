package interfaces

import (
	"github.com/gorilla/sessions"
	"github.com/souther1407/blog/internal/database"
)

type ApiConfig struct {
	DB          *database.Queries
	CookieStore *sessions.CookieStore
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func DBUserToUser(u database.User) User {
	return User{
		Name:  u.Name,
		Email: u.Email,
	}
}
