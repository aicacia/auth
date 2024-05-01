package test

import (
	"log"
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
		log.Printf("could not load env: %v", err)
	}

	err = repository.InitDB()
	if err != nil {
		log.Fatalf("could not init db: %v", err)
	}

	InitTestUser()

	Playwright, err = playwright.Run(&playwright.RunOptions{
		Verbose: true,
	})
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	Browser, err = Playwright.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Timeout:  playwright.Float(10000),
		Headless: playwright.Bool(os.Getenv("CI") == "true"),
	})
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	Page, err = Browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
}

func TeardownTest() {
	if err := repository.CloseDB(); err != nil {
		log.Fatalf("could not close db: %v", err)
	}
	if err := Browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err := Playwright.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}
}

func Goto(path string) {
	if _, err := Page.Goto(BaseUrl + path); err != nil {
		log.Fatalf("could not goto: %v", err)
	}
	if err := Page.Locator("body.hydrated").WaitFor(); err != nil {
		log.Fatalf("could not wait for body.hydrated: %v", err)
	}
}

var DefaultUser *repository.UserRowST
var DefaultUserInfo *repository.UserInfoRowST
var DefaultUserPassword = "password"

func InitTestUser() {
	result, err := repository.CreateUserWithPassword(1, "user-"+uuid.NewString(), DefaultUserPassword)
	if err != nil {
		log.Fatalf("could not create user: %v", err)
	}
	DefaultUser = &result.User
	DefaultUserInfo = &result.UserInfo
}

func Signin() {
	Goto("/signin")

	if err := Page.Locator("input[name=\"username\"]").Fill(DefaultUser.Username); err != nil {
		log.Fatalf("could not fill email: %v", err)
	}
	if err := Page.Locator("input[name=\"password\"]").Fill(DefaultUserPassword); err != nil {
		log.Fatalf("could not fill password: %v", err)
	}
	if err := Page.Locator("button[type=\"submit\"]").Click(); err != nil {
		log.Fatalf("could not click submit: %v", err)
	}
}
