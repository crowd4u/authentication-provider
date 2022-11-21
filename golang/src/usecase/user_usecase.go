package usecase

import (
	"notchman8600/authentication-provider/domain"
)

// TODO domainを定義する
type UserRepository interface {
	FindByUserId(string) (domain.User, error)
}

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) FindByUserId(userId string) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindByUserId(userId)
	return
}
