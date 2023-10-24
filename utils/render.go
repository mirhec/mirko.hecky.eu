package utils

import (
	t "html/template"
	"io"
	"net/http"

	"site/models"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/tools/template"
)

func RenderPage(app *pocketbase.PocketBase, registry *template.Registry, c echo.Context, collection string, slug string) error {
	record, err := app.Dao().FindFirstRecordByData(collection, "slug", slug)
	if err != nil {
		return apis.NewNotFoundError("No page found", err)
	}

	songs := []models.Song{}
	for _, songId := range record.GetStringSlice("songs") {
		songRecord, err := app.Dao().FindRecordById("songs", songId)
		if err != nil {
			return err
		}
		songs = append(songs, models.Song{
			Title: songRecord.GetString("title"),
			Url:   "/songs/" + songRecord.GetString("title"),
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

func GetFileContent(app *pocketbase.PocketBase, registry *template.Registry, collection string, id string, fieldName string) ([]byte, error) {
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

func GetFileContentByTitle(app *pocketbase.PocketBase, registry *template.Registry, collection string, title string, fieldName string) ([]byte, error) {
	record, err := app.Dao().FindFirstRecordByData(collection, "title", title)
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
