package gopatterns

type Component interface {
    GetName() string
}

type Composite interface {
    Add(c Component)
    Remove(c Component)
    GetChildren() []Component
}