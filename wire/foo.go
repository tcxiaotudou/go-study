package main

type Fooer interface {
	HelloWorld()
}

type Foo struct{}

func (f Foo) HelloWorld() {}

type Bar struct {
	foo Fooer
}

func NewBar(fooer Fooer) *Bar {
	return &Bar{foo: fooer}
}
