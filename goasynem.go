package goasynem

import (
	"errors"
	"sync"
)

type AsyncEmitter interface {
	create()
	Subscribe(event string, f func(interface{})) (chan interface{}, error)
	on(event string) (chan interface{}, error)
	Emit(event string, payload interface{}) (chan struct{}, error)
}

type Goasynem struct {
	once      sync.Once
	mx        sync.Mutex
	listeners map[string]chan interface{}
}

func (e *Goasynem) create() {
	e.once.Do(func() {
		e.listeners = make(map[string]chan interface{})
	})
}

func (e *Goasynem) on(event string) (chan interface{}, error) {
	e.mx.Lock()
	defer e.mx.Unlock()
	e.create()
	ch := make(chan interface{})
	if _, ok := e.listeners[event]; ok {
		return nil, errors.New("listener already exisits")
	}
	e.listeners[event] = ch
	return ch, nil
}

func (e *Goasynem) Subscribe(event string, f func(interface{})) (chan interface{}, error) {
	ch, err := e.on(event)
	if err != nil {
		return ch, err
	}
	go func() {
		for p := range ch {
			f(p)
		}
	}()
	return ch, nil
}

func (e *Goasynem) Emit(event string, payload interface{}) (chan struct{}, error) {
	e.mx.Lock()
	defer e.mx.Unlock()
	e.create()
	done := make(chan struct{}, 1)
	var wg sync.WaitGroup
	var err error

	if ch, ok := e.listeners[event]; !ok {
		return nil, errors.New("event not registered")
	} else {
		wg.Add(1)
		go func(ch chan interface{}, d *interface{}) {
			e.mx.Lock()
			defer func() {
				wg.Done()
				e.mx.Unlock()
			}()
			select {
			case <-done:
				return
			case ch <- payload:
				return
			default:
				return
			}
		}(ch, &payload)
	}

	go func(done chan struct{}) {
		defer func() {
			recover()
		}()
		wg.Wait()
		close(done)
	}(done)

	return done, err
}
