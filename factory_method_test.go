package gopatterns

import (
    "testing"
)

type Creator1 struct{
}

func (c *Creator1) NewProduct() AbstractProduct {
    return &ProductA1{}
}

type Creator2 struct{
}

func (c *Creator2) NewProduct() AbstractProduct {
    return &ProductA2{}
}

func TestFactoryMethod(t *testing.T) {
    c1 := Creator1{}
    c2 := Creator2{}
    
    expectedResult1 := "ProductA1"
    if c1.NewProduct().GetName() != expectedResult1 {
        t.Errorf("Expected %v but got %v", expectedResult1, c1.NewProduct().GetName())
    }
    
    expectedResult2 := "ProductA2"
    if c2.NewProduct().GetName() != expectedResult2 {
        t.Errorf("Expected %v but got ", expectedResult2, c1.NewProduct().GetName())
    }
}
