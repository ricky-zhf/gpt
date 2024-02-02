package model

import (
	"context"
	"fmt"
	"time"
)

type Spin struct {
	ch     chan int
	ctx    context.Context
	cancel context.CancelFunc
}

func NewSpin() Spin {
	ctx, cancelFunc := context.WithCancel(context.Background())
	s := Spin{ch: make(chan int), ctx: ctx, cancel: cancelFunc}
	s.Start()
	return s
}

func (s Spin) Start() {
	go func() {
		for {
			select {
			case <-s.ctx.Done():
				return
			default:
				s.print()
			}
		}
	}()
}

func (s Spin) Stop() {
	s.cancel()
}

func (s Spin) print() {
	for _, r := range "-\\|/" {
		fmt.Printf("\r%c", r)
		time.Sleep(200 * time.Millisecond)
	}
}
