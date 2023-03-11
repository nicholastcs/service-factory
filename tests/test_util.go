package test

import (
	"fmt"

	"github.com/hashicorp/go-hclog"
	servicefactory "github.com/nicholastcs/service-factory"
)

var s1Err string = "unable to start service name=`s1` cause=`error emitted, because service.emitErrorAtStart=true`\n"
var s2Err string = "unable to start service name=`s2` cause=`error emitted, because service.emitErrorAtStart=true`\n"

// Example use of service-factory package.

// Native use of go-hclog. Create logger wrapper if not compatible.
var logger hclog.Logger = hclog.New(&hclog.LoggerOptions{
	Name:       "example",
	Level:      hclog.LevelFromString("debug"),
	JSONFormat: false,
})

type exampleService struct {
	logger servicefactory.Logger

	emitErrorAtStart bool
	emitErrorAtStop  bool

	Started bool
}

func newExampleService(logger servicefactory.Logger) *exampleService {
	return &exampleService{
		logger:  logger,
		Started: false,
	}
}

func (service *exampleService) Start() error {
	service.Started = true
	service.logger.Debug("service starts...")

	if service.emitErrorAtStart {
		return fmt.Errorf("error emitted, because service.emitErrorAtStart=%v", service.emitErrorAtStart)
	}

	return nil
}

func (service *exampleService) Stop() error {
	service.Started = false
	service.logger.Debug("service stops...")

	if service.emitErrorAtStart {
		return fmt.Errorf("error emitted, because service.emitErrorAtStop=%v", service.emitErrorAtStop)
	}

	return nil
}
