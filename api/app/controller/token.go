package controller

import (
	"log"
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
		log.Printf("invalid request body: %v\n", err)
		return model.NewError(http.StatusBadRequest).AddError("request", "invalid").Send(c)
	}
	switch tokenRequest.GrantType {
	case model.PasswordGrantType:
		return passwordToken(c, tokenRequest)
	case model.ServieAccountGrantType:
		return serviceAccountToken(c, tokenRequest)
	case model.RefreshTokenGrantType:
		return refreshToken(c, tokenRequest)
	}
	return model.NewError(http.StatusBadRequest).AddError("grant_type", "invalid").Send(c)
}

func passwordToken(c *fiber.Ctx, tokenRequest model.TokenRequestST) error {
	application := middleware.GetApplication(c)
	user, err := repository.GetUserByUsernameOrEmail(application.Id, strings.TrimSpace(tokenRequest.Username))
	if err != nil {
		log.Printf("failed to get user: %v\n", err)
		return model.NewError(http.StatusUnauthorized).AddError("username", "invalid").AddError("password", "invalid").Send(c)
	}
	if user == nil {
		log.Printf("failed to get user: %v\n", err)
		return model.NewError(http.StatusUnauthorized).AddError("username", "invalid").AddError("password", "invalid").Send(c)
	}
	verified, err := util.VerifyPassword(strings.TrimSpace(tokenRequest.Password), user.EncryptedPassword)
	if !verified || err != nil {
		if err != nil {
			log.Printf("failed to verify password: %v\n", err)
		}
		return model.NewError(http.StatusUnauthorized).AddError("username", "invalid").AddError("password", "invalid").Send(c)
	}
	return sendToken(c, tokenRequest.GrantType, tokenRequest.Scope, application, middleware.GetTenent(c), user, nil)
}

func serviceAccountToken(c *fiber.Ctx, tokenRequest model.TokenRequestST) error {
	key, err := uuid.Parse(strings.TrimSpace(tokenRequest.Key))
	if err != nil {
		log.Printf("failed to parse key: %v\n", err)
		return model.NewError(http.StatusUnauthorized).AddError("key", "invalid").AddError("secret", "invalid").Send(c)
	}
	serviceAccount, err := repository.GetServiceAccountByKey(key)
	if err != nil {
		log.Printf("failed to get user: %v\n", err)
		return model.NewError(http.StatusUnauthorized).AddError("key", "invalid").AddError("secret", "invalid").Send(c)
	}
	verified, err := util.VerifyPassword(strings.TrimSpace(tokenRequest.Secret), serviceAccount.EncryptedSecret)
	if !verified || err != nil {
		if err != nil {
			log.Printf("failed to verify secret: %v\n", err)
		}
		return model.NewError(http.StatusUnauthorized).AddError("key", "invalid").AddError("secret", "invalid").Send(c)
	}
	return sendToken(c, tokenRequest.GrantType, tokenRequest.Scope, middleware.GetApplication(c), middleware.GetTenent(c), nil, serviceAccount)
}

func refreshToken(c *fiber.Ctx, tokenRequest model.TokenRequestST) error {
	tenent := middleware.GetTenent(c)
	claims, err := jwt.ParseClaimsFromToken(tokenRequest.RefreshToken, tenent)
	if err != nil {
		log.Printf("failed to get refresh token claims: %v\n", err)
		return model.NewError(http.StatusUnauthorized).AddError("refresh_token", "invalid").Send(c)
	}
	user, err := repository.GetUserById(claims.Subject)
	if err != nil {
		log.Printf("failed to get user: %v\n", err)
		return model.NewError(http.StatusUnauthorized).AddError("refresh_token", "invalid").Send(c)
	}
	return sendToken(c, tokenRequest.GrantType, tokenRequest.Scope, middleware.GetApplication(c), tenent, user, nil)
}

func sendToken(
	c *fiber.Ctx,
	issuedTokenType, scope string,
	application *repository.ApplicationRowST,
	tenent *repository.TenentRowST,
	user *repository.UserRowST,
	serviceAccount *repository.ServiceAccountRowST,
) error {
	now := time.Now().UTC()
	scopes := jwt.ParseScopes(scope)
	audiences := []string{
		application.Uri,
	}
	if application.Website != nil {
		audiences = append(audiences, *application.Website)
	}
	var subject int32
	var subjectType string
	if user != nil {
		subject = user.Id
		subjectType = jwt.UserSubject
	} else if serviceAccount != nil {
		subject = serviceAccount.Id
		subjectType = jwt.ServiceAccountSubject
	}
	claims := jwt.Claims{
		Subject:          subject,
		SubjectType:      subjectType,
		Type:             jwt.BearerTokenType,
		ClientId:         tenent.ClientId,
		Audiences:        audiences,
		NotBeforeSeconds: now.Unix(),
		IssuedAtSeconds:  now.Unix(),
		ExpiresAtSeconds: now.Unix() + int64(tenent.ExpiresInSeconds),
		Issuer:           config.Get().URL,
		Scope:            scopes,
	}
	accessToken, err := jwt.CreateToken(&claims, tenent)
	if err != nil {
		log.Printf("failed to create access token: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application").Send(c)
	}
	refreshToken, err := jwt.CreateToken(claims.ToRefreshClaims(application, tenent), tenent)
	if err != nil {
		log.Printf("failed to create refresh token: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application").Send(c)
	}
	var idToken *string
	if slices.Contains(scopes, "openid") {
		if subjectType == jwt.UserSubject {
			openIdClaims, err := jwt.OpenIdClaimsForUser(&claims, user.Id)
			if err != nil {
				log.Printf("failed to create id claims: %v\n", err)
				return model.NewError(http.StatusInternalServerError).Send(c)
			}
			token, err := jwt.CreateToken(openIdClaims, tenent)
			if err != nil {
				log.Printf("failed to create id token: %v\n", err)
				return model.NewError(http.StatusInternalServerError).Send(c)
			}
			idToken = &token
		} else {
			return model.NewError(http.StatusBadRequest).AddError("scope", "invalid").Send(c)
		}
	}
	return c.JSON(model.TokenST{
		AccessToken:           accessToken,
		TokenType:             jwt.BearerTokenType,
		IssuedTokenType:       issuedTokenType,
		ExpiresIn:             tenent.ExpiresInSeconds,
		Scope:                 scopes,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresIn: tenent.RefreshExpiresInSeconds,
		IdToken:               idToken,
	})
}
