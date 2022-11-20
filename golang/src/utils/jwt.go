package utils

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
	"notchman8600/authentication-provider/domain"
	"notchman8600/authentication-provider/infra"
	"notchman8600/authentication-provider/persistence"
	"time"
)

func readPrivateKey() (*rsa.PrivateKey, error) {
	envHandler := infra.NewEnvHandler()
	// 環境変数の読み込み
	fileName, err := envHandler.ReadEnv("FILE-PRIVATE-KEY")
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
