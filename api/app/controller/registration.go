package controller

import (
	"log"
	"net/http"
	"strings"

	"github.com/aicacia/auth/api/app/middleware"
	"github.com/aicacia/auth/api/app/model"
	"github.com/aicacia/auth/api/app/repository"
	"github.com/gofiber/fiber/v2"
)

// PostRegistration
//
//	@Summary		Registration as a new user
//	@ID				register-user
//	@Tags			registration
//	@Accept			json
//	@Produce		json
//	@Param			registrationRequest	body	model.RegistrationRequestST	true	"token request body"
//	@Success		201	{object}	model.TokenST
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		404	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/register [post]
//
//	@Security		TenentId
func PostRegistration(c *fiber.Ctx) error {
	tenent := middleware.GetTenent(c)
	if tenent.RegistrationWebsite == nil {
		return model.NewError(http.StatusForbidden).AddError("signup", "disabled", "application")
	}
	var registrationRequest model.RegistrationRequestST
	if err := c.BodyParser(&registrationRequest); err != nil {
		log.Printf("failed to parse body: %v\n", err)
		return model.NewError(http.StatusBadRequest).AddError("request", "invalid")
	}
	username := strings.TrimSpace(registrationRequest.Username)
	password := strings.TrimSpace(registrationRequest.Password)
	passwordConfirmation := strings.TrimSpace(registrationRequest.PasswordConfirmation)
	errors := model.NewError(http.StatusBadRequest)
	if username == "" {
		errors.AddError("username", "required")
	}
	if password != passwordConfirmation {
		errors.AddError("password_confirmation", "mismatch")
	}
	if errors.HasErrors() {
		return errors
	}
	application := middleware.GetApplication(c)
	createResult, err := repository.CreateUserWithPassword(application.Id, registrationRequest.Username, password)
	if err != nil {
		log.Printf("failed to create user: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	_, err = repository.AddUserToApplication(application.Id, createResult.User.Id)
	if err != nil {
		log.Printf("failed to add user to application: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	return sendToken(c, model.PasswordGrantType, "openid", application, tenent, &createResult.User, nil)
}
