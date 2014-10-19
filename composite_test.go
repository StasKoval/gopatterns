package gopatterns

import (
    "testing"
)

type ConcreteComponent struct{
    Name string
    Children []Component
}

func (c *ConcreteComponent) GetName() string {
    return c.Name
}

func (c *ConcreteComponent) Add(child Component) {
    c.Children = append(c.Children, child)
}

func (c *ConcreteComponent) Remove(child Component) {
    children := []Component{}
    for _, ch := range c.Children {
        if ch != child {
            children = append(children, ch)
        }
    }
    c.Children = children
}

func (c *ConcreteComponent) GetChildren() []Component {
    return c.Children
}

func TestComposite(t *testing.T) {
    root := &ConcreteComponent{Name: "Parent"}
    
    child1 := &ConcreteComponent{Name: "Child1"}
    root.Add(child1)
    
    child2 := &ConcreteComponent{Name: "Child2"}
    root.Add(child2)
    
    root.Add(&ConcreteComponent{Name: "Child3"})
    
    root.Remove(child2)
    
    if len(root.GetChildren()) != 2 {
        t.Errorf("Expected %v but got %v", 2, len(root.GetChildren()))
    }
    
    root.Remove(child1)
    
    if root.GetChildren()[0].GetName() != "Child3" {
        t.Errorf("Expected %v but got %v", "Child3", root.GetChildren()[0].GetName())
    }
}