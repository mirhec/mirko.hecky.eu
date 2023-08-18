package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `[
			{
				"id": "_pb_users_auth_",
				"created": "2023-08-18 12:03:48.474Z",
				"updated": "2023-08-18 12:03:48.500Z",
				"name": "users",
				"type": "auth",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "users_name",
						"name": "name",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "users_avatar",
						"name": "avatar",
						"type": "file",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"maxSize": 5242880,
							"mimeTypes": [
								"image/jpeg",
								"image/png",
								"image/svg+xml",
								"image/gif",
								"image/webp"
							],
							"thumbs": null,
							"protected": false
						}
					}
				],
				"indexes": [],
				"listRule": "id = @request.auth.id",
				"viewRule": "id = @request.auth.id",
				"createRule": "",
				"updateRule": "id = @request.auth.id",
				"deleteRule": "id = @request.auth.id",
				"options": {
					"allowEmailAuth": true,
					"allowOAuth2Auth": true,
					"allowUsernameAuth": true,
					"exceptEmailDomains": null,
					"manageRule": null,
					"minPasswordLength": 8,
					"onlyEmailDomains": null,
					"requireEmail": false
				}
			},
			{
				"id": "hbz4jbr57m263k0",
				"created": "2023-08-18 12:13:14.847Z",
				"updated": "2023-08-18 12:15:47.247Z",
				"name": "pages",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "6gb3m0uj",
						"name": "title",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "veshzmiu",
						"name": "slug",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "aczqh0sa",
						"name": "seoDescription",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "qxqorhwz",
						"name": "content",
						"type": "editor",
						"required": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "bqnqigda",
						"name": "slide",
						"type": "relation",
						"required": false,
						"unique": false,
						"options": {
							"collectionId": "f71dcbkfkt3wu2p",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": []
						}
					},
					{
						"system": false,
						"id": "oayolxwt",
						"name": "songs",
						"type": "relation",
						"required": false,
						"unique": false,
						"options": {
							"collectionId": "5yqg24w7y6jy337",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": null,
							"displayFields": []
						}
					}
				],
				"indexes": [],
				"listRule": null,
				"viewRule": null,
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "f71dcbkfkt3wu2p",
				"created": "2023-08-18 12:14:15.328Z",
				"updated": "2023-08-18 12:22:38.971Z",
				"name": "slides",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "l1tihkkw",
						"name": "title",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "b5qvxhed",
						"name": "header_image",
						"type": "file",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"maxSize": 5242880,
							"mimeTypes": [],
							"thumbs": [],
							"protected": false
						}
					}
				],
				"indexes": [],
				"listRule": null,
				"viewRule": null,
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "5yqg24w7y6jy337",
				"created": "2023-08-18 12:14:53.189Z",
				"updated": "2023-08-18 12:14:53.189Z",
				"name": "songs",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "e8ta5wu6",
						"name": "title",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "xrzda9zg",
						"name": "mp3",
						"type": "file",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"maxSize": 52428800,
							"mimeTypes": [],
							"thumbs": [],
							"protected": false
						}
					}
				],
				"indexes": [],
				"listRule": null,
				"viewRule": null,
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "0b4lgo66j33slv7",
				"created": "2023-08-18 12:36:48.199Z",
				"updated": "2023-08-18 12:36:48.199Z",
				"name": "posts",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "rqkyf4fq",
						"name": "title",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "7sgqd3sc",
						"name": "slug",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "brviu1pi",
						"name": "seoDescription",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "gtafzaqd",
						"name": "date",
						"type": "date",
						"required": false,
						"unique": false,
						"options": {
							"min": "",
							"max": ""
						}
					},
					{
						"system": false,
						"id": "a29c8omn",
						"name": "content",
						"type": "editor",
						"required": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "dio3sfkd",
						"name": "slide",
						"type": "relation",
						"required": false,
						"unique": false,
						"options": {
							"collectionId": "f71dcbkfkt3wu2p",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": []
						}
					},
					{
						"system": false,
						"id": "ttoxfdkb",
						"name": "songs",
						"type": "relation",
						"required": false,
						"unique": false,
						"options": {
							"collectionId": "5yqg24w7y6jy337",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": null,
							"displayFields": []
						}
					}
				],
				"indexes": [],
				"listRule": null,
				"viewRule": null,
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			}
		]`

		collections := []*models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collections); err != nil {
			return err
		}

		return daos.New(db).ImportCollections(collections, true, nil)
	}, func(db dbx.Builder) error {
		return nil
	})
}
