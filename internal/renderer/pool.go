package renderer

import (
	"log"
	"sync"

	"github.com/rosbit/go-quickjs"
)

type pool struct {
	requestChan  chan string
	receiverChan chan string
	workers      int
	wg           sync.WaitGroup
}

type PoolOpts struct {
	RequestChan  chan string
	ReceiverChan chan string
	Workers      int
}

func NewPool(opts PoolOpts) *pool {
	return &pool{
		requestChan:  opts.RequestChan,
		receiverChan: opts.ReceiverChan,
		workers:      opts.Workers,
	}
}

func (p *pool) Start() {
	p.wg.Add(p.workers)
	for i := 0; i < p.workers; i++ {
		go p.receiveRequest()
	}
	p.wg.Wait()
}

func (p *pool) receiveRequest() {
	defer p.wg.Done()

	qjsCtx, err := quickjs.NewContext()
	if err != nil {
		log.Panicln(err)
	}

	renderer, err := NewRenderer(qjsCtx, nil)
	if err != nil {
		log.Panicln(err)
	}

	for range p.requestChan {
		v, _ := renderer.Render()
		p.receiverChan <- v
	}
}
