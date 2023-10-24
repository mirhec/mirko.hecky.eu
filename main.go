package main

import (
	"log"
	"net/http"
	"os"

	"site/models"
	"site/utils"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/template"
)

func main() {
	app := pocketbase.New()

	utils.RegisterMigrations(app)

	// serves static files from the provided public dir (if exists)
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// this is safe to be used by multiple goroutines
		// (it acts as store for the parsed templates)
		registry := template.NewRegistry()

		e.Router.GET("/slides/:id", func(c echo.Context) error {
			id := c.PathParam("id")
			content, err := utils.GetFileContent(app, registry, "slides", id, "header_image")
			if err != nil {
				return err
			}
			return c.Blob(http.StatusOK, "image/jpeg", content)
		})

		e.Router.GET("/songs/:title", func(c echo.Context) error {
			title := c.PathParam("title")
			content, err := utils.GetFileContentByTitle(app, registry, "songs", title, "mp3")
			if err != nil {
				return err
			}
			return c.Blob(http.StatusOK, "audio/mpeg", content)
		})

		e.Router.GET("/:slug", func(c echo.Context) error {
			slug := c.PathParam("slug")
			return utils.RenderPage(app, registry, c, "pages", slug)
		} /* optional middlewares */)

		e.Router.GET("/", func(c echo.Context) error {
			return utils.RenderPage(app, registry, c, "pages", "home")
		} /* optional middlewares */)

		e.Router.GET("/blog/:slug", func(c echo.Context) error {
			slug := c.PathParam("slug")
			return utils.RenderPage(app, registry, c, "posts", slug)
		})

		e.Router.GET("/blog", func(c echo.Context) error {
			records, err := app.Dao().FindRecordsByExpr("posts")
			if err != nil {
				return apis.NewNotFoundError("No posts found", err)
			}

			posts := []models.Post{}
			for _, postRecord := range records {
				posts = append(posts, models.Post{
					Title: postRecord.GetString("title"),
					Url:   "/blog/" + postRecord.GetString("slug"),
					Date:  postRecord.GetDateTime("date").Time().Format("02.01.2006"),
				})
			}

			homeRecord, err := app.Dao().FindFirstRecordByData("pages", "slug", "home")
			if err != nil {
				return apis.NewNotFoundError("No home page found", err)
			}
			slide := homeRecord.GetString("slide")

			html, err := registry.LoadFiles(
				"views/blog-layout.html",
			).Render(map[string]any{
				"posts": posts,
				"slide": slide,
			})
			if err != nil {
				return err
			}

			return c.HTML(http.StatusOK, html)
		} /* optional middlewares */)

		e.Router.GET("/static/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))

		return nil
	})

	println("starting pocketbase server...")

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
