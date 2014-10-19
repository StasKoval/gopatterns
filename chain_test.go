package gopatterns

import (
    "testing"
)

var lastHandler string

type ConcreteHandlerA struct {
    nextHandler Handler
}

func (h *ConcreteHandlerA) Handle(msg string) {
    if msg == "A" {
        // handle msg
        lastHandler = "handler A"
    } else if next := h.nextHandler; next != nil {
        next.Handle(msg)
    }
}

type ConcreteHandlerB struct {
    nextHandler Handler
}

func (h *ConcreteHandlerB) Handle(msg string) {
    if msg == "B" {
        // handle msg
        lastHandler = "handler B"
    } else if next := h.nextHandler; next != nil {
        next.Handle(msg)
    }
}

type ConcreteHandlerC struct {
    nextHandler Handler
}

func (h *ConcreteHandlerC) Handle(msg string) {
    if msg == "C" {
        // handle msg
        lastHandler = "handler C"
    } else if next := h.nextHandler; next != nil {
        next.Handle(msg)
    }
}

func TestChain(t *testing.T) {
    handlerA := ConcreteHandlerA{}
    handlerB := ConcreteHandlerB{&handlerA}
    handlerC := ConcreteHandlerC{&handlerB}
    
    handlerC.Handle("C")
    if lastHandler != "handler C" {
        t.Errorf("Expected %v but got %v", "handler C", lastHandler )
    }
    
    handlerC.Handle("A")
    if lastHandler != "handler A" {
        t.Errorf("Expected %v but got %v", "handler A", lastHandler)
    }
    
    lastHandler = ""
    handlerC.Handle("X")
    if lastHandler != "" {
        t.Errorf("Expected %v but got %v", "empty string", lastHandler )
    }
}