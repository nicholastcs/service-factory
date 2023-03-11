package servicefactory

import "fmt"

// ServiceFactory is a struct that encapsulates multiple
// services that implements interface Service.
//
// ServiceFactory implements interface Service too.
type ServiceFactory struct {
	// Service by name.
	servicesByName map[string]Service

	// Logger interface.
	Logger Logger
}

// NewServiceFactory news a ServiceFactory.
func NewServiceFactory(logger Logger) *ServiceFactory {
	return &ServiceFactory{
		Logger:         logger,
		servicesByName: map[string]Service{},
	}
}

// Starts all service. Blocking if one of the underlying
// services is blocking in nature. You may use goroutine to
// run your services.
func (factory *ServiceFactory) Start() error {
	err := factory.forEachService(
		func(name string, service Service) error {
			factory.Logger.Info("starting service...", "service", name)
			return service.Start()
		},
	)
	return err
}

// Stops all service.
func (factory *ServiceFactory) Stop() error {
	err := factory.forEachService(
		func(name string, service Service) error {
			factory.Logger.Info("stopping service...", "service", name)
			return service.Stop()
		},
	)
	return err
}

func (factory *ServiceFactory) Add(name string, service Service) {
	_, exists := factory.servicesByName[name]
	if exists {
		factory.Logger.Error("service name already exists!", "duplicate", name)
		panic("application is in invalid state...")
	}

	factory.servicesByName[name] = service
}

// forEachService performs lifecycle actions and aggregates
// errors into multiError.
func (factory *ServiceFactory) forEachService(delegate func(name string, service Service) error) error {
	var errLiterals []string

	for name, service := range factory.servicesByName {
		err := delegate(name, service)
		if err != nil {
			errLiterals = append(
				errLiterals,
				fmt.Sprintf("unable to start service name=`%s` cause=`%v`",
					name, err.Error(),
				),
			)
		}

	}

	if len(errLiterals) > 0 {
		return newErrors(errLiterals)
	}
	return nil
}
