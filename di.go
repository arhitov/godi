package di

import "go.uber.org/dig"

type ContainerInterface interface {
	Scope(name string, opts ...dig.ScopeOption) ContainerInterface
	Bind(constructor any)
	Invoke(constructor any) error
	InvokeOrFail(constructor any)
	Close() error
}

type Container struct {
	container        *dig.Container
	constructorClose func() error
}

type Scope struct {
	scope            *dig.Scope
	constructorClose func() error
}

func NewContainer(dryRun bool, constructorClose func() error) *Container {
	return &Container{
		container:        dig.New(dig.DryRun(dryRun)),
		constructorClose: constructorClose,
	}
}

func (c Container) Scope(name string, opts ...dig.ScopeOption) ContainerInterface {
	return &Scope{
		scope: c.container.Scope(name, opts...),
	}
}

func (c Scope) Scope(name string, opts ...dig.ScopeOption) ContainerInterface {
	return &Scope{
		scope: c.scope.Scope(name, opts...),
	}
}

func (c Container) Dig() *dig.Container {
	return c.container
}

func (c Scope) Dig() *dig.Scope {
	return c.scope
}

func (c Container) Bind(constructor any) {
	if err := c.container.Provide(constructor); err != nil {
		panic(err)
	}
}

func (c Scope) Bind(constructor any) {
	if err := c.scope.Provide(constructor); err != nil {
		panic(err)
	}
}

func (c Container) Invoke(constructor any) error {
	return c.container.Invoke(constructor)
}

func (c Scope) Invoke(constructor any) error {
	return c.scope.Invoke(constructor)
}

func (c Container) InvokeOrFail(constructor any) {
	err := c.Invoke(constructor)
	if err != nil {
		panic(err)
	}
}

func (c Scope) InvokeOrFail(constructor any) {
	err := c.Invoke(constructor)
	if err != nil {
		panic(err)
	}
}

func (c Container) Close() error {
	if c.constructorClose != nil {
		return c.constructorClose()
	}
	return nil
}

func (c Scope) Close() error {
	if c.constructorClose != nil {
		return c.constructorClose()
	}
	return nil
}

// Instance Register an existing instance as shared in the container.
func Instance[T any](c ContainerInterface, instance interface{}) {
	c.Bind(func() T {
		return instance.(T)
	})
}

func IsRegistered[T any](c ContainerInterface) bool {
	err := c.Invoke(func(dep T) {})
	return err == nil
}

func Resolve[T any](c ContainerInterface) T {
	var res T
	c.InvokeOrFail(func(r T) {
		res = r
	})
	return res
}
