package idgenerator

type newService func() interface{}

var serviceFactories = make(map[string]newService)

func RegisterService(name string, factory newService) {
	serviceFactories[name] = factory
}

func ServiceInstance(name string) newService {
	return serviceFactories[name]
}
