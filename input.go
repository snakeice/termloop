package termloop

import (
	"github.com/gdamore/tcell"
)

type input struct {
	endKey tcell.Key
	eventQ chan interface{}
	ctrl   chan bool
}

func newInput() *input {
	i := input{eventQ: make(chan interface{}),
		ctrl:   make(chan bool, 2),
		endKey: tcell.KeyCtrlC}
	return &i
}

func (i *input) start(render tcell.Screen) {
	go poll(i, render)
}

func (i *input) stop() {
	i.ctrl <- true
}

func poll(i *input, render tcell.Screen) {
loop:
	for {
		select {
		case <-i.ctrl:
			break loop
		default:
			i.eventQ <-  render.PollEvent()
		}
	}
}
