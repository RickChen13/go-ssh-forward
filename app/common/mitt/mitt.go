package mitt

import (
	"reflect"
	"sync"
)

var Mitt = New()

// Handler is the event callback type. It receives any payload.
type Handler func(payload any)

// Emitter is a lightweight event emitter similar to JS mitt.
type Emitter struct {
	mu       sync.RWMutex
	handlers map[string][]Handler
}

// New creates a new Emitter.
func New() *Emitter {
	return &Emitter{
		handlers: make(map[string][]Handler),
	}
}

// On registers handler for event. It returns an unsubscribe function.
func (e *Emitter) On(event string, h Handler) (unsubscribe func()) {
	e.mu.Lock()
	e.handlers[event] = append(e.handlers[event], h)
	e.mu.Unlock()

	return func() {
		e.Off(event, h)
	}
}

// Once registers a handler that will be called at most once.
func (e *Emitter) Once(event string, h Handler) (unsubscribe func()) {
	var wrapper Handler
	wrapper = func(payload any) {
		// ensure unsubscribe is called before calling original handler
		e.Off(event, wrapper)
		h(payload)
	}
	return e.On(event, wrapper)
}

// Off removes the given handler for event. If h is nil, remove all handlers for event.
func (e *Emitter) Off(event string, h Handler) {
	e.mu.Lock()
	defer e.mu.Unlock()
	if h == nil {
		delete(e.handlers, event)
		return
	}
	list, ok := e.handlers[event]
	if !ok || len(list) == 0 {
		return
	}
	// filter out handlers equal to h
	newList := list[:0]
	for _, hh := range list {
		if reflect.ValueOf(hh).Pointer() != reflect.ValueOf(h).Pointer() {
			newList = append(newList, hh)
		}
	}
	if len(newList) == 0 {
		delete(e.handlers, event)
	} else {
		e.handlers[event] = newList
	}
}

// Emit emits event with payload. Handlers are called synchronously in the caller goroutine.
// Panics from handlers are recovered to avoid crashing the emitter.
func (e *Emitter) Emit(event string, payload any) {
	// copy handlers to avoid holding lock while calling user code and avoid races with Off/On.
	e.mu.RLock()
	list, ok := e.handlers[event]
	if !ok || len(list) == 0 {
		e.mu.RUnlock()
		return
	}
	handlers := make([]Handler, len(list))
	copy(handlers, list)
	e.mu.RUnlock()

	for _, h := range handlers {
		func(h Handler) {
			defer func() {
				if r := recover(); r != nil {
					// swallow panic from handler to keep emitter robust;
					// if you want to log it, you can extend this package.
				}
			}()
			h(payload)
		}(h)
	}
}

// Clear removes all handlers for all events.
func (e *Emitter) Clear() {
	e.mu.Lock()
	e.handlers = make(map[string][]Handler)
	e.mu.Unlock()
}
