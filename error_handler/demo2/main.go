package main

import (
	"fmt"

	"github.com/pkg/errors"
)

var ErrNameEmpty = errors.New("Name can't be empty!")

func (s *Student) SetName(newName string) (err error) {
	if newName == "" || s.Name == "" {
		return ErrNameEmpty
	} else {
		s.Name = newName
		return nil
	}
}

type Student struct {
	Name string
	Age  int
}

func NewStu() (err error) {
	stu := &Student{Age: 19}
	e := stu.SetName("")

	if e != nil {
		return errors.Wrap(e, "set name failed!")
	} else {
		return nil
	}
}

func main() {
	e := NewStu()
	fmt.Printf("%+v\n", e)
}
