package router

import (
	"github.com/aicacia/auth/api/app/config"
	"github.com/aicacia/auth/api/app/controller"
	"github.com/aicacia/auth/api/app/middleware"
	"github.com/aicacia/auth/api/app/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func InstallRouter(fiberApp *fiber.App) {
	root := fiberApp.Group("", cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			return true
		},
		AllowCredentials: true,
	}))

	if config.Get().OpenAPI.Enabled {
		root.Get("/openapi.json", controller.GetOpenAPI)
	}

	root.Get("/health", controller.GetHealthCheck)
	root.Get("/version", controller.GetVersion)

	token := root.Group("/token")
	token.Use(middleware.TenentMiddleware())
	token.Post("", controller.PostToken)

	registration := root.Group("/registration")
	registration.Use(middleware.TenentMiddleware())
	registration.Post("", controller.PostRegistration)

	mfa := root.Group("/mfa")
	mfa.Use(middleware.MFAAuthorizedMiddleware())
	mfa.Post("", controller.PostValidateMFA)

	wellKnown := root.Group("/.well-known")
	wellKnown.Use(middleware.TenentMiddleware())
	wellKnown.Get("/openid-configuration", controller.GetOpenIDConfiguration)

	user := root.Group("/user")
	user.Use(middleware.AuthorizedMiddleware(), middleware.IsUserMiddleware())
	user.Get("", controller.GetCurrentUser)
	user.Patch("", controller.PatchUpdateCurrentUser)
	user.Patch("/reset-password", controller.PatchResetPassword)

	userEmails := user.Group("/emails")
	userEmails.Patch("/:id/send-confirmation", controller.PatchCurrentUserEmailSendConfirmation)
	userEmails.Patch("/:id/confirm", controller.PatchCurrentUserEmailConfirm)
	userEmails.Patch("/:id/set-primary", controller.PatchCurrentUserEmailSetPrimary)
	userEmails.Post("", controller.PostCurrentUserCreateEmail)
	userEmails.Delete("/:id", controller.DeleteCurrentUserEmail)

	userPhoneNumbers := user.Group("/phone-numbers")
	userPhoneNumbers.Patch("/:id/send-confirmation", controller.PatchCurrentUserPhoneNumberSendConfirmation)
	userPhoneNumbers.Patch("/:id/confirm", controller.PatchCurrentUserPhoneNumberConfirm)
	userPhoneNumbers.Patch("/:id/set-primary", controller.PatchCurrentUserPhoneNumberSetPrimary)
	userPhoneNumbers.Post("", controller.PostCurrentUserCreatePhoneNumber)
	userPhoneNumbers.Delete("/:id", controller.DeleteCurrentUserPhoneNumber)

	openid := user.Group("/info")
	openid.Use(middleware.OpenIdMiddleware())
	openid.Get("", controller.GetCurrentUserInfo)
	openid.Patch("", controller.PatchCurrentUserInfo)

	admin := root.Group("")
	admin.Use(middleware.AuthorizedMiddleware(), middleware.AdminApplicationMiddleware())

	applications := admin.Group("/applications")
	applications.Get("", controller.GetApplications)
	applications.Get("/:id", controller.GetApplicationById)
	applications.Post("", controller.PostCreateApplication)
	applications.Patch("/:id", controller.PatchUpdateApplication)
	applications.Delete("/:id", controller.DeleteApplication)

	tenents := applications.Group("/:applicationId/tenents")
	tenents.Get("", controller.GetTenents)
	tenents.Get("/:id", controller.GetTenentById)
	tenents.Post("", controller.PostCreateTenent)
	tenents.Patch("/:id", controller.PatchUpdateTenent)
	tenents.Delete("/:id", controller.DeleteTenent)

	users := applications.Group("/:applicationId/users")
	users.Get("", controller.GetUsers)
	users.Get("/:id", controller.GetUserById)
	users.Patch("/:id", controller.PatchUpdateUserById)
	users.Post("", controller.PostCreateUser)
	users.Delete("/:id", controller.DeleteUserById)
	users.Get("/:id/info", controller.GetUserInfo)
	users.Patch("/:id/info", controller.PatchUserInfo)
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	if e, ok := err.(*model.ErrorST); ok {
		return e.Send(c)
	}
	return fiber.DefaultErrorHandler(c, err)
}
