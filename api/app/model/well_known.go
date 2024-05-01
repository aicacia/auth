package model

type OpenIDConfigurationST = struct {
	Issuer                            string   `json:"issuer" validate:"required"`
	AuthorizationEndpoint             string   `json:"authorization_endpoint"`
	TokenEndpoint                     string   `json:"token_endpoint" validate:"required"`
	UserInfoEndpoint                  string   `json:"userinfo_endpoint" validate:"required"`
	JwksUri                           string   `json:"jwks_uri"`
	RegistrationEndpoint              *string  `json:"registration_endpoint"`
	ScopesSupported                   []string `json:"scopes_supported" validate:"required"`
	ResponseTypesSupported            []string `json:"response_types_supported" validate:"required"`
	GrantTypesSupported               []string `json:"grant_types_supported" validate:"required"`
	SubjectTypesSupported             []string `json:"subject_types_supported" validate:"required"`
	IdTokenSigningAlgValuesSupported  []string `json:"id_token_signing_alg_values_supported" validate:"required"`
	TokenEndpointAuthMethodsSupported []string `json:"token_endpoint_auth_methods_supported" validate:"required"`
	ClaimsSupported                   []string `json:"claims_supported" validate:"required"`
	CodeChallengeMethodsSupported     []string `json:"code_challenge_methods_supported" validate:"required"`
} // @name OpenIDConfiguration
