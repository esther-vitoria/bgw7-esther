package application

// Application is an interface that represents an application.
type Application interface {
	TearDown()
	// Run runs the application.
	Run() (err error)
	// SetUp sets up the application.
	SetUp() (err error)
}
