package gopatterns

import (
    "testing"
)

func TestSingleton(t *testing.T) {
    instance1 := GetSingletonInstance()
    instance2 := GetSingletonInstance()
    
    if instance1 != instance2 {
        t.Errorf("Expected equal instances")
    }
}
