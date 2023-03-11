package test

import (
	"testing"

	servicefactory "github.com/nicholastcs/service-factory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExampleService(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		require.NotPanics(t, func() {
			_ = newExampleService(logger)
		})
	})

	t.Run("lifecycle", func(t *testing.T) {
		assert.NotPanics(t, func() {
			service := newExampleService(logger)
			service.Start()
			service.Stop()
		})
	})
}

func TestFactory(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		require.NotPanics(t, func() {
			_ = servicefactory.NewServiceFactory(logger)
		})
	})

	t.Run("add_service", func(t *testing.T) {
		assert.NotPanics(t, func() {
			factory := servicefactory.NewServiceFactory(logger)
			factory.Add("s1", newExampleService(logger))
		})
	})

	t.Run("full_lifecycle_service_factory", func(t *testing.T) {
		assert.NotPanics(t, func() {
			service := newExampleService(logger)

			factory := servicefactory.NewServiceFactory(logger)
			factory.Add("s1", service)

			assert.NoError(t, factory.Start())
			assert.True(t, service.Started)

			assert.NoError(t, factory.Stop())
			assert.False(t, service.Started)
		})
	})

	t.Run("add_same_name_will_panic", func(t *testing.T) {
		assert.Panics(t, func() {
			factory := servicefactory.NewServiceFactory(logger)
			factory.Add("s1", newExampleService(logger))
			factory.Add("s1", newExampleService(logger))
		})
	})
}

func TestFactoryErrorEmit(t *testing.T) {
	t.Run("emit_single_error", func(t *testing.T) {
		service := newExampleService(logger)
		service.emitErrorAtStart = true

		factory := servicefactory.NewServiceFactory(logger)
		factory.Add("s1", service)

		err := factory.Start()
		assert.Error(t, err)
		assert.Equal(t, s1Err, err.Error())
	})

	t.Run("emit_two_error", func(t *testing.T) {
		service := newExampleService(logger)
		service.emitErrorAtStart = true

		service2 := newExampleService(logger)
		service2.emitErrorAtStart = true

		factory := servicefactory.NewServiceFactory(logger)
		factory.Add("s1", service)
		factory.Add("s2", service2)

		err := factory.Start()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), s1Err)
		assert.Contains(t, err.Error(), s2Err)
	})
}
