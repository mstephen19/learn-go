package lib

import (
	"sync"
	"time"
)

// A basic event emitter implementation.
type emitter struct {
	callbacks map[string][]*func(val byte)
}

func NewDataEmitter(content string) (em *emitter, waitGroup *sync.WaitGroup) {
	waitGroup = &sync.WaitGroup{}
	byteSlice := []byte(content)

	em = &emitter{
		callbacks: make(map[string][](*func(val byte))),
	}

	waitGroup.Add(len(byteSlice))

	go func() {
		for _, value := range byteSlice {
			em.Emit("data", value)
			time.Sleep(1 * time.Second)
			waitGroup.Done()
		}
	}()

	return
}

// Listen for messages sent on the emitter.
func (emitter *emitter) On(name string, callback *func(val byte)) {
	emitter.callbacks[name] = append(emitter.callbacks[name], callback)
}

// Emit an event on the emitter.
func (emitter *emitter) Emit(name string, val byte) {
	for _, callback := range emitter.callbacks[name] {
		cb := *callback
		cb(val)
	}
}
