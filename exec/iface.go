package exec

//go:generate mockery -name=Execer -case=underscore -dir=. -output=../z_mocks -outpkg=z_mocks

// Execer ...
type Execer interface {
	Start() error
	Stop()
}
