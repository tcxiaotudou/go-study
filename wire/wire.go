//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func InitRoot(name string) Root {
	wire.Build(NewLeaf, NewBranch, NewRoot)
	return Root{}
}

var NewBranchSet = wire.NewSet(NewLeaf, NewBranch)

func InitRootWithSet(name string) Root {
	wire.Build(NewBranchSet, NewRoot)
	return Root{}
}

var bind = wire.Bind(new(Fooer), new(Foo))

func InitBar(foo Foo) *Bar {
	wire.Build(NewBar, bind)
	return nil
}
