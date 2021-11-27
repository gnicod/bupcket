package api

import (
	"log"
	"os"

	"github.com/gnicod/bupcket/config"

	"github.com/gnicod/bupcket/storage"
	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
)

type App struct {
	router          *iris.Application
	storageProvider storage.Provider
}

func NewApp(storageProvider storage.Provider) *App {
	router := iris.New()
	app := &App{
		router:          router,
		storageProvider: storageProvider,
	}
	return app
}

func (app *App) Upload(ctx iris.Context) {
	_, info, err := ctx.FormFile("file")
	tag := ctx.URLParam("tag")
	config, err := config.NewTagConfig(tag)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON("Error while uploading")
		return
	}
	f, _ := os.Open(info.Filename)
	defer f.Close()
	path, err := app.storageProvider.Upload(storage.UploadRequest{
		Bucket: config.Bucket,
		Key:    uuid.NewString(),
		Body:   *f,
	})
	if err != nil {
		ctx.JSON("fail")
	}
	ctx.JSON(path)
}

// Run starts the APIs
func (app *App) Run() {
	app.router.Post("/upload", app.Upload)
	log.Fatal(app.router.Listen(":8090"))
}
