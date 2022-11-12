package controller

import (
	"fmt"
	"net/http"
	"notchman8600/authentication-provider/domain"
	"notchman8600/authentication-provider/usecase"

	"github.com/google/uuid"
)

type AuthController struct {
	OAuthInteractor usecase.OAuthInteractor
}

var SessionList = make(map[string]domain.Session)

// refer: https://github.com/sat0ken/goauth-server/blob/main/main.go
func (controller *AuthController) Auth(req *http.Request) (session domain.Session, err error) {

	query := req.URL.Query()
	session = domain.Session{
		ClientId:    query.Get("client_id"),
		State:       query.Get("state"),
		Scopes:      query.Get("scope"),
		RedirectURL: query.Get("redirect_url"),
	}

	requiredParameters := []string{"response_type", "client_id", "redirect_uri"}
	for _, param := range requiredParameters {
		if !query.Has(param) {
			//TODO エラー処理の文章を変える
			err = fmt.Errorf("error! There is no required parameters")
			return
		}
	}

	// ClientIdの取得
	_, err = controller.OAuthInteractor.FindByClientId(session.ClientId)
	if err != nil {
		err = fmt.Errorf("error! client_id is not match")
		return
	}
	if query.Get("scope") == "openid profile" {
		session.OIDC = true
	} else {
		session.CodeChallenge = query.Get("code_challenge")
		session.CodeChallengeMethod = query.Get("code_challenge_method")
	}
	// セッションIDを生成
	sessionId := uuid.New().String()
	// セッション情報を保存しておく
	// TODO いや、これはDBに保存するべき？
	SessionList[sessionId] = session
	return
}
