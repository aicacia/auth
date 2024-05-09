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

	wellKnown := root.Group("/.well-known")
	wellKnown.Use(middleware.TenentMiddleware())
	wellKnown.Get("/openid-configuration", controller.GetOpenIDConfiguration)

	user := root.Group("/user")
	user.Use(middleware.AuthorizedMiddleware(), middleware.IsUserMiddleware())
	user.Get("", controller.GetCurrentUser)
	user.Patch("", controller.PatchUpdateCurrentUser)
	user.Patch("/reset-password", controller.PatchResetPassword)

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

	permissions := applications.Group("/:applicationId/permissions")
	permissions.Get("", controller.GetPermissions)
	permissions.Get("/:id", controller.GetPermissionById)
	permissions.Post("", controller.PostCreatePermission)
	permissions.Patch("/:id", controller.PatchUpdatePermission)
	permissions.Delete("/:id", controller.DeletePermission)

	users := applications.Group("/:applicationId/users")
	users.Get("", controller.GetUsers)
	users.Get("/:id", controller.GetUserById)
	users.Patch("/:id", controller.PatchUpdateUserById)
	users.Post("", controller.PostCreateUser)
	users.Delete("/:id", controller.DeleteUserById)
	users.Get("/:id/info", controller.GetUserInfo)
	users.Patch("/:id/info", controller.PatchUserInfo)

	emails := users.Group("/:userId/emails")
	emails.Patch("/:id/send-confirmation", controller.PatchUserEmailSendConfirmation)
	emails.Patch("/:id/confirm", controller.PatchUserEmailConfirm)
	emails.Patch("/:id/set-primary", controller.PatchUserEmailSetPrimary)
	emails.Post("", controller.PostUserCreateEmail)
	emails.Delete("/:id", controller.DeleteUserEmail)

	phoneNumbers := users.Group("/:userId/phone-numbers")
	phoneNumbers.Patch("/:id/send-confirmation", controller.PatchUserPhoneNumberSendConfirmation)
	phoneNumbers.Patch("/:id/confirm", controller.PatchUserPhoneNumberConfirm)
	phoneNumbers.Patch("/:id/set-primary", controller.PatchUserPhoneNumberSetPrimary)
	phoneNumbers.Post("", controller.PostUserCreatePhoneNumber)
	phoneNumbers.Delete("/:id", controller.DeleteUserPhoneNumber)
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	if e, ok := err.(*model.ErrorST); ok {
		return e.Send(c)
	}
	return fiber.DefaultErrorHandler(c, err)
}
