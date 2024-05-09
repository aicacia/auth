package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/aicacia/auth/api/app"
	"github.com/aicacia/auth/api/app/config"
	"github.com/aicacia/auth/api/app/repository"
	"github.com/aicacia/auth/api/app/router"
	"github.com/aicacia/auth/api/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	fiberRecover "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

var (
	Version string = "0.1.0"
	Build   string = fmt.Sprint(time.Now().UnixMilli())
)

// @title Auth API
// @description Auth API API
// @contact.name Nathan Faucett
// @contact.email nathanfaucett@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /
// @securityDefinitions.apikey Authorization
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
// @securityDefinitions.apikey TenentId
// @in header
// @name Tenent-Id
func main() {
	defer func() {
		rec := recover()
		if rec != nil {
			log.Fatalf("application panic: %v\n", rec)
		}
	}()
	var envs []string
	env := os.Getenv("APP_ENV")
	if env != "" {
		envs = append(envs, ".env."+env)
	}
	envs = append(envs, ".env")
	err := godotenv.Load(envs...)
	log.Printf("error loading .env file: %v\n", err)
	err = repository.InitDB()
	if err != nil {
		log.Fatalf("error initializing database: %v\n", err)
	}
	defer repository.CloseDB()
	err = config.InitConfig()
	if err != nil {
		log.Fatalf("error initializing config: %v\n", err)
	}
	defer config.CloseConfigListener()

	app.Version.Version = Version
	app.Version.Build = Build

	docs.SwaggerInfo.Version = Version
	uri, err := url.Parse(config.Get().URL)
	if err != nil {
		log.Fatalf("error parsing URI: %v\n", err)
	}
	docs.SwaggerInfo.Host = uri.Host

	logWriter := os.Stdout
	log.SetOutput(logWriter)
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime | log.LUTC)

	// https://docs.gofiber.io/api/fiber#config
	fiberApp := fiber.New(fiber.Config{
		Prefork:       false,
		Network:       "",
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "",
		AppName:       "",
		ErrorHandler:  router.ErrorHandler,
	})
	fiberApp.Use(fiberRecover.New())
	fiberApp.Use(logger.New(logger.Config{
		Output:     logWriter,
		TimeZone:   "UTC",
		TimeFormat: "2006/01/02 15:04:05",
		Format:     "${time} ${status} - ${ip} ${latency} ${method} ${path}\n",
	}))
	if config.Get().Dashboard.Enabled {
		fiberApp.Use("/dashboard", monitor.New())
	}
	router.InstallRouter(fiberApp)

	addr := fmt.Sprintf("%s:%d", config.Get().Host, config.Get().Port)
	log.Printf("Listening on %s\n", addr)

	log.Fatal(fiberApp.Listen(addr))
}
