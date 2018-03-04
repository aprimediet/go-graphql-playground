package main

import (
	"flag"
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/graphql-go/handler"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

var (
	dir = flag.String("dir", "./static", "Directory to serve graphql")
)

func main() {
	h := handler.New(&handler.Config{
		Schema: &Schema,
		Pretty: true,
	})

	log.Printf("Starting to listen in :8080")

	// Trying fasthttp again
	router := fasthttprouter.New()
	adaptH := fasthttpadaptor.NewFastHTTPHandler(h)

	fs := &fasthttp.FS{
		Root:               *dir,
		IndexNames:         []string{"index.html"},
		GenerateIndexPages: true,
		Compress:           true,
	}

	fsHandler := fs.NewRequestHandler()

	router.ServeFiles("/css/*filepath", "./static/css")
	router.ServeFiles("/js/*filepath", "./static/js")
	router.GET("/graphql", adaptH)
	router.POST("/graphql", adaptH)
	router.GET("/", fsHandler)

	panic(fasthttp.ListenAndServe(":8080", router.Handler))

}
