package reflect

import (
	"fmt"
	"reflect"
)

type Container struct {
	s reflect.Value
}

func NewContainer(t reflect.Type, size int) *Container {
	if size < 0 {
		size = 64
	}
	return &Container{s: reflect.MakeSlice(reflect.SliceOf(t), 0, size)}
}

func (c *Container) Put(val interface{}) error {
	if reflect.ValueOf(val).Type() != c.s.Type().Elem() {
		return fmt.Errorf("put: cannot put a %T into a slice of %s",
			val, c.s.Type().Elem())
	}
	c.s = reflect.Append(c.s, reflect.ValueOf(val))
	return nil
}

func (c *Container) Get(refval interface{}) error {
	if reflect.ValueOf(refval).Kind() != reflect.Ptr ||
		reflect.ValueOf(refval).Elem().Type() != c.s.Type().Elem() {
		return fmt.Errorf("get: needs *%s but got %T", c.s.Type().Elem(), refval)
	}
	reflect.ValueOf(refval).Elem().Set(c.s.Index(0))
	c.s = c.s.Slice(1, c.s.Len())
	return nil
}
