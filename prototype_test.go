package gopatterns

import (
    "testing"
)

type Prototype1 struct{
    CutomField1 string
}

func (p *Prototype1) Clone() Prototype {
    cloned := *p
    return &cloned
}

type Prototype2 struct{
    CutomField2 string
}

func (p *Prototype2) Clone() Prototype {
    cloned := *p
    return &cloned
}

func TestPrototype(t *testing.T) {
    p1 := &Prototype1{}
    p2 := &Prototype2{}
    
    c1 := p1.Clone()
    switch v := c1.(type) {
    default:
        t.Errorf("Expected Prototype1 but got %v", v)
    case *Prototype1:
    }
    
    c2 := p2.Clone()
    switch v := c2.(type) {
    default:
        t.Errorf("Expected Prototype2 but got %v", v)
    case *Prototype2:
    }
}
