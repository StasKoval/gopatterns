package gopatterns

import (
    "testing"
    "runtime"
    "time"
)


func TestRaii(t *testing.T) {
    r := NewResource()
    id := r.Id
    
    if  state := (*Resources)[id]; state != "initialized" {
        t.Errorf("Expected initialized but got %v", state)
    }
    
    r = nil
    
    runtime.GC()
    time.Sleep(time.Second)
    runtime.GC()
    
    if  state := (*Resources)[id]; state != "deinitialized" {
        t.Errorf("Expected deinitialized but got %v", state)
    }
}
