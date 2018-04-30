package providers

import "github.com/facebookgo/inject"

type Injector func(graph *inject.Graph)

func NewNamedObject(name string, val interface{}) (obj *inject.Object) {
	obj = NewObject(val)
	obj.Name = name
	return
}

func NewObject(val interface{}) *inject.Object {
	return &inject.Object{
		Value: val,
	}
}
