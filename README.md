
```go
// Initialise
container := di.NewContainer(dryRun, func() error {
    log.Infof("NewContainer close")
})


func YourRealisation() YourInterfaces {
    some code
}

// Bind
container.Bind(func() YourInterfaces {
    return YourRealisation()
})
// Or Bind
container.Bind(YourRealisation)
// Or Bind
di.Instance[YourInterfaces](container, YourRealisation{})

// Scope
containerScope := container.Scope("scope_name")
containerScope.Bind(func() YourInterfacesTwo {
    return YourRealisationTwo()
})

// container
di.IsRegistered[YourInterfaces](container) // true
di.IsRegistered[YourInterfacesTwo](container) // false
di.Resolve[YourInterfaces](container) // YourInterfaces
di.Resolve[YourInterfacesTwo](container) // panic
err := container.Invoke(func(y YourInterfaces) {
    y // YourInterfaces
})
err // nil
err := container.Invoke(func(y YourInterfaces, yTwo YourInterfacesTwo) {
})
err // error

// containerScope
di.IsRegistered[YourInterfaces](containerScope) // true
di.IsRegistered[YourInterfacesTwo](containerScope) // true
di.Resolve[YourInterfaces](containerScope) // YourInterfaces
di.Resolve[YourInterfacesTwo](containerScope) // YourInterfacesTwo
err := container.Invoke(func(y YourInterfaces, yTwo YourInterfacesTwo) {
    y // YourInterfaces
	yTwo // YourInterfacesTwo
})
err // nil
```