package test

import (
	"log"
	"log/slog"
	"os"

	"github.com/aicacia/auth/api/app/repository"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/playwright-community/playwright-go"
)

var Playwright *playwright.Playwright
var Browser playwright.Browser
var Page playwright.Page
var BaseUrl = "http://localhost:5173"

func SetupTest() {
	err := godotenv.Load("../.env", "../.env.test")
	if err != nil {
		slog.Error("could not load env: %s\n", "error", err)
	}

	err = repository.InitDB()
	if err != nil {
		log.Fatalf("could not init db: %s\n", err)
	}

	InitTestUser()

	Playwright, err = playwright.Run(&playwright.RunOptions{
		Verbose: true,
	})
	if err != nil {
		log.Fatalf("could not start playwright: %s\n", err)
	}
	Browser, err = Playwright.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Timeout:  playwright.Float(10000),
		Headless: playwright.Bool(os.Getenv("CI") == "true"),
	})
	if err != nil {
		log.Fatalf("could not launch browser: %s\n", err)
	}
	Page, err = Browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %s\n", err)
	}
}

func TeardownTest() {
	if err := repository.CloseDB(); err != nil {
		log.Fatalf("could not close db: %s\n", err)
	}
	if err := Browser.Close(); err != nil {
		log.Fatalf("could not close browser: %s\n", err)
	}
	if err := Playwright.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %s\n", err)
	}
}

func Goto(path string) {
	if _, err := Page.Goto(BaseUrl + path); err != nil {
		log.Fatalf("could not goto: %s\n", err)
	}
	if err := Page.Locator("body.hydrated").WaitFor(); err != nil {
		log.Fatalf("could not wait for body.hydrated: %s\n", err)
	}
}

var DefaultUser *repository.UserRowST
var DefaultUserInfo *repository.UserInfoRowST
var DefaultUserPassword = "password"

func InitTestUser() {
	result, err := repository.CreateUserWithPassword(1, "user-"+uuid.NewString(), DefaultUserPassword)
	if err != nil {
		log.Fatalf("could not create user: %s\n", err)
	}
	DefaultUser = &result.User
	DefaultUserInfo = &result.UserInfo
}

func Signin() {
	Goto("/signin")

	if err := Page.Locator("input[name=\"username\"]").Fill(DefaultUser.Username); err != nil {
		log.Fatalf("could not fill email: %s\n", err)
	}
	if err := Page.Locator("input[name=\"password\"]").Fill(DefaultUserPassword); err != nil {
		log.Fatalf("could not fill password: %s\n", err)
	}
	if err := Page.Locator("button[type=\"submit\"]").Click(); err != nil {
		log.Fatalf("could not click submit: %s\n", err)
	}
}
