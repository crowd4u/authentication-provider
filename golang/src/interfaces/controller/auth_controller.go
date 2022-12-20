package controller

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"notchman8600/authentication-provider/domain"
	"notchman8600/authentication-provider/interfaces/database"
	"notchman8600/authentication-provider/persistence"
	"notchman8600/authentication-provider/usecase"
	"os"
	"time"

	"github.com/google/uuid"
)

type AuthController struct {
	OAuthInteractor usecase.OAuthInteractor
	UserInteractor  usecase.UserInteractor
}

func base64URLEncode(verifier string) string {
	hash := sha256.Sum256([]byte(verifier))
	return base64.RawURLEncoding.EncodeToString(hash[:])
}

func NewAuthController(sqlHandler database.DBHandler) *AuthController {
	return &AuthController{
		OAuthInteractor: usecase.OAuthInteractor{
			OAuthRepository: &database.OAuthRepository{
				DBHandler: sqlHandler,
			},
		},
		UserInteractor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				DBHandler: sqlHandler,
			},
		},
	}
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

	requiredParameters := []string{"client_id", "redirect_url"}
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
		print(fmt.Println(err))
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
	session.Id = sessionId
	// TODO いや、これはDBに保存するべき？
	persistence.SessionList[sessionId] = session

	return
}

func (controller *AuthController) AuthCheck(w http.ResponseWriter, req *http.Request) {
	clientId := req.FormValue("client_id")
	userId := req.FormValue("user_id")

	user, err := controller.UserInteractor.FindByUserId(userId)
	if err != nil {
		w.Write([]byte("login failed"))
		return
	}

	_, err = controller.OAuthInteractor.FindByClientId(clientId)

	if err != nil {
		w.Write([]byte("login failed"))
	} else {

		cookie, _ := req.Cookie("session")
		fmt.Println(cookie)
		http.SetCookie(w, cookie)
		v := persistence.SessionList[cookie.Value]

		authCodeString := uuid.New().String()
		authData := domain.AuthCode{
			UserId:      user.Id,
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

	requiredParameter := []string{"grant_type", "code", "client_id", "redirect_uri", "client_secret"}
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
	// TODO:バグ対応のため一旦オフに
	// if v.RedirectURL != query.Get("redirect_uri") {
	// 	log.Println("redirect_uri not match")
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte(fmt.Sprintf("invalid_request. redirect_uri not match.\n")))
	// 	return
	// }

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
	// if session.OIDC == false && session.CodeChallenge != base64URLEncode(query.Get("code_verifier")) {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte("PKCE check is err..."))
	// 	return
	// }
	user, err := controller.UserInteractor.FindByUserId(v.UserId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("user is not found"))
		return
	}
	tokenString := uuid.New().String()
	expireTime := time.Now().Unix() + persistence.ACCESS_TOKEN_DURATION

	tokenInfo := domain.TokenCode{
		UserId:    user.Id,
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
	// TODO放り込むパラメーターが合っているか確認

	if session.OIDC {
		tokenResp.IdToken, _ = makeJWT(user)
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

func (controller *AuthController) CheckToken(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	query := req.Form

	requiredParameter := []string{"token"}
	// 必須パラメータのチェック
	for _, v := range requiredParameter {
		if !query.Has(v) {
			log.Printf("%s is missing", v)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("invalid_request. %s is missing\n", v)))
			return
		}
	}

	// 保存していた認可コードのデータを取得。なければエラーを返す
	_, ok := persistence.TokenCodeList[query.Get("token")]
	if !ok {
		log.Println("auth code isn't exist")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("no authrization token")))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("success")))
}

func readPrivateKey() (*rsa.PrivateKey, error) {

	// 環境変数の読み込み
	fileName, err := readEnv("FILE-PRIVATE-KEY")
	if err != nil {
		return nil, err
	}
	//秘密鍵の読み込み
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	//秘密鍵のファイルのデコード
	keyblock, _ := pem.Decode(data)
	if keyblock == nil {
		return nil, fmt.Errorf("invalid private key data")
	}
	//RSA_PRIVATE_KEYか検査
	if keyblock.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("invalid private key type : %s", keyblock.Type)
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(keyblock.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func makeHeaderPayload(clientInfo domain.User) string {
	// ヘッダー
	var header = []byte(`{"alg":"RS256","kid": "12345678","typ":"JWT"}`)

	// ペイロード
	var payload = domain.Payload{
		Iss:        "https://oreore.oidc.com",
		Azp:        clientInfo.Id,
		Aud:        clientInfo.Id,
		Sub:        clientInfo.Sub,
		AtHash:     "PRzSZsEPQVqzY8xyB2ls5A",
		Nonce:      "abc",
		Name:       clientInfo.Name,
		GivenName:  clientInfo.GivenName,
		FamilyName: clientInfo.FamilyName,
		Locale:     clientInfo.Locale,
		Iat:        time.Now().Unix(),
		Exp:        time.Now().Unix() + persistence.ACCESS_TOKEN_DURATION,
	}
	payload_json, _ := json.Marshal(payload)

	// base64エンコード
	b64header := base64.RawURLEncoding.EncodeToString(header)
	b64payload := base64.RawURLEncoding.EncodeToString(payload_json)

	return fmt.Sprintf("%s.%s", b64header, b64payload)
}

// JWTの作成
func makeJWT(clientInfo domain.User) (string, error) {
	jwtString := makeHeaderPayload(clientInfo)
	privateKey, err := readPrivateKey()
	if err != nil {
		return "", err
	}

	//バリデーションチェック
	err = privateKey.Validate()
	if err != nil {
		return "", fmt.Errorf("private key validate err:%s", err)
	}
	hasher := sha256.New()
	hasher.Write([]byte(jwtString))
	tokenHash := hasher.Sum(nil)

	// 署名を作成
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, tokenHash)
	if err != nil {
		return "", fmt.Errorf("sign by private key is err : %s", err)
	}
	enc := base64.RawURLEncoding.EncodeToString(signature)

	// "ヘッダー.ペイロード.署名"を作成
	return fmt.Sprintf("%s.%s", jwtString, enc), nil
}

func readEnv(statement string) (string, error) {
	envs_as_string := os.Getenv(statement)
	if len(envs_as_string) < 1 {
		//エラー処理
		return "No Value", fmt.Errorf("error, there is no value in env")
	}
	return envs_as_string, nil
}
