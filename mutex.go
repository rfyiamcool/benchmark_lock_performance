package main

import (
	"sync"
)

type notLocked struct {
	i int
}

type locked struct {
	i int
	l sync.Mutex
}

func (c *notLocked) add() {
	c.i++
}

func (c *locked) add() {
	c.l.Lock()
	c.i++
	c.l.Unlock()
}
