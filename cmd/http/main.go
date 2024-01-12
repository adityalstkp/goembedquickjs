package main

import (
	"log"
	"net/http"
	"runtime"
	"text/template"

	"github.com/adityalstkp/goembedquickjs/internal/renderer"
)

func main() {
	requestChan := make(chan string)
	receiverChan := make(chan string)

	indexTemplate, err := template.New("index").ParseGlob("template/*.html")
	if err != nil {
		log.Panicln(err)
	}

	pool := renderer.NewPool(renderer.PoolOpts{
		RequestChan:  requestChan,
		ReceiverChan: receiverChan,
		Workers:      runtime.NumCPU() - 1,
	})
	go pool.Start()

	rendererHandler := NewRenderHandler(RenderHandlerOpts{
		RequestChan:  requestChan,
		ReceiverChan: receiverChan,
		Template:     indexTemplate,
	})

	r := newRouter(routerConfig{
		reactRenderer: rendererHandler,
	})

	log.Panicln(http.ListenAndServe(":3111", r))
}
