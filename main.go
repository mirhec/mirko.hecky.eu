package main

import (
	t "html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/template"
)

type Song struct {
	Title string
	Url   string
}

type Post struct {
	Title string
	Url   string
	Date  string
}

func main() {
	println("Starting ...")

	app := pocketbase.New()

	// serves static files from the provided public dir (if exists)
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// this is safe to be used by multiple goroutines
		// (it acts as store for the parsed templates)
		registry := template.NewRegistry()

		renderPage := func(c echo.Context, collection string, slug string) error {
			record, err := app.Dao().FindFirstRecordByData(collection, "slug", slug)
			if err != nil {
				return apis.NewNotFoundError("No page found", err)
			}

			songs := []Song{}
			for _, songId := range record.GetStringSlice("songs") {
				songRecord, err := app.Dao().FindRecordById("songs", songId)
				if err != nil {
					return err
				}
				songs = append(songs, Song{
					Title: songRecord.GetString("title"),
					Url:   "/songs/" + songId,
				})
			}

			html, err := registry.LoadFiles(
				"views/layout.html",
			).Render(map[string]any{
				"content": t.HTML(record.GetString("content")),
				"slide":   record.GetString("slide"),
				"title":   record.GetString("title"),
				"songs":   songs,
			})
			if err != nil {
				return err
			}

			return c.HTML(http.StatusOK, html)
		}

		getFileContent := func(collection string, id string, fieldName string) ([]byte, error) {
			record, err := app.Dao().FindRecordById(collection, id)
			if err != nil {
				return nil, err
			}

			fs, err := app.NewFilesystem()
			if err != nil {
				return nil, err
			}
			defer fs.Close()

			fileKey := record.BaseFilesPath() + "/" + record.GetString(fieldName)
			br, err := fs.GetFile(fileKey)
			if err != nil {
				return nil, err
			}
			defer br.Close()

			content, err := io.ReadAll(br)
			if err != nil {
				return nil, err
			}

			return content, nil
		}

		e.Router.GET("/slides/:id", func(c echo.Context) error {
			id := c.PathParam("id")
			content, err := getFileContent("slides", id, "header_image")
			if err != nil {
				return err
			}
			return c.Blob(http.StatusOK, "image/jpeg", content)
		})

		e.Router.GET("/songs/:id", func(c echo.Context) error {
			id := c.PathParam("id")
			content, err := getFileContent("songs", id, "mp3")
			if err != nil {
				return err
			}
			return c.Blob(http.StatusOK, "audio/mpeg", content)
		})

		e.Router.GET("/:slug", func(c echo.Context) error {
			slug := c.PathParam("slug")
			return renderPage(c, "pages", slug)
		} /* optional middlewares */)

		e.Router.GET("/", func(c echo.Context) error {
			return renderPage(c, "pages", "home")
		} /* optional middlewares */)

		e.Router.GET("/blog/:slug", func(c echo.Context) error {
			slug := c.PathParam("slug")
			return renderPage(c, "posts", slug)
		})

		e.Router.GET("/blog", func(c echo.Context) error {
			records, err := app.Dao().FindRecordsByExpr("posts")
			if err != nil {
				return apis.NewNotFoundError("No posts found", err)
			}

			posts := []Post{}
			for _, postRecord := range records {
				posts = append(posts, Post{
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
