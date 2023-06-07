package goasynem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test__RegisterSubscriber(t *testing.T) {
	e := &Goasynem{}
	e.Subscribe("event", func(payload interface{}) error {
		fmt.Println("Data received", payload)
		return nil
	})

	_, ok := e.listeners["event"]
	assert.Equal(t, ok, true)
}
