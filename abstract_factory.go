package gopatterns

type AbstractFactory interface {
    CreateProduct1() AbstractProduct
    CreateProduct2() AbstractProduct
}
