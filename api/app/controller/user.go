package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/aicacia/auth/api/app/access"
	"github.com/aicacia/auth/api/app/middleware"
	"github.com/aicacia/auth/api/app/model"
	"github.com/aicacia/auth/api/app/repository"
	"github.com/gofiber/fiber/v2"
)

// GetUsers
//
//	@Summary		Get users
//	@ID				users
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			applicationId	path		int	true	"application id"
//	@Param			limit	query		int	false	"limit"
//	@Param			offset	query		int	false	"offset"
//	@Success		200	{object}   	model.PaginationST[model.UserST]
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/applications/{applicationId}/users [get]
//
//	@Security		Authorization
func GetUsers(c *fiber.Ctx) error {
	if err := access.IsAdmin(c); err != nil {
		return err
	}
	limit, offset, err := GetLimitAndOffset(c, 20)
	if err != nil {
		if err == errParseLimitOffset {
			return nil
		}
		return err
	}
	applicationId, err := strconv.Atoi(c.Params("applicationId"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("applicationId", "invalid")
	}
	userRows, err := repository.GetUsers(int32(applicationId), limit, offset)
	if err != nil {
		log.Printf("failed to get users: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	emails, err := repository.GetUsersEmails(int32(applicationId), limit, offset)
	if err != nil {
		log.Printf("failed to get users emails: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	emailsByUserId := make(map[int32][]repository.EmailRowST, len(emails))
	for _, email := range emails {
		emailsByUserId[email.UserId] = append(emailsByUserId[email.UserId], email)
	}
	phoneNumbers, err := repository.GetUsersPhoneNumbers(int32(applicationId), limit, offset)
	if err != nil {
		log.Printf("failed to get users phone numbers: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	phoneNumbersByUserId := make(map[int32][]repository.PhoneNumberRowST, len(phoneNumbers))
	for _, phoneNumber := range phoneNumbers {
		phoneNumbersByUserId[phoneNumber.UserId] = append(phoneNumbersByUserId[phoneNumber.UserId], phoneNumber)
	}
	users := make([]model.UserST, 0, len(userRows))
	for _, userRow := range userRows {
		users = append(users, model.UserFromRow(userRow, emailsByUserId[userRow.Id], phoneNumbersByUserId[userRow.Id]))
	}
	return c.JSON(model.PaginationST[model.UserST]{
		HasMore: len(users) == limit,
		Items:   users,
	})
}

// GetUserById
//
//	@Summary		Get user by id
//	@ID				user-by-id
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			applicationId	path		int	true	"application id"
//	@Param			id	path		int	true	"user id"
//	@Success		200	{object}   	model.UserST
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/applications/{applicationId}/users/{id} [get]
//
//	@Security		Authorization
func GetUserById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("id", "invalid")
	}
	application := middleware.GetApplication(c)
	user, emails, phoneNumbers, err := getUserById(application.Id, int32(id))
	if err != nil {
		log.Printf("failed to get user: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	return c.JSON(model.UserFromRow(*user, emails, phoneNumbers))
}

// PostCreateUser
//
//	@Summary		Create user
//	@ID				create-user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			applicationId	path		int	true	"application id"
//	@Param			createUser	body    model.CreateUserST	true	"create user"
//	@Success		201	{object}   	model.UserST
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/applications/{applicationId}/users [post]
//
//	@Security		Authorization
func PostCreateUser(c *fiber.Ctx) error {
	if err := access.IsAdmin(c); err != nil {
		return err
	}
	applicationId, err := strconv.Atoi(c.Params("applicationId"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("applicationId", "invalid")
	}
	var createUser model.CreateUserST
	if err := c.BodyParser(&createUser); err != nil {
		return model.NewError(http.StatusBadRequest).AddError("request", "invalid")
	}
	result, err := repository.CreateUserFromUsername(int32(applicationId), createUser.Username)
	if err != nil {
		log.Printf("failed to create user: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	emails, phoneNumbers, err := getUserEmailsAndPhoneNumbersById(result.User.Id)
	if err != nil {
		log.Printf("failed to get user emails and phone numbers: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	return c.JSON(model.UserFromRow(result.User, emails, phoneNumbers))
}

// PatchUpdateUserById
//
//	@Summary		Updates a user's username
//	@ID				update-user-by-id
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			applicationId	path		int	true	"application id"
//	@Param			id	path		int	true	"user id"
//	@Param			updateUser	body    model.UpdateUserST	true	"update user"
//	@Success		200	{object}   	model.UserST
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/applications/{applicationId}/users/{id} [patch]
//
//	@Security		Authorization
func PatchUpdateUserById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("id", "invalid")
	}
	var updateUser model.UpdateUserST
	if err := c.BodyParser(&updateUser); err != nil {
		return model.NewError(http.StatusBadRequest).AddError("request", "invalid")
	}
	application := middleware.GetApplication(c)
	user, err := repository.UpdateUsername(application.Id, int32(id), updateUser.Username)
	if err != nil {
		log.Printf("failed to create user: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	if user == nil {
		return model.NewError(http.StatusNotFound).AddError("id", "invalid")
	}
	emails, phoneNumbers, err := getUserEmailsAndPhoneNumbersById(user.Id)
	if err != nil {
		log.Printf("failed to get user emails and phone numbers: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	return c.JSON(model.UserFromRow(*user, emails, phoneNumbers))
}

// DeleteUserById
//
//	@Summary		Delets a user by id
//	@ID				delete-user-by-id
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			applicationId	path		int	true	"application id"
//	@Param			id	path		int	true	"user id"
//	@Success		204
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/applications/{applicationId}/users/{id} [delete]
//
//	@Security		Authorization
func DeleteUserById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("id", "invalid")
	}
	application := middleware.GetApplication(c)
	deleted, err := repository.DeleteUserById(application.Id, int32(id))
	if err != nil {
		log.Printf("failed to create user: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	if !deleted {
		return model.NewError(http.StatusNotFound).AddError("id", "invalid")
	}
	c.Status(http.StatusNoContent)
	return c.Send(nil)
}

// GetUserInfo
//
//	@Summary		Get user info
//	@ID				user-info
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			applicationId	path		int	true	"application id"
//	@Param			id	path		int	true	"user id"
//	@Success		200	{object}	model.UserInfoST
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/applications/{applicationId}/users/{id}/info [get]
//
//	@Security		Authorization
func GetUserInfo(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("id", "invalid")
	}
	application := middleware.GetApplication(c)
	user, err := repository.GetUserById(application.Id, int32(id))
	if err != nil {
		log.Printf("failed to get user: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	if user == nil {
		return model.NewError(http.StatusNotFound).AddError("id", "invalid")
	}
	userInfoRow, err := repository.GetUserInfoByUserId(int32(id))
	if err != nil {
		log.Printf("failed to fetch user info: %v\n", err)
		return model.NewError(http.StatusInternalServerError)
	}
	userInfo := model.UserInfoFromRow(user, userInfoRow)
	return c.JSON(userInfo)
}

// PatchUserInfo
//
//	@Summary		Updates the user's info
//	@ID				update-user-info
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			applicationId	path		int	true	"application id"
//	@Param			id	path		int	true	"user id"
//	@Param			userinfoUpdates	body    model.UpdateUserInfoRequestST	true	"User info updates"
//	@Success		200	{object}	model.UserInfoST
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/applications/{applicationId}/users/{id}/info [patch]
//
//	@Security		Authorization
func PatchUserInfo(c *fiber.Ctx) error {
	var userinfoUpdates model.UpdateUserInfoRequestST
	if err := c.BodyParser(&userinfoUpdates); err != nil {
		log.Printf("invalid request body: %v\n", err)
		return model.NewError(http.StatusBadRequest).AddError("request", "invalid")
	}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("id", "invalid")
	}
	application := middleware.GetApplication(c)
	user, err := repository.GetUserById(application.Id, int32(id))
	if err != nil {
		log.Printf("failed to get user: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	if user == nil {
		return model.NewError(http.StatusNotFound).AddError("id", "invalid")
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
	userInfoRow, err := repository.UpdateUserInfoByUserId(user.Id, updates)
	if err != nil {
		log.Printf("failed to fetch user info: %v\n", err)
		return model.NewError(http.StatusInternalServerError)
	}
	userInfo := model.UserInfoFromRow(user, userInfoRow)
	return c.JSON(userInfo)
}

func getUserById(applicationId, userId int32) (*repository.UserRowST, []repository.EmailRowST, []repository.PhoneNumberRowST, error) {
	user, err := repository.GetUserById(applicationId, userId)
	if err != nil {
		return nil, nil, nil, err
	}
	if user == nil {
		return nil, nil, nil, err
	}
	emails, phoneNumbers, err := getUserEmailsAndPhoneNumbersById(user.Id)
	if err != nil {
		return nil, nil, nil, err
	}
	return user, emails, phoneNumbers, nil
}

func getUserEmailsAndPhoneNumbersById(userId int32) ([]repository.EmailRowST, []repository.PhoneNumberRowST, error) {
	emails, err := repository.GetEmailsByUserId(userId)
	if err != nil {
		return nil, nil, err
	}
	phoneNumbers, err := repository.GetPhoneNumbersByUserId(userId)
	if err != nil {
		return nil, nil, err
	}
	return emails, phoneNumbers, nil
}
