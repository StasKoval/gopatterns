package gopatterns

import (
    "testing"
)

type ComplexProduct struct {
    Product1 AbstractProduct
    Product2 AbstractProduct
}

func (cp *ComplexProduct) GetName() string {
    return cp.Product1.GetName() + " " + cp.Product2.GetName()
}

type Builder struct{
    Product ComplexProduct
}

func (b *Builder) BuildPart1() {
    b.Product.Product1 = &ProductA1{}
}

func (b *Builder) BuildPart2() {
    b.Product.Product2 = &ProductB2{}
}

func (b *Builder) GetResult() AbstractProduct {
    return &b.Product
}

func TestBuilder(t *testing.T) {
    b := Builder{ComplexProduct{}}
    
    b.BuildPart1()
    b.BuildPart2()
    
    expectedResult := "ProductA1 ProductB2"
    if b.GetResult().GetName() != expectedResult {
        t.Errorf("Expected %v but got ", expectedResult, b.GetResult().GetName())
    }
}
