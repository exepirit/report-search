package index

import (
	"github.com/exepirit/report-search/internal/data"
)

func MapUserToIndex(user data.User) User {
	return User{
		ID:        user.ID,
		ShortName: user.ShortName,
	}
}

func MapUserFromIndex(user User) data.User {
	return data.User{
		ID:        user.ID,
		ShortName: user.ShortName,
	}
}
