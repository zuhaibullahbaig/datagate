package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/zuhaibullahbaig/datagate/cli"
	"github.com/zuhaibullahbaig/datagate/middleware"
	"github.com/zuhaibullahbaig/datagate/server"
	"github.com/zuhaibullahbaig/datagate/utils"
)

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		Views:                 engine,
	})

	config := cli.ParseFlags()

	var token string
	var homepageURL string

	info := color.New(color.FgCyan).SprintFunc()
	success := color.New(color.FgGreen, color.Bold).SprintFunc()
	warning := color.New(color.FgYellow).SprintFunc()
	errorMsg := color.New(color.FgRed, color.Bold).SprintFunc()

	if config.Token {
		token = utils.GenerateToken()
		fmt.Println(info("\nGenerating Token..."))
		fmt.Println(success("Generated Successfully:"), token)
		app.Use(middleware.AuthMiddleware(token))
		server.SetupRoutes(app, config.Dir, token, true)
		homepageURL = fmt.Sprintf("http://%s:%d/?token=%s", utils.GetLocalIP(), config.Port, token)
	} else {
		server.SetupRoutes(app, config.Dir, "", false)
		homepageURL = fmt.Sprintf("http://%s:%d/?token=%s", utils.GetLocalIP(), config.Port, token)

	}

	hostIP := utils.GetLocalIP()
	addr := fmt.Sprintf("%s:%d", hostIP, config.Port)

	fmt.Println(info("\nStarting Server:"))
	fmt.Println(success("Serving directory:"), config.Dir)
	fmt.Println(success("Server Started at:"), homepageURL)

	fmt.Println(warning("DataGate is not recommended to be installed or used on critical servers\n"))
	if err := app.Listen(addr); err != nil {
		fmt.Println(errorMsg("Error: Failed to start server on"), warning(addr))
		fmt.Println(errorMsg("It might be in use or unavailable."))
	}
}
