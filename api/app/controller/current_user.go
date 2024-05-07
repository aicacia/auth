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

// GetCurrentUser
//
//	@ID				current-user
//	@Summary		Get current user
//	@Tags			current-user
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}   	model.UserWithPermissionsST
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/user [get]
//
//	@Security		Authorization
func GetCurrentUser(c *fiber.Ctx) error {
	user := middleware.GetUser(c)
	application := middleware.GetApplication(c)
	emails, phoneNumbers, err := getUserEmailsAndPhoneNumbersById(user.Id)
	if err != nil {
		log.Printf("failed to get user emails and phone numbers: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application").Send(c)
	}
	permissionRows, err := repository.GetUserPermissions(user.Id, application.Id)
	if err != nil {
		log.Printf("failed to get user permissions: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application").Send(c)
	}
	permissions := make([]string, 0, len(permissionRows))
	for _, permissionRow := range permissionRows {
		permissions = append(permissions, permissionRow.Uri)
	}
	userWithPermissions := model.UserWithPermissionsST{
		UserST:      model.UserFromRow(*user, emails, phoneNumbers),
		Permissions: permissions,
	}
	return c.JSON(userWithPermissions)
}

// PatchResetPassword
//
//	@Summary		Resets a user's password
//	@ID				reset-password
//	@Tags			current-user
//	@Accept			json
//	@Produce		json
//	@Param			resetPassword	body    model.ResetPasswordST	true	"reset user's password"
//	@Success		204
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/user/reset-password [patch]
//
//	@Security		Authorization
func PatchResetPassword(c *fiber.Ctx) error {
	var resetPassword model.ResetPasswordST
	if err := c.BodyParser(&resetPassword); err != nil {
		log.Printf("failed to parse reset password: %v\n", err)
		return model.NewError(http.StatusBadRequest).AddError("request", "invalid").Send(c)
	}
	password := strings.TrimSpace(resetPassword.Password)
	passwordConfirmation := strings.TrimSpace(resetPassword.PasswordConfirmation)
	errors := model.NewError(http.StatusBadRequest)
	if password != passwordConfirmation {
		errors.AddError("password_confirmation", "mismatch", "body")
	}
	if errors.HasErrors() {
		return errors.Send(c)
	}
	user := middleware.GetUser(c)
	_, err := repository.UpdateUserPassword(user.ApplicationId, user.Id, password)
	if err != nil {
		log.Printf("failed to update user password: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application").Send(c)
	}
	c.Status(http.StatusNoContent)
	return c.Send(nil)
}

// PatchUpdateCurrentUser
//
//	@Summary		Updates current user's username
//	@ID				update-username
//	@Tags			current-user
//	@Accept			json
//	@Produce		json
//	@Param			updateUser	body    model.UpdateUserST	true	"update user"
//	@Success		204
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/user [patch]
//
//	@Security		Authorization
func PatchUpdateCurrentUser(c *fiber.Ctx) error {
	var updateUser model.UpdateUserST
	if err := c.BodyParser(&updateUser); err != nil {
		return model.NewError(http.StatusBadRequest).AddError("request", "invalid").Send(c)
	}
	user := middleware.GetUser(c)
	user, err := repository.UpdateUsername(user.ApplicationId, user.Id, updateUser.Username)
	if err != nil {
		log.Printf("failed to create user: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application").Send(c)
	}
	if user == nil {
		return model.NewError(http.StatusNotFound).AddError("id", "invalid").Send(c)
	}
	c.Status(http.StatusNoContent)
	return c.Send(nil)
}

// GetCurrentUserInfo
//
//	@Summary		Get user info
//	@ID				get-current-user-info
//	@Tags			current-user
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.UserInfoST
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/user/info [get]
//
//	@Security		Authorization
func GetCurrentUserInfo(c *fiber.Ctx) error {
	user := middleware.GetUser(c)
	userInfoRow, err := repository.GetUserInfoByUserId(user.Id)
	if err != nil {
		log.Printf("failed to fetch user info: %v\n", err)
		return model.NewError(http.StatusInternalServerError).Send(c)
	}
	userInfo := model.UserInfoFromRow(user, userInfoRow)
	return c.JSON(userInfo)
}

// PatchCurrentUserInfo
//
//	@Summary		Updates the user's info
//	@ID				update-current-user-info
//	@Tags			current-user
//	@Accept			json
//	@Produce		json
//	@Param			userinfoUpdates	body    model.UpdateUserInfoRequestST	true	"User info updates"
//	@Success		200	{object}	model.UserInfoST
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/user/info [patch]
//
//	@Security		Authorization
func PatchCurrentUserInfo(c *fiber.Ctx) error {
	var userinfoUpdates model.UpdateUserInfoRequestST
	if err := c.BodyParser(&userinfoUpdates); err != nil {
		log.Printf("invalid request body: %v\n", err)
		return model.NewError(http.StatusBadRequest).AddError("request", "invalid").Send(c)
	}
	updates := repository.UpdateUserInfoST{
		Name:       userinfoUpdates.Name,
		GivenName:  userinfoUpdates.GivenName,
		FamilyName: userinfoUpdates.FamilyName,
		MiddleName: userinfoUpdates.MiddleName,
		Nickname:   userinfoUpdates.Nickname,
		Profile:    userinfoUpdates.Profile,
		Picture:    userinfoUpdates.Picture,
		Website:    userinfoUpdates.Website,
		Gender:     userinfoUpdates.Gender,
		Birthdate:  userinfoUpdates.Birthdate,
		Zoneinfo:   userinfoUpdates.Zoneinfo,
		Locale:     userinfoUpdates.Locale,
	}
	if userinfoUpdates.Address != nil {
		updates.Region = userinfoUpdates.Address.Region
		updates.Locality = userinfoUpdates.Address.Locality
		updates.PostalCode = userinfoUpdates.Address.PostalCode
		updates.Country = userinfoUpdates.Address.Country
		updates.StreetAddress = userinfoUpdates.Address.StreetAddress
	}
	user := middleware.GetUser(c)
	userInfoRow, err := repository.UpdateUserInfoByUserId(user.Id, updates)
	if err != nil {
		log.Printf("failed to fetch user info: %v\n", err)
		return model.NewError(http.StatusInternalServerError).Send(c)
	}
	userInfo := model.UserInfoFromRow(user, userInfoRow)
	return c.JSON(userInfo)
}
