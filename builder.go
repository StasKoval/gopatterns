package gopatterns

type AbstractBuilder interface {
    GetResult() AbstractProduct
    BuildPart1()
    BuildPart2()
}
