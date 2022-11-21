package controller

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"notchman8600/authentication-provider/domain"
	"notchman8600/authentication-provider/persistence"
	"notchman8600/authentication-provider/usecase"
	"time"

	"github.com/google/uuid"
)

type AuthController struct {
	OAuthInteractor usecase.OAuthInteractor
}

var SessionList = make(map[string]domain.Session)

func base64URLEncode(verifier string) string {
	hash := sha256.Sum256([]byte(verifier))
	return base64.RawURLEncoding.EncodeToString(hash[:])
}

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

func (controller *AuthController) AuthCheck(w http.ResponseWriter, req *http.Request) {
	clientId := req.FormValue("client_id")

	_, err := controller.OAuthInteractor.FindByClientId(clientId)

	if err != nil {
		w.Write([]byte("login failed"))
	} else {

		cookie, _ := req.Cookie("session")
		http.SetCookie(w, cookie)
		v, _ := SessionList[cookie.Value]

		authCodeString := uuid.New().String()
		authData := domain.AuthCode{
			UserId:      clientId,
			ClientId:    v.ClientId,
			Scopes:      v.Scopes,
			RedirectURL: v.RedirectURL,
			ExpriresAt:  time.Now().Unix() + 300,
		}
		// 認可コードを保存
		persistence.AuthCodeList[authCodeString] = authData

		log.Printf("auth code accepet : %s\n", authData)

		location := fmt.Sprintf("%s?code=%s&state=%s", v.RedirectURL, authCodeString, v.State)
		w.Header().Add("Location", location)
		w.WriteHeader(302)
	}
}

// トークンを発行するエンドポイント
func (controller *AuthController) Token(w http.ResponseWriter, req *http.Request) {

	cookie, _ := req.Cookie("session")
	req.ParseForm()
	query := req.Form
	session := persistence.SessionList[cookie.Value]

	requiredParameter := []string{"grant_type", "code", "client_id", "redirect_uri"}
	// 必須パラメータのチェック
	for _, v := range requiredParameter {
		if !query.Has(v) {
			log.Printf("%s is missing", v)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("invalid_request. %s is missing\n", v)))
			return
		}
	}

	// 認可コードフローだけサポート
	if "authorization_code" != query.Get("grant_type") {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("invalid_request. not support type.\n")))
		return
	}

	// 保存していた認可コードのデータを取得。なければエラーを返す
	v, ok := persistence.AuthCodeList[query.Get("code")]
	if !ok {
		log.Println("auth code isn't exist")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("no authrization code")))
		return
	}

	// 認可リクエスト時のクライアントIDと比較
	if v.ClientId != query.Get("client_id") {
		log.Println("client_id not match")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("invalid_request. client_id not match.\n")))
		return
	}

	// 認可リクエスト時のリダイレクトURIと比較
	if v.RedirectURL != query.Get("redirect_uri") {
		log.Println("redirect_uri not match")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("invalid_request. redirect_uri not match.\n")))
		return
	}

	// 認可コードの有効期限を確認
	if v.ExpriresAt < time.Now().Unix() {
		log.Println("authcode expire")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("invalid_request. auth code time limit is expire.\n")))
		return
	}

	// クライアント情報の取得
	clientInfo, err := controller.OAuthInteractor.FindByClientId(session.ClientId)
	if err != nil {
		log.Println("user not found")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("invalid_request. user account is not found.\n")))
	}

	// clientシークレットの確認
	if clientInfo.Secret != query.Get("client_secret") {
		log.Println("client_secret is not match.")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("invalid_request. client_secret is not match.\n")))
		return
	}

	// PKCEのチェック
	// clientから送られてきたverifyをsh256で計算&base64urlエンコードしてから
	// 認可リクエスト時に送られてきてセッションに保存しておいたchallengeと一致するか確認
	if session.OIDC == false && session.CodeChallenge != base64URLEncode(query.Get("code_verifier")) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("PKCE check is err..."))
		return
	}

	tokenString := uuid.New().String()
	expireTime := time.Now().Unix() + persistence.ACCESS_TOKEN_DURATION

	tokenInfo := domain.TokenCode{
		UserId:    v.ClientId,
		ClientId:  v.ClientId,
		Scopes:    v.Scopes,
		ExpiresAt: expireTime,
	}
	// 払い出したトークン情報を保存
	persistence.TokenCodeList[tokenString] = tokenInfo
	// 認可コードを削除
	delete(persistence.AuthCodeList, query.Get("code"))

	tokenResp := domain.TokenResponse{
		AccessToken: tokenString,
		TokenType:   "Bearer",
		ExpiresIn:   expireTime,
	}
	if session.OIDC {
		tokenResp.IdToken, _ = utils.makeJWT()
	}
	resp, err := json.Marshal(tokenResp)
	if err != nil {
		log.Println("json marshal err")
	}

	log.Printf("token ok to client %s, token is %s", v.ClientId, string(resp))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)

}
