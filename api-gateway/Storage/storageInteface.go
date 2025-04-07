package Storage

import "api-gateway/domain"

type Repo interface {
	CreateUser(user domain.User) (domain.User, error)
	GetUserByEmail(email string) (domain.User, error)
}
