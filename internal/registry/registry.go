package registry

var handlerConstructors []interface{}

func Register(ctor interface{}) {
	handlerConstructors = append(handlerConstructors, ctor)
}

func List() []interface{} {
	return handlerConstructors
}
