package main

import "fmt"

func same[T int | float64 | string](a, b T) bool {
	return a == b
}

type Person[T any] struct {
	Name string
	Sex  T
}

type TMap[K string | int, V string | int] map[K]V
type TSlice[S any] []S
type MyType interface {
	~int | float64 | string | int32
}

func Test[T MyType](s T) {
	fmt.Println(s)
}

func main() {
	fmt.Println("go-study world")
	fmt.Println(same(1, 1))
	p := Person[string]{
		Name: "123",
		Sex:  "tang",
	}
	fmt.Println(p.Sex)
	m := make(TMap[int, string], 10)
	m[1] = "123"
	s := make(TSlice[string], 6)
	s[5] = "456"

}
