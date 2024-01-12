package renderer

import (
	"github.com/rosbit/go-quickjs"
)

type RendererOpts struct {
	FileEntry   string
	Environment map[string]interface{}
}

type renderer struct {
	quickjsCtx *quickjs.JsContext
}

func NewRenderer(ctx *quickjs.JsContext, opts *RendererOpts) (*renderer, error) {
	if opts == nil {
		_, err := ctx.EvalFile("react-embed/index.js", nil)
		if err != nil {
			return nil, err
		}
		return &renderer{quickjsCtx: ctx}, nil
	}

	_, err := ctx.EvalFile(opts.FileEntry, opts.Environment)
	if err != nil {
		return nil, err
	}
	return &renderer{quickjsCtx: ctx}, err

}

func (r *renderer) Render() (string, error) {
	res, err := r.quickjsCtx.CallFunc("renderToStringEmbed")
	if err != nil {
		return "", err
	}
	return res.(string), nil
}
