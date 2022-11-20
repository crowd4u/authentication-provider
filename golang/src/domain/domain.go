package domain

type Session struct {
	ClientId            string `json:"client_id"`
	State               string `json:"state"`
	Scopes              string `json:"scope"`
	RedirectURL         string `json:"redirect_url"`
	CodeChallenge       string `json:"code_challenge"`
	CodeChallengeMethod string `json:"code_challenge_method"`
	Nonce               string `json:"nonce"`
	OIDC                bool   `json:"oidc"`
}
type Client struct {
	Id          string
	Name        string
	Email       string
	RedirectURL string
	Secret      string
	ExpiresAt   int64
}

type User struct {
	Id         string
	Name       string
	Password   string
	Sub        string
	GivenName  string
	FamilyName string
	Locale     string
}

type AuthCode struct {
	UserId      string
	ClientId    string
	Scopes      string
	RedirectURL string
	ExpriresAt  int64
}

type TokenCode struct {
	UserId    string
	ClientId  string
	Scopes    string
	ExpiresAt int64
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
	IdToken     string `json:"id_token,omitempty"`
}

type Payload struct {
	Iss        string `json:"iss"`
	Azp        string `json:"azp"`
	Aud        string `json:"aud"`
	Sub        string `json:"sub"`
	AtHash     string `json:"at_hash"`
	Nonce      string `json:"nonce"`
	Name       string `json:"name"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Locale     string `json:"locale"`
	Iat        int64  `json:"iat"`
	Exp        int64  `json:"exp"`
}

type JwkKey struct {
	Kid string `json:"kid"`
	N   string `json:"n"`
	Alg string `json:"alg"`
	Kty string `json:"kty"`
	E   string `json:"e"`
	Use string `json:"use"`
}

type Env struct {
	Value string
}

type Envs []Env

var ERROR_RESPONSE Env = Env{Value: "Error!"}
