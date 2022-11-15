package plugin

import "reflect"

type Config struct {
	name string
}

type Factory interface {
	Create(conf Config) Plugin
}

type InputFactory struct{}

func (i *InputFactory) Create(config Config) Plugin {
	t := inputNames[config.name]
	return reflect.New(t).Interface().(Plugin)
}

type FilterFactory struct{}

func (f *FilterFactory) Create(config Config) Plugin {
	t := filterNames[config.name]
	return reflect.New(t).Interface().(Plugin)
}

type OutputFactory struct{}

func (o *OutputFactory) Create(conf Config) Plugin {
	t := outputNames[conf.name]
	return reflect.New(t).Interface().(Plugin)
}
