package gopatterns

type Handler interface{
    Handle(msg string)
}
