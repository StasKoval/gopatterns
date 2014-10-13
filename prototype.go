package gopatterns

type Prototype interface{
    Clone() Prototype
}