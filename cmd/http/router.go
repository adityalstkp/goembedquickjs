package main

import (
	"net/http"
	"text/template"

	"github.com/go-chi/chi/v5"
)

type rendererHandlerRouter interface {
	renderReact(w http.ResponseWriter, r *http.Request)
}

type renderHandler struct {
	requestChan  chan string
	receiverChan chan string
	template     *template.Template
}

type RenderHandlerOpts struct {
	RequestChan  chan string
	ReceiverChan chan string
	Template     *template.Template
}

func NewRenderHandler(opts RenderHandlerOpts) *renderHandler {
	return &renderHandler{
		requestChan:  opts.RequestChan,
		receiverChan: opts.ReceiverChan,
		template:     opts.Template,
	}
}

func (h *renderHandler) renderReact(w http.ResponseWriter, r *http.Request) {
	h.requestChan <- r.URL.Path
	res := <-h.receiverChan
	d := make(map[string]string)
	d["App"] = res
	h.template.ExecuteTemplate(w, "index.html", d)
}

type routerConfig struct {
	reactRenderer rendererHandlerRouter
}

func newRouter(rC routerConfig) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", rC.reactRenderer.renderReact)

	return r
}
