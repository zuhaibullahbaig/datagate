package server

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gofiber/fiber/v2"
	"github.com/zuhaibullahbaig/datagate/handlers"
)

type FileData struct {
	Name          string
	Path          string
	IsDir         bool
	Extension     string
	FormattedSize string
}

type TemplateData struct {
	Files              []FileData
	TotalDirectorySize string
}

func SetupRoutes(app *fiber.App, baseDir string, token string, authEnabled bool) {

	app.Static("/static", "./public")

	app.Get("/status", func(c *fiber.Ctx) error {
		return c.Render("status", fiber.Map{})
	})

	app.Get("/", func(c *fiber.Ctx) error {
		files, err := handlers.ListFiles(baseDir)
		if err != nil {
			return c.Status(500).SendString("Error Reading Directory")
		}

		var fileList []FileData
		var totalSize int64

		for _, file := range files {
			filePath := filepath.Join(baseDir, file.Name())

			fileInfo, err := os.Stat(filePath)
			if err != nil {
				continue
			}

			link := handlers.ConstructPath("", file.Name(), file.IsDir(), token, authEnabled)

			fileList = append(fileList, FileData{
				Name:          file.Name(),
				Path:          link,
				IsDir:         fileInfo.IsDir(),
				Extension:     handlers.GetFileExtension(file.Name()),
				FormattedSize: handlers.FormatFileSize(fileInfo.Size()),
			})

			if !fileInfo.IsDir() {
				totalSize += fileInfo.Size()
			}
		}

		tmpl, err := template.ParseFiles("views/index.html")
		if err != nil {
			return c.Status(500).SendString("Error loading template")
		}

		c.Set("Content-Type", "text/html")
		return tmpl.Execute(c.Response().BodyWriter(), TemplateData{
			Files:              fileList,
			TotalDirectorySize: handlers.FormatFileSize(totalSize),
		})
	})

	app.Get("/open", func(c *fiber.Ctx) error {
		dirQuery := c.Query("dir", "")
		targetDir := filepath.Join(baseDir, dirQuery)

		absBase, _ := filepath.Abs(baseDir)
		absTarget, _ := filepath.Abs(targetDir)
		if !strings.HasPrefix(absTarget, absBase) {
			return c.Status(403).SendString("Access Denied")
		}

		dirInfo, err := os.Stat(targetDir)
		if err != nil || !dirInfo.IsDir() {
			return c.Status(404).SendString("Directory not found")
		}

		files, err := handlers.ListFiles(targetDir)
		if err != nil {
			return c.Status(500).SendString("Error reading directory")
		}

		var fileList []FileData
		var totalSize int64

		for _, file := range files {
			fileInfo, err := file.Info()
			if err != nil {
				continue
			}

			link := handlers.ConstructPath(dirQuery, file.Name(), file.IsDir(), token, authEnabled)

			fileList = append(fileList, FileData{
				Name:          file.Name(),
				Path:          link,
				IsDir:         file.IsDir(),
				Extension:     handlers.GetFileExtension(file.Name()),
				FormattedSize: handlers.FormatFileSize(fileInfo.Size()),
			})

			if !file.IsDir() {
				totalSize += fileInfo.Size()
			}
		}

		tmpl, err := template.ParseFiles("views/index.html")
		if err != nil {
			return c.Status(500).SendString("Error loading template")
		}

		c.Set("Content-Type", "text/html")
		return tmpl.Execute(c.Response().BodyWriter(), TemplateData{
			Files:              fileList,
			TotalDirectorySize: handlers.FormatFileSize(totalSize),
		})
	})

	app.Get("/download", func(c *fiber.Ctx) error {
		fileQuery := c.Query("file", "")
		if fileQuery == "" {
			return c.Status(400).SendString("Missing file parameter")
		}

		targetPath := filepath.Join(baseDir, fileQuery)

		absBase, _ := filepath.Abs(baseDir)
		absTarget, _ := filepath.Abs(targetPath)
		if !strings.HasPrefix(absTarget, absBase) {
			return c.Status(403).SendString("Access Denied")
		}

		fileInfo, err := os.Stat(targetPath)
		if err != nil {
			if os.IsNotExist(err) {
				return c.Status(404).SendString("File not found")
			}
			return c.Status(500).SendString("Error accessing file")
		}

		if fileInfo.IsDir() {
			return c.Status(400).SendString("Cannot download a directory")
		}

		return c.Download(targetPath, fileInfo.Name())
	})
}
