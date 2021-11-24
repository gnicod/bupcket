package api

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gnicod/bupcket/storage"
	"github.com/kataras/iris/v12"
)

type App struct {
	router *iris.Application
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
    file, info, err := ctx.FormFile("file")
    if err != nil {
        ctx.StatusCode(iris.StatusInternalServerError)
        ctx.JSON("Error while uploading")
        return
    }
    defer file.Close()
    fn := info.Filename
    ctx.Writef("File Name: " + fn)
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", info.Filename)
	fmt.Printf("File Size: %+v\n", info.Size)
	fmt.Printf("MIME Header: %+v\n", info.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile("/tmp/", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	ctx.JSON("ok")
	app.storageProvider.Upload(storage.UploadRequest{})
}

// Run starts the APIs
func (app *App) Run() {
	app.router.Post("/upload", app.Upload)
	log.Fatal(app.router.Listen(":8090"))
   }