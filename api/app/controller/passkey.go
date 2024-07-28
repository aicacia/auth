package controller

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/aicacia/auth/api/app/middleware"
	"github.com/aicacia/auth/api/app/model"
	"github.com/aicacia/auth/api/app/repository"
	"github.com/aicacia/auth/api/app/service"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/gofiber/fiber/v2"
)

// PostPassKeyBeginRegistration
//
//	@Summary		Begin registering a new passkey
//	@ID				  passkey-begin-registration
//	@Tags			  passkey
//	@Accept			json
//	@Produce		json
//	@Success		201	{object}	protocol.PublicKeyCredentialCreationOptions
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		404	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/user/passkeys/begin-registration [post]
//
//	@Security		Authorization
func PostPassKeyBeginRegistration(c *fiber.Ctx) error {
	tenent := middleware.GetTenent(c)
	webauthn, err := service.WebAuthnFromTenent(tenent)
	if err != nil {
		slog.Error("failed to get webauthn", "error", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	user := middleware.GetUser(c)
	passkeys, err := repository.GetUserPassKeys(user.Id)
	if err != nil {
		slog.Error("failed to get user passkeys", "error", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	webauthUser := service.NewWebAuthnUser(*user, passkeys)
	options, session, err := webauthn.BeginRegistration(webauthUser, func(opts *protocol.PublicKeyCredentialCreationOptions) {
	})
	if err != nil {
		slog.Error("failed to begin registration", "error", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	service.WebAuthnSessions.Set(user.Id, session, time.Now().UTC().Add(time.Minute))

	c.Status(http.StatusCreated)
	return c.JSON(options.Response)
}

// PostPassKeyFinishRegistration
//
//	@Summary		Finish registering a new passkey
//	@ID				  passkey-finish-registration
//	@Tags			  passkey
//	@Accept			json
//	@Produce		json
//	@Success		204
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		404	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/user/passkeys/finish-registration [post]
//
//	@Security		Authorization
func PostPassKeyFinishRegistration(c *fiber.Ctx) error {
	tenent := middleware.GetTenent(c)
	webauthn, err := service.WebAuthnFromTenent(tenent)
	if err != nil {
		slog.Error("failed to get webauthn", "error", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	user := middleware.GetUser(c)
	passkeys, err := repository.GetUserPassKeys(user.Id)
	if err != nil {
		slog.Error("failed to get user passkeys", "error", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	webauthUser := service.NewWebAuthnUser(*user, passkeys)
	session, ok := service.WebAuthnSessions.Get(user.Id)
	service.WebAuthnSessions.Delete(user.Id)
	if !ok {
		slog.Error("failed to get passkey session", "error", err)
		return model.NewError(http.StatusNotFound).AddError("notFound", "session")
	}
	var data protocol.ParsedCredentialCreationData
	if err := c.BodyParser(&data); err != nil {
		slog.Error("failed to parse request body", "error", err)
		return model.NewError(http.StatusBadRequest).AddError("badRequest", "body")
	}
	credential, err := webauthn.CreateCredential(webauthUser, *session, &data)
	if err != nil {
		slog.Error("failed to finish registration", "error", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}

	if _, err := repository.UpsertUserPassKey(service.WebAuthnCredentialsToUpsert(user.ApplicationId, user.Id, *credential)); err != nil {
		slog.Error("failed to upsert user passkey", "error", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}

	c.Status(http.StatusNoContent)
	return c.Send(nil)
}

// PostPassKeyBeginLogin
//
//	@Summary		Begin logining with a passkey
//	@ID				  passkey-begin-login
//	@Tags			  passkey
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	protocol.PublicKeyCredentialRequestOptions
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		404	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/user/passkeys/begin-login [post]
//
//	@Security		Authorization
func PostPassKeyBeginLogin(c *fiber.Ctx) error {
	tenent := middleware.GetTenent(c)
	webauthn, err := service.WebAuthnFromTenent(tenent)
	if err != nil {
		slog.Error("failed to get webauthn", "error", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	user := middleware.GetUser(c)
	passkeys, err := repository.GetUserPassKeys(user.Id)
	if err != nil {
		slog.Error("failed to get user passkeys", "error", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	webauthUser := service.NewWebAuthnUser(*user, passkeys)
	options, session, err := webauthn.BeginLogin(webauthUser, func(opts *protocol.PublicKeyCredentialRequestOptions) {
	})
	if err != nil {
		slog.Error("failed to begin registration", "error", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	service.WebAuthnSessions.Set(user.Id, session, time.Now().UTC().Add(time.Minute))

	c.Status(http.StatusOK)
	return c.JSON(options.Response)
}

// PostPassKeyFinishLogin
//
//	@Summary		Finish logining with a passkey
//	@ID				  passkey-finish-login
//	@Tags			  passkey
//	@Accept			json
//	@Produce		json
//	@Success		204
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		404	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/user/passkeys/finish-login [post]
//
//	@Security		Authorization
func PostPassKeyFinishLogin(c *fiber.Ctx) error {
	tenent := middleware.GetTenent(c)
	webauthn, err := service.WebAuthnFromTenent(tenent)
	if err != nil {
		slog.Error("failed to get webauthn", "error", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	user := middleware.GetUser(c)
	passkeys, err := repository.GetUserPassKeys(user.Id)
	if err != nil {
		slog.Error("failed to get user passkeys", "error", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	webauthUser := service.NewWebAuthnUser(*user, passkeys)
	session, ok := service.WebAuthnSessions.Get(user.Id)
	service.WebAuthnSessions.Delete(user.Id)
	if !ok {
		slog.Error("failed to get passkey session", "error", err)
		return model.NewError(http.StatusNotFound).AddError("notFound", "session")
	}
	var data protocol.ParsedCredentialAssertionData
	if err := c.BodyParser(&data); err != nil {
		slog.Error("failed to parse request body", "error", err)
		return model.NewError(http.StatusBadRequest).AddError("badRequest", "body")
	}
	credential, err := webauthn.ValidateLogin(webauthUser, *session, &data)
	if err != nil {
		slog.Error("failed to finish registration", "error", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}

	if _, err := repository.UpsertUserPassKey(service.WebAuthnCredentialsToUpsert(user.ApplicationId, user.Id, *credential)); err != nil {
		slog.Error("failed to upsert user passkey", "error", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}

	c.Status(http.StatusNoContent)
	return c.Send(nil)
}
