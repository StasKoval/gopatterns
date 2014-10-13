package gopatterns

type RequiredInterface interface {
    MethodA()
}

type Adaptee struct {
}

func (a *Adaptee) MethodB() {
}

type Adapter struct{
    Impl Adaptee
}

func (a *Adapter) MethodA9() {
    a.Impl.MethodB()
}