package model

var (
	PasswordGrantType      = "password"
	ServieAccountGrantType = "service-account"
	RefreshTokenGrantType  = "refresh-token"
)

type TokenRequestST struct {
	GrantType          string `json:"grant_type" validate:"required"`
	Code               string `json:"code"`
	RefreshToken       string `json:"refresh_token"`
	Key                string `json:"key"`
	Secret             string `json:"secret"`
	Username           string `json:"username"`
	Password           string `json:"password"`
	Scope              string `json:"scope"`
	Assertion          string `json:"assertion"`
	CodeVerifier       string `json:"code_verifier"`
	Resource           string `json:"resource"`
	Audience           string `json:"audience"`
	RequestedTokenType string `json:"requested_token_type"`
	SubjectToken       string `json:"subject_token"`
	SubjectTokenType   string `json:"subject_token_type"`
	ActorToken         string `json:"actor_token"`
	ActorTokenType     string `json:"actor_token_type"`
} // @name TokenRequest

type TokenST struct {
	AccessToken           string   `json:"access_token" validate:"required"`
	TokenType             string   `json:"token_type" validate:"required"`
	IssuedTokenType       string   `json:"issued_token_type" validate:"required"`
	ExpiresIn             int64    `json:"expires_in" validate:"required"`
	Scope                 []string `json:"scope" validate:"required"`
	RefreshToken          *string  `json:"refresh_token,omitempty" validate:"required"`
	RefreshTokenExpiresIn *int64   `json:"refresh_token_expires_in,omitempty" validate:"required"`
	IdToken               *string  `json:"id_token,omitempty"`
} // @name Token
