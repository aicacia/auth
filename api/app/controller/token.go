package controller

import (
	"fmt"
	"log/slog"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/aicacia/auth/api/app/config"
	"github.com/aicacia/auth/api/app/jwt"
	"github.com/aicacia/auth/api/app/middleware"
	"github.com/aicacia/auth/api/app/model"
	"github.com/aicacia/auth/api/app/repository"
	"github.com/aicacia/auth/api/app/util"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// PostToken
//
//	@Summary		Create JWT Token
//	@ID				create-token
//	@Tags			token
//	@Accept			json
//	@Produce		json
//	@Param			tokenRequest	body	model.TokenRequestST	true	"token request body"
//	@Success		200	{object}	model.TokenST
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		404	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/token [post]
//
//	@Security		TenentId
func PostToken(c *fiber.Ctx) error {
	var tokenRequest model.TokenRequestST
	if err := c.BodyParser(&tokenRequest); err != nil {
		slog.Error("invalid request body", "error", err)
		return model.NewError(http.StatusBadRequest).AddError("request", "invalid")
	}
	switch tokenRequest.GrantType {
	case model.PasswordGrantType:
		return passwordToken(c, tokenRequest)
	case model.ServieAccountGrantType:
		return serviceAccountToken(c, tokenRequest)
	case model.RefreshTokenGrantType:
		return refreshToken(c, tokenRequest)
	}
	return model.NewError(http.StatusBadRequest).AddError("grant_type", "invalid")
}

func passwordToken(c *fiber.Ctx, tokenRequest model.TokenRequestST) error {
	application := middleware.GetApplication(c)
	user, err := repository.GetUserByUsernameOrEmail(application.Id, strings.TrimSpace(tokenRequest.Username))
	if err != nil {
		slog.Error("failed to get user", "error", err)
		return model.NewError(http.StatusUnauthorized).AddError("username", "invalid").AddError("password", "invalid")
	}
	if user == nil {
		slog.Error("failed to get user", "error", err)
		return model.NewError(http.StatusUnauthorized).AddError("username", "invalid").AddError("password", "invalid")
	}
	verified, err := util.VerifyPassword(strings.TrimSpace(tokenRequest.Password), user.EncryptedPassword)
	if !verified || err != nil {
		if err != nil {
			slog.Error("failed to verify password", "error", err)
		}
		return model.NewError(http.StatusUnauthorized).AddError("username", "invalid").AddError("password", "invalid")
	}
	mfa, err := repository.GetMFA(user.Id)
	if err != nil {
		slog.Error("failed to get mfa", "error", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	return sendToken(c, sendTokenST{
		mfa:             mfa,
		issuedTokenType: tokenRequest.GrantType,
		scope:           tokenRequest.Scope,
		application:     application,
		tenent:          middleware.GetTenent(c),
		user:            user,
	})
}

func serviceAccountToken(c *fiber.Ctx, tokenRequest model.TokenRequestST) error {
	key, err := uuid.Parse(strings.TrimSpace(tokenRequest.Key))
	if err != nil {
		slog.Error("failed to parse key", "error", err)
		return model.NewError(http.StatusUnauthorized).AddError("key", "invalid").AddError("secret", "invalid")
	}
	serviceAccount, err := repository.GetServiceAccountByKey(key)
	if err != nil {
		slog.Error("failed to get user", "error", err)
		return model.NewError(http.StatusUnauthorized).AddError("key", "invalid").AddError("secret", "invalid")
	}
	verified, err := util.VerifyPassword(strings.TrimSpace(tokenRequest.Secret), serviceAccount.EncryptedSecret)
	if !verified || err != nil {
		if err != nil {
			slog.Error("failed to verify secret", "error", err)
		}
		return model.NewError(http.StatusUnauthorized).AddError("key", "invalid").AddError("secret", "invalid")
	}
	return sendToken(c, sendTokenST{
		issuedTokenType: tokenRequest.GrantType,
		scope:           tokenRequest.Scope,
		application:     middleware.GetApplication(c),
		tenent:          middleware.GetTenent(c),
		serviceAccount:  serviceAccount,
	})
}

func refreshToken(c *fiber.Ctx, tokenRequest model.TokenRequestST) error {
	tenent := middleware.GetTenent(c)
	claims, err := jwt.ParseClaimsFromToken[jwt.Claims](tokenRequest.RefreshToken, tenent)
	if err != nil {
		slog.Error("failed to get refresh token claims", "error", err)
		return model.NewError(http.StatusUnauthorized).AddError("refresh_token", "invalid")
	}
	user, err := repository.GetUserById(tenent.ApplicationId, claims.Subject)
	if err != nil {
		slog.Error("failed to get user", "error", err)
		return model.NewError(http.StatusUnauthorized).AddError("refresh_token", "invalid")
	}
	return sendToken(c, sendTokenST{
		issuedTokenType: tokenRequest.GrantType,
		scope:           tokenRequest.Scope,
		application:     middleware.GetApplication(c),
		tenent:          tenent,
		user:            user,
	})
}

type sendTokenST struct {
	mfa             *repository.MFARowST
	issuedTokenType string
	scope           string
	application     *repository.ApplicationRowST
	tenent          *repository.TenentRowST
	user            *repository.UserRowST
	serviceAccount  *repository.ServiceAccountRowST
}

func (sendToken *sendTokenST) MFAEnabled() bool {
	return sendToken.mfa != nil && sendToken.mfa.Enabled
}

func sendToken(
	c *fiber.Ctx,
	params sendTokenST,
) error {
	now := time.Now().UTC()
	scopes := jwt.ParseScopes(params.scope)
	audiences := []string{
		params.application.URI,
	}
	if params.application.Website != nil {
		audiences = append(audiences, *params.application.Website)
	}
	var subject int32
	var subjectType string
	if params.user != nil {
		subject = params.user.Id
		subjectType = jwt.UserSubject
	} else if params.serviceAccount != nil {
		subject = params.serviceAccount.Id
		subjectType = jwt.ServiceAccountSubject
	}
	baseClaims := jwt.Claims{
		Subject:          subject,
		SubjectType:      subjectType,
		Type:             jwt.BearerTokenType,
		ClientId:         params.tenent.ClientId,
		Audiences:        audiences,
		NotBeforeSeconds: now.Unix(),
		IssuedAtSeconds:  now.Unix(),
		ExpiresAtSeconds: now.Unix() + int64(params.tenent.ExpiresInSeconds),
		Issuer:           config.Get().URL,
		Scope:            scopes,
	}
	tokenType := jwt.BearerTokenType
	var claims jwt.ToMapClaims = &baseClaims
	if params.MFAEnabled() {
		baseClaims.Type = jwt.MFATokenType
		tokenType = fmt.Sprintf("%s:%s", jwt.MFATokenType, params.mfa.Type)
		mfaClaims := jwt.MFAClaims{
			Claims:    baseClaims,
			GrantType: params.issuedTokenType,
		}
		claims = &mfaClaims
	}
	accessToken, err := jwt.CreateToken(claims, params.tenent)
	if err != nil {
		slog.Error("failed to create access token", "error", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	var refreshToken *string
	var refreshTokenExpiresIn *int64
	if !params.MFAEnabled() {
		token, err := jwt.CreateToken(baseClaims.ToRefreshClaims(params.application, params.tenent), params.tenent)
		if err != nil {
			slog.Error("failed to create refresh token", "error", err)
			return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
		}
		refreshToken = &token
		refreshTokenExpiresIn = &params.tenent.RefreshExpiresInSeconds
	}
	var idToken *string
	if !params.MFAEnabled() && slices.Contains(scopes, "openid") {
		if subjectType == jwt.UserSubject {
			openIdClaims, err := jwt.OpenIdClaimsForUser(&baseClaims, params.user.Id)
			if err != nil {
				slog.Error("failed to create id claims", "error", err)
				return model.NewError(http.StatusInternalServerError)
			}
			token, err := jwt.CreateToken(openIdClaims, params.tenent)
			if err != nil {
				slog.Error("failed to create id token", "error", err)
				return model.NewError(http.StatusInternalServerError)
			}
			idToken = &token
		} else {
			return model.NewError(http.StatusBadRequest).AddError("scope", "invalid")
		}
	}
	c.Status(http.StatusOK)
	return c.JSON(model.TokenST{
		AccessToken:           accessToken,
		TokenType:             tokenType,
		IssuedTokenType:       params.issuedTokenType,
		ExpiresIn:             params.tenent.ExpiresInSeconds,
		Scope:                 scopes,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresIn: refreshTokenExpiresIn,
		IdToken:               idToken,
	})
}
