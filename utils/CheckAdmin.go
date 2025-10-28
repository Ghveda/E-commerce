package utils

import (
	"e-commerce/models"
)

func IsAdmin(currentUser any) (bool bool, user models.User) {
	user, ok := currentUser.(models.User)

	return ok && user.Role == "admin", user
}
