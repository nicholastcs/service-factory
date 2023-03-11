package servicefactory

// Service describes the simple lifecycle control of an
// struct.
type Service interface {
	// Starts service. Functions that implements it
	// should not be blocking in nature. You may use
	// goroutine to run your services.
	Start() error

	// Stops service.
	Stop() error
}

// Logger describes the interface to be used by an struct
// that implementing Service. You might want to create
// logging wrappers for this interface.
type Logger interface {
	// Logs TRACE message.
	Trace(msg string, args ...interface{})

	// Logs DEBUG message.
	Debug(msg string, args ...interface{})

	// Logs INFO message.
	Info(msg string, args ...interface{})

	// Logs WARN message.
	Warn(msg string, args ...interface{})

	// Logs ERROR message.
	Error(msg string, args ...interface{})
}
