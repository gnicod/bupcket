package api

import (
	"log"

	"github.com/kataras/iris/v12"
)

type App struct {
	router *iris.Application
}
func NewApp() *App {
	router := iris.New()
	app := &App{router: router}
	app.router.Post("/upload", app.Upload) 
	return app
}

func (app *App) Upload(ctx iris.Context) {
	ctx.Writef("Hello<br/>")
    file, info, err := ctx.FormFile("file")
    if err != nil {
        ctx.StatusCode(iris.StatusInternalServerError)
        ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
        return
    }
    defer file.Close()
    fn := info.Filename
    ctx.Writef("File Name: " + fn)
	ctx.JSON("ok")
}

// Run starts the APIs
func (app *App) Run() {
	log.Fatal(app.router.Listen(":8090"))
   }