package gopatterns

import (
    "runtime"
    "math/rand"
)

// the system resource table emulation 
var Resources *map[int]string

type Resource struct{
    Id int
}

func NewResource() *Resource {
    r := &Resource{Id: Initialize()}
    
    runtime.SetFinalizer(r, Deinitialize)
    
    return r
}

func Initialize() int {
    id := rand.Int()
    
    if Resources == nil {
        Resources = &map[int]string{}
    }
    
    (*Resources)[id] = "initialized"
    return id
}

func Deinitialize(r *Resource) {
    (*Resources)[r.Id] = "deinitialized"
}
