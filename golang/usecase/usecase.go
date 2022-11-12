package usecase

import (
	"notchman8600/authentication-provider/domain"
)

// TODO domainを定義する
type OAuthRepository interface {
	FindByClientId(string) (domain.Client, error)
}

type OAuthInteractor struct {
	OAuthRepository OAuthRepository
}

func (interactor *OAuthInteractor) FindByClientId(clientId string) (client domain.Client, err error) {
	client, err = interactor.OAuthRepository.FindByClientId(clientId)
	return
}
