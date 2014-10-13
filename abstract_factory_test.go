package gopatterns

import (
    "testing"
)

type ProductA1 struct{}

func (p *ProductA1) GetName() string {
    return "ProductA1"
}

type ProductA2 struct{}

func (p *ProductA2) GetName() string {
    return "ProductA2"
}

type ProductB1 struct{}

func (p *ProductB1) GetName() string {
    return "ProductB1"
}

type ProductB2 struct{}

func (p *ProductB2) GetName() string {
    return "ProductB2"
}

type FactoryA struct{}

func (f *FactoryA) CreateProduct1() AbstractProduct {
    return &ProductA1{}
}

func (f *FactoryA) CreateProduct2() AbstractProduct {
    return &ProductA2{}
}

type FactoryB struct{}

func (f *FactoryB) CreateProduct1() AbstractProduct {
    return &ProductB1{}
}

func (f *FactoryB) CreateProduct2() AbstractProduct {
    return &ProductB2{}
}

func TestAbstractFactory(t *testing.T) {
    getResult := func(f AbstractFactory) string {
        p1 := f.CreateProduct1()
        p2 := f.CreateProduct2()
        return p1.GetName() + " " + p2.GetName()
    }
    
    expectedResultA := "ProductA1 ProductA2"
    if getResult(&FactoryA{}) != expectedResultA {
        t.Errorf("Expected %v but got %v", expectedResultA, getResult(&FactoryA{}))
    }
    
    expectedResultB := "ProductB1 ProductB2"
    if getResult(&FactoryB{}) != expectedResultB {
        t.Errorf("Expected %v but got ", expectedResultB, getResult(&FactoryB{}))
    }
}