package goasynem

import (
	"errors"
	"fmt"
	"sync"
)

type fn = func(interface{}) error

type Emitter struct {
	mu        sync.Mutex
	once      sync.Once
	listeners map[string]fn
}

func (e *Emitter) create() {
	e.once.Do(func() {
		e.listeners = make(map[string]fn)
	})
}

func (e *Emitter) on(evt string, f fn) error {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.create()
	if _, ok := e.listeners[evt]; ok {
		return errors.New("listener already exisits")
	}
	e.listeners[evt] = f
	return nil
}

func (e *Emitter) Subscribe(event string, f func(interface{}) error) error {
	err := e.on(event, f)
	if err != nil {
		return fmt.Errorf("error while registering subscriber: %s", err.Error())
	}
	return nil
}

func (e *Emitter) Emit(evt string, d interface{}) chan error {
	e.mu.Lock()
	defer e.mu.Unlock()

	listener, ok := e.listeners[evt]
	if !ok {
		return nil
	}
	err := make(chan error, 1)
	var wg sync.WaitGroup
	wg.Add(1)

	go func(f fn, errCh chan error) {
		defer wg.Done()
		err := f(d)
		if err != nil {
			errCh <- err
		}
		close(errCh)
	}(listener, err)

	wg.Wait()
	return err
}
